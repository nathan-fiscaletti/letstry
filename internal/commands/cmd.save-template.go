package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func SaveTemplateHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: save -- Saves the current session as a template

Usage: 

    ` + cmdName + ` save [template-name]

Description:

    This command must be run from within a session. It will save the current
    session as a template with the specified name. If no name is provided, and
    the session was created from a template, the template's name will be used.

Arguments:

    template-name (optional) - The name to use for the template. If not provided,
                               and the session was created from a template, the
                               template's name will be used.

Run '` + cmdName + ` help' for information on additional commands.
`
}

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
