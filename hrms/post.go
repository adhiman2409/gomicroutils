package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `bson:"_id"`
	Topic       string             `bson:"topic"`
	Content     string             `bson:"content"`
	Image       string             `bson:"image"`
	PostedBy    string             `bson:"posted_by"`
	Designation string             `bson:"designation"`
	Date        string             `bson:"date"`
}
