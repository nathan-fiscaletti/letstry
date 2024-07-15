package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
)

type CommandExecutor interface {
	Execute(ctx context.Context) error
}

func GetCommandExecutor(ctx context.Context, args arguments.Parameters) (CommandExecutor, error) {
	var cmd CommandExecutor

	switch arguments.CommandName(args.Arguments.Name()) {
	case arguments.CommandNameNewSession:
		cmd = NewSessionCommand{
			Arguments: args,
		}
	case arguments.CommandNameListSessions:
		cmd = ListSessionsCommand{
			Arguments: args,
		}
	case arguments.CommandNameMonitorSession:
		cmd = MonitorSessionCommand{
			Arguments: args,
		}
	case arguments.CommandNameSaveSession:
		cmd = SaveSessionCommand{
			Arguments: args,
		}
	case arguments.CommandNameListTemplates:
		cmd = ListTemplatesCommand{
			Arguments: args,
		}
	case arguments.CommandNameHelp:
		cmd = HelpCommand{
			Arguments: args,
		}
	}

	if cmd == nil {
		return nil, errors.New("invalid command")
	}

	return cmd, nil
}
