package session_manager

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/nathan-fiscaletti/letstry/internal/arguments"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/otiai10/copy"
)

type template string

func (t template) String() string {
	return string(t)
}

func (s *sessionManager) SaveTemplate(ctx context.Context, arg arguments.SaveSessionArguments) (template, error) {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return "", err
	}

	sessions, err := s.ListSessions(ctx, arguments.ListSessionsArguments{})
	if err != nil {
		return "", err
	}

	var sess *session
	for _, session := range sessions {
		if session.Arguments.SessionName == arg.SessionName {
			sess = &session
			break
		}
	}

	if sess == nil {
		return "", fmt.Errorf("session with name %s does not exist", arg.SessionName)
	}

	logger.Printf("found session %s\n", sess.Name)

	err = s.createTemplatesDirectoryIfNotExists()
	if err != nil {
		return "", err
	}

	templatePath := filepath.Join("templates", arg.TemplateName)

	logger.Printf("storing template in %s\n", templatePath)

	if s.storage.DirectoryExists(templatePath) {
		return "", fmt.Errorf("template with name %s already exists", arg.TemplateName)
	}

	err = s.storage.CreateDirectory(templatePath)
	if err != nil {
		return "", err
	}

	// Copy the session contents to the template directory
	logger.Printf("copying %s to %s\n", sess.Location, s.storage.GetPath(templatePath))

	err = copy.Copy(sess.Location, s.storage.GetPath(templatePath))
	if err != nil {
		return "", err
	}

	return template(templatePath), nil
}

func (s *sessionManager) ListTemplates(ctx context.Context) ([]template, error) {
	templates, err := s.storage.ListDirectories("templates")
	if err != nil {
		return nil, err
	}

	var result []template
	for _, t := range templates {
		result = append(result, template(t))
	}

	return result, nil
}

func (s *sessionManager) DeleteTemplate(ctx context.Context, t template) error {
	return s.storage.DeleteDirectory(t.String())
}

func (s *sessionManager) createTemplatesDirectoryIfNotExists() error {
	if !s.storage.DirectoryExists("templates") {
		err := s.storage.CreateDirectory("templates")
		if err != nil {
			return err
		}
	}

	return nil
}
