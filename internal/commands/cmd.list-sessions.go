package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func ListSessionsHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: list -- List running sessions

Usage: 

    ` + cmdName + ` list

Description:

    This command will list all currently running sessions.

Run '` + cmdName + ` help' for information on additional commands.
`
}

func ListSessions(ctx context.Context, args []string) error {
	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	sessions, err := mgr.ListSessions(ctx)
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
