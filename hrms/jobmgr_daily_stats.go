package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecruiterDailyStats struct {
	ID            primitive.ObjectID `bson:"_id"`
	RecruiterId   string             `bson:"recruiter_id"`
	Name          string             `bson:"name"`
	Day           int                `bson:"day"`
	Weekday       int                `bson:"weekday"`
	Month         int                `bson:"month"`
	Year          int                `bson:"year"`
	HH            int                `bson:"hour"`
	MM            int                `bson:"minute"`
	ApplicationId string             `bson:"application_id"`
	JobId         string             `bson:"job_id"`
	StartStatus   JobAppStatus       `bson:"start_status"`
	EndStatus     JobAppStatus       `bson:"end_status"`
}
