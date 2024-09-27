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

type PolicyViewStats struct {
	ID         primitive.ObjectID `bson:"_id"`
	PolicyId   string             `bson:"policy_id"`
	PolicyName string             `bson:"policy_name"`
	EmployeeId string             `bson:"employee_id"`
	ViewedAt   time.Time          `bson:"viewed_at"`
}
