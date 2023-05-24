package logger

import (
	"bytes"
	"log"
	"testing"
)

type TestLogger struct {
	Logger
	now  string
	buff *bytes.Buffer
}

func NewTest(t *testing.T) TestLogger {
	yellow.DisableColor()
	red.DisableColor()
	blue.DisableColor()
	magenta.DisableColor()

	out := []byte{}

	buff := bytes.NewBuffer(out)
	logger := log.New(buff, "test", 0)

	tl := TestLogger{
		buff: buff,
		Logger: Logger{
			log: logger,
		},
	}

	t.Cleanup(func() {
		buff.Reset()
	})

	return tl
}

func (tl TestLogger) SetClock(clock Clock) TestLogger {
	tl.clock = clock
	return tl
}

func (tl TestLogger) GetOutput() string {
	return tl.buff.String()
}
