package main

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Drafteame/mgorepo"
)

const collection = "users"

type UserRepository struct {
	mgorepo.Repository[UserModel, UserDao, UserSearchFilters, SearchOrders, SearchOptions, UserUpdateFields]
}

func NewUserRepository(db mgorepo.Driver) UserRepository {
	return UserRepository{
		Repository: mgorepo.NewRepository[
			UserModel,
			UserDao,
			UserSearchFilters,
			SearchOrders,
			SearchOptions,
			UserUpdateFields,
		](
			db,
			collection,
			[]func(UserSearchFilters) (*bson.E, error){
				buildNameFilter,
				buildLastNameFilter,
				buildGreaterThanAgeFilter,
				buildIDsFilter,
			},
			[]func(UserUpdateFields) (*bson.E, error){
				buildNameUpdate,
				buildLastNameUpdate,
				buildAgeUpdate,
			},
		),
	}
}
