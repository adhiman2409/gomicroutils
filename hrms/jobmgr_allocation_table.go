package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecruiterInfo struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type AllocationEntry struct {
	ID         primitive.ObjectID `bson:"_id"`
	NextIndex  int64              `bson:"index"`
	Recruiters []RecruiterInfo    `bson:"recruiters"`
}
