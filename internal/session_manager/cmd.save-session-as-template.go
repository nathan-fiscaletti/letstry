package session_manager

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/logging"
	"github.com/otiai10/copy"
)

var (
	ErrMissingTemplateName = errors.New("missing template name")
)

type SaveSessionAsTemplateArguments struct {
	TemplateName string `json:"template_name"`
}

func (s *sessionManager) SaveSessionAsTemplate(ctx context.Context, arg SaveSessionAsTemplateArguments) (Template, error) {
	logger, err := logging.LoggerFromContext(ctx)
	if err != nil {
		return "", err
	}

	session, err := s.GetCurrentSession(ctx)
	if err != nil {
		return "", err
	}

	err = s.createTemplatesDirectoryIfNotExists()
	if err != nil {
		return "", err
	}

	var template Template

	if arg.TemplateName != "" {
		template = Template(arg.TemplateName)
	} else {
		if session.Source.SourceType == SessionSourceTypeTemplate {
			template = Template(session.Source.Value)
		}
	}

	if template == "" {
		return "", ErrMissingTemplateName
	}

	if s.storage.DirectoryExists(template.StoragePath()) {
		logger.Printf("deleting existing template %s\n", template.StoragePath())
		err = s.storage.DeleteDirectory(template.StoragePath())
		if err != nil {
			return "", err
		}
	}

	logger.Printf("storing template in %s\n", template.AbsolutePath(ctx))

	err = s.storage.CreateDirectory(template.StoragePath())
	if err != nil {
		return "", err
	}

	templateAbsolutePath := template.AbsolutePath(ctx)

	// Copy the session contents to the template directory
	logger.Printf("copying %s to %s\n", session.Location, templateAbsolutePath)

	err = copy.Copy(session.Location, templateAbsolutePath)
	if err != nil {
		return "", err
	}

	return template, nil
}
