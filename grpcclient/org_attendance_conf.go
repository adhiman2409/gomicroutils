package grpcclient

type OrgAttendanceConf struct {
	OrgId                   string
	OrgName                 string
	CaptureCheckInLocation  bool
	EnforceLocationCheckIn  bool
	EnforceLocationCheckOut bool
	RequiredLat             float64
	RequiredLng             float64
	FlexiHoursEnabled       bool
	FlexiHourWindowInMin    int32
	OrgCheckInTime          string
	OrgCheckOutTime         string
	CheckinMarginInMin      int32
	DailyWorkingHours       float32
	WorkingDaysPerWeek      float32
	WeeklyOffDays           []string
}
