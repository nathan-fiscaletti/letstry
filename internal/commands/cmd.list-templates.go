package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

func ListTemplatesHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: templates -- List templates

Usage: 

    ` + cmdName + ` templates

Description:

    Lists available templates.

Run '` + cmdName + ` help' for information on additional commands.
`
}

func ListTemplates(ctx context.Context, args []string) error {
	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	templates, err := manager.ListTemplates(ctx)
	if err != nil {
		return err
	}

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	if len(templates) < 1 {
		logger.Println("no templates found")
		return nil
	}

	for _, template := range templates {
		logger.Printf("template: %s\n", template.FormattedString(ctx))
	}

	return nil
}
