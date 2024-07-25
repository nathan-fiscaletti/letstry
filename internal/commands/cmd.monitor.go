package commands

import (
	"context"
	"errors"
	"strconv"

	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

var (
	ErrMissingArgumentPID = errors.New("monitor: missing required argument PID")
)

func Monitor(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingArgumentPID
	}

	manager, err := session_manager.GetSessionManager(ctx)
	if err != nil {
		return err
	}

	intVal, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	return manager.MonitorSession(ctx, session_manager.MonitorSessionArguments{
		PID: intVal,
	})
}
