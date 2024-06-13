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

type Arguments struct {
	SessionName  string        `json:"session_name"`
	WithArgument *WithArgument `json:"with"`
	FromArgument *FromArgument `json:"from"`
}

// ParseArguments parses the command line arguments and returns an Arguments struct
func ParseArguments() (Arguments, error) {
	args := os.Args[1:]
	if len(args) < 2 {
		return Arguments{}, errors.New("session name is required")
	}

	sessionName := args[0]

	var withValue string
	var fromTemplate string
	var withValueType WithArgumentType

	for i := 1; i < len(args); i++ {
		if args[i] == "-with" && i+1 < len(args) {
			withValue = args[i+1]
			i++
		} else if args[i] == "-from" && i+1 < len(args) {
			fromTemplate = args[i+1]
			i++
		}
	}

	if withValue == "" && fromTemplate == "" {
		return Arguments{}, errors.New("either -with or -from argument is required")
	}

	if withValue != "" && fromTemplate != "" {
		return Arguments{}, errors.New("only one of -with or -from argument can be passed")
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

	return Arguments{
		SessionName:  sessionName,
		WithArgument: withPathValue,
		FromArgument: fromTemplateValue,
	}, nil
}
