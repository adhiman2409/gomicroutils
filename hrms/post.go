package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id"`
	Topic       string             `bson:"topic"`
	Content     string             `bson:"content"`
	Image       string             `bson:"image"`
	Important   bool               `bson:"important"`
	PostedBy    string             `bson:"posted_by"`
	Designation string             `bson:"designation"`
	Date        string             `bson:"date"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
