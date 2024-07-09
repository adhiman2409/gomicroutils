package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Policy struct {
	ID             primitive.ObjectID `bson:"_id"`
	PolicyName     string             `bson:"policy_name"`
	PolicyDocument string             `bson:"policy_document"`
	UploadedAt     time.Time          `bson:"uploaded_at"`
	Views          []string           `bson:"views"`
}
