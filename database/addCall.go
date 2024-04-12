package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddCall(client *mongo.Client, db_name string, key string) {
	keyCollection := client.Database(db_name).Collection("keys")
	filter := bson.D{{Key: "key", Value: key}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "calls", Value: 1}}}}
	_, err := keyCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error while incrementing calls for key: " + key)
	}
}
