package arguments

import (
	"flag"
	"fmt"
	"os"
)

type CommandArguments interface {
	Name() string
	Scan(args []string) error
	FlagSet() *flag.FlagSet
}

func GetPublicCommandArguments() []CommandArguments {
	return []CommandArguments{
		&CreateSessionArguments{},
		&ListSessionsArguments{},
		&HelpArguments{},
	}
}

func GetPrivateCommandArguments() []CommandArguments {
	return []CommandArguments{
		&MonitorSessionArguments{},
	}
}

type ParsedArguments struct {
	Command                  CommandName
	CreateSessionArguments   *CreateSessionArguments
	ListSectionsArguments    *ListSessionsArguments
	MonitorSessionsArguments *MonitorSessionArguments
	HelpArguments            *HelpArguments
}

func (a ParsedArguments) IsPrivate() bool {
	for _, args := range GetPrivateCommandArguments() {
		if a.Command == CommandName(args.FlagSet().Name()) {
			return true
		}
	}

	return false
}

// ParseArguments parses the command line arguments and returns an Arguments struct
func ParseArguments() (ParsedArguments, error) {
	args := os.Args[1:]

	var cmdArgs ParsedArguments
	var err error

	switch args[0] {

	// New Session Command
	case CommandNewSession.String():
		sessionArgs := &CreateSessionArguments{}
		err = sessionArgs.Scan(args[1:])

		if err == nil {
			cmdArgs = ParsedArguments{
				Command:                CommandNewSession,
				CreateSessionArguments: sessionArgs,
			}
		}

	// List Sessions Command
	case CommandListSessions.String():
		listArgs := &ListSessionsArguments{}
		err = listArgs.Scan(args[1:])

		if err == nil {
			cmdArgs = ParsedArguments{
				Command:               CommandListSessions,
				ListSectionsArguments: listArgs,
			}
		}

	case CommandMonitorSession.String():
		// TODO: make this private, it should only be executable
		// TODO: from within this process.
		monitorArgs := &MonitorSessionArguments{}
		err = monitorArgs.Scan(args[1:])
		if err == nil {
			cmdArgs = ParsedArguments{
				Command:                  CommandMonitorSession,
				MonitorSessionsArguments: monitorArgs,
			}
		}

	case CommandHelp.String():
		helpArgs := &HelpArguments{}
		err = helpArgs.Scan(args[1:])
		if err == nil {
			cmdArgs = ParsedArguments{
				Command:       CommandHelp,
				HelpArguments: helpArgs,
			}
		}

	default:
		err = fmt.Errorf("unknown command: %s", args[0])
	}

	return cmdArgs, err
}
