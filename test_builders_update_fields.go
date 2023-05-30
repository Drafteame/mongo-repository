package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func buildIdentifierUpdate(u updateFields) (*bson.E, error) {
	if u.Identifier == nil {
		return nil, nil
	}

	return &bson.E{Key: identifierField, Value: *u.Identifier}, nil
}

func buildSortableUpdate(u updateFields) (*bson.E, error) {
	if u.Sortable == nil {
		return nil, nil
	}

	return &bson.E{Key: sortableField, Value: *u.Sortable}, nil
}

func buildUpdatedAtUpdate(u updateFields) (*bson.E, error) {
	if u.UpdatedAt == nil {
		return nil, nil
	}

	updateValue := primitive.NewDateTimeFromTime(*u.UpdatedAt)

	return &bson.E{Key: updatedAtField, Value: updateValue}, nil
}
