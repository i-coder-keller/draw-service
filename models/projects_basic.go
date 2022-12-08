package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectBasic struct {
	Identity  string `bson:"_id" json:"identity"`
	Name      string `bson:"name" json:"name"`
	Info      string `bson:"info" json:"info"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}

func (ProjectBasic) CollectionName() string {
	return "project_basic"
}

// FindAllProjectByIdentity 根据项目Id查询项目信息
func FindAllProjectByIdentity(identity string) (*ProjectBasic, error) {
	filter := bson.D{{"_id", identity}}
	result := new(ProjectBasic)
	err := Mongo.Collection(ProjectBasic{}.CollectionName()).FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
