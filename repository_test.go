package mgorepo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/domain/options"
	"github.com/Drafteame/mgorepo/testdata/repository/daos"
)

func TestNewRepository(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	repository := NewRepository[
		domain.TestModel,
		daos.TestDAO,
		options.SearchFilters,
		options.SearchOrders,
		options.SearchOptions,
		options.UpdateFields,
	](
		d,
		collection,
		getFilterBuilders(),
		getOrderBuilders(),
		getUpdateBuilders(),
	)

	assert.NotNil(t, repository.clock, "clock should not be nil")
	assert.NotEmpty(t, repository.db, "db should not be empty")
	assert.NotEmpty(t, repository.collectionName, "collectionName should not be empty")
	assert.NotEmpty(t, repository.filterBuilders, "filterBuilders should not be empty")
	assert.NotEmpty(t, repository.orderBuilders, "orderBuilders should not be empty")
	assert.NotEmpty(t, repository.updateBuilders, "updateBuilders should not be empty")
}
