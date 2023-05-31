package mgorepo

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/internal/env"
	"github.com/Drafteame/mgorepo/logger"
)

type Repository[M Model, D Dao, SF SearchFilters, UF UpdateFields] struct {
	db             Driver
	clock          Clock
	log            Logger
	logLevel       string
	searchLimit    int
	collectionName string
	withTimestamps bool
	updatedAtField string
	createdAtField string
	deletedAtField string
	filterBuilders []func(SF) (*bson.E, error)
	updateBuilders []func(UF) (*bson.E, error)
}

func NewRepository[
	M Model,
	D Dao,
	SF SearchFilters,
	UF UpdateFields,
](
	db Driver,
	collectionName string,
	filterBuilders []func(SF) (*bson.E, error),
	updateBuilders []func(UF) (*bson.E, error),
) Repository[M, D, SF, UF] {
	return Repository[M, D, SF, UF]{
		db:             db,
		clock:          clock.New(),
		log:            logger.New(),
		logLevel:       strings.ToUpper(env.GetString(driver.MongoLogLevelEnv)),
		searchLimit:    DefaultSearchLimit,
		collectionName: collectionName,
		withTimestamps: true,
		updatedAtField: DefaultUpdatedAtField,
		createdAtField: DefaultCreatedAtField,
		deletedAtField: DefaultDeletedAtField,
		filterBuilders: filterBuilders,
		updateBuilders: updateBuilders,
	}
}

func (r Repository[M, D, SF, UF]) Db() Driver {
	return r.db
}

func (r Repository[M, D, SF, UF]) Clock() Clock {
	return r.clock
}

func (r Repository[M, D, SF, UF]) Logger() Logger {
	return r.log
}

func (r Repository[M, D, SF, UF]) CollectionName() string {
	return r.collectionName
}

func (r Repository[M, D, SF, UF]) Collection() *mongo.Collection {
	return r.db.Client().Database(r.db.DbName()).Collection(r.collectionName)
}

func (r Repository[M, D, SF, UF]) SetUpdatedAtField(updatedAtField string) Repository[M, D, SF, UF] {
	if updatedAtField == "" {
		updatedAtField = DefaultUpdatedAtField
	}

	r.updatedAtField = updatedAtField

	return r
}

func (r Repository[M, D, SF, UF]) SetCreatedAtField(createdAtField string) Repository[M, D, SF, UF] {
	if createdAtField == "" {
		createdAtField = DefaultCreatedAtField
	}

	r.createdAtField = createdAtField

	return r
}

func (r Repository[M, D, SF, UF]) SetDeletedAtField(deletedAtField string) Repository[M, D, SF, UF] {
	if deletedAtField == "" {
		deletedAtField = DefaultDeletedAtField
	}

	r.deletedAtField = deletedAtField
	return r
}

func (r Repository[M, D, SF, UF]) SetLogger(log Logger) Repository[M, D, SF, UF] {
	r.log = log
	return r
}

func (r Repository[M, D, SF, UF]) SetClock(clock Clock) Repository[M, D, SF, UF] {
	r.clock = clock
	return r
}

func (r Repository[M, D, SF, UF]) SetLogLevel(logLevel string) Repository[M, D, SF, UF] {
	r.logLevel = strings.ToUpper(logLevel)
	return r
}

func (r Repository[M, D, SF, UF]) SetDefaultSearchLimit(searchLimit int) Repository[M, D, SF, UF] {
	if searchLimit <= 0 {
		searchLimit = DefaultSearchLimit
	}

	r.searchLimit = searchLimit
	return r
}

func (r Repository[M, D, SF, UF]) WithTimestamps(withTimestamps bool) Repository[M, D, SF, UF] {
	r.withTimestamps = withTimestamps
	return r
}

func (r Repository[M, D, SF, UF]) logErrorf(err error, action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelError {
		r.log.Errorf(err, action, message, args...)
	}
}

func (r Repository[M, D, SF, UF]) logDebugf(action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelDebug {
		r.log.Debugf(action, message, args...)
	}
}
