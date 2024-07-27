package commands

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func NewSessionHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: new -- Create a new session

Usage: 

    ` + cmdName + ` new [source]

Description:

    Create a new session using the specified source.

Arguments:

    source (optional) - The source to use for the new session. Can be
                        a git repository URL, a path to a directory, or
                        the name of a letstry template.

                        If source is not provided, the session will be
                        created with a blank source.

Run '` + cmdName + ` help' for information on additional commands.
`
}

func NewSession(ctx context.Context, args []string) error {
	var source string

	if len(args) >= 1 {
		source = args[0]
	}

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	_, err = mgr.CreateSession(ctx, manager.CreateSessionArguments{
		Source: source,
	})
	if err != nil {
		return err
	}

	return nil
}
