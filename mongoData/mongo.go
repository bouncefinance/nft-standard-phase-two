package mongoData

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nft_standard/config"
	"nft_standard/logger"
	"time"
)

var MongoDB *MongoDatabase

func init() {
	MongoDB = &MongoDatabase{
		MongoClient: SetConnect(),
	}
}

type MongoDatabase struct {
	MongoClient *mongo.Client
}

func (m *MongoDatabase) GetCollection(database string, collection string) *mongo.Collection {
	return m.MongoClient.Database(database).Collection(collection)
}

func (m *MongoDatabase) FindOne(database string, collection_ string, filter bson.D, object interface{}) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	err := collection.FindOne(context.Background(), filter).Decode(object)
	for err != nil && err != mongo.ErrNoDocuments && time < config.RETRY_TIME {
		err = collection.FindOne(context.Background(), filter).Decode(object)
		time++
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.Wrap(err, "mongo FindOne error")
	} else if err == mongo.ErrNoDocuments {
		return err
	}
	return nil
}

func (m *MongoDatabase) InsertOne(database string, collection_ string, data bson.M) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	_, err := collection.InsertOne(context.Background(), data)
	for err != nil && time < config.RETRY_TIME {
		_, err = collection.InsertOne(context.Background(), data)
		time++
	}
	if err != nil {
		return errors.Wrap(err, "mongo InsertOne error")
	}
	return nil
}

func (m *MongoDatabase) UpdateOne(database string, collection_ string, filter bson.D, data bson.D) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	_, err := collection.UpdateOne(context.Background(), filter, data)
	for err != nil && time < config.RETRY_TIME {
		_, err = collection.UpdateOne(context.Background(), filter, data)
		time++
	}
	if err != nil {
		return errors.Wrap(err, "mongo UpdateOne error")
	}
	return nil
}

func SetConnect() *mongo.Client {
	uri := config.MongoURL
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20))
	if err != nil {
		panic(err)
	}
	logger.Logger.Info().Str("mongodb link", uri).Msg("连接上mongodb")
	return client
}
