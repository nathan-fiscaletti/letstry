package commands

import (
	"context"
	"os"
)

func HelpCommand(app CliApp) Command {
	return Command{
		Name: CommandHelp,
		Executor: func(ctx context.Context, args []string) error {
			var inputCmd string

			if len(args) > 0 {
				inputCmd = args[0]
			}

			if inputCmd != "" {
				command, err := GetCommandName(inputCmd)
				if err != nil {
					return err
				}

				if !app.IsCommand(command) {
					return ErrUnknownCommand
				}

				cmd, err := app.Command(command)
				if err != nil {
					return err
				}
				return cmd.Write(app, os.Stdout)
			}

			return app.Write(os.Stdout)
		},
	}
}
