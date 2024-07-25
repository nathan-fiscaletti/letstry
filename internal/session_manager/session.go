package session_manager

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
	"github.com/shirou/gopsutil/process"
)

type sessionSource struct {
	SourceType SessionSourceType `json:"sourceType"`
	Value      string            `json:"value"`
}

func (s sessionSource) String() string {
	switch s.SourceType {
	case SessionSourceTypeBlank:
		return fmt.Sprintf("[%s]", s.SourceType)
	default:
		return fmt.Sprintf("[%s, %s]", s.SourceType, s.Value)
	}
}

func (s sessionSource) FormattedString() string {
	var colorWrapper func(format string, a ...interface{}) string = color.WhiteString

	switch s.SourceType {
	case SessionSourceTypeDirectory:
		fallthrough
	case SessionSourceTypeRepository:
		colorWrapper = color.HiBlueString
	case SessionSourceTypeTemplate:
		colorWrapper = color.HiMagentaString
	case SessionSourceTypeBlank:
		colorWrapper = color.HiWhiteString
	}

	return colorWrapper("%s", s.String())
}

type session struct {
	ID       identifier.ID `json:"id"`
	PID      int32         `json:"pid"`
	Location string        `json:"location"`
	Source   sessionSource `json:"source"`
	Editor   config.Editor `json:"editor"`
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
	src := s.Source.FormattedString()
	id := s.FormattedID()
	editor := color.BlueString("(%s, PID %d)", s.Editor.Name, s.PID)

	return fmt.Sprintf("id=%s, editor=%s, src=%s", id, editor, src)
}

func (s *session) FormattedID() string {
	return color.HiGreenString(s.ID.String())
}
