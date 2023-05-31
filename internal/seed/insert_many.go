package seed

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

func InsertMany(t *testing.T, md *mgo.Database, collection string, data ...any) {
	res, err := md.Collection(collection).InsertMany(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	newIDs := res.InsertedIDs

	t.Cleanup(func() {
		filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: newIDs}}}}

		_, errDel := md.Collection(collection).DeleteMany(context.Background(), filter)
		if errDel != nil {
			t.Fatal(errDel)
		}
	})
}
