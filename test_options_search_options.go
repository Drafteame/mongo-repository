package mgorepo

type searchOptions struct {
	filters searchFilters
	orders  SearchOrders
	limit   int64
	skip    int64
}

var _ SearchOptions[searchFilters] = searchOptions{}

func newSearchOptions() searchOptions {
	return searchOptions{
		filters: newSearchFilters(),
		orders:  NewSearchOrders(),
	}
}

func (o searchOptions) GetLimit() int64 {
	return o.limit
}

func (o searchOptions) GetSkip() int64 {
	return o.skip
}

func (o searchOptions) GetOrders() SearchOrders {
	return o.orders
}

func (o searchOptions) GetFilters() searchFilters {
	return o.filters
}

func (o searchOptions) WithOrder(field string, order int) searchOptions {
	o.orders = o.orders.Add(field, order)
	return o
}

func (o searchOptions) WithLimit(limit int64) searchOptions {
	o.limit = limit
	return o
}

func (o searchOptions) WithSkip(skip int64) searchOptions {
	o.skip = skip
	return o
}

func (o searchOptions) WithIDFilter(id string) searchOptions {
	o.filters = o.filters.WithID(id)
	return o
}

func (o searchOptions) WithIdentifierFilter(identifier string) searchOptions {
	o.filters = o.filters.WithIdentifier(identifier)
	return o
}
