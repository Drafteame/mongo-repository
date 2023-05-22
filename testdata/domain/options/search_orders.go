package options

type SearchOrders struct {
	Sortable  int
	CreatedAt int
}

func (o SearchOptions) WithSortSorter(sort int) SearchOptions {
	o.orders.Sortable = sort
	return o
}

func (o SearchOptions) WithSortCreatedAt(sort int) SearchOptions {
	o.orders.CreatedAt = sort
	return o
}
