package session_manager

import (
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
