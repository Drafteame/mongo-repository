package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/internal/seed"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func TestRepository_Delete(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	db := d.Client().Database(d.DbName())

	t.Run("success delete", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		data := testDAO{
			ID:         oid,
			Sortable:   0,
			Identifier: "asd",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, data)

		repo := newTestRepository(d).SetClock(c)

		deletedCount, err := repo.Delete(context.Background(), oid.Hex())

		assert.Nil(t, err)
		assert.Equal(t, int64(1), deletedCount)

		errFind := db.Collection(collection).FindOne(context.Background(), bson.D{{Key: "_id", Value: oid}}).Decode(&data)
		if errFind != nil {
			t.Fatal(errFind)
		}

		ptesting.AssertDate(t, c.Now(), data.DeletedAt.Time().UTC())
		ptesting.AssertDate(t, c.Now(), data.UpdatedAt.Time().UTC())
	})

	t.Run("success delete with no affected", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		pnow := primitive.NewDateTimeFromTime(c.Now())

		data := testDAO{
			ID:         oid,
			Sortable:   0,
			Identifier: "asd",
			CreatedAt:  pnow,
			UpdatedAt:  pnow,
		}

		seed.InsertOne(t, db, collection, data)

		repo := newTestRepository(d).SetClock(c)

		deletedCount, err := repo.Delete(context.Background(), primitive.NewObjectID().Hex())

		assert.Nil(t, err)
		assert.Equal(t, int64(0), deletedCount)

		errFind := db.Collection(collection).FindOne(context.Background(), bson.D{{Key: "_id", Value: oid}}).Decode(&data)
		if errFind != nil {
			t.Fatal(errFind)
		}

		ptesting.AssertEmptyDate(t, data.DeletedAt)
		ptesting.AssertDate(t, pnow, data.UpdatedAt)
	})

	t.Run("success delete with no timestamps", func(t *testing.T) {
		oid := primitive.NewObjectID()

		data := testDAO{
			ID:         oid,
			Sortable:   0,
			Identifier: "asd",
		}

		seed.InsertOne(t, db, collection, data)

		repo := newTestRepository(d)

		deletedCount, err := repo.Delete(context.Background(), oid.Hex())

		assert.Nil(t, err)
		assert.Equal(t, int64(1), deletedCount)

		total, errCount := repo.Count(context.Background(), newSearchFilters())
		if errCount != nil {
			t.Fatal(errCount)
		}

		assert.Equal(t, int64(0), total)
	})
}
