package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProjectBasic struct {
	Identity  string `bson:"_id" json:"identity"`
	Name      string `bson:"name" json:"name"`
	Info      string `bson:"info" json:"info"`
	CreateAt  int64  `bson:"create_at" json:"create_at"`
	UpdatedAt int64  `bson:"update_at" json:"update_at"`
}

func (ProjectBasic) CollectionName() string {
	return "project_basic"
}

// FindAllProjectByIdentity 根据项目Id查询项目信息
func FindAllProjectByIdentity(identity string) (ProjectBasic, error) {
	objectId, _ := primitive.ObjectIDFromHex(identity)
	filter := bson.D{{"_id", objectId}}
	var result ProjectBasic
	err := Mongo.Collection(ProjectBasic{}.CollectionName()).FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// 创建新项目
func InsertProject(name, info string) (interface{}, error) {
	createAt := time.Now().UnixMilli()
	updateAt := time.Now().UnixMilli()
	doc := bson.D{{"name", name}, {"info", info}, {"create_at", createAt}, {"update_at", updateAt}}
	result, err := Mongo.Collection(ProjectBasic{}.CollectionName()).InsertOne(context.Background(), doc)
	return result.InsertedID, err
}
