package arguments

import (
	"os"
)

type Command string

func (c Command) String() string {
	return string(c)
}

const (
	CommandNewSession   Command = "new"
	CommandListSessions Command = "list"
)

type Arguments struct {
	Command                Command
	CreateSessionArguments *CreateSessionArguments
	ListSectionsArguments  *ListSessionsArguments
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
	}

	return cmdArgs, err
}
