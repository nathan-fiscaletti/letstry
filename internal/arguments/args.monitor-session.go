package arguments

import (
	"fmt"
	"strconv"
)

type MonitorSessionsArguments struct {
	PID int `json:"pid"`
}

func (a *MonitorSessionsArguments) Scan(args []string) error {
	var pid int

	for i := 0; i < len(args); i++ {
		if args[i] == "-pid" && i+1 < len(args) {
			_pid, err := strconv.Atoi(args[i+1])
			if err != nil {
				return err
			}

			pid = _pid
			break
		}
	}

	if pid == 0 {
		return fmt.Errorf("pid is required")
	}

	a.PID = pid
	return nil
}
