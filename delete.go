package mgorepo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, SORD, SO, UF]) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logError(err, actionDelete, "error converting %s to ObjectID", id)
		return 0, errors.Join(ErrInvalidIDFilter, err)
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
	}

	data := bson.M{
		"$set": bson.M{
			r.deletedAtField: primitive.NewDateTimeFromTime(r.clock.Now()),
		},
	}

	r.logDebug(actionDelete, "filters: %+v data: %+v", filters, data)

	res, deleteErr := r.Collection().UpdateOne(ctx, &filters, data)
	if deleteErr != nil {
		r.logError(deleteErr, actionDelete, "error deleting %s document", r.collectionName)
		return 0, deleteErr
	}

	return res.ModifiedCount, nil
}
