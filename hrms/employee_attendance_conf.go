package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DelayNotification struct {
	EmailSent   bool   `bson:"email_sent"`
	Day         string `bson:"day"`
	CheckInTime string `bson:"check_in_time"`
	DelayInMins int32  `bson:"delay_in_mins"`
}

type DelayedCheckInStats struct {
	DelayedCheckInDays  []DelayNotification `bson:"days"`
	Month               string              `bson:"month"`
	Year                string              `bson:"year"`
	DelayedCheckinCount int32               `bson:"delayed_check_in_count"`
}

type EmpAttendanceConf struct {
	ID                               primitive.ObjectID    `bson:"_id"`
	EID                              string                `bson:"eid"`
	Department                       string                `bson:"department"`
	ApplyOrgDefaultRules             bool                  `bson:"apply_org_default_rules"`
	CaptureCheckInLocation           bool                  `bson:"capture_check_in_location"`
	EnforceLocationCheckIn           bool                  `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut          bool                  `bson:"enforce_location_check_out"`
	EnforceAttendanceSource          bool                  `bson:"enforce_attendance_source"`
	AttendanceSource                 string                `bson:"attendance_source"`
	RequiredLat                      float64               `bson:"required_lat"`
	RequiredLng                      float64               `bson:"required_lng"`
	RequiredRadiusInMeters           float64               `bson:"required_radius_in_meters"`
	GeoTrackingEnabled               bool                  `bson:"geo_tracking_enabled"`
	GeoTrackingIntervalInMin         int32                 `bson:"geo_tracking_interval_in_min"`
	FlexiHoursEnabled                bool                  `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin             int32                 `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime                   string                `bson:"org_check_in_time"`
	OrgCheckOutTime                  string                `bson:"org_check_out_time"`
	CheckinMarginInMin               int32                 `bson:"check_in_margin_in_min"`
	WorkingDaysPerWeek               float32               `bson:"working_days_per_week"`
	WeeklyOffDays                    []string              `bson:"weekly_off_days"`
	DailyWorkingHours                float32               `bson:"daily_working_hours"`
	EmploymentStatus                 string                `bson:"employment_status"`
	GeoLocationScreenShot            string                `bson:"geo_location_screen_shot"`
	ShiftTimings                     []EmpShiftTimings     `bson:"shift_timings"`
	DelayedCheckInStats              []DelayedCheckInStats `bson:"delayed_check_in_stats"`
	AllowedDelayedCheckInCount       int32                 `bson:"allowed_delayed_check_in_count"`
	SendWarningMailOnDelayedCheckIn  bool                  `bson:"send_warning_mail_on_delayed_check_in"`
	ApplyLOPAfterMaxDelayedCheckIn   bool                  `bson:"apply_lop_after_max_delayed_check_in"`
	FullDayLOPAfterMaxDelayedCheckIn bool                  `bson:"full_day_lop_after_max_delayed_check_in"`
	NextShiftTimingsIndex            int32                 `bson:"next_shift_timings_index"`
	Country                          string                `bson:"country"`
	TimeZone                         string                `bson:"time_zone"`
	CreatedAt                        time.Time             `bson:"created_at"`
	UpdatedAt                        time.Time             `bson:"updated_at"`
}

type EmpShiftTimings struct {
	Day               string    `bson:"day"`
	Month             string    `bson:"month"`
	Year              string    `bson:"year"`
	Weekday           string    `bson:"weekday"`
	ShiftType         ShiftType `bson:"shift_type"`
	IsOverlappedShift bool      `bson:"is_overlapped_shift"`
	OrgCheckInTime    string    `bson:"org_check_in_time"`
	OrgCheckOutTime   string    `bson:"org_check_out_time"`
	IsWorkingDay      bool      `bson:"is_working_day"`
	IsWeeklyOffDay    bool      `bson:"is_weekly_off_day"`
	IsHoliday         bool      `bson:"is_holiday"`
	Color             string    `bson:"color"`
	TextColor         string    `bson:"text_color"`
}
