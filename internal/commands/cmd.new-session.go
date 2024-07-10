package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type NewSessionCommand struct {
	Arguments *arguments.CreateSessionArguments
}

func (c NewSessionCommand) Execute(ctx context.Context) error {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	manager := session_manager.GetSessionManager()

	session, err := manager.CreateSession(ctx, *c.Arguments)
	if err != nil {
		return err
	}

	logger.Printf("%s\n", session.String())
	return nil
}
