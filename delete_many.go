package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteMany deletes many documents from the collection. It returns the number of deleted documents and an error.
// If the repository has timestamps enabled, it will soft delete the documents. Otherwise, it will hard delete them.
func (r Repository[M, D, SF, UF]) DeleteMany(ctx context.Context, filters SF) (int64, error) {
	if !r.withTimestamps {
		return r.HardDeleteMany(ctx, filters)
	}

	bf, err := r.deleteManyFilters(filters)
	if err != nil {
		return 0, err
	}

	data := bson.D{{Key: "$set", Value: bson.D{{Key: r.deletedAtField, Value: r.Now()}}}}

	r.logDebugf(actionDeleteMany, "filters: %+v data: %+v", bf, data)

	res, err := r.Collection().UpdateMany(ctx, bf, data)
	if err != nil {
		r.logErrorf(err, actionDeleteMany, "error deleting %s documents", r.collectionName)
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (r Repository[M, D, SF, UF]) deleteManyFilters(filters SF) (bson.D, error) {
	if r.IsSearchFiltersEmpty(filters) {
		r.logErrorf(ErrEmptyFilters, actionDeleteMany, "error deleting many %s document", r.collectionName)
		return nil, ErrEmptyFilters
	}

	bf, err := r.BuildSearchFilters(filters)
	if err != nil {
		r.logErrorf(err, actionDeleteMany, "error deleting many %s document", r.collectionName)
		return nil, err
	}

	return bf, nil
}
