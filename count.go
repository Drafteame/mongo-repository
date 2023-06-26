package mgorepo

import (
	"context"
)

// Count returns the number of documents in the collection that match the given search options.
func (r Repository[M, D, SF, SORD, SO, UF]) Count(ctx context.Context, opts SF) (int64, error) {
	filters, err := r.BuildSearchFilters(opts)
	if err != nil {
		return 0, err
	}

	r.logDebugf(actionCount, "filters: %+v", filters)

	return r.Collection().CountDocuments(ctx, filters)
}
