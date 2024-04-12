package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RemoveKey(client *mongo.Client, db_name string, key_name string) bool {
	keyCollection := client.Database(db_name).Collection("keys")
	filter := bson.D{{Key: "name", Value: key_name}}
	_, err := keyCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true
}
