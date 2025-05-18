package helpers

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// SetupLogger sets up the log output to a timestamped file inside the existing "./log" folder.
// It assumes "./log" folder is already created.
func SetupLogger() (*os.File, error) {
	logFilename := "logfile_" + time.Now().Format("20060102_150405") + ".log"
	logFilePath := filepath.Join("./log", logFilename)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return file, nil
}
