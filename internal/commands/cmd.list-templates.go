package commands

import (
	"context"

	"github.com/fatih/color"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type ListTemplatesCommand struct {
	Arguments arguments.Parameters
}

func (c ListTemplatesCommand) Execute(ctx context.Context) error {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	manager := session_manager.GetSessionManager()

	templates, err := manager.ListTemplates(ctx)
	if err != nil {
		return err
	}

	logger.Printf("Templates:\n")

	if len(templates) == 0 {
		logger.Printf(color.RedString("No templates found"))
		return nil
	}

	for idx, template := range templates {
		logger.Printf("%d: %s\n", idx+1, template.String())
	}

	return nil
}
