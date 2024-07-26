package session_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/nathan-fiscaletti/letstry/internal/config"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/nathan-fiscaletti/letstry/internal/util/identifier"
	"github.com/otiai10/copy"
	"github.com/shirou/gopsutil/process"
)

type CreateSessionArguments struct {
	Source string `json:"source"`
}

func (s *sessionManager) CreateSession(ctx context.Context, args CreateSessionArguments) (session, error) {
	var zeroValue session

	sourceType, err := s.GetSessionSourceType(ctx, args.Source)
	if err != nil {
		return zeroValue, err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return zeroValue, err
	}

	editor, err := cfg.GetDefaultEditor()
	if err != nil {
		return zeroValue, err
	}

	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return zeroValue, err
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "letstry")
	if err != nil {
		return zeroValue, fmt.Errorf("failed to create temporary directory: %v", err)
	}

	logger.Printf("found source type: %s\n", sourceType.FormattedString())

	// Handle "With" arguments
	switch sourceType {
	case SessionSourceTypeBlank:
		// Do nothing
	case SessionSourceTypeDirectory:
		if _, err := os.Stat(args.Source); err != nil {
			return zeroValue, fmt.Errorf("directory %s does not exist", args.Source)
		}

		// Copy the directory to the temporary directory
		err = copy.Copy(args.Source, tempDir)
		if err != nil {
			return zeroValue, fmt.Errorf("failed to copy directory: %v", err)
		}
	case SessionSourceTypeRepository:
		logger.Printf("cloning repository %s\n", args.Source)
		_, err := git.PlainClone(tempDir, false, &git.CloneOptions{
			URL: args.Source,
		})
		if err != nil {
			return zeroValue, fmt.Errorf("failed to clone repository: %v", err)
		}
	case SessionSourceTypeTemplate:
		// Check if the specified template exists.
		template, err := s.GetTemplate(ctx, args.Source)
		if err != nil {
			return zeroValue, err
		}

		// Copy the template to the temporary directory
		err = copy.Copy(template.AbsolutePath(ctx), tempDir)
		if err != nil {
			return zeroValue, fmt.Errorf("failed to load template %s: %s", args.Source, err)
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

	logger.Printf("searching for editor process\n")
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

	logger.Printf("found editor process with PID %v\n", editorProcess.Pid)

	// Create the session
	newSession := session{
		ID:       identifier.NewID(),
		PID:      int32(editorProcess.Pid),
		Location: tempDir,
		Source:   sessionSource{SourceType: sourceType, Value: args.Source},
		Editor:   editor,
	}

	logger.Printf("persisting session %s\n", newSession.FormattedID())

	// Save the session
	err = s.addSession(ctx, newSession)
	if err != nil {
		return zeroValue, err
	}

	// Call this application again, but start it in the background as it's own process.
	// This will allow the user to continue using the current terminal session.
	logger.Printf("starting monitor process for session %s\n", newSession.FormattedID())
	cmd = exec.Command(os.Args[0], "monitor", fmt.Sprintf("%d", editorProcess.Pid))
	err = cmd.Start()
	if err != nil {
		return zeroValue, fmt.Errorf("failed to start monitor process: %v", err)
	}
	logger.Printf("monitor process started with PID %v\n", cmd.Process.Pid)

	logger.Printf("session created: %s\n", newSession.String())

	return newSession, nil
}

func (s *sessionManager) addSession(ctx context.Context, sess session) error {
	sessions, err := s.ListSessions(ctx)
	if err != nil {
		return err
	}

	// check if the session already exists by the same name
	for _, session := range sessions {
		if session.ID == sess.ID {
			return fmt.Errorf("session with ID %s already exists", sess.ID)
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
