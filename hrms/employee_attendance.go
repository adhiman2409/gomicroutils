package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GeoTracking struct {
	Lat                        float64   `bson:"lat"`
	Lng                        float64   `bson:"lng"`
	Time                       time.Time `bson:"time"`
	Location                   string    `bson:"location"`
	IsStationary               bool      `bson:"is_stationary"`
	StationaryDurationInMins   int64     `bson:"stationary_duration_in_mins"`
	TransitDurationInMins      int64     `bson:"transit_duration_in_mins"`
	TotalDistanceTraveledInKMs float64   `bson:"total_distance_traveled_in_kms"`
	EventName                  string    `bson:"event_name"` // e.g., "checkin", "checkout", "location_update"
	EventId                    string    `bson:"event_id"`   // Unique identifier for the event
	Remarks                    string    `bson:"remarks"`    // Additional information about the event
}

type EmployeeAttendance struct {
	ID                            primitive.ObjectID `bson:"_id"`
	EID                           string             `bson:"eid"`
	Domain                        string             `bson:"domain"`
	Department                    string             `bson:"department"`
	Day                           string             `bson:"day"`
	Month                         string             `bson:"month"`
	Year                          string             `bson:"year"`
	CaptureCheckInLocation        bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn        bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut       bool               `bson:"enforce_location_check_out"`
	RequiredLat                   float64            `bson:"required_lat"`
	RequiredLng                   float64            `bson:"required_lng"`
	RequiredRadiusInMeters        float64            `bson:"required_radius_in_meters"`
	GeoTrackingEnabled            bool               `bson:"geo_tracking_enabled"`
	GeoTrackingIntervalInMin      int32              `bson:"geo_tracking_interval_in_min"`
	CheckInLat                    float64            `bson:"check_in_lat"`
	CheckInLng                    float64            `bson:"check_in_lng"`
	CheckOutLat                   float64            `bson:"check_out_lat"`
	CheckOutLng                   float64            `bson:"check_out_lng"`
	FlexiHoursEnabled             bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin          int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime                string             `bson:"org_check_in_time"`
	OrgCheckOutTime               string             `bson:"org_check_out_time"`
	CheckinMarginInMin            int32              `bson:"check_in_margin_in_min"`
	IsCheckedIn                   bool               `bson:"is_checked_in"`
	CheckInTime                   string             `bson:"check_in_time"`
	UTCCheckInTime                time.Time          `bson:"utc_check_in_time"`
	IsCheckedOut                  bool               `bson:"is_checked_out"`
	CheckOutTime                  string             `bson:"check_in_out"`
	UTCCheckOutTime               time.Time          `bson:"utc_check_out_time"`
	EnforceAttendanceSource       bool               `bson:"enforce_attendance_source"`
	AttendanceSource              string             `bson:"attendance_source"`
	CheckInSource                 string             `bson:"check_in_source"`
	CheckOutSource                string             `bson:"check_out_source"`
	TotalWorkingHours             float32            `bson:"total_working_hours"`
	YesterdayNightWorkingHours    float32            `bson:"yesterday_night_working_hours"`
	TodayNightWorkingHours        float32            `bson:"today_night_working_hours"`
	TodayMorningWorkingHours      float32            `bson:"today_morning_working_hours"`
	IsDelayedCheckIn              bool               `bson:"is_delayed_check_in"`
	IsEarlyCheckOut               bool               `bson:"is_early_check_out"`
	IsOnLeave                     bool               `bson:"is_on_leave"`
	IsHalfDayLeave                bool               `bson:"is_half_day_leave"`
	IsWorkingDay                  bool               `bson:"is_working_day"`
	IsHoliday                     bool               `bson:"is_holiday"`
	IsWeeklyOffDay                bool               `bson:"is_weekly_off_day"`
	IsFullDayLOP                  bool               `bson:"is_full_day_lop"`
	IsHalfDayLOP                  bool               `bson:"is_half_day_lop"`
	ShiftType                     ShiftType          `bson:"shift_type"`
	SentCheckInMail               bool               `bson:"sent_check_in_mail"`
	SentCheckOutMail              bool               `bson:"sent_check_out_mail"`
	OrgCheckInTimeInTicks         int64              `bson:"org_check_in_time_in_ticks"`
	OrgCheckOutTimeInTicks        int64              `bson:"org_check_out_time_in_ticks"`
	DailyCheckInStats             []DailyCheckInStat `bson:"daily_checkin_stats"`
	TotalDistanceTraveledInKMs    float64            `bson:"total_distance_traveled_in_kms"`
	TotalStationaryDurationInMins int64              `bson:"total_stationary_duration_in_mins"`
	TotalTransitDurationInMins    int64              `bson:"total_transit_duration_in_mins"`
	GeoTracking                   []GeoTracking      `bson:"geo_tracking"`
	IsOnBreak                     bool               `bson:"is_on_break"`
	TotalBreakTimeInMins          int64              `bson:"total_break_time_in_mins"`
	DailyBreakStats               []DailyBreakStat   `bson:"daily_break_stats"`
	IsLocked                      bool               `bson:"is_locked"`
	IsExpired                     bool               `bson:"is_expired"`
	IsRegularized                 bool               `bson:"is_regularized"`
	RegularizedOn                 string             `bson:"regularized_on"`
	RegularizedBy                 string             `bson:"regularized_by"`
	Remarks                       string             `bson:"remarks"`
	Country                       string             `bson:"country"`
	TimeZone                      string             `bson:"time_zone"`
	CreatedAt                     int64              `bson:"createdAt"`
}

type DailyAttendanceObject struct {
	EID         string `bson:"eid"`
	IsProcessed bool   `bson:"is_processed"`
}

type DailyCheckInStat struct {
	CheckInTime    string  `bson:"check_in_time"`
	CheckOutTime   string  `bson:"check_out_time"`
	CheckInSource  string  `bson:"check_in_source"`
	CheckOutSource string  `bson:"check_out_source"`
	CheckInLat     float64 `bson:"check_in_lat"`
	CheckInLng     float64 `bson:"check_in_lng"`
	CheckOutLat    float64 `bson:"check_out_lat"`
	CheckOutLng    float64 `bson:"check_out_lng"`
	WorkingHours   float32 `bson:"working_hours"`
}

type DailyBreakStat struct {
	StartTime      time.Time `bson:"start_time"`
	EndTime        time.Time `bson:"end_time"`
	Remarks        string    `bson:"remarks"`
	DurationInMins int64     `bson:"duration_in_mins"`
}
