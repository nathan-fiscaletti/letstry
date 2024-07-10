package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
)

type Command interface {
	Execute(ctx context.Context) error
}

func GetCommand(ctx context.Context, args arguments.Arguments) (Command, error) {
	var cmd Command

	switch args.Command {
	case arguments.CommandNewSession:
		cmd = NewSessionCommand{
			Arguments: args.CreateSessionArguments,
		}
	case arguments.CommandListSessions:
		cmd = ListSessionsCommand{
			Arguments: args.ListSectionsArguments,
		}
	case arguments.CommandMonitorSession:
		cmd = MonitorSessionCommand{
			Arguments: args.MonitorSessionsArguments,
		}
	}

	if cmd == nil {
		return nil, errors.New("invalid command")
	}

	return cmd, nil
}
