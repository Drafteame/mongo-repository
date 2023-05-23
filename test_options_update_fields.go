package mgorepo

import (
	"time"
)

type updateFields struct {
	Identifier *string
	Sortable   *int
	UpdatedAt  *time.Time
}

func newUpdateFields() updateFields {
	return updateFields{}
}

func (u updateFields) WithUpdatedAt(updatedAt time.Time) updateFields {
	u.UpdatedAt = &updatedAt
	return u
}

func (u updateFields) WithSortable(sortable int) updateFields {
	u.Sortable = &sortable
	return u
}

func (u updateFields) WithIdentifier(identifier string) updateFields {
	u.Identifier = &identifier
	return u
}
