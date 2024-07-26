package session_manager

import (
	"context"
	"fmt"
	"os"

	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
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
