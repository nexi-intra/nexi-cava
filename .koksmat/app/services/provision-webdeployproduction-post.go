// -------------------------------------------------------------------
// Generated by 365admin-publish 
// Service wrapper  v1
// -------------------------------------------------------------------
/*
---
title: Web deploy to production
---
*/
package cmds
import (
"context"
"encoding/json"
"os"
"path"
"github.com/nats-io/nats.go"
"github.com/nats-io/nats.go/micro"
"github.com/spf13/cobra"
"github.com/365admin/nexi-cava/schemas"
"github.com/365admin/nexi-cava/execution"
"github.com/365admin/nexi-cava/utils"
)
func ProvisionWebdeployproductionPost(ctx context.Context, args  []string)  ( *string, error) {

_, pwsherr := execution.ExecutePowerShell("john","*","nexi-cava","60-provision","20-web.ps1","" )
if (pwsherr != nil) {
	return nil,pwsherr
}
return nil, nil
	
// end result mapping

func init(){
	
}
}
