package arguments

import "flag"

type helpArgumentsFlags struct {
	*flag.FlagSet
}

func (c *helpArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type HelpArguments struct{}

func (a *HelpArguments) Name() string {
	return CommandHelp.String()
}

func (a *HelpArguments) FlagSet() *flag.FlagSet {
	return a.Flags().FlagSet
}

func (a *HelpArguments) Flags() *helpArgumentsFlags {
	cmd := flag.NewFlagSet(CommandHelp.String(), flag.ExitOnError)

	return &helpArgumentsFlags{
		FlagSet: cmd,
	}
}

func (a *HelpArguments) Scan(args []string) error {
	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	// There are no flags to parse for this command

	return nil
}
