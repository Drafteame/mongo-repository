package builders

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/domain/options"
)

func BuildIDFilter(f options.SearchFilters) (*bson.E, error) {
	if f.ID == nil {
		return nil, nil
	}

	id, err := primitive.ObjectIDFromHex(*f.ID)
	if err != nil {
		return nil, err
	}

	return &bson.E{Key: domain.IDField, Value: id}, nil
}

func BuildIdentifierFilter(f options.SearchFilters) (*bson.E, error) {
	if f.Identifier == nil {
		return nil, nil
	}

	return &bson.E{Key: domain.IdentifierField, Value: *f.Identifier}, nil
}

func BuildSortableGraterThanFilter(f options.SearchFilters) (*bson.E, error) {
	if f.SortableGT == nil {
		return nil, nil
	}

	return &bson.E{Key: domain.SortableField, Value: bson.M{"$gt": *f.SortableGT}}, nil
}
