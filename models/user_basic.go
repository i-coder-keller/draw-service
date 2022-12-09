package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity  string `bson:"_id" json:"identity"`
	Account   string `bson:"account" json:"account"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Avatar    string `bson:"avatar" json:"avatar"`
	Email     string `bson:"email" json:"email"`
	CreatedAt int64  `bson:"created_at" json:"createdAt"`
	UpdatedAt int64  `bson:"updated_at" json:"updatedAt"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountAndPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByAccountAndEmail(account, email string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"email", email}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByEmail(email string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

// GetUserBasicByIdentity 根据用户Id查询用户信息
func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	objectId, _ := primitive.ObjectIDFromHex(identity)
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", objectId}}).
		Decode(ub)
	return ub, err
}

func InsertUserBasic(account, email, nickname, password, avatar string, createdAt, updatedAt int64) error {
	doc := bson.D{{"account", account}, {"email", email}, {"nickname", nickname}, {"avatar", avatar}, {"created_at", createdAt}, {"updated_at", updatedAt}, {"password", password}}
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).InsertOne(context.Background(), doc)
	return err
}
