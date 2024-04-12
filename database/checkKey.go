package database

import (
	"context"

	"github.com/maximierung/http-openai-tts/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckKey(client *mongo.Client, db_name string, key string) bool {
	userCollection := client.Database(db_name).Collection("keys")
	var result structs.Key
	filter := bson.D{{Key: "key", Value: key}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	if err == nil {
		return true
	}
	return false
}
