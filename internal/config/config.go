package config

import "fmt"

type Editor struct {
	Name     string `json:"name"`
	ExecPath string `json:"path"`
	Args     string `json:"args"`
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
