package workspace_repo

import (
	"context"

	"github.com/chanzuckerberg/happy/pkg/options"
	"github.com/hashicorp/go-tfe"
)

type WorkspaceRepoIface interface {
	GetWorkspace(ctx context.Context, workspaceName string) (Workspace, error)
}

type Workspace interface {
	GetWorkspaceID() string
	WorkspaceName() string
	GetCurrentRunID() string
	GetLatestConfigVersionID() (string, error)
	Run(isDestroy bool) error
	SetVars(key string, value string, description string, sensitive bool) error
	RunConfigVersion(configVersionId string, isDestroy bool) error
	Wait(ctx context.Context) error
	WaitWithOptions(ctx context.Context, waitOptions options.WaitOptions) error
	ResetCache()
	GetTags() (map[string]string, error)
	GetWorkspaceId() string
	GetOutputs() (map[string]string, error)
	GetCurrentRunStatus() string
	UploadVersion(targzFilePath string) (string, error)
	SetOutputs(map[string]string)          // For testing purposes only
	SetClient(tfc *tfe.Client)             // For testing purposes only
	SetWorkspace(workspace *tfe.Workspace) // For testing purposes only
}
