package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/logger"
	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/domain/options"
	"github.com/Drafteame/mgorepo/testdata/repository/builders"
	"github.com/Drafteame/mgorepo/testdata/repository/daos"
)

const collection = "model"

func newTestRepository(driver Driver) Repository[domain.TestModel, daos.TestDAO, options.SearchFilters, options.SearchOrders, options.SearchOptions, options.UpdateFields] {
	return NewRepository[
		domain.TestModel,
		daos.TestDAO,
		options.SearchFilters,
		options.SearchOrders,
		options.SearchOptions,
		options.UpdateFields,
	](
		driver,
		collection,
		getFilterBuilders(),
		getOrderBuilders(),
		getUpdateBuilders(),
	).
		SetClock(clock.New().ForceUTC()).
		SetLogLevel(logger.LevelDebug)
}

func getFilterBuilders() []func(options.SearchFilters) (*bson.E, error) {
	return []func(options.SearchFilters) (*bson.E, error){
		builders.BuildIDFilter,
		builders.BuildIdentifierFilter,
		builders.BuildSortableGraterThanFilter,
	}
}

func getOrderBuilders() []func(options.SearchOrders) (*bson.E, error) {
	return []func(options.SearchOrders) (*bson.E, error){
		builders.BuildSortableOrder,
		builders.BuildCreatedAtOrder,
	}
}

func getUpdateBuilders() []func(options.UpdateFields) (*bson.E, error) {
	return []func(options.UpdateFields) (*bson.E, error){
		builders.BuildIdentifierUpdate,
		builders.BuildSortableUpdate,
		builders.BuildUpdatedAtUpdate,
	}
}
