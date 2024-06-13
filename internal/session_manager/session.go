package session_manager

import (
	"github.com/nathan-fiscaletti/letstry/internal/arguments"
)

type session struct {
	arguments arguments.Arguments
	// config
}

// GetArguments returns the session arguments
func (s *session) GetArguments() arguments.Arguments {
	return s.arguments
}

// Start starts the session
func (s *session) Start() {
	// do something
}
