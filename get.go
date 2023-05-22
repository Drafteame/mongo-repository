package mgorepo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

func (r Repository[M, D, SF, SORD, SO, UF]) Get(ctx context.Context, id string) (M, error) {
	var dao D
	var zeroM M

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logError(err, actionGet, "invalid ObjectId %s", id)
		return zeroM, errors.Join(ErrInvalidIDFilter, ErrNotFound, err)
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
		{Key: r.deletedAtField, Value: nil},
	}

	r.logDebug(actionGet, "filters: %+v", filters)

	res := r.Collection().FindOne(ctx, filters)

	if errRes := res.Err(); errRes != nil {
		if errors.Is(errRes, mgo.ErrNoDocuments) {
			r.logError(errRes, actionGet, "%s document not found by id %s", r.collectionName, id)
			return zeroM, errors.Join(ErrNotFound, errRes)
		}

		r.logError(err, actionGet, "error getting %s document by id %s", r.collectionName, id)
		return zeroM, err
	}

	if errDecode := res.Decode(&dao); errDecode != nil {
		r.logError(errDecode, actionGet, "error decoding %s document by id %s", r.collectionName, id)
		return zeroM, errDecode
	}

	filler, ok := any(&dao).(DaoFiller[M])
	if !ok {
		r.logError(ErrInvalidDaoFiller, actionGet, "error building model on get for %s in document id %s", r.collectionName, id)
		return zeroM, ErrInvalidDaoFiller
	}

	return filler.ToModel(), nil
}
