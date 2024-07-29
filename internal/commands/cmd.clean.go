package commands

import (
	"context"

	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
)

var (
	ErrMissingSessionID = errors.New("missing session id")
)

func CleanCommand() Command {
	return Command{
		Name:             CommandClean,
		ShortDescription: "Remove a dangling session",
		Description:      "This command removes all files for the specified session. This command is useful for cleaning up a session that was not properly closed.",
		Arguments: []Argument{
			{
				Name:        "session-id",
				Description: "The session to remove files for.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return ErrMissingSessionID
			}

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			return mgr.CleanSession(ctx, manager.CleanSessionArguments{
				SessionID: identifier.ID(args[0]),
			})
		},
	}
}
