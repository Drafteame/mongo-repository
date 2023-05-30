package main

import "go.mongodb.org/mongo-driver/bson"

type UserSearchFilters struct {
	Name           *string
	LastName       *string
	GreaterThanAge *int
}

func buildNameFilter(filters UserSearchFilters) (*bson.E, error) {
	if filters.Name == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "name",
		Value: *filters.Name,
	}, nil
}

func buildLastNameFilter(filters UserSearchFilters) (*bson.E, error) {
	if filters.LastName == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "last_name",
		Value: *filters.LastName,
	}, nil
}

func buildGreaterThanAgeFilter(filters UserSearchFilters) (*bson.E, error) {
	if filters.GreaterThanAge == nil {
		return nil, nil
	}

	return &bson.E{
		Key:   "age",
		Value: bson.M{"$gt": *filters.GreaterThanAge},
	}, nil
}
