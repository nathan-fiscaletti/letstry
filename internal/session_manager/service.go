package session_manager

import (
	"context"
	"fmt"

	"github.com/nathan-fiscaletti/letstry/internal/storage"
)

var (
	ErrSessionManagerNotFound = fmt.Errorf("session manager not found in context")
)

type sessionMgrCtxKey struct {
	value string
}

var (
	sessionMgrKey = sessionMgrCtxKey{"session_manager"}
)

type sessionManager struct {
	storage *storage.Storage
}

// ContextWithSessionManager returns a new context with the session manager
func ContextWithSessionManager(ctx context.Context) context.Context {
	return context.WithValue(ctx, sessionMgrKey, sessionManager{
		storage: storage.GetStorage(),
	})
}

// GetSessionManager returns the session manager
func GetSessionManager(ctx context.Context) (*sessionManager, error) {
	mgr := ctx.Value(sessionMgrKey)
	if mgr != nil {
		if sessionMgr, ok := mgr.(sessionManager); ok {
			return &sessionMgr, nil
		}
	}

	return nil, ErrSessionManagerNotFound
}
