package session_manager

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type SessionSourceType string

func (s SessionSourceType) String() string {
	return string(s)
}

func (s SessionSourceType) FormattedString() string {
	return color.HiYellowString(s.String())
}

const (
	SessionSourceTypeRepository SessionSourceType = "repository"
	SessionSourceTypeDirectory  SessionSourceType = "directory"
	SessionSourceTypeTemplate   SessionSourceType = "template"
	SessionSourceTypeBlank      SessionSourceType = "blank"
)

// GetSessionSourceType returns the type of session source for the given value.
func (s *sessionManager) GetSessionSourceType(ctx context.Context, value string) (SessionSourceType, error) {
	// Check for blank value.
	if value == "" {
		return SessionSourceTypeBlank, nil
	}

	// Check for template.
	_, err := s.GetTemplate(ctx, value)
	if err == nil {
		return SessionSourceTypeTemplate, nil
	}

	// Check for repository.
	_, err = git.NewRemote(nil, &config.RemoteConfig{
		URLs: []string{value},
	}).List(&git.ListOptions{})
	if err == nil {
		return SessionSourceTypeRepository, nil
	}

	// Check if directory exists and is a directory.
	stat, err := os.Stat(value)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", os.ErrNotExist
	}

	return SessionSourceTypeDirectory, nil
}
