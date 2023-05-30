package mgorepo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

func (r Repository[M, D, SF, UF]) Get(ctx context.Context, id string) (M, error) {
	var dao D
	var zeroM M

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logErrorf(err, actionGet, "invalid ObjectId %s", id)
		return zeroM, errors.Join(ErrInvalidIDFilter, ErrNotFound, err)
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
		{Key: r.deletedAtField, Value: nil},
	}

	r.logDebugf(actionGet, "filters: %+v", filters)

	res := r.Collection().FindOne(ctx, filters)

	if errRes := res.Err(); errRes != nil {
		if errors.Is(errRes, mgo.ErrNoDocuments) {
			r.logErrorf(errRes, actionGet, "%s document not found by id %s", r.collectionName, id)
			return zeroM, errors.Join(ErrNotFound, errRes)
		}

		r.logErrorf(err, actionGet, "error getting %s document by id %s", r.collectionName, id)
		return zeroM, err
	}

	if errDecode := res.Decode(&dao); errDecode != nil {
		r.logErrorf(errDecode, actionGet, "error decoding %s document by id %s", r.collectionName, id)
		return zeroM, errDecode
	}

	filler, ok := any(&dao).(DaoFiller[M])
	if !ok {
		r.logErrorf(ErrInvalidDaoFiller, actionGet, "error building model on get for %s in document id %s", r.collectionName, id)
		return zeroM, ErrInvalidDaoFiller
	}

	return filler.ToModel(), nil
}
