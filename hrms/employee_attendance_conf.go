package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpAttendanceConf struct {
	ID                      primitive.ObjectID `bson:"_id"`
	EID                     string             `bson:"eid"`
	Department              string             `bson:"department"`
	ApplyOrgDefaultRules    bool               `bson:"apply_org_default_rules"`
	CaptureCheckInLocation  bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn  bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut bool               `bson:"enforce_location_check_out"`
	EnforceAttendanceSource bool               `bson:"enforce_attendance_source"`
	AttendanceSource        string             `bson:"attendance_source"`
	RequiredLat             float64            `bson:"required_lat"`
	RequiredLng             float64            `bson:"required_lng"`
	FlexiHoursEnabled       bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin    int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime          string             `bson:"org_check_in_time"`
	OrgCheckOutTime         string             `bson:"org_check_out_time"`
	CheckinMarginInMin      int32              `bson:"check_in_margin_in_min"`
	WorkingDaysPerWeek      float32            `bson:"working_days_per_week"`
	WeeklyOffDays           []string           `bson:"weekly_off_days"`
	DailyWorkingHours       float32            `bson:"daily_working_hours"`
	EmploymentStatus        string             `bson:"employment_status"`
	ShiftsEnabled           bool               `bson:"shifts_enabled"`
	CurrentShiftTimings     []EmpShiftTimings  `bson:"current_shift_timings"`
	NextMonthShiftTimings   []EmpShiftTimings  `bson:"next_month_shift_timings"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}

type EmpShiftTimings struct {
	Day                       int    `bson:"day"`
	Month                     int    `bson:"month"`
	Year                      int    `bson:"year"`
	IsOverlappedShift         bool   `bson:"is_overlapped_shift"`
	OrgCheckInTimeFirstHalf   string `bson:"org_check_in_time_first_half"`
	OrgCheckOutTimeFirstHalf  string `bson:"org_check_out_time_first_half"`
	OrgCheckInTimeSecondHalf  string `bson:"org_check_in_time_second_half"`
	OrgCheckOutTimeSecondHalf string `bson:"org_check_out_time_second_half"`
	IsWorkingDay              bool   `bson:"is_working_day"`
	IsWeeklyOffDay            bool   `bson:"is_weekly_off_day"`
	IsHoliday                 bool   `bson:"is_holiday"`
}
