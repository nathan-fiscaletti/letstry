package commands

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Command func(context.Context, []string) error

func GetCallerName() string {
	cmd := filepath.Base(os.Args[0])

	if runtime.GOOS == "windows" {
		cmd = strings.TrimSuffix(cmd, ".exe")
	}

	return cmd
}
