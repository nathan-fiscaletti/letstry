package editors

import (
	"context"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/letstry/internal/application/commands"
	"github.com/nathan-fiscaletti/letstry/internal/cli"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

func GetEditorCommand() cli.Command {
	return cli.Command{
		Name:             commands.CommandGetEditor.String(),
		ShortDescription: "Get the default editor",
		Description:      "This command gets the default editor to use for new sessions.",
		Executor: func(ctx context.Context, args []string) error {
			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			logger, err := logging.LoggerFromContext(ctx)
			if err != nil {
				return err
			}

			editor, err := mgr.GetDefaultEditor(ctx)
			if err != nil {
				return err
			}

			logger.Printf("%s: [%s]\n", color.HiWhiteString("default editor"), editor.FullString())

			return nil
		},
	}
}
