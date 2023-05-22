package mgorepo

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/seed"
	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/repository/daos"
)

func TestRepository_Get(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	db := d.Client().Database(d.DbName())

	t.Run("get error not found", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := daos.TestDAO{
			ID:         &oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)

		model, err := repo.Get(context.Background(), primitive.NewObjectID().Hex())

		assert.Empty(t, model)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNotFound))
	})

	t.Run("get error not found on deleted doc", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := daos.TestDAO{
			ID:         &oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			DeletedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)

		model, err := repo.Get(context.Background(), oid.Hex())

		assert.Empty(t, model)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNotFound))
	})

	t.Run("get success", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := daos.TestDAO{
			ID:         &oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)
		model, err := repo.Get(context.Background(), oid.Hex())

		assert.NotNil(t, model)
		assert.NoError(t, err)
		assert.IsType(t, domain.TestModel{}, model)
	})
}
