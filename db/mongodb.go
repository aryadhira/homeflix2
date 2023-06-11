package db

import (
	"context"
	"fmt"
	"homeflix2/helper"
	"homeflix2/settings"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type Mongo struct{}

func (m *Mongo) Init() (*mongo.Database, error) {
	err := new(settings.Setting).GetConfig()
	helper.ErrorHandler(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DATABASE")
	mongourl := fmt.Sprintf("mongodb://%v:%v", host, port)

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongourl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(db), nil

}
