package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

type Editor struct {
	Name                string        `json:"name"`
	ExecPath            string        `json:"path"`
	Args                string        `json:"args"`
	ProcessCaptureDelay time.Duration `json:"process_capture_delay"`
}

func (e Editor) GetExecName() string {
	return filepath.Base(e.ExecPath)
}

func (e Editor) String() string {
	return color.BlueString(fmt.Sprintf("(%s, %s)", e.Name, e.GetExecName()))
}

type Config struct {
	DefaultEditorName string   `json:"default_editor"`
	AvailableEditors  []Editor `json:"editors"`
}

func (cfg Config) GetDefaultEditor() (Editor, error) {
	for _, editor := range cfg.AvailableEditors {
		if editor.Name == cfg.DefaultEditorName {
			return editor, nil
		}
	}

	return Editor{}, fmt.Errorf("editor %s not found", cfg.DefaultEditorName)
}
