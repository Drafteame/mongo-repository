package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository[M, D, SF, SO, UF]) DeleteMany(ctx context.Context, filters SF) (int64, error) {
	bf, err := r.deleteManyFilters(filters)
	if err != nil {
		return 0, err
	}

	data := bson.D{{Key: "$set", Value: bson.D{{Key: r.deletedAtField, Value: r.Now()}}}}

	r.logDebug(actionDeleteMany, "filters: %+v data: %+v", bf, data)

	res, err := r.Collection().UpdateMany(ctx, bf, data)
	if err != nil {
		r.logError(err, actionDeleteMany, "error deleting %s documents", r.collectionName)
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (r Repository[M, D, SF, SO, UF]) deleteManyFilters(filters SF) (bson.D, error) {
	if r.IsSearchFiltersEmpty(filters) {
		r.logError(ErrEmptyFilters, actionDeleteMany, "error deleting many %s document", r.collectionName)
		return nil, ErrEmptyFilters
	}

	bf, err := r.BuildSearchFilters(filters)
	if err != nil {
		r.logError(err, actionDeleteMany, "error deleting many %s document", r.collectionName)
		return nil, err
	}

	if bf == nil {
		r.logError(ErrEmptyFilters, actionDeleteMany, "error deleting many %s document", r.collectionName)
		return nil, ErrEmptyFilters
	}

	return bf, nil
}
