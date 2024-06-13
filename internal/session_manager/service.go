package session_manager

type sessionManager struct{}

var mgr *sessionManager = &sessionManager{}

// GetSessionManager returns the session manager
func GetSessionManager() *sessionManager {
	return mgr
}
