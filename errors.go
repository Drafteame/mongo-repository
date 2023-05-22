package mgorepo

import "errors"

var (
	ErrNotFound         = errors.New("mongo-repository: model not found")
	ErrInvalidIDFilter  = errors.New("mongo-repository: invalid id filter")
	ErrEmptyFilters     = errors.New("mongo-repository: empty filters")
	ErrCreatingDAO      = errors.New("mongo-repository: error creating dao")
	ErrCreatingModel    = errors.New("mongo-repository: error creating model")
	ErrEmptyUpdate      = errors.New("mongo-repository: empty update")
	ErrInvalidDaoFiller = errors.New("mongo-repository: invalid dao does not implement DaoFiller")
)
