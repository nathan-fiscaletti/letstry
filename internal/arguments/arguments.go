package arguments

import (
	"fmt"
	"os"
)

type Command string

func (c Command) String() string {
	return string(c)
}

const (
	CommandNewSession     Command = "new"
	CommandListSessions   Command = "list"
	CommandMonitorSession Command = "monitor"
)

type Arguments struct {
	Command                  Command
	CreateSessionArguments   *CreateSessionArguments
	ListSectionsArguments    *ListSessionsArguments
	MonitorSessionsArguments *MonitorSessionsArguments
}

// ParseArguments parses the command line arguments and returns an Arguments struct
func ParseArguments() (Arguments, error) {
	args := os.Args[1:]

	var cmdArgs Arguments
	var err error

	switch args[0] {

	// New Session Command
	case CommandNewSession.String():
		sessionArgs := &CreateSessionArguments{}
		err = sessionArgs.Scan(args[1:])

		if err == nil {
			cmdArgs = Arguments{
				Command:                CommandNewSession,
				CreateSessionArguments: sessionArgs,
			}
		}

	// List Sessions Command
	case CommandListSessions.String():
		listArgs := &ListSessionsArguments{}
		err = listArgs.Scan(args[1:])

		if err == nil {
			cmdArgs = Arguments{
				Command:               CommandListSessions,
				ListSectionsArguments: listArgs,
			}
		}

	case CommandMonitorSession.String():
		// TODO: make this private, it should only be executable
		// TODO: from within this process.
		monitorArgs := &MonitorSessionsArguments{}
		err = monitorArgs.Scan(args[1:])
		if err == nil {
			cmdArgs = Arguments{
				Command:                  CommandMonitorSession,
				MonitorSessionsArguments: monitorArgs,
			}
		}

	default:
		err = fmt.Errorf("unknown command: %s", args[0])
	}

	return cmdArgs, err
}
