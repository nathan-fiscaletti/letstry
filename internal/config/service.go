package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/nathan-fiscaletti/letstry/internal/storage"
)

var config *Config

func GetConfig() (*Config, error) {
	store := storage.GetStorage()

	var file *os.File

	if !store.Exists("config.json") {
		defaultConfig := getDefaultConfig()

		var defaultValue []byte
		defaultValue, err := json.Marshal(defaultConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %v", err)
		}

		file, err = store.OpenFileWithDefaultContent("config.json", defaultValue)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %v", err)
		}
	} else {
		var err error
		file, err = store.OpenFile("config.json")
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %v", err)
		}
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}

func getDefaultConfig() *Config {
	currentUser, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("failed to get current user: %v", err))
	}

	var vsCodePath string
	switch os := runtime.GOOS; os {
	case "darwin":
		vsCodePath = filepath.Join(currentUser.HomeDir, "Applications", "Visual Studio Code.app", "Contents", "Resources", "app", "bin", "code")
	case "linux":
		vsCodePath = filepath.Join(currentUser.HomeDir, "bin", "code")
	case "windows":
		vsCodePath = filepath.Join(currentUser.HomeDir, "AppData", "Local", "Programs", "Microsoft VS Code", "Code.exe")
	}

	editors := []Editor{
		{
			Name:     "vscode",
			ExecPath: vsCodePath,
			Args:     "",
		},
	}

	return &Config{
		DefaultEditorName: "vscode",
		AvailableEditors:  editors,
	}
}
