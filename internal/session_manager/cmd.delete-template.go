package session_manager

import "context"

func (s *sessionManager) DeleteTemplate(ctx context.Context, t Template) error {
	return s.storage.DeleteDirectory(t.StoragePath())
}
