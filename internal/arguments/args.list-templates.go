package arguments

import "flag"

type listTemplatesArgumentsFlags struct {
	*flag.FlagSet
}

func (c *listTemplatesArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type ListTemplatesArguments struct{}

func (a *ListTemplatesArguments) Name() string {
	return CommandNameListTemplates.String()
}

func (a *ListTemplatesArguments) Aliases() []string {
	return []string{}
}

func (a *ListTemplatesArguments) FlagSet() *flag.FlagSet {
	return a.Flags().FlagSet
}

func (a *ListTemplatesArguments) Flags() *listTemplatesArgumentsFlags {
	cmd := flag.NewFlagSet(CommandNameListTemplates.String(), flag.ExitOnError)

	return &listTemplatesArgumentsFlags{
		FlagSet: cmd,
	}
}

func (a *ListTemplatesArguments) Scan(args []string) error {
	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	// There are no flags to parse for this command

	return nil
}
