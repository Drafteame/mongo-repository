package mgorepo

import (
	"time"
)

type testModel struct {
	ID         string
	Identifier string
	Sortable   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
