package logging

import (
	"log"
	"os"
)

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stderr, "letstry: ", log.LstdFlags)
	}
	return logger
}
