package logging

import (
	"os"
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s", "log/")
}

// getLogFileName get the save name of the log file
func getLogFileName() string {

	env := os.Getenv("ENV")

	if env == "" {
		env = "deploy"
	}

	return fmt.Sprintf("%s%s.%s",
		env,
		time.Now().Format("2006-01-02"),
		"log",
	)
}