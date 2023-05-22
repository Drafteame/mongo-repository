package main

import "go.mongodb.org/mongo-driver/bson"

type UserUpdateFields struct {
	Name     *string
	LastName *string
	Age      *int
}

func buildNameUpdate(fields UserUpdateFields) (*bson.E, error) {
	if fields.Name == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "name",
		Value: *fields.Name,
	}, nil
}

func buildLastNameUpdate(fields UserUpdateFields) (*bson.E, error) {
	if fields.LastName == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "last_name",
		Value: *fields.LastName,
	}, nil
}

func buildAgeUpdate(fields UserUpdateFields) (*bson.E, error) {
	if fields.Age == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "age",
		Value: *fields.Age,
	}, nil
}
