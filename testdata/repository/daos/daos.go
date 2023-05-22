package daos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/testdata/domain"
)

type TestDAO struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	Sortable   int                 `bson:"sortable"`
	Identifier string              `bson:"identifier"`
	CreatedAt  primitive.DateTime  `bson:"createdAt"`
	UpdatedAt  primitive.DateTime  `bson:"updatedAt"`
	DeletedAt  primitive.DateTime  `bson:"deletedAt,omitempty"`
}

// var _ mgorepo.DaoFiller[domain.TestModel] = (*TestDAO)(nil)

func (d *TestDAO) ToModel() domain.TestModel {
	return domain.TestModel{
		ID:         d.ID.Hex(),
		Identifier: d.Identifier,
		Sortable:   d.Sortable,
		CreatedAt:  d.CreatedAt.Time().UTC(),
		UpdatedAt:  d.UpdatedAt.Time().UTC(),
		DeletedAt:  d.DeletedAt.Time().UTC(),
	}
}

func (d *TestDAO) FromModel(m domain.TestModel) error {
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
