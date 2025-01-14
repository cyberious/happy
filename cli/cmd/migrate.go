package cmd

import (
	backend "github.com/chanzuckerberg/happy/cli/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/cli/pkg/cmd"
	"github.com/chanzuckerberg/happy/cli/pkg/config"
	"github.com/chanzuckerberg/happy/cli/pkg/orchestrator"
	stackservice "github.com/chanzuckerberg/happy/cli/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/cli/pkg/util"
	"github.com/chanzuckerberg/happy/cli/pkg/workspace_repo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var reset bool

func init() {
	rootCmd.AddCommand(migrateCmd)
	config.ConfigureCmdWithBootstrapConfig(migrateCmd)

	migrateCmd.Flags().BoolVar(&reset, "reset", false, "Resetting the task")
}

var migrateCmd = &cobra.Command{
	Use:          "migrate STACK_NAME",
	Short:        "Migrate stack",
	Long:         "Run migration tasks for stack with given name",
	SilenceUsage: true,
	PreRunE:      cmd.Validate(cobra.ExactArgs(1), cmd.CheckStackName),
	RunE: func(cmd *cobra.Command, args []string) error {
		stackName := args[0]
		return runMigrate(cmd, stackName)
	},
}

func runMigrate(cmd *cobra.Command, stackName string) error {
	ctx := cmd.Context()
	isDryRun := util.DryRunType(dryRun)
	bootstrapConfig, err := config.NewBootstrapConfig(cmd)
	if err != nil {
		return err
	}
	happyConfig, err := config.NewHappyConfig(bootstrapConfig)
	if err != nil {
		return err
	}

	b, err := backend.NewAWSBackend(ctx, happyConfig)
	if err != nil {
		return err
	}

	taskOrchestrator := orchestrator.NewOrchestrator().WithBackend(b).WithDryRun(isDryRun)

	url := b.Conf().GetTfeUrl()
	org := b.Conf().GetTfeOrg()

	workspaceRepo := workspace_repo.NewWorkspaceRepo(url, org).WithDryRun(isDryRun)
	stackService := stackservice.NewStackService().WithBackend(b).WithWorkspaceRepo(workspaceRepo)

	stacks, err := stackService.GetStacks(ctx)
	if err != nil {
		return err
	}
	stack, ok := stacks[stackName]
	if !ok {
		return errors.Errorf("stack %s not found", stackName)
	}

	if reset {
		err = taskOrchestrator.RunTasks(ctx, stack, backend.TaskTypeDelete)
		if err != nil {
			return err
		}
	}

	return taskOrchestrator.RunTasks(ctx, stack, backend.TaskTypeMigrate)
}
