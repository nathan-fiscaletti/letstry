package environment

import (
	"context"
	"errors"
	"os"
)

type appEnvironmentCtxKey struct {
	Name string
}

var (
	appEnvironmentKey appEnvironmentCtxKey = appEnvironmentCtxKey{"app_environment"}
)

var (
	ErrAppEnvironmentNotFound = errors.New("app environment not found")
)

type AppEnvironment struct {
	DebuggerAttached bool
}

func AppEnvironmentFromContext(ctx context.Context) (AppEnvironment, error) {
	if env, ok := ctx.Value(appEnvironmentKey).(AppEnvironment); ok {
		return env, nil
	}

	return AppEnvironment{}, ErrAppEnvironmentNotFound
}

func ContextWithAppEnvironment(ctx context.Context) context.Context {
	return context.WithValue(ctx, appEnvironmentKey, AppEnvironment{
		DebuggerAttached: os.Getenv("DEBUGGER_ATTACHED") == "true",
	})
}
