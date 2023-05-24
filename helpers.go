package mgorepo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository[M, D, SF, UF]) IsSearchFiltersEmpty(opts SF) bool {
	return reflect.DeepEqual(*new(SF), opts)
}

func (r Repository[M, D, SF, UF]) IsSortOptionsEmpty(opts SearchOrders) bool {
	return len(opts) == 0
}

func (r Repository[M, D, SF, UF]) IsSearchOptionsEmpty(opt SearchOptions[SF]) bool {
	return reflect.DeepEqual(SearchOptions[SF]{}, opt) || reflect.DeepEqual(NewSearchOptions(*new(SF)), opt)
}

func (r Repository[M, D, SF, UF]) IsUpdateFieldsEmpty(fields UF) bool {
	return reflect.DeepEqual(*new(UF), fields)
}

func (r Repository[M, D, SF, UF]) Now() primitive.DateTime {
	return primitive.NewDateTimeFromTime(r.clock.Now())
}
