package builders

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Drafteame/mgorepo/testdata/domain"
	"github.com/Drafteame/mgorepo/testdata/domain/options"
)

func BuildSortableOrder(o options.SearchOrders) (*bson.E, error) {
	if o.Sortable == 0 {
		return nil, nil
	}

	return &bson.E{Key: domain.SortableField, Value: normalizeOrder(o.Sortable)}, nil
}

func BuildCreatedAtOrder(o options.SearchOrders) (*bson.E, error) {
	if o.CreatedAt == 0 {
		return nil, nil
	}

	return &bson.E{Key: domain.CreatedAtField, Value: normalizeOrder(o.CreatedAt)}, nil
}

func normalizeOrder(order int) int {
	if order < -1 {
		return -1
	}

	if order > 1 {
		return 1
	}

	return order
}
