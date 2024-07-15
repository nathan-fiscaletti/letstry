package arguments

type CommandName string

func (c CommandName) String() string {
	return string(c)
}

const (
	CommandNameNewSession     CommandName = "new"
	CommandNameListSessions   CommandName = "list"
	CommandNameMonitorSession CommandName = "monitor"
	CommandNameHelp           CommandName = "help"
	CommandNameSaveSession    CommandName = "save"
	CommandNameListTemplates  CommandName = "list-templates"
)
