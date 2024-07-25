package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
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
