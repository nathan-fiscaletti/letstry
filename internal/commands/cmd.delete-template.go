package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingTemplateName = errors.New("missing template name")
)

func DeleteTemplate(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingTemplateName
	}

	templateName := args[0]

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	template, err := mgr.GetTemplate(ctx, templateName)
	if err != nil {
		return err
	}

	err = mgr.DeleteTemplate(ctx, template)
	if err != nil {
		return err
	}

	return nil
}
