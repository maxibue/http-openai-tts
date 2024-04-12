package database

import (
	"context"

	"github.com/maximierung/http-openai-tts/structs"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddKey(client *mongo.Client, db_name string, key_name string, key string) bool {
	keyCollection := client.Database(db_name).Collection("keys")
	newKey := structs.Key{Name: key_name, Key: key, Calls: 0}
	_, err := keyCollection.InsertOne(context.TODO(), newKey)
	if err != nil {
		return false
	}
	return true
}
