package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

func ListSessions(ctx context.Context, args []string) error {
	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	sessions, err := manager.ListSessions(ctx)
	if err != nil {
		return err
	}

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	if len(sessions) < 1 {
		logger.Println("no sessions found")
		return nil
	}

	for _, session := range sessions {
		logger.Printf("session: %s\n", session.String())
	}

	return nil
}
