package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type EmailValidation struct {
	Identity string `bson:"_id" json:"identity"`
	Email    string `bson:"email" json:"email"`
	Code     string `bson:"code" json:"code"`
	Expires  int64  `bson:"expires" json:"expires"`
	SendTime int64  `bson:"send_time" json:"sendTime"`
}

func (EmailValidation) CollectionName() string {
	return "email_validation"
}

func GetEmailValidationByEmail(email string) (*EmailValidation, error) {
	ev := new(EmailValidation)
	err := Mongo.Collection(EmailValidation{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"email", email}}).Decode(ev)
	return ev, err
}

func UpdateCodeAndExpiresByEmail(code string, expires int64, email string) error {
	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"code", code}, {"expires", expires}}}}
	_, err := Mongo.Collection(EmailValidation{}.CollectionName()).UpdateOne(context.Background(), filter, update)
	return err
}

func InsertCodeAndEmailAndExpires(code string, email string, expires int64, sendTime int64) error {
	doc := bson.D{{"email", email}, {"code", code}, {"expires", expires}, {"send_time", sendTime}}
	_, err := Mongo.Collection(EmailValidation{}.CollectionName()).InsertOne(context.Background(), doc)
	return err
}
