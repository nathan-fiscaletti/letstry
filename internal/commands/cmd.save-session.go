package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type SaveSessionCommand struct {
	Arguments arguments.Parameters
}

func (c SaveSessionCommand) Execute(ctx context.Context) error {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	manager := session_manager.GetSessionManager()

	args := *c.Arguments.Arguments.(*arguments.SaveSessionArguments)
	_, err = manager.SaveTemplate(ctx, args)

	if err != nil {
		return err
	}

	logger.Printf("Session %s saved to template %s\n", args.SessionName, args.TemplateName)

	return nil
}
