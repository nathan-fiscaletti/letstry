package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingExportPath = errors.New("missing export path")
)

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
