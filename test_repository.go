package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/logger"
)

const collection = "model"

func newTestRepository(driver Driver) Repository[testModel, testDAO, searchFilters, SearchOrders, SearchOptions[searchFilters, SearchOrders], updateFields] {
	return NewRepository[
		testModel,
		testDAO,
		searchFilters,
		SearchOrders,
		SearchOptions[searchFilters, SearchOrders],
		updateFields,
	](
		driver,
		collection,
		getFilterBuilders(),
		getUpdateBuilders(),
	).
		SetClock(clock.New().ForceUTC()).
		SetLogLevel(logger.LevelDebug)
}

func getFilterBuilders() []func(searchFilters) (*bson.E, error) {
	return []func(searchFilters) (*bson.E, error){
		buildIDFilter,
		buildIdentifierFilter,
		buildSortableGraterThanFilter,
	}
}

func getUpdateBuilders() []func(updateFields) (*bson.E, error) {
	return []func(updateFields) (*bson.E, error){
		buildIdentifierUpdate,
		buildSortableUpdate,
		buildUpdatedAtUpdate,
	}
}
