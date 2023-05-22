package main

import "go.mongodb.org/mongo-driver/bson"

type UserSearchOrders struct {
	Name     int
	LastName int
	Age      int
}

func buildNameOrder(orders UserSearchOrders) (*bson.E, error) {
	if orders.Name == 0 {
		return nil, nil
	}

	return &bson.E{
		Key:   "name",
		Value: orders.Name,
	}, nil
}

func buildLastNameOrder(orders UserSearchOrders) (*bson.E, error) {
	if orders.LastName == 0 {
		return nil, nil
	}

	return &bson.E{
		Key:   "last_name",
		Value: orders.LastName,
	}, nil
}

func buildAgeOrder(orders UserSearchOrders) (*bson.E, error) {
	if orders.Age == 0 {
		return nil, nil
	}

	return &bson.E{
		Key:   "age",
		Value: orders.Age,
	}, nil
}
