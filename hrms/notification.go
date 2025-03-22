package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId   string             `bson:"deviceId"`
	DeviceType string             `bson:"deviceType"`
	EmployeeId string             `bson:"employeeId"`
}
