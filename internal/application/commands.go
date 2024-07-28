package application

import (
	"os"

	"github.com/nathan-fiscaletti/letstry/internal/commands"
)

func (a *Application) registerCli() {
	a.CliApp = commands.CliApp{
		Config: commands.Config{
			DescriptionMaxWidth: 60,
		},

		Name:             commands.MainName(),
		ShortDescription: "a powerful tool for creating temporary workspaces",
		Description:      commands.MainName() + " provides a temporary workspace for you to work in, and then destroys it when you are done.",

		// =========================
		// Commands
		// =========================

		Commands: []commands.Command{
			// =========================
			// lt new [source]
			// =========================
			{
				Name:             commands.CommandNewSession,
				ShortDescription: "Create a new session",
				Description:      "Create a new session using the specified source.",
				Executor:         commands.NewSession,
				Arguments: []commands.Argument{
					{
						Name:        "source",
						Description: "The source to use for the new session. Can be a git repository URL, a path to a directory, or the name of a letstry template.\n\nIf source is not provided, the session will be created from a blank source.",
						Optional:    true,
					},
				},
			},

			// =========================
			// lt list
			// =========================
			{
				Name:             commands.CommandListSessions,
				ShortDescription: "List running sessions",
				Description:      "This command will list all currently running sessions.",
				Executor:         commands.ListSessions,
			},

			// =========================
			// lt templates
			// =========================
			{
				Name:             commands.CommandListTemplates,
				ShortDescription: "List all templates",
				Description:      "This command will list all available templates that can be used when creating a new session.",
				Executor:         commands.ListTemplates,
			},

			// =========================
			// lt delete-template <name>
			// =========================
			{
				Name:             commands.CommandDeleteTemplate,
				ShortDescription: "Delete a template",
				Description:      "Delete a template by name.",
				Executor:         commands.DeleteTemplate,
				Arguments: []commands.Argument{
					{
						Name:        "name",
						Description: "The name of the template to delete.",
						Optional:    false,
					},
				},
			},

			// =========================
			// lt save <template-name>
			// =========================
			{
				Name:                 commands.CommandSaveTemplate,
				ShortDescription:     "Saves the current session as a template",
				Description:          "This command must be run from within a session. It will save the current session as a template with the specified name. If no name is provided, and the session was created from a template, the template's name will be used.",
				Executor:             commands.SaveTemplate,
				MustBeRunFromSession: true,
				Arguments: []commands.Argument{
					{
						Name:        "template-name",
						Description: "The name to use for the template. If not provided, and the session was created from a template, the template's name will be used.",
						Optional:    true,
					},
				},
			},

			// =========================
			// lt export <path>
			// =========================
			{
				Name:                 commands.CommandExportSession,
				ShortDescription:     "Export the current session",
				Description:          "This command must be run from within a session. It will export the current session to the specified path.",
				Executor:             commands.ExportSession,
				MustBeRunFromSession: true,
				Arguments: []commands.Argument{
					{
						Name:        "path",
						Description: "The path to export the session to.",
						Optional:    false,
					},
				},
			},

			// =========================
			// lt editors
			// =========================
			{
				Name:             commands.CommandListEditors,
				ShortDescription: "Lists all available editors",
				Description:      "This command will list all available editors that can be used when creating a new session.",
				Executor:         commands.ListEditors,
			},

			// =========================
			// lt set-editor <editor-name>
			// =========================
			{
				Name:             commands.CommandSetEditor,
				ShortDescription: "Set the default editor",
				Description:      "This command sets the default editor to use for new sessions. You can run 'lt editors' for a list of available editors.\n\nAdd new editors by editing the configuration file directly.",
				Executor:         commands.SetEditor,
				Arguments: []commands.Argument{
					{
						Name:        "editor-name",
						Description: "The name of the editor to use.",
						Optional:    false,
					},
				},
			},

			// =========================
			// lt version
			// =========================
			{
				Name:             commands.CommandVersion,
				ShortDescription: "Display the version of " + commands.MainName(),
				Description:      "This command will display the version of " + commands.MainName() + ".",
				Executor:         commands.Version,
			},

			// =========================
			// lt monitor <delay> <location>
			// =========================
			{
				Name:             commands.CommandMonitor,
				ShortDescription: "Monitor a session",
				Description:      "This command is used to monitor a session. It is not intended to be run directly by the user.",
				Executor:         commands.Monitor,
				LogToFile:        true,
				Arguments: []commands.Argument{
					{
						Name:        "delay",
						Description: "The delay before initiating the monitor, formatted as a duration string. For example, '5s' for 5 seconds.",
						Optional:    false,
					},
					{
						Name:        "location",
						Description: "The location to monitor. This should be the path to a letstry session directory.",
						Optional:    false,
					},
				},
			},
		},
	}

	// Register help command
	a.RegisterHelpCommand()
}

type ParsedCommand struct {
	Command   commands.Command
	Arguments []string
}

func (a *Application) parseCommand() (*ParsedCommand, error) {
	var err error
	var command commands.Command
	var cmdArgs []string = []string{string(commands.CommandHelp)}

	args := os.Args

	if len(args) > 1 {
		cmdArgs = args[1:]
	}

	commandName, err := commands.GetCommandName(cmdArgs[0])
	if err != nil {
		return nil, err
	}

	command, err = a.Command(commandName)
	if err != nil {
		return nil, err
	}

	return &ParsedCommand{
		Command:   command,
		Arguments: cmdArgs[1:],
	}, nil
}
