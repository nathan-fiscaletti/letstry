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
	args := a.GetArguments()

	switch args.Command {
	case arguments.CommandNewSession:
		session, err := manager.CreateSession(*args.CreateSessionArguments)
		if err != nil {
			logging.GetLogger().Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		logging.GetLogger().Printf("Session %s created with PID %d\n", session.Arguments.SessionName, session.PID)

	case arguments.CommandListSessions:
		sessions, err := manager.ListSessions(*args.ListSectionsArguments)
		if err != nil {
			logging.GetLogger().Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		for _, session := range sessions {
			logging.GetLogger().Printf("Session %s with PID %d\n", session.Arguments.SessionName, session.PID)
		}

	case arguments.CommandMonitorSession:
		err := manager.MonitorSession(*args.MonitorSessionsArguments)
		if err != nil {
			logging.GetLogger().Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

// GetArguments returns the application arguments
func (a *application) GetArguments() arguments.Arguments {
	return a.arguments
}
