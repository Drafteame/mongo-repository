package driver

import "errors"

var (
	ErrEmptyClient = errors.New("mgorepo-driver: client is nil")
)
