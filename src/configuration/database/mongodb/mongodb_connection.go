package mongodb

import (
	"context"
	"os"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	MONGODB_URL   = "MONGODB_URL"
	MONGO_USER_DB = "MONGO_USER_DB"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {

	journey := zap.String("journey", "MongoDB")
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGO_USER_DB)
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}

	logger.Info("Init MongoDB", journey)
	client, err := mongo.Connect(ctx, options.Client().SetBSONOptions(bsonOpts).ApplyURI(mongodb_uri))
	if err != nil {
		logger.Error("Error trying to connect to MongoDB", err, journey)
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying testing connection to MongoDB", err, journey)
		return nil, err
	}
	logger.Info("MongoDB Connected sucessfully", journey)

	return client.Database(mongodb_database), nil
}
