package session_manager

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/storage"
)

type sessionManager struct {
	storage *storage.Storage
}

var mgr *sessionManager

func init() {
	mgr = &sessionManager{
		storage: storage.GetStorage(),
	}
}

// GetSessionManager returns the session manager
func GetSessionManager() *sessionManager {
	return mgr
}

func (s *sessionManager) GetRunningSessions() ([]session, error) {
	var sessions []session = make([]session, 0)

	var defaultSessions []byte
	defaultSessions, err := json.Marshal(sessions)
	if err != nil {
		return sessions, fmt.Errorf("failed to marshal default sessions: %v", err)
	}

	file, err := s.storage.OpenFileWithDefaultContent("sessions.json", defaultSessions)
	if err != nil {
		return sessions, fmt.Errorf("failed to open sessions file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&sessions)
	if err != nil {
		return sessions, fmt.Errorf("failed to decode sessions file: %v", err)
	}

	return sessions, nil
}

func (s *sessionManager) CreateSession(name string, args arguments.Arguments) (session, error) {
	var zeroValue session

	sessions, err := s.GetRunningSessions()
	if err != nil {
		return zeroValue, err
	}

	for _, session := range sessions {
		if session.Arguments.SessionName == name {
			return zeroValue, fmt.Errorf("session with name %s already exists", name)
		}
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return zeroValue, err
	}

	editor, err := cfg.GetDefaultEditor()
	if err != nil {
		return zeroValue, err
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "letstry")
	if err != nil {
		return zeroValue, fmt.Errorf("failed to create temporary directory: %v", err)
	}

	// Launch the editor
	editorArgs := fmt.Sprintf("%s %s", editor.Args, tempDir)
	cmd := exec.Command(editor.ExecPath, editorArgs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to run editor: %v", err)
	}

	// Create the session
	newSession := session{
		Name:      name,
		PID:       int32(cmd.Process.Pid),
		Location:  tempDir,
		Arguments: args,
	}

	// Persist the session to the sessions file
	sessions = append(sessions, newSession)

	file, err := s.storage.OpenFile("sessions.json")
	if err != nil {
		return zeroValue, fmt.Errorf("failed to open sessions file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(sessions)
	if err != nil {
		return zeroValue, fmt.Errorf("failed to encode sessions file: %v", err)
	}

	return newSession, nil
}
