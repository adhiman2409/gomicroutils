package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgAttendanceConf struct {
	ID                               primitive.ObjectID `bson:"_id"`
	OrgId                            string             `bson:"org_id"`
	OrgName                          string             `bson:"org_name"`
	CaptureCheckInLocation           bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn           bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut          bool               `bson:"enforce_location_check_out"`
	RequiredLat                      float64            `bson:"required_lat"`
	RequiredLng                      float64            `bson:"required_lng"`
	RequiredRadiusInMeters           float64            `bson:"required_radius_in_meters"`
	GeoTrackingEnabled               bool               `bson:"geo_tracking_enabled"`
	GeoTrackingIntervalInMin         int32              `bson:"geo_tracking_interval_in_min"`
	FlexiHoursEnabled                bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin             int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime                   string             `bson:"org_check_in_time"`
	OrgCheckOutTime                  string             `bson:"org_check_out_time"`
	CheckinMarginInMin               int32              `bson:"check_in_margin_in_min"`
	WorkingDaysPerWeek               float32            `bson:"working_days_per_week"`
	WeeklyOffDays                    []string           `bson:"weekly_off_days"`
	AlternateEvenWeeklyOffDays       []string           `bson:"alternate_even_weekly_off_days"`
	DailyWorkingHours                float32            `bson:"daily_working_hours"`
	AllowedDelayedCheckInCount       int32              `bson:"allowed_delayed_check_in_count"`
	SendWarningMailOnDelayedCheckIn  bool               `bson:"send_warning_mail_on_delayed_check_in"`
	ApplyLOPAfterMaxDelayedCheckIn   bool               `bson:"apply_lop_after_max_delayed_check_in"`
	FullDayLOPAfterMaxDelayedCheckIn bool               `bson:"full_day_lop_after_max_delayed_check_in"`
	Country                          string             `bson:"country"`
	TimeZone                         string             `bson:"time_zone"`
	State                            string             `bson:"state"`
	CreatedAt                        time.Time          `bson:"created_at"`
	UpdatedAt                        time.Time          `bson:"updated_at"`
}
