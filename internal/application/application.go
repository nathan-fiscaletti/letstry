package application

import (
	"context"
	"os"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

// NewApplication creates a new application instance
func NewApplication(ctx context.Context) *application {
	// Initialize base line logging, writing to the console.
	logger, err := logging.New(&logging.LoggerConfig{
		LogMode: logging.LogModeConsole,
	})
	if err != nil {
		panic(err)
	}

	args, err := arguments.ParseArguments()
	if err != nil {
		logger.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	// Update the logging based on the command passed.
	// If it is a private command, write to a file only.
	for _, cmd := range arguments.PrivateCommands() {
		if args.Command == cmd {
			if logFile := logger.File(); logFile != nil {
				logFile.Close()
			}

			logger, err = logging.New(&logging.LoggerConfig{
				LogMode: logging.LogModeFile,
			})
			if err != nil {
				panic(err)
			}
			break
		}
	}

	// Update the context with the logger.
	ctx = logging.ContextWithLogger(ctx, logger)

	return &application{
		arguments: args,
		context:   ctx,
	}
}

type application struct {
	context   context.Context
	arguments arguments.Arguments
}

// Start starts the application
func (a *application) Start() {
	logger, err := logging.LoggerFromContext(a.GetContext())
	if err != nil {
		panic(err)
	}

	// Close the logger when the application closes.
	defer func() {
		if logFile := logger.File(); logFile != nil {
			logFile.Close()
		}
	}()

	manager := session_manager.GetSessionManager()
	args := a.GetArguments()

	switch args.Command {
	case arguments.CommandNewSession:
		session, err := manager.CreateSession(*args.CreateSessionArguments)
		if err != nil {
			logger.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		logger.Printf("Session %s created with PID %d\n", session.Arguments.SessionName, session.PID)

	case arguments.CommandListSessions:
		sessions, err := manager.ListSessions(*args.ListSectionsArguments)
		if err != nil {
			logger.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		for _, session := range sessions {
			logger.Printf("Session %s with PID %d\n", session.Arguments.SessionName, session.PID)
		}

	case arguments.CommandMonitorSession:
		err := manager.MonitorSession(a.GetContext(), *args.MonitorSessionsArguments)
		if err != nil {
			logger.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

// GetArguments returns the application arguments
func (a *application) GetArguments() arguments.Arguments {
	return a.arguments
}

// GetContext returns the application context
func (a *application) GetContext() context.Context {
	return a.context
}
