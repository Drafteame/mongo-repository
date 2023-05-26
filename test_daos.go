package mgorepo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type testDAO struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	Sortable   int                 `bson:"sortable,omitempty"`
	Identifier string              `bson:"identifier,omitempty"`
	CreatedAt  primitive.DateTime  `bson:"createdAt,omitempty"`
	UpdatedAt  primitive.DateTime  `bson:"updatedAt,omitempty"`
	DeletedAt  primitive.DateTime  `bson:"deletedAt,omitempty"`
}

var _ DaoFiller[testModel] = (*testDAO)(nil)

func (d *testDAO) ToModel() testModel {
	var id string

	if d.ID != nil {
		id = d.ID.Hex()
	}

	return testModel{
		ID:         id,
		Identifier: d.Identifier,
		Sortable:   d.Sortable,
		CreatedAt:  d.CreatedAt.Time().UTC(),
		UpdatedAt:  d.UpdatedAt.Time().UTC(),
		DeletedAt:  d.DeletedAt.Time().UTC(),
	}
}

func (d *testDAO) FromModel(m testModel) error {
	if m.ID != "" {
		id, err := primitive.ObjectIDFromHex(m.ID)
		if err != nil {
			return err
		}

		d.ID = &id
	}

	d.Identifier = m.Identifier
	d.Sortable = m.Sortable
	d.CreatedAt = primitive.NewDateTimeFromTime(m.CreatedAt)
	d.UpdatedAt = primitive.NewDateTimeFromTime(m.UpdatedAt)
	d.DeletedAt = primitive.NewDateTimeFromTime(m.DeletedAt)

	return nil
}
