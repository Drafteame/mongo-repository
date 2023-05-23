package mgorepo

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/env"
	"github.com/Drafteame/mgorepo/logger"
)

type Repository[M Model, D Dao, SF SearchFilters, SO SearchOptions[SF], UF UpdateFields] struct {
	db             Driver
	clock          Clock
	log            Logger
	logLevel       string
	searchLimit    int64
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
	SO SearchOptions[SF],
	UF UpdateFields,
](
	db Driver,
	collectionName string,
	filterBuilders []func(SF) (*bson.E, error),
	updateBuilders []func(UF) (*bson.E, error),
) Repository[M, D, SF, SO, UF] {
	return Repository[M, D, SF, SO, UF]{
		db:             db,
		clock:          clock.New(),
		log:            logger.New(),
		logLevel:       strings.ToUpper(env.GetString(env.MongoLogLevelEnv)),
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

func (r Repository[M, D, SF, SO, UF]) Db() Driver {
	return r.db
}

func (r Repository[M, D, SF, SO, UF]) Clock() Clock {
	return r.clock
}

func (r Repository[M, D, SF, SO, UF]) Log() Logger {
	return r.log
}

func (r Repository[M, D, SF, SO, UF]) CollectionName() string {
	return r.collectionName
}

func (r Repository[M, D, SF, SO, UF]) Collection() *mongo.Collection {
	return r.db.Client().Database(r.db.DbName()).Collection(r.collectionName)
}

func (r Repository[M, D, SF, SO, UF]) SetUpdatedAtField(updatedAtField string) Repository[M, D, SF, SO, UF] {
	if updatedAtField == "" {
		updatedAtField = DefaultUpdatedAtField
	}

	r.updatedAtField = updatedAtField

	return r
}

func (r Repository[M, D, SF, SO, UF]) SetCreatedAtField(createdAtField string) Repository[M, D, SF, SO, UF] {
	if createdAtField == "" {
		createdAtField = DefaultCreatedAtField
	}

	r.createdAtField = createdAtField

	return r
}

func (r Repository[M, D, SF, SO, UF]) SetDeletedAtField(deletedAtField string) Repository[M, D, SF, SO, UF] {
	if deletedAtField == "" {
		deletedAtField = DefaultDeletedAtField
	}

	r.deletedAtField = deletedAtField
	return r
}

func (r Repository[M, D, SF, SO, UF]) SetLogger(log Logger) Repository[M, D, SF, SO, UF] {
	r.log = log
	return r
}

func (r Repository[M, D, SF, SO, UF]) SetClock(clock Clock) Repository[M, D, SF, SO, UF] {
	r.clock = clock
	return r
}

func (r Repository[M, D, SF, SO, UF]) SetLogLevel(logLevel string) Repository[M, D, SF, SO, UF] {
	r.logLevel = strings.ToUpper(logLevel)
	return r
}

func (r Repository[M, D, SF, SO, UF]) SetSearchLimit(searchLimit int64) Repository[M, D, SF, SO, UF] {
	if searchLimit <= 0 {
		searchLimit = DefaultSearchLimit
	}

	r.searchLimit = searchLimit
	return r
}

func (r Repository[M, D, SF, SO, UF]) WithTimestamps(withTimestamps bool) Repository[M, D, SF, SO, UF] {
	r.withTimestamps = withTimestamps
	return r
}

func (r Repository[M, D, SF, SO, UF]) logError(err error, action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelError {
		r.log.Errorf(err, action, message, args...)
	}
}

func (r Repository[M, D, SF, SO, UF]) logDebug(action, message string, args ...any) {
	if r.log != nil && r.logLevel == logger.LevelDebug {
		r.log.Debugf(action, message, args...)
	}
}
