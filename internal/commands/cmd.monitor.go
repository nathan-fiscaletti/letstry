package commands

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingArgumentDelay    = errors.New("monitor: missing required argument 'delay'")
	ErrMissingArgumentLocation = errors.New("monitor: missing required argument 'location'")
)

func Monitor(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingArgumentDelay
	}

	if len(args) < 2 {
		return ErrMissingArgumentLocation
	}

	delay, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}

	location := args[1]

	_, err = os.Stat(location)
	if err != nil {
		return err
	}

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	return mgr.MonitorSession(ctx, manager.MonitorSessionArguments{
		Delay:    delay,
		Location: location,
	})
}
