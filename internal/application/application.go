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
	app := &Application{
		context: ctx,
		CliApp: commands.CliApp{
			Config: commands.Config{
				DescriptionMaxWidth: 60,
			},

			Name:             commands.MainName(),
			ShortDescription: "a powerful tool for creating temporary workspaces",
			Description:      commands.MainName() + " provides a temporary workspace for you to work in, and then destroys it when you are done.",

			// =========================
			// Commands
			// =========================

			Commands: []commands.Command{
				commands.NewSessionCommand(),
				commands.ListSessionsCommand(),
				commands.ListTemplatesCommand(),
				commands.SaveTemplateCommand(),
				commands.DeleteTemplateCommand(),
				commands.ExportSessionCommand(),
				commands.ListEditorsCommand(),
				commands.SetEditorCommand(),
				commands.VersionCommand(),
				commands.MonitorCommand(),
			},
		},
	}

	// Add help command
	app.RegisterHelpCommand()

	return app
}

// Start starts the application
func (app *Application) Start() {
	// Parse the command line
	invocation, err := app.GetInvocation()
	if err != nil {
		logger, _ := logging.LoggerFromContext(app.GetContext())
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}

	// Configure logging
	err = invocation.UpdateLogger(app)
	if err != nil {
		logger, _ := logging.LoggerFromContext(app.GetContext())
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}
	defer logging.CloseLog(app.GetContext())

	// Run the command
	err = invocation.Execute(app)
	if err != nil {
		logger, _ := logging.LoggerFromContext(app.GetContext())
		logger.Printf("Error: %s\n", color.RedString(err.Error()))
		os.Exit(1)
	}
}

// GetContext returns the application context
func (a *Application) GetContext() context.Context {
	return a.context
}
