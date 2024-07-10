package commands

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
)

type HelpCommand struct {
	Arguments arguments.Parameters
}

func (c HelpCommand) Execute(ctx context.Context) error {
	publicCmds := arguments.AllArguments.GetPublic()

	var sb strings.Builder

	sb.WriteString("letstry - A simple CLI tool for trying out new things.\n\n")

	// Generate usage line
	sb.WriteString("usage: letstry <command> [options]\n\n")
	sb.WriteString("--------\n\n")

	// Generate detailed help for each flag
	for _, fs := range publicCmds {
		var name string = fs.Name()
		if len(fs.Aliases()) > 0 {
			name = fmt.Sprintf("[%s, %s]", fs.Name(), strings.Join(fs.Aliases(), ", "))
		}

		sb.WriteString(fmt.Sprintf("letstry %s [options]\n\n", name))

		flagLen := 0
		fs.FlagSet().VisitAll(func(f *flag.Flag) {
			flagLen++
		})

		if flagLen > 0 {
			fs.FlagSet().VisitAll(func(f *flag.Flag) {
				fmt.Fprintf(&sb, "  -%s: %s (default: %v)\n", f.Name, f.Usage, f.DefValue)
			})
			sb.WriteString("\n")
		} else {
			sb.WriteString("  No options available.\n\n")
		}
	}

	fmt.Printf("%s", sb.String())

	return nil
}
