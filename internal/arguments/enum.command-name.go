package arguments

type CommandName string

func (c CommandName) String() string {
	return string(c)
}

const (
	CommandNewSession     CommandName = "new"
	CommandListSessions   CommandName = "list"
	CommandMonitorSession CommandName = "monitor"
	CommandHelp           CommandName = "help"
)
