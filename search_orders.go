package mgorepo

import "go.mongodb.org/mongo-driver/bson"

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
	return append(so, orderField{Name: name, Order: NormalizeOrder(order)})
}

func (so SearchOrders) Build() (bson.D, error) {
	if len(so) == 0 {
		return bson.D{{Key: "_id", Value: 1}}, nil
	}

	var orders bson.D

	for _, field := range so {
		if field.Order == 0 {
			continue
		}

		orders = append(orders, bson.E{Key: field.Name, Value: field.Order})
	}

	return orders, nil
}
