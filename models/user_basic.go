package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
