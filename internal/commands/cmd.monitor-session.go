package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type MonitorSessionCommand struct {
	Arguments arguments.Parameters
}

func (c MonitorSessionCommand) Execute(ctx context.Context) error {
	manager := session_manager.GetSessionManager()
	arguments := *c.Arguments.Arguments.(*arguments.MonitorSessionArguments)
	return manager.MonitorSession(ctx, arguments)
}
