package mgorepo

import "context"

func (r Repository[M, D, SF, SO, UF]) HardDeleteMany(ctx context.Context, filters SF) (int64, error) {
	if r.IsSearchFiltersEmpty(filters) {
		r.logErrorf(ErrEmptyFilters, actionHardDeleteMany, "empty search filters for %s hard delete", r.collectionName)
		return 0, ErrEmptyFilters
	}

	bf, err := r.BuildSearchFilters(filters)
	if err != nil {
		r.logErrorf(err, actionHardDeleteMany, "error building search filters for %s hard delete", r.collectionName)
		return 0, err
	}

	r.logDebugf(actionHardDeleteMany, "filters: %+v", bf)

	res, deleteErr := r.Collection().DeleteMany(ctx, bf)
	if deleteErr != nil {
		r.logErrorf(deleteErr, actionHardDeleteMany, "error deleting %s documents", r.collectionName)
		return 0, deleteErr
	}

	return res.DeletedCount, nil
}
