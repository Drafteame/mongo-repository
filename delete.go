package mgorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes a document by id. If the repository is configured to not use timestamps, this operation will be
// applied as a hard delete.
func (r Repository[M, D, SF, UF]) Delete(ctx context.Context, id string) (int64, error) {
	if !r.withTimestamps {
		return r.HardDelete(ctx, id)
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logErrorf(err, actionDelete, "error converting %s to ObjectID", id)
		return 0, nil
	}

	filters := bson.D{
		{Key: "_id", Value: oid},
	}

	now := r.Now()

	data := bson.M{
		"$set": bson.M{
			r.updatedAtField: now,
			r.deletedAtField: now,
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
