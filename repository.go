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

type Repository[M Model, D Dao, SF SearchFilters, SORD SearchOrderer, SO SearchOptioner[SF, SORD], UF UpdateFields] struct {
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
	SORD SearchOrderer,
	SO SearchOptioner[SF, SORD],
	UF UpdateFields,
](
	db Driver,
	collectionName string,
	filterBuilders []func(SF) (*bson.E, error),
	updateBuilders []func(UF) (*bson.E, error),
) Repository[M, D, SF, SORD, SO, UF] {
	return Repository[M, D, SF, SORD, SO, UF]{
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

func (r Repository[M, D, SF, SORD, SO, UF]) Db() Driver {
	return r.db
}

func (r Repository[M, D, SF, SORD, SO, UF]) Clock() Clock {
	return r.clock
}

func (r Repository[M, D, SF, SORD, SO, UF]) Logger() Logger {
	return r.log
}

func (r Repository[M, D, SF, SORD, SO, UF]) CollectionName() string {
	return r.collectionName
}

func (r Repository[M, D, SF, SORD, SO, UF]) Collection() *mongo.Collection {
	return r.db.Client().Database(r.db.DbName()).Collection(r.collectionName)
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetUpdatedAtField(updatedAtField string) Repository[M, D, SF, SORD, SO, UF] {
	if updatedAtField == "" {
		updatedAtField = DefaultUpdatedAtField
	}

	r.updatedAtField = updatedAtField

	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetCreatedAtField(createdAtField string) Repository[M, D, SF, SORD, SO, UF] {
	if createdAtField == "" {
		createdAtField = DefaultCreatedAtField
	}

	r.createdAtField = createdAtField

	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetDeletedAtField(deletedAtField string) Repository[M, D, SF, SORD, SO, UF] {
	if deletedAtField == "" {
		deletedAtField = DefaultDeletedAtField
	}

	r.deletedAtField = deletedAtField
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetLogger(log Logger) Repository[M, D, SF, SORD, SO, UF] {
	r.log = log
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetClock(clock Clock) Repository[M, D, SF, SORD, SO, UF] {
	r.clock = clock
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetLogLevel(logLevel string) Repository[M, D, SF, SORD, SO, UF] {
	r.logLevel = strings.ToUpper(logLevel)
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) SetDefaultSearchLimit(searchLimit int) Repository[M, D, SF, SORD, SO, UF] {
	if searchLimit <= 0 {
		searchLimit = DefaultSearchLimit
	}

	r.searchLimit = searchLimit
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) WithTimestamps(withTimestamps bool) Repository[M, D, SF, SORD, SO, UF] {
	r.withTimestamps = withTimestamps
	return r
}

func (r Repository[M, D, SF, SORD, SO, UF]) logErrorf(err error, action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelError {
		r.log.Errorf(err, action, message, args...)
	}
}

func (r Repository[M, D, SF, SORD, SO, UF]) logDebugf(action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelDebug {
		r.log.Debugf(action, message, args...)
	}
}
