package mgorepo

const (
	FieldAdd    = 1
	FieldRemove = 0
)

type SearchOptions[SF SearchFilters] struct {
	filters    SF
	orders     SearchOrders
	limit      int64
	skip       int64
	projection map[string]int
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

func (so SearchOptions[SF]) Projection() map[string]int {
	return so.projection
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

func (so SearchOptions[SF]) Project(field string, project int) SearchOptions[SF] {
	if field == "" {
		return so
	}

	if so.projection == nil {
		so.projection = make(map[string]int)
	}

	so.projection[field] = so.normalizeProjection(project)

	return so
}

func (so SearchOptions[SF]) ProjectFields(project map[string]int) SearchOptions[SF] {
	if project == nil {
		return so
	}

	for field, val := range project {
		so = so.Project(field, val)
	}

	return so
}

func (so SearchOptions[SF]) normalizeProjection(val int) int {
	if val <= 0 {
		return 0
	}

	return 1
}
