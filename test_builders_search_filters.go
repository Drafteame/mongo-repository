package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func buildIDFilter(f searchFilters) (*bson.E, error) {
	if f.ID == nil {
		return nil, nil
	}

	id, err := primitive.ObjectIDFromHex(*f.ID)
	if err != nil {
		return nil, err
	}

	return &bson.E{Key: idField, Value: id}, nil
}

func buildIdentifierFilter(f searchFilters) (*bson.E, error) {
	if f.Identifier == nil {
		return nil, nil
	}

	return &bson.E{Key: identifierField, Value: *f.Identifier}, nil
}

func buildSortableGraterThanFilter(f searchFilters) (*bson.E, error) {
	if f.SortableGT == nil {
		return nil, nil
	}

	return &bson.E{Key: sortableField, Value: bson.M{"$gt": *f.SortableGT}}, nil
}
