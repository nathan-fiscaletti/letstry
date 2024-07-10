package arguments

import (
	"flag"
	"fmt"
)

type monitorSessionArgumentsFlags struct {
	*flag.FlagSet

	pid *int
}

func (c *monitorSessionArgumentsFlags) Parse(args []string) error {
	return c.FlagSet.Parse(args)
}

type MonitorSessionArguments struct {
	PID int `json:"pid"`
}

func (a *MonitorSessionArguments) Name() string {
	return CommandNameMonitorSession.String()
}

func (a *MonitorSessionArguments) IsPrivate() bool {
	return true
}

func (a *MonitorSessionArguments) Aliases() []string {
	return []string{}
}

func (a *MonitorSessionArguments) FlagSet() *flag.FlagSet {
	return a.Flags().FlagSet
}

func (a *MonitorSessionArguments) Flags() *monitorSessionArgumentsFlags {
	var pid *int

	cmd := flag.NewFlagSet(CommandNameMonitorSession.String(), flag.ExitOnError)

	pid = cmd.Int("pid", -1, "PID of the process to monitor")

	return &monitorSessionArgumentsFlags{
		FlagSet: cmd,
		pid:     pid,
	}
}

func (a *MonitorSessionArguments) Scan(args []string) error {
	flags := a.Flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	if flags.pid == nil || *flags.pid == -1 {
		return fmt.Errorf("pid is required")
	}

	a.PID = *flags.pid
	return nil
}
