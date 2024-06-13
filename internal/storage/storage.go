package storage

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var storage *Storage

func init() {
	currentUser, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("failed to get current user: %v", err))
	}

	path := filepath.Join(currentUser.HomeDir, ".letstry")

	err = os.MkdirAll(path, 0755)
	if err != nil {
		panic(fmt.Errorf("failed to create storage directory: %v", err))
	}

	storage = &Storage{dir: path}
}

type Storage struct {
	dir string
}

func GetStorage() *Storage {
	return storage
}

func (s *Storage) OpenFile(name string) (*os.File, error) {
	return s.OpenFileWithDefaultContent(name, []byte{})
}

func (s *Storage) Exists(name string) bool {
	_, err := os.Stat(filepath.Join(s.dir, name))
	return !os.IsNotExist(err)
}

func (s *Storage) OpenFileWithDefaultContent(name string, defaultContent []byte) (*os.File, error) {
	filePath := filepath.Join(s.dir, name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		_, err = file.Write(defaultContent)
		if err != nil {
			return nil, err
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	return os.OpenFile(filePath, os.O_RDWR, 0644)
}
