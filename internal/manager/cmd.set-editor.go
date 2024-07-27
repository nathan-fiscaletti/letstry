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

	var found bool
	var editor config.Editor
	for _, availableEditor := range cfg.AvailableEditors {
		if availableEditor.Name == editorName {
			found = true
			editor = availableEditor
			break
		}
	}

	if !found {
		return ErrEditorNotFound
	}

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	logger.Printf("Setting default editor to: %s\n", editor.String())
	cfg.DefaultEditorName = editorName
	return config.SaveConfig(cfg)
}
