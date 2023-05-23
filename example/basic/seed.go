package main

import (
	"context"
	"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(collection string, docs int, db *mongo.Database) error {
	for i := 0; i < docs; i++ {
		dao := UserDao{
			Name:     fmt.Sprintf("name_%d", i),
			LastName: fmt.Sprintf("last_name_%d", i),
			Age:      randomAge(),
		}

		if _, err := db.Collection(collection).InsertOne(context.Background(), dao); err != nil {
			return err
		}
	}

	return nil
}

func randomAge() int {
	return 18 + rand.Intn(100-18)
}
