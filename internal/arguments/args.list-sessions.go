package arguments

import "flag"

type listSessionsArgumentsFlags struct {
	*flag.FlagSet
}

func (c *listSessionsArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type ListSessionsArguments struct{}

func (a *ListSessionsArguments) Flags() *listSessionsArgumentsFlags {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)

	return &listSessionsArgumentsFlags{
		FlagSet: cmd,
	}
}

func (a *ListSessionsArguments) Scan(args []string) error {
	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	// There are no flags to parse for this command

	return nil
}