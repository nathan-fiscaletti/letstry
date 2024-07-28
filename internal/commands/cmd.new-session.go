package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func NewSession(ctx context.Context, args []string) error {
	var source string

	if len(args) >= 1 {
		source = args[0]
	}

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	_, err = mgr.CreateSession(ctx, manager.CreateSessionArguments{
		Source: source,
	})
	if err != nil {
		return err
	}

	return nil
}
