package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	JobId   string             `bson:"job_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}
