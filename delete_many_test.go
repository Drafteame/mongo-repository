package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/seed"
	testoptions "github.com/Drafteame/mgorepo/testdata/domain/options"
	testdaos "github.com/Drafteame/mgorepo/testdata/repository/daos"
)

func TestRepository_DeleteMany(t *testing.T) {
	d, errDriver := driver.NewTest(t)
	if errDriver != nil {
		t.Fatal(errDriver)
	}

	db := d.Client().Database(d.DbName())

	t.Run("success delete many", func(t *testing.T) {
		c := clock.NewTest(time.Now()).ForceUTC()

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()
			dao := testdaos.TestDAO{
				ID:         &oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		filters := testoptions.NewSearchFilters().WithSortableGreaterThan(50)

		repo := newTestRepository(d).SetClock(c)

		total, err := repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		deleted, errDelete := repo.DeleteMany(context.Background(), filters)
		if errDelete != nil {
			t.Fatal(errDelete)
		}

		assert.Equal(t, total, deleted)
	})

	t.Run("error delete many with empty filters", func(t *testing.T) {
		repo := newTestRepository(d)

		_, err := repo.DeleteMany(context.Background(), testoptions.NewSearchFilters())

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrEmptyFilters)
	})
}
