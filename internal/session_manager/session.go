package session_manager

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/shirou/gopsutil/v3/process"
)

type session struct {
	Name      string                           `json:"name"`
	PID       int32                            `json:"pid"`
	Location  string                           `json:"location"`
	Arguments arguments.CreateSessionArguments `json:"arguments"`
	Editor    config.Editor                    `json:"editor"`
}

// GetProcess returns the process for the session
func (s *session) GetProcess() (*process.Process, error) {
	return process.NewProcess(s.PID)
}

// IsRunning returns true if the session is running
func (s *session) IsRunning() bool {
	_, err := process.NewProcess(s.PID)
	return err == nil
}

func (s *session) Kill() {
	proc, err := s.GetProcess()
	if err != nil {
		return
	}

	proc.Kill()
}

func (s *session) String() string {
	var src string
	if s.Arguments.WithArgument != nil {
		src = color.HiBlueString("[%s, %s]", s.Arguments.WithArgument.ArgumentType, s.Arguments.WithArgument.Value)
	} else if s.Arguments.FromArgument != nil {
		src = color.HiMagentaString("[template, %s]", s.Arguments.FromArgument.TemplateName)
	} else {
		src = color.HiRedString("unknown")
	}

	name := color.HiGreenString(s.Name)
	editor := color.BlueString("(%s, PID %d)", s.Editor.Name, s.PID)

	return fmt.Sprintf("name=%s, editor=%s, src=%s", name, editor, src)
}

func (s *session) FormattedName() string {
	return color.HiGreenString(s.Name)
}
