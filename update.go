package mgorepo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, UF]) Update(ctx context.Context, id string, fields UF) (int64, error) {
	filters, errFilters := r.updateFilters(id)
	if errFilters != nil {
		r.logErrorf(errFilters, actionUpdate, "error updating %s document", r.collectionName)
		return 0, nil
	}

	data, errData := r.updateData(fields, false)
	if errData != nil {
		r.logErrorf(errData, actionUpdate, "error updating %s document", r.collectionName)
		return 0, errData
	}

	r.logDebugf(actionUpdate, "filters: %+v data: %+v", filters, data)

	res, updateErr := r.Collection().UpdateOne(ctx, &filters, data)
	if updateErr != nil {
		r.logErrorf(updateErr, actionUpdate, "error updating %s document", r.collectionName)
		return 0, updateErr
	}

	return res.ModifiedCount, nil
}

func (r Repository[M, D, SF, UF]) updateFilters(id string) (bson.D, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Join(ErrInvalidIDFilter, err)
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
	}

	return filters, nil
}

func (r Repository[M, D, SF, UF]) updateData(fields UF, allowEmpty bool) (bson.D, error) {
	if !allowEmpty && r.IsUpdateFieldsEmpty(fields) {
		return nil, ErrEmptyUpdate
	}

	data, err := r.BuildUpdateFields(fields)
	if err != nil {
		return nil, err
	}

	return bson.D{{Key: "$set", Value: data}}, nil
}
