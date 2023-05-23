package mgorepo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, SO, UF]) IsSearchFiltersEmpty(opts SF) bool {
	return reflect.DeepEqual(*new(SF), opts)
}

func (r Repository[M, D, SF, SO, UF]) IsSortOptionsEmpty(opts SearchOrders) bool {
	return len(opts) == 0
}

func (r Repository[M, D, SF, SO, UF]) IsSearchOptionsEmpty(opt SO) bool {
	return reflect.DeepEqual(*new(SO), opt)
}

func (r Repository[M, D, SF, SO, UF]) IsUpdateFieldsEmpty(fields UF) bool {
	return reflect.DeepEqual(*new(UF), fields)
}

func (r Repository[M, D, SF, SO, UF]) Now() primitive.DateTime {
	return primitive.NewDateTimeFromTime(r.clock.Now())
}

func NormalizeOrder(order int) int {
	if order < -1 {
		return -1
	}

	if order > 1 {
		return 1
	}

	return order
}
