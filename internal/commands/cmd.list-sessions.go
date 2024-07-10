package commands

import (
	"context"

	"github.com/fatih/color"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type ListSessionsCommand struct {
	Arguments *arguments.ListSessionsArguments
}

func (c ListSessionsCommand) Execute(ctx context.Context) error {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	manager := session_manager.GetSessionManager()

	sessions, err := manager.ListSessions(*c.Arguments)
	if err != nil {
		return err
	}

	logger.Printf("Sessions:\n")

	if len(sessions) == 0 {
		logger.Printf(color.RedString("No sessions found"))
		return nil
	}

	for idx, session := range sessions {
		logger.Printf("%d: %s\n", idx+1, session.String())
	}

	return nil
}
