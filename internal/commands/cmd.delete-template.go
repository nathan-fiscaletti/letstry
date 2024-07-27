package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingTemplateName = errors.New("missing template name")
)

func DeleteTemplateHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: delete-template -- Delete a template

Usage: 

    ` + cmdName + ` delete-template <name>

Description:

    Delete a template by name.

Arguments:

	name - The name of the template to delete.

Run '` + cmdName + ` help' for information on additional commands.
`
}

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
