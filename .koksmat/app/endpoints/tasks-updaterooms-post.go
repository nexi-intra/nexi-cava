// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Update Rooms
---
*/
package endpoints

import (
	"context"

	"github.com/365admin/nexi-cava/execution"
	"github.com/swaggest/usecase"
)

func TasksUpdateroomsPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "nexi-cava", "20-tasks", "10-step1.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Update Rooms")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Tasks")
	return u
}
