package logger

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	Prefix     = "[mongo-repository]"
	DateFormat = "2006-01-02T15:04:05.000Z07:00"

	debugFormat = "%s %s%s %s%s %s%s\n"
	errorFormat = "%s %s%s %s%s %s%s %s%s\n"

	debugPrefix    = "[DEBUG]"
	errorPrefix    = "[ERROR]"
	actionSection  = "action="
	messageSection = "message="
	errorSection   = "error="
	timeSection    = "time="
)

var (
	magenta = color.New(color.FgMagenta)
	yellow  = color.New(color.FgYellow)
	red     = color.New(color.FgRed)
	blue    = color.New(color.FgBlue)
)

type Logger struct {
	log     *log.Logger
	colored bool
}

func New() Logger {
	yellow.DisableColor()
	red.DisableColor()
	blue.DisableColor()
	magenta.DisableColor()

	return Logger{
		log: log.New(log.Writer(), "", 0),
	}
}

func (l Logger) Colored() Logger {
	yellow.EnableColor()
	red.EnableColor()
	blue.EnableColor()
	magenta.EnableColor()

	return l
}

func (l Logger) Debugf(action, message string, args ...any) {
	timeVal := time.Now().Format(DateFormat)
	message = strings.ReplaceAll(fmt.Sprintf(message, args...), `"`, `\"`)

	l.log.SetPrefix(magenta.Sprint(Prefix) + " ")

	l.log.Printf(
		debugFormat,
		yellow.Sprint(debugPrefix),
		blue.Sprint(timeSection),
		fmt.Sprintf(`"%s"`, timeVal),
		blue.Sprint(actionSection),
		fmt.Sprintf(`"%s"`, action),
		blue.Sprint(messageSection),
		fmt.Sprintf(`"%s"`, message),
	)
}

func (l Logger) Errorf(err error, action, message string, args ...any) {
	timeVal := time.Now().Format(DateFormat)
	message = strings.ReplaceAll(fmt.Sprintf(message, args...), `"`, `\"`)

	l.log.SetPrefix(magenta.Sprint(Prefix) + " ")

	l.log.Printf(
		errorFormat,
		red.Sprint(errorPrefix),
		blue.Sprint(timeSection),
		fmt.Sprintf(`"%s"`, timeVal),
		blue.Sprint(actionSection),
		fmt.Sprintf(`"%s"`, action),
		blue.Sprint(messageSection),
		fmt.Sprintf(`"%s"`, message),
		blue.Sprint(errorSection),
		fmt.Sprintf(`"%v"`, err),
	)
}
