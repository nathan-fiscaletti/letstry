package commands

import (
	"context"
	"debug/buildinfo"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/letstry/internal/logging"
)

func VersionCommand() Command {
	return Command{
		Name:             CommandVersion,
		ShortDescription: "Display the version of " + MainName(),
		Description:      "This command will display the version of " + MainName() + ".",

		Executor: func(ctx context.Context, args []string) error {
			exe, err := os.Executable()
			if err != nil {
				return fmt.Errorf("failed to get executable path: %v", err)
			}

			info, err := buildinfo.ReadFile(exe)
			if err != nil {
				return fmt.Errorf("failed to read build info: %v", err)
			}

			logger, err := logging.LoggerFromContext(ctx)
			if err != nil {
				return err
			}

			logger.Println("module path:", info.Path)

			latestVersion, err := getLatestVersion(info.Path)
			if err != nil {
				return fmt.Errorf("failed to get latest version: %v", err)
			}

			logger.Println("version:", info.Main.Version)
			if info.Main.Version != latestVersion && info.Main.Version != "(devel)" {
				logger.Println(color.HiWhiteString("!! new version (" + color.HiGreenString(latestVersion) + ") available"))
				logger.Println("run 'go install", info.Path+"@"+latestVersion+"' to update")
			} else {
				logger.Println(color.HiWhiteString("you are running the latest version"))
			}

			return nil
		},
	}
}

// getRootModulePath returns the root module path by removing the subdirectory.
func getRootModulePath(modulePath string) string {
	parts := strings.Split(modulePath, "/")
	if len(parts) > 3 {
		return strings.Join(parts[:3], "/")
	}
	return modulePath
}

func getLatestVersion(modulePath string) (string, error) {
	rootModulePath := getRootModulePath(modulePath)
	cmd := exec.Command("go", "list", "-m", "-versions", rootModulePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	versions := strings.Fields(string(output))
	if len(versions) < 2 {
		return "", fmt.Errorf("no versions found for module %s", rootModulePath)
	}

	latestVersion := versions[len(versions)-1]
	return latestVersion, nil
}
