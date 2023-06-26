package main

import "go.mongodb.org/mongo-driver/bson"

type UserSearchFilters struct {
	Name           *string
	LastName       *string
	GreaterThanAge *int
	IDs            []string
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

func buildIDsFilter(filters UserSearchFilters) (*bson.E, error) {
	if len(filters.IDs) == 0 {
		return nil, nil
	}

	return &bson.E{
		Key:   "_id",
		Value: bson.M{"$in": filters.IDs},
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
