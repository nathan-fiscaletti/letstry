package arguments

import (
	"flag"
	"fmt"
)

type saveSessionArgumentsFlags struct {
	*flag.FlagSet

	sessionName  *string
	templateName *string
}

func (c *saveSessionArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type SaveSessionArguments struct {
	SessionName  string
	TemplateName string
}

func (a *SaveSessionArguments) Name() string {
	return CommandNameSaveSession.String()
}

func (a *SaveSessionArguments) Aliases() []string {
	return []string{}
}

func (a *SaveSessionArguments) FlagSet() *flag.FlagSet {
	return a.Flags().FlagSet
}

func (a *SaveSessionArguments) Flags() *saveSessionArgumentsFlags {
	var sessionName, templateName *string

	cmd := flag.NewFlagSet(CommandNameSaveSession.String(), flag.ExitOnError)

	sessionName = cmd.String("session", "", "Name of the session to save")
	templateName = cmd.String("as", "", "Name of the template to save the session as")

	return &saveSessionArgumentsFlags{
		FlagSet: cmd,

		sessionName:  sessionName,
		templateName: templateName,
	}
}

func (a *SaveSessionArguments) Scan(args []string) error {
	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	if flags.sessionName == nil || *flags.sessionName == "" {
		return fmt.Errorf("session is required")
	}

	if flags.templateName == nil || *flags.templateName == "" {
		return fmt.Errorf("template name (-as) is required")
	}

	a.SessionName = *flags.sessionName
	a.TemplateName = *flags.templateName

	return nil
}
