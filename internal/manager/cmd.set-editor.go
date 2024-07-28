package manager

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

var (
	ErrEditorNotFound = errors.New("editor not found")
)

func (s *manager) SetDefaultEditor(ctx context.Context, editorName string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	editor, err := cfg.GetEditor(editorName)
	if err != nil {
		return err
	}

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	logger.Printf("Setting default editor to: %s\n", editor.String())
	cfg.DefaultEditorName = editor.Name
	return config.SaveConfig(cfg)
}
