package options

type SearchFilters struct {
	ID         *string
	Identifier *string
	SortableGT *int
}

// var _ repository.SearchFilters = SearchFilters{}

func NewSearchFilters() SearchFilters {
	return SearchFilters{}
}

func (f SearchFilters) WithID(id string) SearchFilters {
	f.ID = &id
	return f
}

func (f SearchFilters) WithIdentifier(identifier string) SearchFilters {
	f.Identifier = &identifier
	return f
}

func (f SearchFilters) WithSortableGreaterThan(sortable int) SearchFilters {
	f.SortableGT = &sortable
	return f
}
