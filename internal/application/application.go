package application

import (
	"context"
	"errors"
	"os"

	"github.com/fatih/color"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrNoCommandProvided = errors.New("no command provided")
)

type Application struct {
	commands.CliApp

	context context.Context
}

// NewApplication creates a new application instance
func NewApplication(ctx context.Context) *Application {
	// Initialize base line logging, writing to the console.
	logger, err := logging.New(&logging.LoggerConfig{
		LogMode: logging.LogModeConsole,
	})
	if err != nil {
		panic(err)
	}

	// Initialize session manager
	ctx = manager.ContextWithManager(ctx)

	// Initialize logging
	ctx = logging.ContextWithLogger(ctx, logger)

	// Initialize the application.
	app := &Application{context: ctx}

	// Register commands
	app.registerCli()

	return app
}

// Start starts the application
func (a *Application) Start() {
	logger, err := logging.LoggerFromContext(a.GetContext())
	if err != nil {
		panic(err)
	}

	// Parse the command line
	cmd, err := a.parseCommand()
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	// Update the logging based on the command passed.
	// If it is a private command, write to a file only.
	if cmd.Command.LogToFile {
		logger, err = logging.New(&logging.LoggerConfig{
			LogMode: logging.LogModeFile,
		})
		if err != nil {
			panic(err)
		}

		a.context = logging.ContextWithLogger(a.GetContext(), logger)
	}

	// Close the logger when the application closes.
	defer func() {
		if logFile := logger.File(); logFile != nil {
			logFile.Close()
		}
	}()

	// Run the command
	err = cmd.Command.Execute(a.GetContext(), cmd.Arguments)
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	os.Exit(0)
}

// GetContext returns the application context
func (a *Application) GetContext() context.Context {
	return a.context
}
