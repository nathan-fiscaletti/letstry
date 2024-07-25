package commands

import (
	"errors"
	"strings"
)

var (
	ErrUnknownCommand = errors.New("unknown command")
)

type CommandName string

func (c CommandName) String() string {
	return string(c)
}

const (
	CommandNewSession     CommandName = "new"
	CommandListSessions   CommandName = "list"
	CommandMonitor        CommandName = "monitor"
	CommandListTemplates  CommandName = "templates"
	CommandDeleteTemplate CommandName = "delete-template"
	CommandSaveTemplate   CommandName = "save"
)

var allCommands = []CommandName{
	CommandNewSession,
	CommandListSessions,
	CommandMonitor,
	CommandListTemplates,
	CommandDeleteTemplate,
	CommandSaveTemplate,
}

// GetCommandName returns the CommandName for the given string value. If the
// value is not a valid CommandName, ErrUnknownCommand will be returned.
func GetCommandName(value string) (CommandName, error) {
	for _, command := range GetCommandNames() {
		if strings.EqualFold(command.String(), value) {
			return command, nil
		}
	}

	return "", ErrUnknownCommand
}

func GetCommandNames() []CommandName {
	return allCommands
}
