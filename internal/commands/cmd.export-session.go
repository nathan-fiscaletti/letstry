package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingExportPath = errors.New("missing export path")
)

func ExportSessionHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: export -- Export the current session

Usage:

	` + cmdName + ` export <path>

Description:

	This command must be run from within a session. It will export the current
	session to the specified path.

Arguments:

	path - The path to export the session to.

Run '` + cmdName + ` help' for information on additional commands.
`
}

func ExportSession(ctx context.Context, args []string) error {
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
}
