package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

var (
	ErrMissingExportPath = errors.New("missing export path")
)

func ExportSession(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingExportPath
	}

	exportPath := args[0]

	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	return manager.ExportSession(ctx, session_manager.ExportSessionArguments{
		Path: exportPath,
	})
}
