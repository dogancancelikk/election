package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

var (
	Database *mongo.Database
)

func init() {
	uri := fmt.Sprintf("mongodb://%s", "127.0.0.1:27017")
	tm := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(
		bsontype.EmbeddedDocument, tm).Build()
	clientOptions := options.Client().SetRegistry(reg)
	clientOptions.ApplyURI(uri)

	clientOptions.Auth = &options.Credential{
		Username:   "root",
		Password:   "rootpassword",
		AuthSource: "election"}

	var err error
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	Database = client.Database("election")
}
