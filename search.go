package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository[M, D, SF, SO, UF]) Search(ctx context.Context, opts SearchOptioner[SF, SO]) ([]M, error) {
	filters, findOpts, err := r.BuildSearchOptions(opts)
	if err != nil {
		r.logErrorf(err, actionSearch, "error building SearchOptions")
		return nil, err
	}

	return r.searchExecute(ctx, filters, findOpts)
}

func (r Repository[M, D, SF, SO, UF]) searchExecute(ctx context.Context, filters bson.D, findOptions *options.FindOptions) ([]M, error) {
	var result []D

	r.printSearchDebug(filters, findOptions)

	cursor, errFind := r.Collection().Find(ctx, filters, findOptions)
	if errFind != nil {
		r.logErrorf(errFind, actionSearch, "error searching %s", r.collectionName)
		return nil, errFind
	}

	if errDecode := cursor.All(ctx, &result); errDecode != nil {
		r.logErrorf(errDecode, actionSearch, "error decoding %s search result", r.collectionName)
		return nil, errDecode
	}

	models := make([]M, len(result))
	for i, m := range result {
		filler, ok := any(&m).(DaoFiller[M])
		if !ok {
			r.logErrorf(ErrInvalidDaoFiller, actionSearch, "error getting DaoFiller from %s search result", r.collectionName)
			return nil, ErrInvalidDaoFiller
		}

		models[i] = filler.ToModel()
	}

	return models, nil
}

func (r Repository[M, D, SF, SO, UF]) printSearchDebug(filters bson.D, findOpts *options.FindOptions) {
	var shallowOpts options.FindOptions
	var skip, limit int64

	if findOpts != nil {
		shallowOpts = *findOpts

		if findOpts.Skip != nil {
			skip = *findOpts.Skip
		}

		if findOpts.Limit != nil {
			limit = *findOpts.Limit
		}
	}

	r.logDebugf(actionSearch, "filters: %+v sort: %+v skip: %+v limit: %+v project: %+v", filters, shallowOpts.Sort, skip, limit, findOpts.Projection)
}
