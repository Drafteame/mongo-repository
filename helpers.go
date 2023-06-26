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

func (r Repository[M, D, SF, SO, UF]) IsSearchOptionsEmpty(opt SearchOptions[SF, SO]) bool {
	return reflect.DeepEqual(SearchOptions[SF, SO]{}, opt) || reflect.DeepEqual(NewSearchOptions(*new(SF), *new(SO)), opt)
}

func (r Repository[M, D, SF, SO, UF]) IsUpdateFieldsEmpty(fields UF) bool {
	return reflect.DeepEqual(*new(UF), fields)
}

func (r Repository[M, D, SF, SO, UF]) Now() primitive.DateTime {
	return primitive.NewDateTimeFromTime(r.clock.Now())
}
