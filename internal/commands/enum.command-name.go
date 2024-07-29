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
	CommandVersion        CommandName = "version"
	CommandMonitor        CommandName = "monitor"
	CommandClean          CommandName = "clean"
	CommandCleanAll       CommandName = "clean-all"
	CommandNewSession     CommandName = "new"
	CommandListSessions   CommandName = "list"
	CommandListTemplates  CommandName = "templates"
	CommandListEditors    CommandName = "editors"
	CommandGetEditor      CommandName = "get-editor"
	CommandSetEditor      CommandName = "set-editor"
	CommandDeleteTemplate CommandName = "delete-template"
	CommandSaveTemplate   CommandName = "save"
	CommandExportSession  CommandName = "export"
)

var allCommands = []CommandName{
	CommandHelp,
	CommandVersion,
	CommandMonitor,
	CommandClean,
	CommandCleanAll,
	CommandNewSession,
	CommandListSessions,
	CommandListTemplates,
	CommandListEditors,
	CommandSetEditor,
	CommandGetEditor,
	CommandDeleteTemplate,
	CommandSaveTemplate,
	CommandExportSession,
}

var commandAliases = map[CommandName][]string{
	CommandHelp:    {"-h", "?", "--help"},
	CommandVersion: {"-v", "--version"},
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
