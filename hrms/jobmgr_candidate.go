package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CandidateInfo struct {
	ID            primitive.ObjectID `bson:"_id"`
	EID           string             `bson:"eid"`
	Domain        string             `bson:"domain"`
	OrgName       string             `bson:"org_name"`
	EmailId       string             `bson:"email_id"`
	PhoneNumber   string             `bson:"phone_number"`
	Name          string             `bson:"name"`
	Department    string             `bson:"department"`
	EmailVerified bool               `bson:"email_verified"`
	Role          string             `bson:"role"`
	Status        string             `bson:"status"`
	CreatedBy     string             `bson:"created_by"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
