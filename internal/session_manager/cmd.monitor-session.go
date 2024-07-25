package session_manager

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

type MonitorSessionArguments struct {
	PID int
}

func (s *sessionManager) MonitorSession(ctx context.Context, args MonitorSessionArguments) error {
	// Start monitoring the session
	return s.monitorProcessClosed(int32(args.PID), func() error {
		session, err := s.GetSessionForPID(ctx, args.PID)
		if err != nil {
			return err
		}

		logger, err := logging.LoggerFromContext(ctx)
		if err != nil {
			return err
		}

		logger.Printf("cleaning up session: %s (process closed, PID %d)\n", session.ID, session.PID)

		err = s.removeSession(ctx, session.ID)
		if err != nil {
			return err
		}

		return nil
	})
}
