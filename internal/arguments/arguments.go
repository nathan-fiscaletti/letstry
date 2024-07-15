package arguments

import (
	"flag"
	"fmt"
	"os"
)

// AllArguments is a list of all available command line arguments
var AllArguments = ArgumentsList{
	&CreateSessionArguments{},
	&ListSessionsArguments{},
	&MonitorSessionArguments{},
	&SaveSessionArguments{},
	&ListTemplatesArguments{},
	&HelpArguments{},
}

// ArgumentsList is a list of Arguments
type ArgumentsList []Arguments

// GetArguments returns the Arguments struct for the given command name
func (l ArgumentsList) GetArguments(commandName string) Arguments {
	for _, args := range l {
		if args.Name() == commandName {
			return args
		}

		for _, alias := range args.Aliases() {
			if alias == commandName {
				return args
			}
		}
	}

	return nil
}

// GetPrivate returns a list of private Arguments
func (l ArgumentsList) GetPrivate() ArgumentsList {
	var private ArgumentsList

	for _, args := range l {
		if privateArgs, ok := args.(PrivateArguments); ok && privateArgs.IsPrivate() {
			private = append(private, args)
		}
	}

	return private
}

// GetPublic returns a list of public Arguments
func (l ArgumentsList) GetPublic() ArgumentsList {
	var public ArgumentsList

	for _, args := range l {
		if privateArgs, ok := args.(PrivateArguments); !ok || !privateArgs.IsPrivate() {
			public = append(public, args)
		}
	}

	return public
}

// Arguments is an interface that defines the methods that must be
// implemented by a struct that holds the command line arguments
type Arguments interface {
	Name() string
	Aliases() []string
	Scan(args []string) error
	FlagSet() *flag.FlagSet
}

// PrivateArguments is an interface that defines the IsPrivate method
type PrivateArguments interface {
	IsPrivate() bool
}

// Parameters is a struct that holds the parsed command line arguments
type Parameters struct {
	Arguments
}

// IsPrivate returns true if the command line arguments are private
func (a Parameters) IsPrivate() bool {
	if privateArgs, ok := a.Arguments.(PrivateArguments); ok {
		return privateArgs.IsPrivate()
	}

	return false
}

// ParseArguments parses the command line arguments and returns an Parameters struct
func ParseArguments(registered ArgumentsList) (Parameters, error) {
	args := os.Args[1:]

	var cmdArgs Parameters
	var err error

	arguments := registered.GetArguments(args[0])
	if arguments == nil {
		err = fmt.Errorf("unknown command: %s", args[0])
	} else {
		err = arguments.Scan(args[1:])
		if err == nil {
			cmdArgs = Parameters{
				Arguments: arguments,
			}
		}
	}

	return cmdArgs, err
}
