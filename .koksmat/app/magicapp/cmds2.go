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
	outlook-addinCmd := &cobra.Command{
	   Use:   "outlook-addin",
	   Short: "outlook-addin",
	   Long:  `Collection of stuff build for nexi "CAVA"`,
   }
 
RootCmd.AddCommand(outlook-addinCmd)
	powerappsCmd := &cobra.Command{
	   Use:   "powerapps",
	   Short: "powerapps",
	   Long:  `Collection of stuff build for nexi "CAVA"`,
   }
 
RootCmd.AddCommand(powerappsCmd)
	reactappsCmd := &cobra.Command{
	   Use:   "reactapps",
	   Short: "reactapps",
	   Long:  `Collection of stuff build for nexi "CAVA"`,
   }
 
RootCmd.AddCommand(reactappsCmd)
}
