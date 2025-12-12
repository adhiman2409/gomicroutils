package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceShifting struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	EmployeeID   string             `bson:"employee_id"`
	EmployeeName string             `bson:"employee_name"`
	OldDeviceID  string             `bson:"old_device_id"`
	NewDeviceID  string             `bson:"new_device_id"`
	ShiftedAt    time.Time          `bson:"shifted_at"` // ISO 8601 format
	Description  string             `bson:"description,omitempty"`
}
