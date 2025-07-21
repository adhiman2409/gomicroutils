package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DepAttendanceConf struct {
	ID                       primitive.ObjectID `bson:"_id"`
	DepId                    string             `bson:"dep_id"`
	Department               string             `bson:"department"`
	CaptureCheckInLocation   bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn   bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut  bool               `bson:"enforce_location_check_out"`
	RequiredLat              float64            `bson:"required_lat"`
	RequiredLng              float64            `bson:"required_lng"`
	RequiredRadiusInMeters   float64            `bson:"required_radius_in_meters"`
	GeoTrackingEnabled       bool               `bson:"geo_tracking_enabled"`
	GeoTrackingIntervalInMin int32              `bson:"geo_tracking_interval_in_min"`
	FlexiHoursEnabled        bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin     int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime           string             `bson:"org_check_in_time"`
	OrgCheckOutTime          string             `bson:"org_check_out_time"`
	CheckinMarginInMin       int32              `bson:"check_in_margin_in_min"`
	WorkingDaysPerWeek       float32            `bson:"working_days_per_week"`
	WeeklyOffDays            []string           `bson:"weekly_off_days"`
	DailyWorkingHours        float32            `bson:"daily_working_hours"`
	CreatedAt                time.Time          `bson:"created_at"`
	UpdatedAt                time.Time          `bson:"updated_at"`
}
