package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MonthlyTimeSheet struct {
	ID                          primitive.ObjectID `bson:"_id"`
	Month                       int                `bson:"month"`
	Year                        int                `bson:"year"`
	EmployeeID                  string             `bson:"employee_id"`
	EmployeeName                string             `bson:"employee_name"`
	ManagerID                   string             `bson:"manager_id"`
	ManagerName                 string             `bson:"manager_name"`
	MonthlyWorkingDays          float32            `bson:"monthly_working_days"`
	RequiredMonthlyWorkingHours float32            `bson:"required_monthly_working_hours"`
	ClaimedMonthlyWorkingHours  float32            `bson:"claimed_monthly_working_hours"`
	ApprovedMonthlyWorkingHours float32            `bson:"approved_monthly_working_hours"`
	DailyTimesheet              []TimeSheet        `bson:"daily_timesheet"`
}

type ProjectTask struct {
	TaskID                   string    `bson:"task_id"`
	TaskName                 string    `bson:"task_name"`
	TaskWorkingHours         float64   `bson:"task_working_hours"`
	AppliedById              string    `bson:"applied_by_id"`
	AppliedByName            string    `bson:"applied_by_name"`
	AppliedOn                time.Time `bson:"applied_on"`
	EmployeeRemarks          string    `bson:"employee_remarks"`
	IsApproved               bool      `bson:"is_approved"`
	ApprovedTaskWorkingHours float64   `bson:"approved_task_working_hours"`
	ApproverRemarks          string    `bson:"approver_remarks"`
	ApproverId               string    `bson:"approver_id"`
	ApproverName             string    `bson:"approver_name"`
	ApprovedOn               time.Time `bson:"approved_on"`
}

type TimeSheetProject struct {
	ProjectID                        string        `bson:"project_id"`
	ProjectName                      string        `bson:"project_name"`
	ProjectHeadId                    string        `bson:"project_head_id"`
	ProjectHeadName                  string        `bson:"project_head_name"`
	ManagerID                        string        `bson:"manager_id"`
	ManagerName                      string        `bson:"manager_name"`
	TotalClaimedProjectWorkingHours  float64       `bson:"total_claimed_project_working_hours"`
	TotalApprovedProjectWorkingHours float64       `bson:"total_approved_project_working_hours"`
	Tasks                            []ProjectTask `bson:"tasks"`
}

type TimeSheet struct {
	Day                  int                `bson:"day"`
	Date                 string             `bson:"date"`
	FirstCheckInTime     time.Time          `bson:"first_check_in_time"`
	LastCheckOutTime     time.Time          `bson:"last_check_out_time"`
	RequiredWorkingHours float64            `bson:"required_working_hours"`
	TotalWorkingHours    float64            `bson:"total_working_hours"`
	ClaimedWorkingHours  float64            `bson:"claimed_working_hours"`
	ApprovedWorkingHours float64            `bson:"approved_working_hours"`
	IsWorkingDay         bool               `bson:"is_working_day"`
	IsHoliday            bool               `bson:"is_holiday"`
	IsOnLeave            bool               `bson:"is_on_leave"`
	IsWeeklyOffDay       bool               `bson:"is_weekly_off_day"`
	Entries              []TimeSheetEntry   `bson:"entries"`
	Projects             []TimeSheetProject `bson:"projects"`
}

type TimeSheetEntry struct {
	ProjectId string  `bson:"project_id"`
	Hours     float64 `bson:"hours"`
	Remark    string  `bson:"remark"`
}
