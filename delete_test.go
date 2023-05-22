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
	"github.com/Drafteame/mgorepo/seed"
	"github.com/Drafteame/mgorepo/testdata/repository/daos"
	ptesting "github.com/Drafteame/mgorepo/testing"
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

		data := daos.TestDAO{
			ID:         &oid,
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
	})

	t.Run("success delete with no affected", func(t *testing.T) {
		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		data := daos.TestDAO{
			ID:         &oid,
			Sortable:   0,
			Identifier: "asd",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
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

		assert.Equal(t, primitive.DateTime(0), data.DeletedAt)
	})
}
