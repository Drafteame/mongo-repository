package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BuildSearchFilters builds filters from a SearchFilters struct. It will return bson.D with
// the filters, or an error if one occurs. If no filters are given, it will return nil.
// If the deletedAtField is not explicitly set, it will filter out deleted instances if
// withTimestamps is set to true.
func (r Repository[M, D, SF, SORD, SO, UF]) BuildSearchFilters(opts SF) (bson.D, error) {
	filters := bson.D{}

	deletedFilter := false

	for _, builder := range r.filterBuilders {
		filter, err := builder(opts)
		if err != nil {
			r.logErrorf(err, buildSearchFilters, "error building search filter for %s", r.collectionName)
			return nil, err
		}

		if filter != nil {
			if filter.Key == r.deletedAtField {
				deletedFilter = true
			}

			filters = append(filters, *filter)
		}
	}

	// Filter out deleted instances if the deletedAtField is not explicitly set
	if !deletedFilter && r.withTimestamps {
		filters = append(filters, bson.E{Key: r.deletedAtField, Value: nil})
	}

	return filters, nil
}

// BuildSearchOptions builds filters, and mongo.FindOptions from a SearchOptions struct.
// If no limit is given, it will default to the repository's search limit. If no orders
// are given, it will default to ascending order by id.
func (r Repository[M, D, SF, SORD, SO, UF]) BuildSearchOptions(opts SO) (bson.D, *options.FindOptions, error) {
	bsonFilters, err := r.BuildSearchFilters(opts.Filters())
	if err != nil {
		r.logErrorf(err, buildSearchOptions, "error building search filters for %s", r.collectionName)
		return nil, nil, err
	}

	findOpts, err := r.buildFindOptions(opts)
	if err != nil {
		r.logErrorf(err, buildSearchOptions, "error building search options for %s", r.collectionName)
		return nil, nil, err
	}

	return bsonFilters, findOpts, nil
}

func (r Repository[M, D, SF, SORD, SO, UF]) buildProjection(opts SO) bson.M {
	projection := bson.M{}

	for field, val := range opts.Projection() {
		projection[field] = val
	}

	return projection
}

func (r Repository[M, D, SF, SORD, SO, UF]) buildFindOptions(opts SO) (*options.FindOptions, error) {
	bsonOrders, err := r.BuildSearchOrders(opts.Orders())
	if err != nil {
		return nil, err
	}

	findOpts := options.Find()

	if opts.Limit() > 0 {
		findOpts.SetLimit(opts.Limit())
	} else {
		findOpts.SetLimit(int64(r.searchLimit))
	}

	if opts.Skip() > 0 {
		findOpts.SetSkip(opts.Skip())
	}

	if len(bsonOrders) > 0 {
		findOpts.SetSort(bsonOrders)
	}

	if projection := r.buildProjection(opts); len(projection) > 0 {
		findOpts.SetProjection(projection)
	}

	return findOpts, nil
}

// BuildUpdateFields builds the update fields from the given options and returns a bson.D
// that can be used to update the document. If repository is configured with timestamps, it
// will automatically add the updatedAtField to the update fields.
func (r Repository[M, D, SF, SORD, SO, UF]) BuildUpdateFields(fields UF) (bson.D, error) {
	bsonFields := bson.D{}

	for _, builder := range r.updateBuilders {
		field, err := builder(fields)
		if err != nil {
			r.logErrorf(err, buildUpdateFields, "error building update fields for %s", r.collectionName)
			return nil, err
		}

		if field != nil {
			bsonFields = append(bsonFields, *field)
		}
	}

	if r.withTimestamps {
		bsonFields = append(bsonFields, bson.E{Key: r.updatedAtField, Value: r.Now()})
	}

	return bsonFields, nil
}

func (r Repository[M, D, SF, SORD, SO, UF]) BuildSearchOrders(so SearchOrderer) (bson.D, error) {
	ordersMap := so.ToMap()

	if len(ordersMap) == 0 {
		return nil, nil
	}

	var orders bson.D

	for field, order := range ordersMap {
		if order == 0 {
			continue
		}

		orders = append(orders, bson.E{Key: field, Value: order})
	}

	return orders, nil
}
