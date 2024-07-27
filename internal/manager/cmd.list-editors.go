package manager

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/config"
)

func (s *manager) ListEditors(ctx context.Context) ([]config.Editor, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	return cfg.AvailableEditors, nil
}
