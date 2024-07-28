package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func ListTemplatesCommand() Command {
	return Command{
		Name:             CommandListTemplates,
		ShortDescription: "List all templates",
		Description:      "This command will list all available templates that can be used when creating a new session.",
		Executor: func(ctx context.Context, args []string) error {
			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			templates, err := mgr.ListTemplates(ctx)
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
		},
	}
}
