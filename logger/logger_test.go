package logger

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/logger/mocks"
)

func TestNew(t *testing.T) {
	t.Run("should create new logger no colored", func(t *testing.T) {
		logger := New()

		assert.NotNil(t, logger)
		assert.NotNil(t, logger.log)
		assert.False(t, logger.colored)
	})

	t.Run("should create new logger colored", func(t *testing.T) {
		logger := New().Colored()

		assert.NotNil(t, logger)
		assert.NotNil(t, logger.log)
		assert.True(t, logger.IsColored())
	})

	t.Run("should create new logger with clock", func(t *testing.T) {
		c := mocks.NewClock(t)
		logger := New().SetClock(c)

		assert.NotNil(t, logger)
		assert.NotNil(t, logger.log)
		assert.NotNil(t, logger.clock)
	})
}

func TestLogger_printMessages(t *testing.T) {
	t.Run("should print debug log", func(t *testing.T) {
		now := time.Now()
		c := mocks.NewClock(t)
		c.On("Now").Return(now)

		tl := NewTest(t).SetClock(c)

		tl.Debugf("action", "message")

		expected := fmt.Sprintf(
			debugFormat,
			debugPrefix,
			timeSection,
			fmt.Sprintf(`"%s"`, now.Format(DateFormat)),
			actionSection,
			`"action"`,
			messageSection,
			`"message"`,
		)

		expected = fmt.Sprintf("%s %s", Prefix, expected)

		assert.Equal(t, expected, tl.GetOutput())
	})

	t.Run("should print error log", func(t *testing.T) {
		now := time.Now()
		c := mocks.NewClock(t)
		c.On("Now").Return(now)

		tl := NewTest(t).SetClock(c)

		err := errors.New("some error")
		tl.Errorf(err, "action", "message")

		expected := fmt.Sprintf(
			errorFormat,
			errorPrefix,
			timeSection,
			fmt.Sprintf(`"%s"`, now.Format(DateFormat)),
			actionSection,
			`"action"`,
			messageSection,
			`"message"`,
			errorSection,
			`"some error"`,
		)

		expected = fmt.Sprintf("%s %s", Prefix, expected)

		assert.Equal(t, expected, tl.GetOutput())
	})
}
