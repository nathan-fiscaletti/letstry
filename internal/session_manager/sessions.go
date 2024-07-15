package session_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/otiai10/copy"
	"github.com/shirou/gopsutil/process"
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

func (s *sessionManager) CreateSession(ctx context.Context, args arguments.CreateSessionArguments) (session, error) {
	var zeroValue session

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return zeroValue, err
	}

	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
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

	// Handle "With" arguments
	if args.WithArgument != nil {
		switch args.WithArgument.ArgumentType {

		// Handle git repository.
		case arguments.WithArgumentTypeRepoPath:
			logger.Printf("cloning repository %s\n", args.WithArgument.Value)
			_, err := git.PlainClone(tempDir, false, &git.CloneOptions{
				URL: args.WithArgument.Value,
			})
			if err != nil {
				return zeroValue, fmt.Errorf("failed to clone repository: %v", err)
			}

		// Handle directory.
		case arguments.WithArgumentTypeDirectory:
			dirPath := args.WithArgument.Value
			if _, err := os.Stat(dirPath); err != nil {
				return zeroValue, fmt.Errorf("directory %s does not exist", dirPath)
			}

			// Copy the directory to the temporary directory
			err = copy.Copy(dirPath, tempDir)
			if err != nil {
				return zeroValue, fmt.Errorf("failed to copy directory: %v", err)
			}
		}

	}

	// Handle "From" arguments
	if args.FromArgument != nil {
		// Check if the specified template exists.
		template, err := s.GetTemplate(ctx, args.FromArgument.TemplateName)
		if err != nil {
			return zeroValue, err
		}

		// Copy the template to the temporary directory
		templatePath := s.storage.GetPath(filepath.Join("templates", template.String()))
		err = copy.Copy(templatePath, tempDir)
		if err != nil {
			return zeroValue, fmt.Errorf("failed to load template %s: %s", args.FromArgument.TemplateName, err)
		}
	}

	startTime := time.Now()

	// Launch the editor
	logger.Printf("launching editor %s\n", editor.String())
	cfgArgs := strings.Split(editor.Args, " ")
	cmdArgs := append(cfgArgs, tempDir)
	cmd := exec.Command(editor.ExecPath, cmdArgs...)
	err = cmd.Run()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to run editor: %v", err)
	}

	// Give the process time to start
	logger.Printf("waiting %v for editor process to start\n", editor.ProcessCaptureDelay)
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
	err = s.addSession(ctx, newSession)
	if err != nil {
		return zeroValue, err
	}

	// Call this application again, but start it in the background as it's own process.
	// This will allow the user to continue using the current terminal session.
	logger.Printf("starting monitor process for session %s\n", newSession.FormattedName())
	cmd = exec.Command(os.Args[0], "monitor", "-pid", fmt.Sprintf("%d", editorProcess.Pid))
	err = cmd.Start()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to start monitor process: %v", err)
	}
	logger.Printf("monitor process started with PID %v\n", cmd.Process.Pid)

	return newSession, nil
}

func (s *sessionManager) ListSessions(ctx context.Context, args arguments.ListSessionsArguments) ([]session, error) {
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

func (s *sessionManager) MonitorSession(ctx context.Context, args arguments.MonitorSessionArguments) error {
	// Start monitoring the session
	return s.monitorProcessClosed(int32(args.PID), func() error {
		session, err := s.GetSessionForPID(ctx, args.PID)
		if err != nil {
			return err
		}

		logger, err := logging.LoggerFromContext(ctx)
		if err != nil {
			return err
		}

		logger.Printf("cleaning up session: %s (process closed, PID %d)\n", session.Name, session.PID)

		err = s.removeSession(ctx, session.Name)
		if err != nil {
			return err
		}

		return nil
	})
}

// GetSessionForPID returns the session with the given PID
func (s *sessionManager) GetSessionForPID(ctx context.Context, pid int) (session, error) {
	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
	if err != nil {
		return session{}, err
	}

	for _, session := range sessions {
		if session.PID == int32(pid) {
			return session, nil
		}
	}

	return session{}, fmt.Errorf("session with PID %d not found", pid)
}

// GetSession returns the session with the given name
func (s *sessionManager) GetSession(ctx context.Context, name string) (session, error) {
	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
	if err != nil {
		return session{}, err
	}

	for _, session := range sessions {
		if session.Name == name {
			return session, nil
		}
	}

	return session{}, fmt.Errorf("session with name %s not found", name)

}

func (s *sessionManager) removeSession(ctx context.Context, name string) error {
	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
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

			err = file.Truncate(0)
			if err != nil {
				return fmt.Errorf("failed to truncate sessions file: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("failed to write sessions: %v", err)
			}

			err = file.Sync()
			if err != nil {
				return fmt.Errorf("failed to sync sessions file: %v", err)
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

func (s *sessionManager) addSession(ctx context.Context, sess session) error {
	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
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

	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to truncate sessions: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write sessions: %v", err)
	}

	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync sessions file: %v", err)
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
			return callback()
		}

		time.Sleep(1 * time.Second) // Check every second
	}
}
