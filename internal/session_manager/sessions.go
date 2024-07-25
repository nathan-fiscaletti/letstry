package session_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
	"github.com/shirou/gopsutil/process"
)

var (
	ErrCurrentDirectoryIsNotASession = fmt.Errorf("current directory is not a session")
	ErrSessionNotFound               = fmt.Errorf("session not found")
)

// GetSession returns the session with the given ID
func (s *sessionManager) GetSession(ctx context.Context, id identifier.ID) (session, error) {
	sessions, err := s.ListSessions(ctx)
	if err != nil {
		return session{}, err
	}

	for _, session := range sessions {
		if session.ID == id {
			return session, nil
		}
	}

	return session{}, fmt.Errorf("session with ID %s not found", id)
}

// GetCurrentSession returns the session for the current working directory
func (s *sessionManager) GetCurrentSession(ctx context.Context) (session, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return session{}, fmt.Errorf("failed to get current working directory: %v", err)
	}

	sess, err := s.GetSessionForPath(ctx, cwd)
	if err != nil && err == ErrSessionNotFound {
		return sess, ErrCurrentDirectoryIsNotASession
	}

	return sess, err
}

// GetSessionForPID returns the session with the given PID
func (s *sessionManager) GetSessionForPID(ctx context.Context, pid int) (session, error) {
	return s.GetSessionForPredicate(ctx, func(sess session) bool {
		return sess.PID == int32(pid)
	})
}

// GetSessionForPath returns the session for the given path
func (s *sessionManager) GetSessionForPath(ctx context.Context, path string) (session, error) {
	return s.GetSessionForPredicate(ctx, func(sess session) bool {
		return sess.Location == path
	})
}

// GetSessionForPredicate returns the session that matches the given predicate
func (s *sessionManager) GetSessionForPredicate(ctx context.Context, predicate func(session) bool) (session, error) {
	// get the list of sessions
	sessions, err := s.ListSessions(ctx)
	if err != nil {
		return session{}, err
	}

	// find the session with the same location
	for _, session := range sessions {
		if predicate(session) {
			return session, nil
		}
	}

	return session{}, ErrSessionNotFound
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

func (s *sessionManager) addSession(ctx context.Context, sess session) error {
	sessions, err := s.ListSessions(ctx)
	if err != nil {
		return err
	}

	// check if the session already exists by the same name
	for _, session := range sessions {
		if session.ID == sess.ID {
			return fmt.Errorf("session with ID %s already exists", sess.ID)
		}
	}

	// add the session to the list of sessions
	sessions = append(sessions, sess)

	// save the sessions
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
		return fmt.Errorf("failed to truncate sessions: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write sessions: %v", err)
	}

	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync sessions file: %v", err)
	}

	return nil
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
