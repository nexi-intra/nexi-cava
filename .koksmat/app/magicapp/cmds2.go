package magicapp

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"github.com/spf13/cobra"
	"github.com/365admin/nexi-cava/endpoints"
	"github.com/365admin/nexi-cava/cmds"

)

func RegisterCmds(){
	magicCmd := &cobra.Command{
	   Use:   "magic",
	   Short: "Magic Buttons",
	   Long:  `Collection of stuff build for nexi "CAVA"`,
   }
	MagicSubmitOrdersPostCmd := &cobra.Command{
		Use:   "submit-orders",
		Short: "Submit Orders",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.MagicSubmitOrdersPost(ctx,args)
		},
	}
	magicCmd.AddCommand(MagicSubmitOrdersPostCmd)
	MagicSubmitOrdersPostCmd := &cobra.Command{
		Use:   "submit-orders",
		Short: "Submit Orders",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.MagicSubmitOrdersPost(ctx,args)
		},
	}
	magicCmd.AddCommand(MagicSubmitOrdersPostCmd)
	MagicInvoicePostCmd := &cobra.Command{
		Use:   "invoice",
		Short: "Invoice",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.MagicInvoicePost(ctx,args)
		},
	}
	magicCmd.AddCommand(MagicInvoicePostCmd)
 
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
	TasksGenerateSampleDataPostCmd := &cobra.Command{
		Use:   "generate-sample-data",
		Short: "Generate Sample Data",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.TasksGenerateSampleDataPost(ctx,args)
		},
	}
	tasksCmd.AddCommand(TasksGenerateSampleDataPostCmd)
	TasksInvoicePostCmd := &cobra.Command{
		Use:   "invoice",
		Short: "Invoice",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.TasksInvoicePost(ctx,args)
		},
	}
	tasksCmd.AddCommand(TasksInvoicePostCmd)
	TasksInvoicePostCmd := &cobra.Command{
		Use:   "invoice",
		Short: "Invoice",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			
			cmds.TasksInvoicePost(ctx,args)
		},
	}
	tasksCmd.AddCommand(TasksInvoicePostCmd)
 
RootCmd.AddCommand(tasksCmd)
	buildCmd := &cobra.Command{
	   Use:   "build",
	   Short: "Build",
	   Long:  `Collection of stuff build for nexi "CAVA"`,
   }
 
RootCmd.AddCommand(buildCmd)
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
			
			cmds.ProvisionWebdeployproductionPost(ctx,args)
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
