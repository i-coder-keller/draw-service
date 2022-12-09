package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectOwner struct {
	OwnerIdentity   string `bson:"owner_identity" json:"ownerIdentity"`
	ProjectIdentity string `bson:"project_identity" json:"projectIdentity"`
}

func (ProjectOwner) CollectionName() string {
	return "project_owner"
}

// FindAllProjectIdentityByOwnerIdentity 查询拥有人的项目列表
func FindAllProjectIdentityByOwnerIdentity(ownerIdentity string) ([]*ProjectOwner, error) {
	filter := bson.D{{"owner_identity", ownerIdentity}}
	cursor, _ := Mongo.Collection(ProjectOwner{}.CollectionName()).Find(context.Background(), filter)
	var result []*ProjectOwner
	err := cursor.All(context.Background(), &result)
	return result, err
}

// InsertProjectIdentityByOwnerIdentity 插入拥有人的项目
func InsertProjectIdentityByOwnerIdentity(ownerIdentity, projectIdentity string) error {
	doc := bson.D{{"project_identity", projectIdentity}, {"owner_identity", ownerIdentity}}
	_, err := Mongo.Collection(ProjectOwner{}.CollectionName()).InsertOne(context.Background(), doc)
	return err
}

// 查询项目是否是当前人所拥有
func ValidationProjectOfOwner(ownerIdentity, projectIdentity string) (ProjectOwner, error) {
	filter := bson.D{{"owner_identity", ownerIdentity}, {"project_identity", projectIdentity}}
	var result ProjectOwner
	err := Mongo.Collection(ProjectOwner{}.CollectionName()).FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
