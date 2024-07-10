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

func (a *MonitorSessionArguments) Flags() *monitorSessionArgumentsFlags {
	var pid *int

	cmd := flag.NewFlagSet("monitor", flag.ExitOnError)

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
