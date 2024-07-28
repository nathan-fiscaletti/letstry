package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

// NewSessionCommand returns a new command for creating a new session.
func NewSessionCommand() Command {
	return Command{
		Name:             CommandNewSession,
		ShortDescription: "Create a new session",
		Description:      "Create a new session using the specified source.",
		Arguments: []Argument{
			{
				Name:        "source",
				Description: "The source to use for the new session. Can be a git repository URL, a path to a directory, or the name of a letstry template.\n\nIf source is not provided, the session will be created from a blank source.",
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			var source string

			if len(args) >= 1 {
				source = args[0]
			}

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			_, err = mgr.CreateSession(ctx, manager.CreateSessionArguments{
				Source: source,
			})
			if err != nil {
				return err
			}

			return nil
		},
	}
}
