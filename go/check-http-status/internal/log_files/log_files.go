package log_files

import (
	"io"
	"os"

	"github.com/supermarine1377/monitoring-scripts/go/check-http-status/timeutil"
)

func New(createLogFile bool) ([]io.Writer, error) {
	files := make([]io.Writer, 0, 2)
	files = append(files, os.Stdout)

	if createLogFile {
		logFile, err := os.Create(fileName())
		if err != nil {
			return nil, err
		}
		files = append(files, logFile)
	}

	return files, nil
}

func fileName() string {
	return "check-http-status_" + timeutil.NowStr() + ".log"
}
