package session_manager

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

func (s *sessionManager) DeleteTemplate(ctx context.Context, t Template) error {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.storage.DeleteDirectory(t.StoragePath())
	if err != nil {
		return err
	}

	logger.Printf("deleted template: %s\n", t.String())
	return nil
}
