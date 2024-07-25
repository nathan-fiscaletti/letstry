package application

import (
	"context"
	"errors"
	"os"

	"github.com/fatih/color"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

var (
	ErrNoCommandProvided = errors.New("no command provided")
)

type application struct {
	context         context.Context
	commands        map[commands.CommandName]commands.Command
	privateCommands []commands.CommandName
}

// NewApplication creates a new application instance
func NewApplication(ctx context.Context) *application {
	// Initialize base line logging, writing to the console.
	logger, err := logging.New(&logging.LoggerConfig{
		LogMode: logging.LogModeConsole,
	})
	if err != nil {
		panic(err)
	}

	// Initialize logging
	ctx = logging.ContextWithLogger(ctx, logger)

	// Initialize the application.
	app := &application{context: ctx}

	// Register commands
	app.registerCommands()

	return app
}

// Start starts the application
func (a *application) Start() {
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
	if cmd.IsPrivate {
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
	err = cmd.Execute(a.GetContext(), cmd.Arguments)
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	os.Exit(0)
}

// GetContext returns the application context
func (a *application) GetContext() context.Context {
	return a.context
}
