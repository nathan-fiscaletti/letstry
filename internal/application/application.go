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
	context   context.Context
	arguments arguments.ParsedArguments
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

	// Parse the command line arguments.
	args, err := arguments.ParseArguments()
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

	return &application{
		arguments: args,
		context:   ctx,
	}
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

	// Retrieve the corresponding command structure based on
	// the command line arguments.
	cmd, err := commands.GetCommand(a.GetContext(), a.GetArguments())
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	// Execute the command.
	err = cmd.Execute(a.GetContext())
	if err != nil {
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	os.Exit(0)
}

// GetArguments returns the application arguments
func (a *application) GetArguments() arguments.ParsedArguments {
	return a.arguments
}

// GetContext returns the application context
func (a *application) GetContext() context.Context {
	return a.context
}
