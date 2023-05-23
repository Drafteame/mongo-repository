package mgorepo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/driver"
)

func TestNewRepository(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	repository := NewRepository[
		testModel,
		testDAO,
		searchFilters,
		searchOptions,
		updateFields,
	](
		d,
		collection,
		getFilterBuilders(),
		getUpdateBuilders(),
	)

	assert.NotNil(t, repository.clock, "clock should not be nil")
	assert.NotEmpty(t, repository.db, "db should not be empty")
	assert.NotEmpty(t, repository.collectionName, "collectionName should not be empty")
	assert.NotEmpty(t, repository.filterBuilders, "filterBuilders should not be empty")
	assert.NotEmpty(t, repository.updateBuilders, "updateBuilders should not be empty")
}
