package main

import "github.com/Drafteame/mgorepo"

type SearchOrders struct {
	mgorepo.SearchOrders
}

func NewSearchOrders() SearchOrders {
	return SearchOrders{
		SearchOrders: mgorepo.NewSearchOrders(),
	}
}

func (so SearchOrders) Add(name string, order int) SearchOrders {
	so.SearchOrders = so.SearchOrders.Add(name, order)
	return so
}

func (so SearchOrders) ToMap() map[string]int {
	return so.SearchOrders.ToMap()
}
