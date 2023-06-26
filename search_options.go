package mgorepo

const (
	FieldAdd    = 1
	FieldRemove = 0
)

type SearchOptions[SF SearchFilters, SO SearchOrderer] struct {
	filters    SF
	orders     SO
	limit      int64
	skip       int64
	projection map[string]int
}

var _ SearchOptioner[SearchFilters, SearchOrderer] = SearchOptions[SearchFilters, SearchOrderer]{}

func NewSearchOptions[SF SearchFilters, SO SearchOrderer](filters SF, orders SO) SearchOptions[SF, SO] {
	return SearchOptions[SF, SO]{
		filters: filters,
		orders:  orders,
		limit:   DefaultSearchLimit,
	}
}

func (so SearchOptions[SF, SO]) Filters() SF {
	return so.filters
}

func (so SearchOptions[SF, SO]) Orders() SO {
	return so.orders
}

func (so SearchOptions[SF, SO]) Limit() int64 {
	return so.limit
}

func (so SearchOptions[SF, SO]) Skip() int64 {
	return so.skip
}

func (so SearchOptions[SF, SO]) Projection() map[string]int {
	return so.projection
}

func (so SearchOptions[SF, SO]) WithLimit(limit int64) SearchOptions[SF, SO] {
	so.limit = limit
	return so
}

func (so SearchOptions[SF, SO]) WithSkip(skip int64) SearchOptions[SF, SO] {
	so.skip = skip
	return so
}

func (so SearchOptions[SF, SO]) WithProject(field string, project int) SearchOptions[SF, SO] {
	if field == "" {
		return so
	}

	if so.projection == nil {
		so.projection = make(map[string]int)
	}

	so.projection[field] = so.normalizeProjection(project)

	return so
}

func (so SearchOptions[SF, SO]) WithProjectFields(project map[string]int) SearchOptions[SF, SO] {
	if project == nil {
		return so
	}

	for field, val := range project {
		so = so.WithProject(field, val)
	}

	return so
}

func (so SearchOptions[SF, SO]) normalizeProjection(val int) int {
	if val <= 0 {
		return 0
	}

	return 1
}
