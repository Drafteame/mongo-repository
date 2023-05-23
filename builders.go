package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BuildSearchFilters builds filters from a SearchFilters struct. It will return bson.D with
// the filters, or an error if one occurs. If no filters are given, it will return nil.
// If the deletedAtField is not explicitly set, it will filter out deleted instances if
// withTimestamps is set to true.
func (r Repository[M, D, SF, SO, UF]) BuildSearchFilters(opts SF) (bson.D, error) {
	filters := bson.D{}

	deletedFilter := false

	for _, builder := range r.filterBuilders {
		filter, err := builder(opts)
		if err != nil {
			r.logError(err, buildSearchFilters, "error building search filter for %s", r.collectionName)
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
func (r Repository[M, D, SF, SO, UF]) BuildSearchOptions(opts SO) (bson.D, *options.FindOptions, error) {
	bsonFilters, err := r.BuildSearchFilters(opts.GetFilters())
	if err != nil {
		r.logError(err, buildSearchOptions, "error building search filters for %s", r.collectionName)
		return nil, nil, err
	}

	bsonOrders, err := opts.GetOrders().Build()
	if err != nil {
		r.logError(err, buildSearchOptions, "error building search orders for %s", r.collectionName)
		return nil, nil, err
	}

	findOpts := options.Find()

	if opts.GetLimit() > 0 {
		findOpts.SetLimit(opts.GetLimit())
	} else {
		findOpts.SetLimit(r.searchLimit)
	}

	if opts.GetSkip() > 0 {
		findOpts.SetSkip(opts.GetSkip())
	}

	if len(bsonOrders) > 0 {
		findOpts.SetSort(bsonOrders)
	}

	return bsonFilters, findOpts, nil
}

// BuildUpdateFields builds the update fields from the given options and returns a bson.D
// that can be used to update the document. If repository is configured with timestamps, it
// will automatically add the updatedAtField to the update fields.
func (r Repository[M, D, SF, SO, UF]) BuildUpdateFields(fields UF) (bson.D, error) {
	bsonFields := bson.D{}

	for _, builder := range r.updateBuilders {
		field, err := builder(fields)
		if err != nil {
			r.logError(err, buildUpdateFields, "error building update fields for %s", r.collectionName)
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
