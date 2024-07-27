package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingEditorName = errors.New("missing editor name")
)

func SetEditorHelp() string {
	cmdName := GetCallerName()

	return `
` + cmdName + `: set-editor -- Sets the default editor to use for new sessions

Usage: 

    ` + cmdName + ` set-editor [editor-name]

Description:

    This command sets the default editor to use for new sessions.
    Run '` + cmdName + ` editors' for a list of available editors.

    You can add new editors by editing the configuration file directly.

Arguments:

    editor-name - The name of the editor to use as the default.

Run '` + cmdName + ` help' for information on additional commands.
`
}

func SetEditor(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return ErrMissingEditorName
	}

	editorName := args[0]

	mgr, err := manager.GetManager(ctx)
	if err != nil {
		return err
	}

	return mgr.SetDefaultEditor(ctx, editorName)
}
