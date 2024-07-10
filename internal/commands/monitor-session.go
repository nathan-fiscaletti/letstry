package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

type MonitorSessionCommand struct {
	Arguments *arguments.MonitorSessionsArguments
}

func (c MonitorSessionCommand) Execute(ctx context.Context) error {
	manager := session_manager.GetSessionManager()
	return manager.MonitorSession(ctx, *c.Arguments)
}
