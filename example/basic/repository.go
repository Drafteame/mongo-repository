package main

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Drafteame/mgorepo"
)

const collection = "users"

type UserRepository struct {
	mgorepo.Repository[UserModel, UserDao, UserSearchFilters, UserSearchOrders, UserSearchOptions, UserUpdateFields]
}

func NewUserRepository(db mgorepo.Driver) UserRepository {
	return UserRepository{
		Repository: mgorepo.NewRepository[
			UserModel,
			UserDao,
			UserSearchFilters,
			UserSearchOrders,
			UserSearchOptions,
			UserUpdateFields,
		](
			db,
			collection,
			[]func(UserSearchFilters) (*bson.E, error){
				buildNameFilter,
				buildLastNameFilter,
				buildGreaterThanAgeFilter,
			},
			[]func(UserSearchOrders) (*bson.E, error){
				buildNameOrder,
				buildLastNameOrder,
				buildAgeOrder,
			},
			[]func(UserUpdateFields) (*bson.E, error){
				buildNameUpdate,
				buildLastNameUpdate,
				buildAgeUpdate,
			},
		),
	}
}
