package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PollOption struct {
	Option    string `bson:"option"`
	OptionId  string `bson:"option_id"`
	VoteCount int    `bson:"vote_count"`
}

type Vote struct {
	EmployeeId  string    `bson:"employee_id"`
	Option      string    `bson:"option"`
	OptionId    string    `bson:"option_id"`
	DateAndTime time.Time `bson:"date_and_time"`
}

type Poll struct {
	ID                primitive.ObjectID `bson:"_id"`
	Title             string             `bson:"title"`
	Content           string             `bson:"content"`
	Options           []PollOption       `bson:"options"`
	Votes             []Vote             `bson:"vote"`
	VoteCount         int                `bson:"vote_count"`
	StartDateAndTime  time.Time          `bson:"start_date_and_time"`
	EndDateAndTime    time.Time          `bson:"end_date_and_time"`
	PostedBy          string             `bson:"posted_by"`
	PostedById        string             `bson:"posted_by_id"`
	IsPublished       bool               `bson:"is_published"`
	IsOpen            bool               `bson:"is_open"`
	IsClosed          bool               `bson:"is_closed"`
	ListDateAndTime   time.Time          `bson:"list_date_and_time"`
	UnlistDateAndTime time.Time          `bson:"unlist_date_and_time"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	Remarks           string             `bson:"remarks"`
}
