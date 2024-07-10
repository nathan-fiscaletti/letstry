package arguments

import (
	"errors"
	"flag"
	"os"
)

type WithArgumentType string

const (
	WithArgumentTypeRepoPath  WithArgumentType = "repository"
	WithArgumentTypeDirectory WithArgumentType = "directory"
)

type WithArgument struct {
	ArgumentType WithArgumentType `json:"type"`
	Value        string           `json:"value"`
}

type FromArgument struct {
	TemplateName string `json:"template_name"`
}

type createSessionArgumentsFlags struct {
	*flag.FlagSet

	sessionName  *string
	withValue    *string
	fromTemplate *string
}

func (c *createSessionArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type CreateSessionArguments struct {
	SessionName  string        `json:"session_name"`
	WithArgument *WithArgument `json:"with"`
	FromArgument *FromArgument `json:"from"`
}

func (a *CreateSessionArguments) Name() string {
	return CommandNewSession.String()
}

func (a *CreateSessionArguments) FlagSet() *flag.FlagSet {
	return a.Flags().FlagSet
}

func (a *CreateSessionArguments) Flags() *createSessionArgumentsFlags {
	var sessionName *string
	var withValue *string
	var fromTemplate *string

	cmd := flag.NewFlagSet(a.Name(), flag.ExitOnError)

	sessionName = cmd.String("name", "", "Name of the session")
	withValue = cmd.String("with", "", "Repository address or directory path")
	fromTemplate = cmd.String("from", "", "Name of the template to use")

	return &createSessionArgumentsFlags{
		FlagSet:      cmd,
		sessionName:  sessionName,
		withValue:    withValue,
		fromTemplate: fromTemplate,
	}
}

func (a *CreateSessionArguments) Scan(args []string) error {
	var withValueType WithArgumentType

	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	withValueFilled := flags.withValue != nil && *flags.withValue != ""
	fromTemplateFilled := flags.fromTemplate != nil && *flags.fromTemplate != ""

	if flags.sessionName == nil || *flags.sessionName == "" {
		return errors.New("session name is required")
	}

	if (flags.withValue == nil || *flags.withValue == "") &&
		(flags.fromTemplate == nil || *flags.fromTemplate == "") {
		return errors.New("either -with or -from argument is required")
	}

	if withValueFilled && fromTemplateFilled {
		return errors.New("only one of -with or -from argument can be passed")
	}

	// Determine path type if withPath is provided
	if withValueFilled {
		if _, err := os.Stat(*flags.withValue); err == nil {
			withValueType = WithArgumentTypeDirectory
		} else {
			withValueType = WithArgumentTypeRepoPath
		}
	}

	// Determine withPathValue or fromTemplateValue
	var withPathValue *WithArgument
	var fromTemplateValue *FromArgument
	if withValueFilled {
		withPathValue = &WithArgument{
			ArgumentType: withValueType,
			Value:        *flags.withValue,
		}
	} else if fromTemplateFilled {
		fromTemplateValue = &FromArgument{
			TemplateName: *flags.fromTemplate,
		}
	}

	*a = CreateSessionArguments{
		SessionName:  *flags.sessionName,
		WithArgument: withPathValue,
		FromArgument: fromTemplateValue,
	}

	return nil
}
