package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, SORD, SO, UF]) HardDelete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logError(err, actionHardDelete, "error converting %s to ObjectID", id)
		return 0, nil
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
	}

	r.logDebug(actionHardDelete, "filters: %+v", filters)

	res, deleteErr := r.Collection().DeleteOne(ctx, &filters)
	if deleteErr != nil {
		r.logError(deleteErr, actionHardDelete, "error deleting %s document", r.collectionName)
		return 0, deleteErr
	}

	return res.DeletedCount, nil
}
