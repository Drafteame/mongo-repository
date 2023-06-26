package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository[M, D, SF, SORD, SO, UF]) UpdateMany(ctx context.Context, opts SF, data UF) (int64, error) {
	filters, err := r.updateManyFilters(opts)
	if err != nil {
		return 0, err
	}

	update, err := r.updateData(data, false)
	if err != nil {
		r.logErrorf(err, actionUpdateMany, "error updating %s document", r.collectionName)
		return 0, err
	}

	r.logDebugf(actionUpdateMany, "filters: %+v data: %+v", filters, update)

	res, err := r.Collection().UpdateMany(ctx, filters, update)
	if err != nil {
		r.logErrorf(err, actionUpdateMany, "error updating %s documents", r.collectionName)
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (r Repository[M, D, SF, SORD, SO, UF]) updateManyFilters(opts SF) (bson.D, error) {
	filters, err := r.BuildSearchFilters(opts)
	if err != nil {
		return nil, err
	}

	if filters == nil {
		r.logErrorf(ErrEmptyFilters, actionUpdateMany, "error updating many %s document", r.collectionName)
		return nil, ErrEmptyFilters
	}

	if len(filters) == 1 && filters[0].Key == r.deletedAtField && filters[0].Value == nil {
		r.logErrorf(ErrEmptyFilters, actionUpdateMany, "error updating many %s document", r.collectionName)
		return nil, ErrEmptyFilters
	}

	return filters, nil
}
