package grpcclient

type OrgLeaveConfiguration struct {
	OrgName                     string
	Title                       string
	IsFixed                     bool    `bson:"is_fixed"`
	Frequency                   string  `bson:"frequency"`
	LeaveCount                  float32 `bson:"leave_count"`
	LeaveUnit                   string  `bson:"leave_unit"`
	MaxLeaveCount               float32 `bson:"max_leave_count"`
	IsCarryForwardAllowed       bool    `bson:"is_carry_forward_allowed"`
	MaxCarryForwardCount        float32 `bson:"max_carry_forward_count"`
	BulkLeaveCount              int     `bson:"bulk_leave_count"`
	BulkLeaveNoticeInDays       int     `bson:"bulk_leave_notice_in_days"`
	IsEncashmentAllowed         bool    `bson:"encashment_allowed"`
	ApplicableAfterWorkingDays  int     `bson:"applicable_after_working_days"`
	DocumentRequired            bool    `bson:"document_required"`
	WeeklyOffAndHolidayIncluded bool    `bson:"weekly_off_and_holiday_included"`
	IsActive                    bool    `bson:"is_active"`
	Description                 string  `bson:"description"`
	YearStartDate               string  `bson:"year_start_date"`
	YearEndDate                 string  `bson:"year_end_date"`
	IncludeWeeklyOffDays        bool    `bson:"include_weekly_off_days"`
	IncludeHolidays             bool    `bson:"include_holidays"`
	MaxLeaveAllowedCount        float32 `bson:"max_leave_allowed_count"`
}

type OrgLeaveObj struct {
	OrgName                string
	OrgLeaveConfigurations []OrgLeaveConfiguration
}
