package mgorepo

type searchFilters struct {
	ID         *string
	Identifier *string
	SortableGT *int
}

func newSearchFilters() searchFilters {
	return searchFilters{}
}

func (f searchFilters) WithID(id string) searchFilters {
	f.ID = &id
	return f
}

func (f searchFilters) WithIdentifier(identifier string) searchFilters {
	f.Identifier = &identifier
	return f
}

func (f searchFilters) WithSortableGreaterThan(sortable int) searchFilters {
	f.SortableGT = &sortable
	return f
}
