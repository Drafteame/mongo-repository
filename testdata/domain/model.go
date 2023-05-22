package domain

import (
	"time"
)

type TestModel struct {
	ID         string
	Identifier string
	Sortable   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// var _ repository.ModelSetter = (*TestModel)(nil)
