package manager

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/config/editors"
)

func (m *manager) GetDefaultEditor(ctx context.Context) (editors.Editor, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return editors.Editor{}, err
	}

	return cfg.GetDefaultEditor()
}
