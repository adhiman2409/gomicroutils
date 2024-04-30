package grpcclient

type DepAttendanceConf struct {
	DepId                   string
	Department              string
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
	OfficeHours             float32
}
