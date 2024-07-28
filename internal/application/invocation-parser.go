package application

import (
	"os"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

type ApplicationInvocation struct {
	Command   commands.Command
	Arguments []string
}

func (i ApplicationInvocation) Execute(app *Application) error {
	return i.Command.Execute(app.GetContext(), i.Arguments)
}

func (i ApplicationInvocation) UpdateLogger(app *Application) error {
	if i.Command.LogToFile {
		logger, err := logging.New(&logging.LoggerConfig{
			LogMode: logging.LogModeFile,
		})
		if err != nil {
			return err
		}

		app.context = logging.ContextWithLogger(app.GetContext(), logger)
	}

	return nil
}

func (a *Application) GetInvocation() (*ApplicationInvocation, error) {
	var err error
	var command commands.Command
	var cmdArgs []string = []string{string(commands.CommandHelp)}

	args := os.Args

	if len(args) > 1 {
		cmdArgs = args[1:]
	}

	commandName, err := commands.GetCommandName(cmdArgs[0])
	if err != nil {
		return nil, err
	}

	command, err = a.Command(commandName)
	if err != nil {
		return nil, err
	}

	return &ApplicationInvocation{
		Command:   command,
		Arguments: cmdArgs[1:],
	}, nil
}
