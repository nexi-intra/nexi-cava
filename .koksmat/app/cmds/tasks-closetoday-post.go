// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Close Today
---
*/
package cmds

import (
	"context"

	"github.com/365admin/nexi-cava/execution"
)

func TasksClosetodayPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "nexi-cava", "20-tasks", "20-close-today.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}