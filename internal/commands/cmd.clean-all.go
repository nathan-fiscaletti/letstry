package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func CleanAllCommand() Command {
	return Command{
		Name:             CommandCleanAll,
		ShortDescription: "Clean any dangling sessions",
		Description:      "This command will clean any dangling sessions that are no longer in use. This command is useful for cleaning up any sessions that were not properly closed.",
		Executor: func(ctx context.Context, args []string) error {
			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			logger, err := logging.LoggerFromContext(ctx)
			if err != nil {
				return err
			}

			sessions, err := mgr.ListSessions(ctx)
			if err != nil {
				return err
			}

			inactiveSessions := []manager.Session{}
			for _, session := range sessions {
				if !session.IsActive() {
					inactiveSessions = append(inactiveSessions, session)
				}
			}

			if len(inactiveSessions) < 1 {
				logger.Printf("%s: no inactive sessions to clean\n", CommandCleanAll)
				return nil
			}

			for _, session := range inactiveSessions {
				err := mgr.CleanSession(ctx, manager.CleanSessionArguments{
					SessionID: session.ID,
				})
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
