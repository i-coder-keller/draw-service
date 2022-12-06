package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
func FindAllProjectByOwnerIdentity(ownerIdentity string) (*[]ProjectBasic, error) {
	filter := bson.D{{"owner_identity", ownerIdentity}}
	var result *[]ProjectBasic
	cursor, err := Mongo.Collection(ProjectBasic{}.CollectionName()).Find(context.Background(), filter)
	err = cursor.All(context.Background(), &result)
	return result, err
}
