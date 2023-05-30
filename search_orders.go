package mgorepo

const (
	OrderAsc  = 1
	OrderDesc = -1
)

type orderField struct {
	Name  string
	Order int
}

type SearchOrders []orderField

func NewSearchOrders() SearchOrders {
	return SearchOrders{}
}

func (so SearchOrders) Add(name string, order int) SearchOrders {
	return append(so, orderField{Name: name, Order: so.normalizeOrder(order)})
}

func (so SearchOrders) normalizeOrder(order int) int {
	if order < -1 {
		return -1
	}

	if order > 1 {
		return 1
	}

	return order
}
