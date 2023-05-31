package testing

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDate interface {
	Time() time.Time
}

type TimeTruncater interface {
	Year() int
	Truncate(d time.Duration) time.Time
}

// AssertEmptyDate compare mongo primitive.DateTime with a native time.Time object truncating nanoseconds to milliseconds.
func AssertEmptyDate(t *testing.T, actual any, msgAndArgs ...any) {
	if actual == nil {
		assert.Nil(t, actual)
		return
	}

	switch a := actual.(type) {
	case MongoDate:
		if aux := int64(primitive.NewDateTimeFromTime(a.Time())); aux == 0 {
			assert.Equal(t, int64(0), aux)
			return
		}

		AssertDate(t, getEmptyDateByYear(a.Time().Year()), actual, msgAndArgs...)
	case TimeTruncater:
		AssertDate(t, getEmptyDateByYear(a.Year()), actual, msgAndArgs...)
	default:
		t.Fatal("can't assert types different than time.Time or primitive.DateTime values")
	}
}

func getEmptyDateByYear(year int) time.Time {
	switch year {
	case 1970:
		return time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local)
	default:
		return time.Time{}
	}
}

// AssertDate compare mongo primitive.DateTime with a native time.Time object truncating nanoseconds to milliseconds to
// avoid
func AssertDate(t *testing.T, exp any, actual any, msgAndArgs ...any) {
	if exp == nil {
		assert.Nil(t, actual)
		return
	}

	normalizedExp, err := normalizeDate(t, exp)
	if err != nil {
		t.Fatal(err)
		return
	}

	normalizedActual, err := normalizeDate(t, actual)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, normalizedExp, normalizedActual, msgAndArgs...)
}

func normalizeDate(t *testing.T, date any) (time.Time, error) {
	if casted, ok := date.(MongoDate); ok {
		return normalizeMongoDate(t, casted), nil
	}

	if casted, ok := date.(TimeTruncater); ok {
		return normalizeTime(t, casted), nil
	}

	return time.Time{}, errors.New("can't assert types different than time.Time or primitive.DateTime values")
}

func normalizeMongoDate(t *testing.T, date MongoDate) time.Time {
	newDate := time.Time{}

	switch d := date.(type) {
	case *primitive.DateTime:
		if d != nil {
			newDate = d.Time().Truncate(time.Millisecond)
		}
	case primitive.DateTime:
		newDate = d.Time().Truncate(time.Millisecond)
	default:
		t.Fatal("can't assert types different than time.Time as newDate values")
	}

	return time.Date(newDate.Year(), newDate.Month(), newDate.Day(), newDate.Hour(), newDate.Minute(), newDate.Second(), newDate.Nanosecond(), time.Local)
}

func normalizeTime(t *testing.T, date any) time.Time {
	var newDate time.Time

	switch d := date.(type) {
	case *time.Time:
		if d != nil {
			newDate = d.Truncate(time.Millisecond)
		}
	case time.Time:
		newDate = d.Truncate(time.Millisecond)
	default:
		t.Fatal("can't assert types different than time.Time as newDate values")
	}

	return time.Date(newDate.Year(), newDate.Month(), newDate.Day(), newDate.Hour(), newDate.Minute(), newDate.Second(), newDate.Nanosecond(), time.Local)
}
