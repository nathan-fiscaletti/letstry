package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

func NewSession(ctx context.Context, args []string) error {
	var source string

	if len(args) >= 1 {
		source = args[0]
	}

	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	_, err = manager.CreateSession(ctx, session_manager.CreateSessionArguments{
		Source: source,
	})
	if err != nil {
		return err
	}

	return nil
}
