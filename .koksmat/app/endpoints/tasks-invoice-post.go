// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Invoice
---
*/
package endpoints

import (
	"context"

	"github.com/365admin/nexi-cava/execution"
	"github.com/swaggest/usecase"
)

func TasksInvoicePost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "nexi-cava", "20-tasks", "30-invoice.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Invoice")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Tasks")
	return u
}
