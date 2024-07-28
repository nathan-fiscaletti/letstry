package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingExportPath = errors.New("missing export path")
)

func ExportSessionCommand() Command {
	return Command{
		Name:                 CommandExportSession,
		ShortDescription:     "Export the current session",
		Description:          "This command must be run from within a session. It will export the current session to the specified path.",
		MustBeRunFromSession: true,
		Arguments: []Argument{
			{
				Name:        "path",
				Description: "The path to export the session to.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return ErrMissingExportPath
			}

			exportPath := args[0]

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			return mgr.ExportSession(ctx, manager.ExportSessionArguments{
				Path: exportPath,
			})
		},
	}
}
