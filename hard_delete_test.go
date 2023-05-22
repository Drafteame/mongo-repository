package mgorepo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/seed"
)

func TestRepository_HardDelete(t *testing.T) {
	d, errDriver := driver.NewTest(t)
	if errDriver != nil {
		t.Fatal(errDriver)
	}

	db := d.Client().Database(d.DbName())

	t.Run("success hard delete", func(t *testing.T) {
		repo := newTestRepository(d)
		oid := primitive.NewObjectID()

		dao := testDAO{
			ID:         &oid,
			Identifier: "test",
			Sortable:   randomNumber(),
		}

		seed.InsertOne(t, db, collection, dao)

		deleted, err := repo.HardDelete(context.Background(), oid.Hex())
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(1), deleted)

		total, errCount := db.Collection(collection).CountDocuments(context.Background(), bson.D{})
		if errCount != nil {
			t.Fatal(errCount)
		}

		assert.Equal(t, int64(0), total)
	})

	t.Run("error hard delete with invalid id", func(t *testing.T) {
		repo := newTestRepository(d)

		deleted, err := repo.HardDelete(context.Background(), "invalid_id")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(0), deleted)
	})
}
