package options

import (
	"time"
)

type UpdateFields struct {
	Identifier *string
	Sortable   *int
	UpdatedAt  *time.Time
}

func NewUpdateFields() UpdateFields {
	return UpdateFields{}
}

func (u UpdateFields) WithUpdatedAt(updatedAt time.Time) UpdateFields {
	u.UpdatedAt = &updatedAt
	return u
}

func (u UpdateFields) WithSortable(sortable int) UpdateFields {
	u.Sortable = &sortable
	return u
}

func (u UpdateFields) WithIdentifier(identifier string) UpdateFields {
	u.Identifier = &identifier
	return u
}
