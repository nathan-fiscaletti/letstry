package application

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
)

func (a *Application) IsCommand(name commands.CommandName) bool {
	_, ok := a.commands[name]
	return ok
}

type registerCommandInput struct {
	Name      commands.CommandName
	Command   commands.Command
	Help      string
	IsPrivate bool
}

func (a *Application) registerCommand(input registerCommandInput) {
	a.commands[input.Name] = input.Command
	if input.IsPrivate {
		a.privateCommands = append(a.privateCommands, input.Name)
	}
	if input.Help != "" {
		a.helpMessages[input.Name] = input.Help
	}
}

func (a *Application) registerCommands() {
	a.commands = make(map[commands.CommandName]commands.Command)
	a.privateCommands = make([]commands.CommandName, 0)
	a.helpMessages = make(map[commands.CommandName]string)

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandNewSession,
		Command:   commands.NewSession,
		Help:      commands.NewSessionHelp(),
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandListSessions,
		Command:   commands.ListSessions,
		Help:      commands.ListSessionsHelp(),
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandListTemplates,
		Command:   commands.ListTemplates,
		Help:      commands.ListTemplatesHelp(),
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandDeleteTemplate,
		Command:   commands.DeleteTemplate,
		Help:      commands.DeleteTemplateHelp(),
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandSaveTemplate,
		Command:   commands.SaveTemplate,
		Help:      commands.SaveTemplateHelp(),
		IsPrivate: false,
	})

	// Private commands
	a.registerCommand(registerCommandInput{
		Name:      commands.CommandMonitor,
		Command:   commands.Monitor,
		IsPrivate: true,
	})

	// Help Command
	a.registerCommand(registerCommandInput{
		Name: commands.CommandHelp,
		Command: func(ctx context.Context, args []string) error {
			var inputCmd string

			if len(args) > 0 {
				inputCmd = args[0]
			}

			if inputCmd != "" {
				command, err := commands.GetCommandName(inputCmd)
				if err != nil {
					return err
				}

				if !a.IsCommand(command) {
					return commands.ErrUnknownCommand
				}

				if helpMessage, ok := a.helpMessages[command]; ok {
					fmt.Printf("%s\n", helpMessage)
					return nil
				}
			}

			cmdName := commands.GetCallerName()
			helpMessage := `
` + cmdName + `: a powerful tool for creating temporary workspaces

Usage:

	` + cmdName + ` <command> [arguments]

Commands:

	new              Create a new session
	list             List all sessions
	templates        List all templates
	delete-template  Delete a template
	save             Save a session as a template
	help             Show this help message

Run '` + cmdName + ` help <command>' for more information on a command.
`

			fmt.Printf("%s\n", helpMessage)
			return nil
		},
		IsPrivate: false,
	})

}

type ParsedCommand struct {
	Name      commands.CommandName
	Execute   commands.Command
	IsPrivate bool
	Arguments []string
}

func (a *Application) parseCommand() (*ParsedCommand, error) {
	// Get argv
	args := os.Args[1:]

	if len(args) < 1 {
		return nil, ErrNoCommandProvided
	}

	commandName, err := commands.GetCommandName(args[0])
	if err != nil {
		return nil, err
	}

	if !a.IsCommand(commandName) {
		return nil, commands.ErrUnknownCommand
	}

	command := a.commands[commandName]

	return &ParsedCommand{
		Name:      commandName,
		Execute:   command,
		Arguments: args[1:],
		IsPrivate: slices.Contains(a.privateCommands, commandName),
	}, nil
}
