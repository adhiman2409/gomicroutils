package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeAttendance struct {
	ID                      primitive.ObjectID `bson:"_id"`
	EID                     string             `bson:"eid"`
	Domain                  string             `bson:"domain"`
	Department              string             `bson:"department"`
	Day                     string             `bson:"day"`
	Month                   string             `bson:"month"`
	Year                    string             `bson:"year"`
	CaptureCheckInLocation  bool               `bson:"capture_check_in_location"`
	EnforceLocationCheckIn  bool               `bson:"enforce_location_check_in"`
	EnforceLocationCheckOut bool               `bson:"enforce_location_check_out"`
	RequiredLat             float64            `bson:"required_lat"`
	RequiredLng             float64            `bson:"required_lng"`
	CheckInLat              float64            `bson:"check_in_lat"`
	CheckInLng              float64            `bson:"check_in_lng"`
	CheckOutLat             float64            `bson:"check_out_lat"`
	CheckOutLng             float64            `bson:"check_out_lng"`
	FlexiHoursEnabled       bool               `bson:"flexi_hours_enabled"`
	FlexiHourWindowInMin    int32              `bson:"flexi_hour_window_in_min"`
	OrgCheckInTime          string             `bson:"org_check_in_time"`
	OrgCheckOutTime         string             `bson:"org_check_out_time"`
	CheckinMarginInMin      int32              `bson:"check_in_margin_in_min"`
	IsCheckedIn             bool               `bson:"is_checked_in"`
	CheckInTime             string             `bson:"check_in_time"`
	IsCheckedOut            bool               `bson:"is_checked_out"`
	CheckOutTime            string             `bson:"check_in_out"`
	WorkingHours            float32            `bson:"working_hours"`
	IsDelayedCheckIn        bool               `bson:"is_delayed_check_in"`
	IsEarlyCheckOut         bool               `bson:"is_early_check_out"`
	IsOnLeave               bool               `bson:"is_on_leave"`
	IsWorkingDay            bool               `bson:"is_working_day"`
	IsWeeklyOffDay          bool               `bson:"is_weekly_off_day"`
	Remarks                 string             `bson:"remarks"`
}

type ProcessedEmployees struct {
	EID       string `bson:"eid"`
	Processed bool   `bson:"processed"`
}
