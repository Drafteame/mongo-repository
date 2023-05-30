package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo"
)

type UserDao struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	LastName string             `bson:"last_name"`
	Age      int                `bson:"age"`
}

var _ mgorepo.DaoFiller[UserModel] = (*UserDao)(nil)

func (d *UserDao) ToModel() UserModel {
	return UserModel{
		ID:       d.ID.Hex(),
		Name:     d.Name,
		LastName: d.LastName,
		Age:      d.Age,
	}
}

func (d *UserDao) FromModel(m UserModel) error {
	d.ID = primitive.NewObjectID()
	d.Name = m.Name
	d.LastName = m.LastName
	d.Age = m.Age
	return nil
}
