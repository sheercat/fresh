package runner

import (
	"fmt"
	logPkg "log"
	"time"

	"github.com/mattn/go-colorable"
)

type logFunc func(string, ...interface{})

var logger = logPkg.New(colorable.NewColorableStderr(), "", 0)

func newLogFunc(prefix string) func(string, ...interface{}) {
	color, clear := "", ""
	if settings.Colors {
		color = fmt.Sprintf("\033[%sm", logColor(prefix))
		clear = fmt.Sprintf("\033[%sm", colors["reset"])
	}
	prefix = fmt.Sprintf("%-11s", prefix)

	return func(format string, v ...interface{}) {
		now := time.Now()
		fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
		timeString := now.Format("15:04:03")
		format = fmt.Sprintf("%s%s %s |%s %s", color, timeString, prefix, clear, format)
		if len(v) == 0 {
			logger.Print(format)
		} else {
			logger.Printf(format, v...)
		}
	}
}

func fatal(err error) {
	logger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog("%s", string(p))

	return len(p), nil
}
