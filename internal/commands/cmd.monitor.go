package commands

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/config/editors"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingArgumentDelay        = errors.New("monitor: missing required argument 'delay'")
	ErrMissingArgumentTrackingType = errors.New("monitor: missing required argument 'tracking-type'")
	ErrMissingArgumentTrackingData = errors.New("monitor: missing required argument 'tracking-data'")
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
				Name:        "tracking-data",
				Description: "The tracking data required for the monitor to run.",
				Required:    true,
			},
			{
				Name:        "tracking-type",
				Description: "The type of tracking data required for the monitor to run. Can be one of the following: 'file_access', 'process'.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return ErrMissingArgumentDelay
			}

			if len(args) < 2 {
				return ErrMissingArgumentTrackingData
			}

			if len(args) < 3 {
				return ErrMissingArgumentTrackingType
			}

			delay, err := time.ParseDuration(args[0])
			if err != nil {
				return err
			}

			trackingData := args[1]

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			trackingType, err := editors.GetTrackingType(args[2])
			if err != nil {
				return err
			}

			var location string
			var pid int

			switch trackingType {
			case editors.TrackingTypeFileAccess:
				location = trackingData
			case editors.TrackingTypeProcess:
				pid, err = strconv.Atoi(trackingData)
				if err != nil {
					return err
				}
			}

			return mgr.MonitorSession(ctx, manager.MonitorSessionArguments{
				Delay:        delay,
				TrackingType: trackingType,
				PID:          pid,
				Location:     location,
			})
		},
	}
}
