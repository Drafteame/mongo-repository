package options

type SearchOptions struct {
	filters SearchFilters
	orders  SearchOrders
	limit   int64
	skip    int64
}

// var _ repository.SearchOptions[SearchOrders, SearchFilters] = SearchOptions{}

func NewSearchOptions() SearchOptions {
	return SearchOptions{
		filters: NewSearchFilters(),
	}
}

func (o SearchOptions) GetLimit() int64 {
	return o.limit
}

func (o SearchOptions) GetSkip() int64 {
	return o.skip
}

func (o SearchOptions) GetOrders() SearchOrders {
	return o.orders
}

func (o SearchOptions) GetFilters() SearchFilters {
	return o.filters
}

func (o SearchOptions) WithIDFilter(id string) SearchOptions {
	o.filters = o.filters.WithID(id)
	return o
}

func (o SearchOptions) WithIdentifierFilter(identifier string) SearchOptions {
	o.filters = o.filters.WithIdentifier(identifier)
	return o
}

func (o SearchOptions) WithLimit(limit int64) SearchOptions {
	o.limit = limit
	return o
}

func (o SearchOptions) WithSkip(skip int64) SearchOptions {
	o.skip = skip
	return o
}
