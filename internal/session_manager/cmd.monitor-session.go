package session_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
)

type MonitorSessionArguments struct {
	Delay    time.Duration
	Location string
}

func (s *sessionManager) MonitorSession(ctx context.Context, args MonitorSessionArguments) error {
	// delay the start of the monitoring
	time.Sleep(args.Delay)

	// Start monitoring the session
	return s.monitorDirectoryAccessible(args.Location, func() error {
		session, err := s.GetSessionForPath(ctx, args.Location)
		if err != nil {
			return err
		}

		logger, err := logging.LoggerFromContext(ctx)
		if err != nil {
			return err
		}

		logger.Printf("cleaning up session: %s (directory no longer being accessed)\n", session.ID)

		err = s.removeSession(ctx, session.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *sessionManager) monitorDirectoryAccessible(path string, callback func() error) error {
	for {
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return callback()
			}

			return err
		}

		// try moving the directory, if we are able to move it, then it is no longer being accessed
		now := time.Now()
		newPath := fmt.Sprintf("%s-%d", path, now.Unix())
		err = os.Rename(path, newPath)
		if err == nil {
			// move the directory back
			err = os.Rename(newPath, path)
			if err != nil {
				return err
			}

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
