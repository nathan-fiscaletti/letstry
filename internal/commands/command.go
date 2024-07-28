package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

type Config struct {
	DescriptionMaxWidth int
}

type CliApp struct {
	Config           Config
	Name             string
	ShortDescription string
	Description      string
	Commands         []Command
}

func (app CliApp) Write(out io.Writer) error {
	return template.Must(
		template.
			New("help").
			Funcs(defaultTemplateFuncs(app)).
			Parse(applicationHelpTemplate),
	).Execute(out, app)
}

func (app CliApp) IsCommand(name CommandName) bool {
	for _, command := range app.Commands {
		if command.Name == name {
			return true
		}
	}
	return false
}

func (app CliApp) Command(name CommandName) (Command, error) {
	for _, command := range app.Commands {
		if command.Name == name {
			return command, nil
		}
	}
	return Command{}, ErrUnknownCommand
}

func (app *CliApp) RegisterHelpCommand() {
	app.Commands = append(app.Commands, GetCommandHelp(*app))
}

type Argument struct {
	Name        string
	Description string
	Optional    bool
}

func (a Argument) Label() string {
	label := color.HiWhiteString(a.Name)
	if a.Optional {
		label = fmt.Sprintf("%s %s", label, color.HiBlueString("(optional)"))
	}
	return label
}

type CommandExecutor func(context.Context, []string) error

type Command struct {
	Name                 CommandName
	ShortDescription     string
	Description          string
	Arguments            []Argument
	Executor             CommandExecutor
	LogToFile            bool
	MustBeRunFromSession bool
}

func (command Command) Execute(ctx context.Context, args []string) error {
	if command.MustBeRunFromSession {
		mgr, err := manager.GetManager(ctx)
		if err != nil {
			return err
		}

		_, err = mgr.GetCurrentSession(ctx)
		if err != nil {
			return err
		}
	}

	return command.Executor(ctx, args)
}

func (command Command) Write(app CliApp, out io.Writer) error {
	return template.Must(
		template.
			New(fmt.Sprintf("help.%s", command.Name)).
			Funcs(defaultTemplateFuncs(app)).
			Parse(commandHelpTemplate),
	).Execute(out, command)
}

func GetCommandHelp(app CliApp) Command {
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

func MainName() string {
	cmd := filepath.Base(os.Args[0])

	if runtime.GOOS == "windows" {
		cmd = strings.TrimSuffix(cmd, ".exe")
	}

	return cmd
}

func defaultTemplateFuncs(app CliApp) template.FuncMap {
	return template.FuncMap{
		"getCallerName": func() string {
			return app.Name
		},
		"wrap": wrap,
		"getMaxWidth": func() int {
			return app.Config.DescriptionMaxWidth
		},
		"longestStringLength": longestStringLength,
		"padEnd":              padEnd,
		"white":               white,
		"whiteCommand":        whiteCommand,
	}
}

func wrap(text string, lineLength int, lineStartPadding int) string {
	if lineLength < 1 {
		return text
	}

	padding := strings.Repeat(" ", lineStartPadding)
	var result strings.Builder
	words := strings.Fields(text)

	currentLineLength := lineStartPadding

	for i, word := range words {
		if currentLineLength+len(word)+1 > lineLength {
			result.WriteString("\n" + padding)
			currentLineLength = lineStartPadding
		} else if i > 0 {
			result.WriteString(" ")
			currentLineLength++
		}
		result.WriteString(word)
		currentLineLength += len(word)
	}

	return result.String()
}

func longestStringLength(cmds []Command) int {
	var max int
	for _, s := range cmds {
		if len(s.Name) > max {
			max = len(s.Name)
		}
	}
	return max
}

func padEnd(val CommandName, length int) string {
	if length <= len(val) {
		return val.String()
	}
	return val.String() + strings.Repeat(" ", length-len(val.String()))
}

func white(s string) string {
	return color.HiWhiteString(s)
}

func whiteCommand(cmd CommandName) string {
	return color.HiWhiteString(cmd.String())
}
