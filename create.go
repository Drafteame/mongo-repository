package mgorepo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Create creates a new record on the collection based on the model.
func (r Repository[M, D, SF, SORD, SO, UF]) Create(ctx context.Context, model M) (M, error) {
	var zeroM M

	data, err := r.createBuildData(model)
	if err != nil {
		return zeroM, err
	}

	r.logDebugf(actionCreate, "data: %+v", data)

	res, err := r.Collection().InsertOne(ctx, data)
	if err != nil {
		r.logErrorf(err, actionCreate, "error inserting %s DAO", r.collectionName)
		return zeroM, err
	}

	data["_id"] = res.InsertedID

	r.logDebugf(actionCreate, "insertedId: %+v", res.InsertedID)

	resModel, err := r.createBuildModel(data)

	return resModel, err
}

func (r Repository[M, D, SF, SORD, SO, UF]) createBuildModel(data bson.M) (M, error) {
	var zeroM M

	dao := new(D)

	bytes, err := bson.Marshal(data)
	if err != nil {
		r.log.Errorf(err, actionCreate, "error creating %s model", r.collectionName)
		return zeroM, errors.Join(ErrCreatingModel, err)
	}

	if err := bson.Unmarshal(bytes, dao); err != nil {
		r.log.Errorf(err, actionCreate, "error creating %s model", r.collectionName)
		return zeroM, errors.Join(ErrCreatingModel, err)
	}

	filler := any(dao).(DaoFiller[M])

	return filler.ToModel(), nil
}

func (r Repository[M, D, SF, SORD, SO, UF]) createBuildData(model M) (bson.M, error) {
	dao := any(new(D)).(DaoFiller[M])

	if errDao := dao.FromModel(model); errDao != nil {
		r.logErrorf(errDao, actionCreate, "error filling %s DAO", r.collectionName)
		return bson.M{}, errors.Join(ErrCreatingDAO, errDao)
	}

	bsonData := bson.M{}

	bsonBytes, err := bson.Marshal(dao)
	if err != nil {
		r.log.Errorf(err, actionCreate, "error filling %s DAO", r.collectionName)
		return bson.M{}, errors.Join(ErrCreatingDAO, err)
	}

	if err := bson.Unmarshal(bsonBytes, &bsonData); err != nil {
		r.log.Errorf(err, actionCreate, "error filling %s DAO", r.collectionName)
		return bson.M{}, errors.Join(ErrCreatingDAO, err)
	}

	if r.withTimestamps {
		now := r.Now()

		bsonData["createdAt"] = now
		bsonData["updatedAt"] = now
	}

	return bsonData, nil
}
