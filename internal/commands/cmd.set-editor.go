package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingEditorName = errors.New("missing editor name")
)

func SetEditor(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingEditorName
	}

	editorName := args[0]

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	return mgr.SetDefaultEditor(ctx, editorName)
}
