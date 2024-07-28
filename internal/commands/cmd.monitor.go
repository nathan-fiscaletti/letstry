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

func MonitorCommand() Command {
	return Command{
		Name:             CommandMonitor,
		ShortDescription: "Monitor a session",
		Description:      "This command is used to monitor a session. It is not intended to be run directly by the user.",
		LogToFile:        true,
		Arguments: []Argument{
			{
				Name:        "delay",
				Description: "The delay before initiating the monitor, formatted as a duration string. For example, '5s' for 5 seconds.",
				Required:    true,
			},
			{
				Name:        "location",
				Description: "The location to monitor. This should be the path to a letstry session directory.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
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
		},
	}
}
