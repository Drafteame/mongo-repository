package builders

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/domain/options"
)

func BuildIdentifierUpdate(u options.UpdateFields) (*bson.E, error) {
	if u.Identifier == nil {
		return nil, nil
	}

	return &bson.E{Key: domain.IdentifierField, Value: *u.Identifier}, nil
}

func BuildSortableUpdate(u options.UpdateFields) (*bson.E, error) {
	if u.Sortable == nil {
		return nil, nil
	}

	return &bson.E{Key: domain.SortableField, Value: *u.Sortable}, nil
}

func BuildUpdatedAtUpdate(u options.UpdateFields) (*bson.E, error) {
	if u.UpdatedAt == nil {
		return nil, nil
	}

	updateValue := primitive.NewDateTimeFromTime(*u.UpdatedAt)

	return &bson.E{Key: domain.UpdatedAtField, Value: updateValue}, nil
}
