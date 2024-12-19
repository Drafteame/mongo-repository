package mgorepo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/internal/seed"
)

func TestRepository_HardDeleteMany(t *testing.T) {
	t.Parallel()

	t.Run("success hard delete many", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

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

		filters := newSearchFilters().WithSortableGreaterThan(50)

		repo := newTestRepository(d)

		total, err := repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		deleted, errDelete := repo.HardDeleteMany(context.Background(), filters)
		if errDelete != nil {
			t.Fatal(errDelete)
		}

		assert.Equal(t, total, deleted)

		total, err = repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(0), total)
	})

	t.Run("error hard delete many with empty filters", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)

		repo := newTestRepository(d)

		deleted, err := repo.HardDeleteMany(context.Background(), newSearchFilters())

		assert.Equal(t, int64(0), deleted)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrEmptyFilters)
	})

	t.Run("no error when filters match 0 documents", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

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

		filters := newSearchFilters().WithSortableGreaterThan(1000)

		repo := newTestRepository(d)

		deleted, err := repo.HardDeleteMany(context.Background(), filters)

		assert.Equal(t, int64(0), deleted)
		assert.NoError(t, err)

		total, err := repo.Count(context.Background(), newSearchFilters())
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(100), total)
	})
}
