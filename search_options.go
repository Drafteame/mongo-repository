package mgorepo

type SearchOptions[SF SearchFilters] struct {
	filters SF
	orders  SearchOrders
	limit   int64
	skip    int64
}

func NewSearchOptions[SF SearchFilters](filters SF) SearchOptions[SF] {
	return SearchOptions[SF]{
		filters: filters,
		orders:  NewSearchOrders(),
		limit:   DefaultSearchLimit,
	}
}

func (so SearchOptions[SF]) Filters() SF {
	return so.filters
}

func (so SearchOptions[SF]) Orders() SearchOrders {
	return so.orders
}

func (so SearchOptions[SF]) Limit() int64 {
	return so.limit
}

func (so SearchOptions[SF]) Skip() int64 {
	return so.skip
}

func (so SearchOptions[SF]) WithOrder(field string, order int) SearchOptions[SF] {
	so.orders = so.orders.Add(field, order)
	return so
}

func (so SearchOptions[SF]) WithLimit(limit int64) SearchOptions[SF] {
	so.limit = limit
	return so
}

func (so SearchOptions[SF]) WithSkip(skip int64) SearchOptions[SF] {
	so.skip = skip
	return so
}
