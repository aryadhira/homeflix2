package repository

import (
	"context"
	"homeflix2/db"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepo struct{}

func (repo *MongoRepo) Insert(collname string, data bson.M) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	_, err = db.Collection(collname).InsertOne(ctx, data)
	if err != nil {
		return err
	}

	defer db.Client().Disconnect(ctx)

	return nil
}

func (repo *MongoRepo) InsertMany(collname string, data []interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	_, err = db.Collection(collname).InsertMany(ctx, data)
	if err != nil {
		return err
	}

	defer db.Client().Disconnect(ctx)

	return nil

}

func (repo *MongoRepo) Update(collname string, filter interface{}, data interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	_, err = db.Collection(collname).UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	defer db.Client().Disconnect(ctx)

	return nil
}

func (repo *MongoRepo) UpdateMany(collname string, filter interface{}, data interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	_, err = db.Collection(collname).UpdateMany(ctx, filter, data)
	if err != nil {
		return err
	}

	defer db.Client().Disconnect(ctx)

	return nil
}

func (repo *MongoRepo) Find(collname string, filter interface{}, data interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	csr, err := db.Collection(collname).Find(ctx, filter)
	if err != nil {
		return err
	}

	defer csr.Close(ctx)

	err = csr.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoRepo) Aggregate(collname string, pipeline interface{}, data interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	csr, err := db.Collection(collname).Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	defer csr.Close(ctx)

	err = csr.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoRepo) DeleteMany(collname string, filter interface{}) error {
	ctx := context.Background()
	db, err := new(db.Mongo).Init()
	if err != nil {
		return err
	}

	_, err = db.Collection(collname).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
