package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func TestRepository_Create(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	c := clock.NewTest(time.Now()).ForceUTC()

	t.Run("success create", func(t *testing.T) {
		expected := testModel{
			Identifier: "identifier",
		}

		repo := newTestRepository(d).SetClock(c)
		model, err := repo.Create(context.Background(), expected)

		assert.Nil(t, err)
		assert.NotEmpty(t, model.ID)
		ptesting.AssertDate(t, c.Now(), model.CreatedAt)
		ptesting.AssertDate(t, c.Now(), model.UpdatedAt)
	})

	t.Run("success create with no timestamps", func(t *testing.T) {
		expected := testModel{
			Identifier: "identifier",
		}

		repo := newTestRepository(d).SetClock(c).WithTimestamps(false)
		model, err := repo.Create(context.Background(), expected)

		assert.Nil(t, err)
		assert.NotEmpty(t, model.ID)
		ptesting.AssertEmptyDate(t, model.CreatedAt)
		ptesting.AssertEmptyDate(t, model.UpdatedAt)
	})
}
