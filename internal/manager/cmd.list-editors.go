package manager

import (
	"context"

	"github.com/letstrygo/letstry/internal/config"
	"github.com/letstrygo/letstry/internal/config/editors"
)

func (s *manager) ListEditors(ctx context.Context) ([]editors.Editor, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	return cfg.AvailableEditors, nil
}
