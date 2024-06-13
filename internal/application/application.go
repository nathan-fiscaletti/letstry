package application

import (
	"os"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

// NewApplication creates a new application instance
func NewApplication() *application {
	args, err := arguments.ParseArguments()
	if err != nil {
		logging.GetLogger().Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	return &application{
		arguments: args,
	}
}

type application struct {
	arguments arguments.Arguments
}

// Start starts the application
func (a *application) Start() {
	manager := session_manager.GetSessionManager()
	arguments := a.GetArguments()

	session, err := manager.CreateSession(arguments.SessionName, arguments)
	if err != nil {
		logging.GetLogger().Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	logging.GetLogger().Printf("Session %s created with PID %d\n", session.Arguments.SessionName, session.PID)
}

// GetArguments returns the application arguments
func (a *application) GetArguments() arguments.Arguments {
	return a.arguments
}
