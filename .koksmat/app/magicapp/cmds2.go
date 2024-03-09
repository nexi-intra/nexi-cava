package magicapp

import (
	"github.com/365admin/nexi-cava/cmds"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic Buttons",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}

	RootCmd.AddCommand(magicCmd)
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}

	RootCmd.AddCommand(setupCmd)
	tasksCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Tasks",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}
	TasksUpdateroomsPostCmd := &cobra.Command{
		Use:   "updaterooms",
		Short: "Update Rooms",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TasksUpdateroomsPost(ctx, args)
		},
	}
	tasksCmd.AddCommand(TasksUpdateroomsPostCmd)

	RootCmd.AddCommand(tasksCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}
	ProvisionWebdeployproductionPostCmd := &cobra.Command{
		Use:   "webdeployproduction",
		Short: "Web deploy to production",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeployproductionPostCmd)

	RootCmd.AddCommand(provisionCmd)
	decommissionCmd := &cobra.Command{
		Use:   "decommission",
		Short: "Decommision",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}

	RootCmd.AddCommand(decommissionCmd)
	legacyCmd := &cobra.Command{
		Use:   "legacy",
		Short: "Legacy",
		Long:  `Collection of stuff build for nexi "CAVA"`,
	}

	RootCmd.AddCommand(legacyCmd)
}
