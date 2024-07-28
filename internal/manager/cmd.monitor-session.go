package manager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/process"

	"github.com/nathan-fiscaletti/letstry/internal/config/editors"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
)

var (
	ErrUnknownTrackingType = errors.New("unknown tracking type")
)

type MonitorSessionArguments struct {
	Delay        time.Duration
	TrackingType editors.TrackingType
	PID          int
	Location     string
}

func (s *manager) MonitorSession(ctx context.Context, args MonitorSessionArguments) error {
	// delay the start of the monitoring
	time.Sleep(args.Delay)

	handler := func() error {
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
	}

	switch args.TrackingType {
	case editors.TrackingTypeProcess:
		return s.monitorProcess(args.PID, handler)
	case editors.TrackingTypeFileAccess:
		_, err := os.Stat(args.Location)
		if err != nil {
			return err
		}
		return s.monitorDirectoryAccessible(args.Location, handler)
	}

	return ErrUnknownTrackingType
}

func (s *manager) monitorProcess(pid int, callback func() error) error {
	for {
		_, err := process.NewProcess(int32(pid))
		if err != nil {
			return callback()
		}

		time.Sleep(1 * time.Second) // Check every second
	}
}

func (s *manager) monitorDirectoryAccessible(path string, callback func() error) error {
	for {
		if !isInUse(path) {
			return callback()
		}

		time.Sleep(1 * time.Second) // Check every second
	}
}

func isInUse(path string) bool {
	switch runtime.GOOS {
	case "windows":
		newPath := fmt.Sprintf("%s-%v", path, time.Now().Unix())
		err := os.Rename(path, newPath)
		if err != nil {
			return true
		}

		_ = os.Rename(newPath, path)
		return false
	default:
		cmd := exec.Command("lsof", path)

		if err := cmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				ec := exitError.ExitCode()
				if ec == 1 {
					return false
				}
			}
		}

		return true
	}
}

func (s *manager) removeSession(ctx context.Context, id identifier.ID) error {
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
