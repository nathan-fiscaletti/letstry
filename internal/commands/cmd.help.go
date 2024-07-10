package commands

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
)

type HelpCommand struct {
	Arguments *arguments.HelpArguments
}

func (c HelpCommand) Execute(ctx context.Context) error {
	publicCmds := arguments.GetPublicCommandArguments()

	var sb strings.Builder

	sb.WriteString("letstry - A simple CLI tool for trying out new things.\n\n")

	// Generate usage line
	sb.WriteString("usage: letstry <command> [options]\n\n")
	sb.WriteString("--------\n\n")

	// Generate detailed help for each flag
	for _, fs := range publicCmds {
		sb.WriteString(fmt.Sprintf("letstry %s [options]\n\n", fs.Name()))

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
