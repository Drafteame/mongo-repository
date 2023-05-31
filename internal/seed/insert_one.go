package seed

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(t *testing.T, md *mgo.Database, collection string, data any) {
	res, err := md.Collection(collection).InsertOne(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	newID := res.InsertedID.(primitive.ObjectID)

	t.Cleanup(func() {
		filter := bson.D{{Key: "_id", Value: newID}}

		_, errDel := md.Collection(collection).DeleteOne(context.Background(), filter)
		if errDel != nil {
			t.Fatal(errDel)
		}
	})
}
