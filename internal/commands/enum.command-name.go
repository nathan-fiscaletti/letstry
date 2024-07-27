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
	CommandHelp           CommandName = "help"
	CommandMonitor        CommandName = "monitor"
	CommandNewSession     CommandName = "new"
	CommandListSessions   CommandName = "list"
	CommandListTemplates  CommandName = "templates"
	CommandListEditors    CommandName = "editors"
	CommandSetEditor      CommandName = "set-editor"
	CommandDeleteTemplate CommandName = "delete-template"
	CommandSaveTemplate   CommandName = "save"
	CommandExportSession  CommandName = "export"
)

var allCommands = []CommandName{
	CommandHelp,
	CommandMonitor,
	CommandNewSession,
	CommandListSessions,
	CommandListTemplates,
	CommandListEditors,
	CommandSetEditor,
	CommandDeleteTemplate,
	CommandSaveTemplate,
	CommandExportSession,
}

var commandAliases = map[CommandName][]string{
	CommandHelp: {"-h", "?", "--help"},
}

// GetCommandName returns the CommandName for the given string value. If the
// value is not a valid CommandName, ErrUnknownCommand will be returned.
func GetCommandName(value string) (CommandName, error) {
	for _, command := range GetCommandNames() {
		if strings.EqualFold(command.String(), value) {
			return command, nil
		}
	}

	for command, aliases := range commandAliases {
		for _, alias := range aliases {
			if strings.EqualFold(alias, value) {
				return command, nil
			}
		}
	}

	return "", ErrUnknownCommand
}

func GetCommandNames() []CommandName {
	return allCommands
}
