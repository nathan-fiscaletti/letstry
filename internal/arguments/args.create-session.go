package arguments

import (
	"errors"
	"os"
)

type WithArgumentType string

const (
	WithArgumentTypeRepoPath      WithArgumentType = "repo"
	WithArgumentTypeFileDirectory WithArgumentType = "file"
)

type WithArgument struct {
	ArgumentType WithArgumentType `json:"type"`
	Value        string           `json:"value"`
}

type FromArgument struct {
	TemplateName string `json:"template_name"`
}

type CreateSessionArguments struct {
	SessionName  string        `json:"session_name"`
	WithArgument *WithArgument `json:"with"`
	FromArgument *FromArgument `json:"from"`
}

func (a *CreateSessionArguments) Scan(args []string) error {
	var sessionName string
	var withValue string
	var fromTemplate string
	var withValueType WithArgumentType

	for i := 0; i < len(args); i++ {
		if args[i] == "-name" && i+1 < len(args) {
			sessionName = args[i+1]
			i++
		} else if args[i] == "-with" && i+1 < len(args) {
			withValue = args[i+1]
			i++
		} else if args[i] == "-from" && i+1 < len(args) {
			fromTemplate = args[i+1]
			i++
		}
	}

	if sessionName == "" {
		return errors.New("session name is required")
	}

	if withValue == "" && fromTemplate == "" {
		return errors.New("either -with or -from argument is required")
	}

	if withValue != "" && fromTemplate != "" {
		return errors.New("only one of -with or -from argument can be passed")
	}

	// Determine path type if withPath is provided
	if withValue != "" {
		if _, err := os.Stat(withValue); err == nil {
			withValueType = WithArgumentTypeFileDirectory
		} else {
			withValueType = WithArgumentTypeRepoPath
		}
	}

	// Determine withPathValue or fromTemplateValue
	var withPathValue *WithArgument
	var fromTemplateValue *FromArgument
	if withValue != "" {
		withPathValue = &WithArgument{
			ArgumentType: withValueType,
			Value:        withValue,
		}
	} else if fromTemplate != "" {
		fromTemplateValue = &FromArgument{
			TemplateName: fromTemplate,
		}
	}

	*a = CreateSessionArguments{
		SessionName:  sessionName,
		WithArgument: withPathValue,
		FromArgument: fromTemplateValue,
	}

	return nil
}
