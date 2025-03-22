package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Token      string             `bson:"token"`
	DeviceId   string             `bson:"device_id"`
	DeviceType string             `bson:"device_type"`
	EmployeeId string             `bson:"employee_id"`
}
