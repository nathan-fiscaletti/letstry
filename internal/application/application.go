package application

import (
	"context"
	"os"

	"github.com/fatih/color"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/commands"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

type application struct {
	context    context.Context
	parameters arguments.Parameters
	commands   arguments.ArgumentsList
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

	app := &application{
		commands: arguments.AllArguments,
	}

	// Parse the command line arguments.
	args, err := arguments.ParseArguments(app.commands)
	if err != nil {
		logger.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	// Update the logging based on the command passed.
	// If it is a private command, write to a file only.
	if args.IsPrivate() {
		if logFile := logger.File(); logFile != nil {
			logFile.Close()
		}

		logger, err = logging.New(&logging.LoggerConfig{
			LogMode: logging.LogModeFile,
		})
		if err != nil {
			panic(err)
		}
	}

	// Update the context with the logger.
	ctx = logging.ContextWithLogger(ctx, logger)

	app.context = ctx
	app.parameters = args

	return app
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

	// Retrieve the corresponding command executor based on
	// the command line arguments.
	executor, err := commands.GetCommandExecutor(a.GetContext(), a.GetParsedArguments())
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	// Execute the command.
	err = executor.Execute(a.GetContext())
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	os.Exit(0)
}

// GetParsedArguments returns the parsed application arguments
func (a *application) GetParsedArguments() arguments.Parameters {
	return a.parameters
}

// GetContext returns the application context
func (a *application) GetContext() context.Context {
	return a.context
}
