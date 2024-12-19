package mgorepo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/internal/seed"
)

func TestRepository_HardDelete(t *testing.T) {
	t.Run("success hard delete", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		repo := newTestRepository(d)
		oid := primitive.NewObjectID()

		dao := testDAO{
			ID:         oid,
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
		d := getTestDriver(t)
		repo := newTestRepository(d)

		deleted, err := repo.HardDelete(context.Background(), "invalid_id")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(0), deleted)
	})
}
