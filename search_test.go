package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/seed"
	testoptions "github.com/Drafteame/mgorepo/testdata/domain/options"
	testdaos "github.com/Drafteame/mgorepo/testdata/repository/daos"
)

func TestRepository_Search(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	db := d.Client().Database(d.DbName())

	t.Run("success search", func(t *testing.T) {
		now := time.Now().UTC()
		dao := testdaos.TestDAO{
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions()
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
	})

	t.Run("empty search", func(t *testing.T) {
		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions()
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 0)
	})

	t.Run("success search with limit", func(t *testing.T) {
		now := time.Now().UTC()
		dao := testdaos.TestDAO{
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions().WithLimit(1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
	})

	t.Run("success search with offset", func(t *testing.T) {
		now := time.Now().UTC()
		dao := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions().WithSkip(1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
		assert.Equal(t, dao2.Sortable, models[0].Sortable)
	})

	t.Run("success search with sort", func(t *testing.T) {
		now := time.Now().UTC()
		dao := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions().WithSortSorter(1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.Equal(t, dao.Sortable, models[0].Sortable)
		assert.Equal(t, dao2.Sortable, models[1].Sortable)
	})

	t.Run("success search with sort desc", func(t *testing.T) {
		now := time.Now().UTC()
		dao := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions().WithSortSorter(-1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.Equal(t, dao2.Sortable, models[0].Sortable)
		assert.Equal(t, dao.Sortable, models[1].Sortable)
	})

	t.Run("success omitting deleted items", func(t *testing.T) {
		now := time.Now().UTC()
		pnow := primitive.NewDateTimeFromTime(now)

		dao := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testdaos.TestDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
			DeletedAt:  pnow,
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := testoptions.NewSearchOptions()
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
		assert.Equal(t, dao.Sortable, models[0].Sortable)
	})
}
