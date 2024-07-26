package session_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
	"github.com/shirou/gopsutil/process"
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

func (s *sessionManager) monitorProcessClosed(pid int32, callback func() error) error {
	p, err := process.NewProcess(pid)
	if err != nil {
		return err
	}

	for {
		exists, err := p.IsRunning()
		if err != nil {
			return err
		}

		if !exists {
			return callback()
		}

		time.Sleep(1 * time.Second) // Check every second
	}
}

func (s *sessionManager) removeSession(ctx context.Context, id identifier.ID) error {
	sessions, err := s.ListSessions(ctx)
	if err != nil {
		return err
	}

	for i, session := range sessions {
		if session.ID == id {
			// Remove the session
			sessions = append(sessions[:i], sessions[i+1:]...)

			file, err := s.storage.OpenFile("sessions.json")
			if err != nil {
				return fmt.Errorf("failed to open sessions file: %v", err)
			}
			defer file.Close()

			data, err := json.MarshalIndent(sessions, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to marshal sessions: %v", err)
			}

			err = file.Truncate(0)
			if err != nil {
				return fmt.Errorf("failed to truncate sessions file: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("failed to write sessions: %v", err)
			}

			err = file.Sync()
			if err != nil {
				return fmt.Errorf("failed to sync sessions file: %v", err)
			}

			// Give the process manager time to settle
			time.Sleep(1 * time.Second)

			// Remove the temporary directory
			err = os.RemoveAll(session.Location)
			if err != nil {
				return fmt.Errorf("failed to remove temporary directory: %v", err)
			}

			return nil
		}
	}

	return fmt.Errorf("session with id %s not found", id)
}
