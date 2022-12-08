package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectParticipant struct {
	ParticipantIdentity string `bson:"participant_identity" json:"participantIdentity"`
	ProjectIdentity     string `bson:"project_identity" json:"projectIdentity"`
}

func (ProjectParticipant) CollectionName() string {
	return "project_participant"
}

// FindAllProjectIdentityByParticipantIdentity 根据项目Id查询项目参与人
func FindAllProjectIdentityByParticipantIdentity(projectIdentity string) ([]*ProjectParticipant, error) {
	filter := bson.D{{"project_identity", projectIdentity}}
	cursor, _ := Mongo.Collection(ProjectParticipant{}.CollectionName()).Find(context.Background(), filter)
	var result []*ProjectParticipant
	err := cursor.All(context.Background(), &result)
	return result, err
}

// FindAllParticipantIdentityByProjectIdentityBy 根据参与人查询项目
func FindAllParticipantIdentityByProjectIdentity(participantIdentity string) ([]*ProjectParticipant, error) {
	filter := bson.D{{"participant_identity", participantIdentity}}
	cursor, _ := Mongo.Collection(ProjectParticipant{}.CollectionName()).Find(context.Background(), filter)
	var result []*ProjectParticipant
	err := cursor.All(context.Background(), &result)
	return result, err
}

// InsertProjectIdentityByParticipantIdentity 插入参与人的项目
func InsertProjectIdentityByParticipantIdentity(participantIdentity, projectIdentity string) error {
	doc := bson.D{{"project_identity", projectIdentity}, {"participant_identity", participantIdentity}}
	_, err := Mongo.Collection(ProjectParticipant{}.CollectionName()).InsertOne(context.Background(), doc)
	return err
}
