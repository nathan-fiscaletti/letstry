package editors

import (
	"context"
	"errors"

	"github.com/letstrygo/letstry/internal/application/commands"
	"github.com/letstrygo/letstry/internal/cli"
	"github.com/letstrygo/letstry/internal/manager"
)

var (
	ErrMissingEditorName = errors.New("missing editor name")
)

func SetEditorCommand() cli.Command {
	return cli.Command{
		Name:             commands.CommandSetEditor.String(),
		ShortDescription: "Set the default editor",
		Description:      "This command sets the default editor to use for new sessions. You can run 'lt editors' for a list of available editors.\n\nAdd new editors by editing the configuration file directly.",
		Arguments: []cli.Argument{
			{
				Name:        "editor-name",
				Description: "The name of the editor to use.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return ErrMissingEditorName
			}

			editorName := args[0]

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			return mgr.SetDefaultEditor(ctx, editorName)
		},
	}
}
