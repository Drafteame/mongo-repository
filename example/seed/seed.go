package seed

import (
	"context"
	"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedCollection(collection string, docs int, db *mongo.Database) error {
	for i := 0; i < docs; i++ {
		_, err := db.Collection(collection).InsertOne(context.Background(), bson.M{
			"name":      fmt.Sprintf("name_%d", i),
			"last_name": fmt.Sprintf("last_name_%d", i),
			"age":       randomAge(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func randomAge() int {
	return 18 + rand.Intn(100-18)
}
