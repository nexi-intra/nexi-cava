// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Publish Roomlist
---
*/
package cmds

import (
	"context"

	"github.com/365admin/nexi-cava/execution"
)

func TasksPublishRoomlistPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "nexi-cava", "20-tasks", "40-publish-roomlist.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
