package main

import "github.com/Drafteame/mgorepo"

type SearchOptions struct {
	mgorepo.SearchOptions[UserSearchFilters, SearchOrders]
}

var _ mgorepo.SearchOptioner[UserSearchFilters, SearchOrders] = SearchOptions{}

func NewSearchOptions(filters UserSearchFilters, orders SearchOrders) SearchOptions {
	return SearchOptions{
		SearchOptions: mgorepo.NewSearchOptions[UserSearchFilters, SearchOrders](filters, orders),
	}
}

func (so SearchOptions) WithLimit(limit int64) SearchOptions {
	so.SearchOptions = so.SearchOptions.WithLimit(limit)
	return so
}

func (so SearchOptions) WithSkip(skip int64) SearchOptions {
	so.SearchOptions = so.SearchOptions.WithSkip(skip)
	return so
}

func (so SearchOptions) WithProject(field string, project int) SearchOptions {
	so.SearchOptions = so.SearchOptions.WithProject(field, project)
	return so
}

func (so SearchOptions) WithProjectFields(projects map[string]int) SearchOptions {
	so.SearchOptions = so.SearchOptions.WithProjectFields(projects)
	return so
}
