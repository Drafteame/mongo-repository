package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/internal/seed"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func TestRepository_DeleteMany(t *testing.T) {
	t.Run("success delete many", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		c := clock.NewTest(time.Now()).ForceUTC()

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()

			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		repo := newTestRepository(d).SetClock(c)

		filters := newSearchFilters().WithSortableGreaterThan(50)
		bsonFilters, err := repo.BuildSearchFilters(filters)
		if err != nil {
			t.Fatal(err)
		}

		total, err := repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		deleted, errDelete := repo.DeleteMany(context.Background(), filters)
		if errDelete != nil {
			t.Fatal(errDelete)
		}

		totalAfterDel, errAfterDel := repo.Count(context.Background(), filters)
		if errAfterDel != nil {
			t.Fatal(errAfterDel)
		}

		assert.Equal(t, totalAfterDel, total-deleted)

		cursor, errFind := db.Collection(collection).Find(context.Background(), bsonFilters)
		if errFind != nil {
			t.Fatal(errFind)
		}

		for cursor.Next(context.Background()) {
			var dao testDAO
			if errDecode := cursor.Decode(&dao); errDecode != nil {
				t.Fatal(errDecode)
			}

			ptesting.AssertDate(t, c.Now(), dao.DeletedAt.Time().UTC())
			ptesting.AssertDate(t, c.Now(), dao.UpdatedAt.Time().UTC())
		}
	})

	t.Run("error delete many with empty filters", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)

		repo := newTestRepository(d)
		_, err := repo.DeleteMany(context.Background(), newSearchFilters())

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrEmptyFilters)
	})

	t.Run("delete many with no timestamps", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		daos := make([]any, 0, 100)
		c := clock.NewTest(time.Now()).ForceUTC()

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()
			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		repo := newTestRepository(d).SetClock(c).WithTimestamps(false)

		filters := newSearchFilters().WithSortableGreaterThan(50)

		bsonFilters, err := repo.BuildSearchFilters(filters)
		if err != nil {
			t.Fatal(err)
		}

		total, err := db.Collection(collection).CountDocuments(context.Background(), bsonFilters)
		if err != nil {
			t.Fatal(err)
		}

		deleted, errDelete := repo.DeleteMany(context.Background(), filters)
		if errDelete != nil {
			t.Fatal(errDelete)
		}

		remaining, err := db.Collection(collection).CountDocuments(context.Background(), bsonFilters)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, total, remaining+deleted)
	})
}
