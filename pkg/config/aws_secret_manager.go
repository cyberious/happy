package config

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/pkg/errors"
)

type Secrets interface {
	GetTfeUrl() string
	GetTfeOrg() string
	GetClusterArn() string
	GetPrivateSubnets() []string
	GetSecurityGroups() []string
	GetServiceUrl(service string) (string, error)
	GetAllServicesUrl() map[string]*RegistryConfig
}

type SecretsBackend interface {
	GetSecrets(secretArn string) (Secrets, error)
	// GetRegistries() (map[string]*Registry, error)
}

type TfeSecrets struct {
	Url string `json:"url"`
	Org string `json:"org"`
}

type AwsSecretMgrSecrets struct {
	ClusterArn     string                     `json:"cluster_arn"`
	PrivateSubnets []string                   `json:"private_subnets"`
	SecurityGroups []string                   `json:"security_groups"`
	Services       map[string]*RegistryConfig `json:"ecrs"`
	Tfe            *TfeSecrets                `json:"tfe"`
}

func (s *AwsSecretMgrSecrets) GetClusterArn() string {
	return s.ClusterArn
}

func (s *AwsSecretMgrSecrets) GetPrivateSubnets() []string {
	return s.PrivateSubnets
}

func (s *AwsSecretMgrSecrets) GetSecurityGroups() []string {
	return s.SecurityGroups
}

func (s *AwsSecretMgrSecrets) GetServiceUrl(serviceName string) (string, error) {
	svc, ok := s.Services[serviceName]
	if !ok {
		return "", fmt.Errorf("Can't find service %q", serviceName)
	}

	return svc.Url, nil
}

func (s *AwsSecretMgrSecrets) GetAllServicesUrl() map[string]*RegistryConfig {
	return s.Services
}

func (s *AwsSecretMgrSecrets) GetTfeUrl() string {
	return s.Tfe.Url
}

func (s *AwsSecretMgrSecrets) GetTfeOrg() string {
	return s.Tfe.Org
}

type AwsSecretMgr struct {
	session   *session.Session
	awsConfig *aws.Config
	// secretmgrClient *secretsmanager.SecretsManager
	secretmgrClient secretsmanageriface.SecretsManagerAPI
	secrets         *AwsSecretMgrSecrets
}

var secretMgrSessInst SecretsBackend
var creatSecretMgeOnce sync.Once

func GetAwsSecretMgr(awsProfile string) SecretsBackend {
	creatSecretMgeOnce.Do(func() {
		awsConfig := &aws.Config{
			Region:     aws.String("us-west-2"),
			MaxRetries: aws.Int(2),
		}
		session := session.Must(session.NewSessionWithOptions(session.Options{
			Profile:           awsProfile,
			Config:            *awsConfig,
			SharedConfigState: session.SharedConfigEnable,
		}))
		secretmgrClient := secretsmanager.New(session)
		secretMgrSessInst = &AwsSecretMgr{
			session:         session,
			awsConfig:       awsConfig,
			secretmgrClient: secretmgrClient,
		}
	})
	return secretMgrSessInst
}

func GetAwsSecretMgrWithClient(client secretsmanageriface.SecretsManagerAPI) SecretsBackend {
	secretMgrSessInst = &AwsSecretMgr{
		secretmgrClient: client,
	}
	return secretMgrSessInst
}

func (s *AwsSecretMgr) GetSecrets(secretArn string) (Secrets, error) {
	if s.secrets != nil {
		return s.secrets, nil
	}
	configInput := secretArn
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &configInput,
	}
	secretOutput, err := s.secretmgrClient.GetSecretValue(input)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading from AWS secrets manager")
	}

	s.secrets = &AwsSecretMgrSecrets{}
	secrets := *secretOutput.SecretString
	json.Unmarshal([]byte(secrets), s.secrets)

	return s.secrets, nil
}

// func (s *AwsSecretMgr) GetRegistries() (map[string]*artifact_builder.Registry, error) {
// 	// _, err := s.GetSecrets()
// 	dummy := map[string]*artifact_builder.Registry{}
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return dummy, nil
// }
