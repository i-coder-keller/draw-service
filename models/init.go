package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "ZHIzhi11",
	}).ApplyURI("mongodb://101.43.78.75:27017"))
	if err != nil {
		log.Println("connect MongoDB Error:", err)
		return nil
	}
	return client.Database("draw-test")
}
