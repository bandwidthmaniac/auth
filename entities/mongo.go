package entities

import "go.mongodb.org/mongo-driver/mongo"

// A MongoInstace contains the Mongo client & database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
