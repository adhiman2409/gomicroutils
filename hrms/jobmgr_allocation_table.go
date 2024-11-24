package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobApplicationErrorRecord struct {
	ID           primitive.ObjectID `bson:"_id"`
	Date         string             `bson:"date"`
	MessageId    string             `bson:"message_id"`
	Subject      string             `bson:"subject"`
	From         string             `bson:"from"`
	ErrorMessage string             `bson:"error_meesage"`
}

type RecruiterInfo struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type AllocationEntry struct {
	ID         primitive.ObjectID `bson:"_id"`
	NextIndex  int64              `bson:"index"`
	Recruiters []RecruiterInfo    `bson:"recruiters"`
}
