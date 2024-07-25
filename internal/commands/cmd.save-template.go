package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

func SaveTemplate(ctx context.Context, args []string) error {
	var templateName string

	if len(args) >= 1 {
		templateName = args[0]
	}

	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	_, err = manager.SaveSessionAsTemplate(ctx, session_manager.SaveSessionAsTemplateArguments{
		TemplateName: templateName,
	})

	if err != nil {
		return err
	}

	return nil
}
