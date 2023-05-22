package main

import "github.com/Drafteame/mgorepo"

type UserSearchOptions struct {
	Filters UserSearchFilters
	Orders  UserSearchOrders
	Limit   int64
	Skip    int64
}

var _ mgorepo.SearchOptions[UserSearchOrders, UserSearchFilters] = UserSearchOptions{}

func (o UserSearchOptions) GetFilters() UserSearchFilters {
	return o.Filters
}

func (o UserSearchOptions) GetOrders() UserSearchOrders {
	return o.Orders
}

func (o UserSearchOptions) GetLimit() int64 {
	return o.Limit
}

func (o UserSearchOptions) GetSkip() int64 {
	return o.Skip
}
