package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func SaveTemplate(ctx context.Context, args []string) error {
	var templateName string

	if len(args) >= 1 {
		templateName = args[0]
	}

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	_, err = mgr.SaveSessionAsTemplate(ctx, manager.SaveSessionAsTemplateArguments{
		TemplateName: templateName,
	})

	if err != nil {
		return err
	}

	return nil
}
