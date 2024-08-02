package sessions

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/application/commands"
	"github.com/nathan-fiscaletti/letstry/internal/cli"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingExportPath = errors.New("missing export path")
)

func ExportSessionCommand() cli.Command {
	return cli.Command{
		Name:                 commands.CommandExportSession.String(),
		ShortDescription:     "Export the current session",
		Description:          "This command must be run from within a session. It will export the current session to the specified path.",
		MustBeRunFromSession: true,
		Arguments: []cli.Argument{
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
