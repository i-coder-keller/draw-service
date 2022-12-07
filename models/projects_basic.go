package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type ProjectBasic struct {
	Identity      string   `bson:"_id" json:"identity"`
	OwnerIdentity string   `bson:"owner_identity" json:"owner_identity"`
	Name          string   `bson:"name" json:"name"`
	Info          string   `bson:"info" json:"info"`
	CreatedAt     int64    `bson:"created_at" json:"created_at"`
	UpdatedAt     int64    `bson:"updated_at" json:"updated_at"`
	Participant   []string `bson:"participant" json:"participant"`
}

func (ProjectBasic) CollectionName() string {
	return "project_basic"
}

func FindAllProjectByOwnerIdentity(ownerIdentity string) ([]*ProjectBasic, error) {
	log.Println(ownerIdentity)
	filter := bson.D{{"owner_identity", "项目所有人唯一标识"}}
	var result []*ProjectBasic
	cursor, _ := Mongo.Collection(ProjectBasic{}.CollectionName()).Find(context.Background(), filter)
	err := cursor.All(context.Background(), &result)
	return result, err
}
