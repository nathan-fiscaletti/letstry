package session_manager

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/storage"
	"github.com/shirou/gopsutil/v3/process"
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

func (s *sessionManager) CreateSession(args arguments.CreateSessionArguments) (session, error) {
	var zeroValue session

	sessions, err := s.ListSessions(arguments.ListSessionsArguments{})
	if err != nil {
		return zeroValue, err
	}

	for _, session := range sessions {
		if session.Arguments.SessionName == args.SessionName {
			return zeroValue, fmt.Errorf("session with name %s already exists", args.SessionName)
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

	startTime := time.Now()

	// Launch the editor
	cfgArgs := strings.Split(editor.Args, " ")
	cmdArgs := append(cfgArgs, tempDir)
	cmd := exec.Command(editor.ExecPath, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to run editor: %v", err)
	}

	// Give the process time to start
	time.Sleep(editor.ProcessCaptureDelay)

	processes, err := process.Processes()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to get processes: %v", err)
	}

	// Find the editor process based on the start time
	var editorProcess *process.Process
	for _, p := range processes {
		name, err := p.Name()
		if err == nil {
			if strings.Contains(strings.ToLower(name), strings.ToLower(editor.GetExecName())) {
				createTime, err := p.CreateTime()
				if err != nil {
					continue
				}

				processStartTime := time.Unix(0, createTime*int64(time.Millisecond))
				if processStartTime.After(startTime) {
					editorProcess = p
					break
				}
			}
		}
	}

	if editorProcess == nil {
		return zeroValue, fmt.Errorf("failed to find editor process")
	}

	// Create the session
	newSession := session{
		Name:      args.SessionName,
		PID:       int32(editorProcess.Pid),
		Location:  tempDir,
		Arguments: args,
		Editor:    editor,
	}

	// Save the session
	err = s.addSession(newSession)
	if err != nil {
		return zeroValue, err
	}

	// Call this application again, but start it in the background as it's own process.
	// This will allow the user to continue using the current terminal session.
	cmd = exec.Command(os.Args[0], "monitor", "-pid", fmt.Sprintf("%d", editorProcess.Pid))
	err = cmd.Start()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to start monitor process: %v", err)
	}

	return newSession, nil
}

func (s *sessionManager) ListSessions(args arguments.ListSessionsArguments) ([]session, error) {
	var sessions []session = make([]session, 0)

	var defaultSessions []byte
	defaultSessions, err := json.MarshalIndent(sessions, "", "    ")
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

func (s *sessionManager) MonitorSession(args arguments.MonitorSessionsArguments) error {
	// Start monitoring the session
	return s.monitorProcessClosed(int32(args.PID), func() error {
		sessions, err := s.ListSessions(arguments.ListSessionsArguments{})
		if err != nil {
			return err
		}

		for _, session := range sessions {
			if session.PID == int32(args.PID) {
				err := s.removeSession(session.Name)
				if err != nil {
					return err
				}
				break
			}
		}

		return nil
	})
}

func (s *sessionManager) removeSession(name string) error {
	sessions, err := s.ListSessions(arguments.ListSessionsArguments{})
	if err != nil {
		return err
	}

	for i, session := range sessions {
		if session.Name == name {
			// Remove the session
			sessions = append(sessions[:i], sessions[i+1:]...)

			file, err := s.storage.OpenFile("sessions.json")
			if err != nil {
				return fmt.Errorf("failed to open sessions file: %v", err)
			}
			defer file.Close()

			data, err := json.MarshalIndent(sessions, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to marshal sessions: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("failed to write sessions: %v", err)
			}

			// Give the process manager time to settle
			time.Sleep(1 * time.Second)

			// Remove the temporary directory
			err = os.RemoveAll(session.Location)
			if err != nil {
				return fmt.Errorf("failed to remove temporary directory: %v", err)
			}

			return nil
		}
	}

	return fmt.Errorf("session with name %s not found", name)
}

func (s *sessionManager) addSession(sess session) error {
	sessions, err := s.ListSessions(arguments.ListSessionsArguments{})
	if err != nil {
		return err
	}

	// check if the session already exists by the same name
	for _, session := range sessions {
		if session.Name == sess.Name {
			return fmt.Errorf("session with name %s already exists", sess.Name)
		}
	}

	// add the session to the list of sessions
	sessions = append(sessions, sess)

	// save the sessions
	file, err := s.storage.OpenFile("sessions.json")
	if err != nil {
		return fmt.Errorf("failed to open sessions file: %v", err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(sessions, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal sessions: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write sessions: %v", err)
	}

	return nil
}

func (s *sessionManager) monitorProcessClosed(pid int32, callback func() error) error {
	p, err := process.NewProcess(pid)
	if err != nil {
		return err
	}

	for {
		exists, err := p.IsRunning()
		if err != nil {
			return err
		}

		if !exists {
			fmt.Printf("Process %v: %s\n", pid, "does not exist")
			return callback()
		} else {
			fmt.Printf("Process %v: %s\n", pid, "exist")
		}

		time.Sleep(1 * time.Second) // Check every second
	}
}
