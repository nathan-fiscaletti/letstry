package application

import (
	"os"
	"slices"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
)

func (a *application) IsCommand(name commands.CommandName) bool {
	_, ok := a.commands[name]
	return ok
}

type registerCommandInput struct {
	Name      commands.CommandName
	Command   commands.Command
	IsPrivate bool
}

func (a *application) registerCommand(input registerCommandInput) {
	a.commands[input.Name] = input.Command
	if input.IsPrivate {
		a.privateCommands = append(a.privateCommands, input.Name)
	}
}

func (a *application) registerCommands() {
	a.commands = make(map[commands.CommandName]commands.Command)
	a.privateCommands = make([]commands.CommandName, 0)

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandNewSession,
		Command:   commands.NewSession,
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandListSessions,
		Command:   commands.ListSessions,
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandListTemplates,
		Command:   commands.ListTemplates,
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandDeleteTemplate,
		Command:   commands.DeleteTemplate,
		IsPrivate: false,
	})

	a.registerCommand(registerCommandInput{
		Name:      commands.CommandSaveTemplate,
		Command:   commands.SaveTemplate,
		IsPrivate: false,
	})

	// Private commands
	a.registerCommand(registerCommandInput{
		Name:      commands.CommandMonitor,
		Command:   commands.Monitor,
		IsPrivate: true,
	})
}

type ParsedCommand struct {
	Name      commands.CommandName
	Execute   commands.Command
	IsPrivate bool
	Arguments []string
}

func (a *application) parseCommand() (*ParsedCommand, error) {
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
