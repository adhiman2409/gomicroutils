package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceConf struct {
	ID                      primitive.ObjectID `bson:"_id"`
	OrgId                   string             `bson:"org_id"`
	OrgName                 string             `bson:"org_name"`
	CaptureCheckInLocation  bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn  bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut bool               `bson:"enforce_location_check_out"`
	RequiredLat             float64            `bson:"required_lat"`
	RequiredLng             float64            `bson:"required_lng"`
	FlexiHoursEnabled       bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin    int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime          string             `bson:"org_check_in_time"`
	OrgCheckOutTime         string             `bson:"org_check_out_time"`
	CheckinMarginInMin      int32              `bson:"check_in_margin_in_min"`
	OfficeHours             float32            `bson:"office_hours"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}
