package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type CandidateIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}
