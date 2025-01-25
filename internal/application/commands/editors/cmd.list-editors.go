package editors

import (
	"context"

	"github.com/fatih/color"
	"github.com/letstrygo/letstry/internal/application/commands"
	"github.com/letstrygo/letstry/internal/cli"
	"github.com/letstrygo/letstry/internal/logging"
	"github.com/letstrygo/letstry/internal/manager"
)

func ListEditorsCommand() cli.Command {
	return cli.Command{
		Name:             commands.CommandListEditors.String(),
		ShortDescription: "Lists all available editors",
		Description:      "This command will list all available editors that can be used when creating a new session.",
		Executor: func(ctx context.Context, args []string) error {
			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			logger, err := logging.LoggerFromContext(ctx)
			if err != nil {
				return err
			}

			editors, err := mgr.ListEditors(ctx)
			if err != nil {
				return err
			}

			for _, editor := range editors {
				logger.Printf("%s: [%s]\n", color.HiWhiteString("editor"), editor.FullString())
			}

			return nil
		},
	}
}
