package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/testdata/domain"
	ptesting "github.com/Drafteame/mgorepo/testing"
)

func TestRepository_Create(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	c := clock.NewTest(time.Now()).ForceUTC()

	expected := domain.TestModel{
		Identifier: "identifier",
	}

	repo := newTestRepository(d).SetClock(c)
	model, err := repo.Create(context.Background(), expected)

	assert.Nil(t, err)
	assert.NotEmpty(t, model.ID)
	ptesting.AssertDate(t, c.Now(), model.CreatedAt)
	ptesting.AssertDate(t, c.Now(), model.UpdatedAt)
}
