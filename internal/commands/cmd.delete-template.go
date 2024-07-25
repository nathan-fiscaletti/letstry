package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

var (
	ErrMissingTemplateName = errors.New("missing template name")
)

func DeleteTemplate(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingTemplateName
	}

	templateName := args[0]

	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	template, err := manager.GetTemplate(ctx, templateName)
	if err != nil {
		return err
	}

	err = manager.DeleteTemplate(ctx, template)
	if err != nil {
		return err
	}

	return nil
}
