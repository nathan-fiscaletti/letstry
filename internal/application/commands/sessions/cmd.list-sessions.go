package sessions

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application/commands"
	"github.com/nathan-fiscaletti/letstry/internal/cli"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func ListSessionsCommand() cli.Command {
	return cli.Command{
		Name:             commands.CommandListSessions.String(),
		ShortDescription: "List running sessions",
		Description:      "This command will list all currently running sessions.",
		Executor: func(ctx context.Context, args []string) error {
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
		},
	}
}
