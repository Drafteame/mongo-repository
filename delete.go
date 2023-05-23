package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, SO, UF]) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logErrorf(err, actionDelete, "error converting %s to ObjectID", id)
		return 0, nil
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
	}

	data := bson.M{
		"$set": bson.M{
			r.deletedAtField: primitive.NewDateTimeFromTime(r.clock.Now()),
		},
	}

	r.logDebugf(actionDelete, "filters: %+v data: %+v", filters, data)

	res, deleteErr := r.Collection().UpdateOne(ctx, &filters, data)
	if deleteErr != nil {
		r.logErrorf(deleteErr, actionDelete, "error deleting %s document", r.collectionName)
		return 0, deleteErr
	}

	return res.ModifiedCount, nil
}
