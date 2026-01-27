package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MonthlyTimeSheet struct {
	ID                          primitive.ObjectID `bson:"_id" json:"_id"`
	Month                       int                `bson:"month" json:"month"`
	Year                        int                `bson:"year" json:"year"`
	EmployeeID                  string             `bson:"employee_id" json:"employee_id"`
	EmployeeName                string             `bson:"employee_name" json:"employee_name"`
	ManagerID                   string             `bson:"manager_id" json:"manager_id"`
	ManagerName                 string             `bson:"manager_name" json:"manager_name"`
	MonthlyWorkingDays          float32            `bson:"monthly_working_days" json:"monthly_working_days"`
	RequiredMonthlyWorkingHours float32            `bson:"required_monthly_working_hours" json:"required_monthly_working_hours"`
	ClaimedMonthlyWorkingHours  float32            `bson:"claimed_monthly_working_hours" json:"claimed_monthly_working_hours"`
	ApprovedMonthlyWorkingHours float32            `bson:"approved_monthly_working_hours" json:"approved_monthly_working_hours"`
	DailyTimesheet              []TimeSheet        `bson:"daily_timesheet" json:"daily_timesheet"`
}

type ProjectTask struct {
	TaskID                   string    `bson:"task_id" json:"task_id"`
	TaskName                 string    `bson:"task_name" json:"task_name"`
	TaskWorkingHours         float64   `bson:"task_working_hours" json:"task_working_hours"`
	AppliedById              string    `bson:"applied_by_id" json:"applied_by_id"`
	AppliedByName            string    `bson:"applied_by_name" json:"applied_by_name"`
	AppliedOn                time.Time `bson:"applied_on" json:"applied_on"`
	EmployeeRemarks          string    `bson:"employee_remarks" json:"employee_remarks"`
	IsApproved               bool      `bson:"is_approved" json:"is_approved"`
	ApprovedTaskWorkingHours float64   `bson:"approved_task_working_hours" json:"approved_task_working_hours"`
	ApproverRemarks          string    `bson:"approver_remarks" json:"approver_remarks"`
	ApproverId               string    `bson:"approver_id" json:"approver_id"`
	EstimatedEffortInHours   float64   `bson:"estimated_effort_in_hours" json:"estimated_effort_in_hours"`
	ActualEffortInHours      float64   `bson:"actual_effort_in_hours" json:"actual_effort_in_hours"`
	ApproverName             string    `bson:"approver_name" json:"approver_name"`
	ApprovedOn               time.Time `bson:"approved_on" json:"approved_on"`
}

type TimeSheetProject struct {
	ProjectID                        string        `bson:"project_id" json:"project_id"`
	ProjectName                      string        `bson:"project_name" json:"project_name"`
	ProjectHeadId                    string        `bson:"project_head_id" json:"project_head_id"`
	ProjectHeadName                  string        `bson:"project_head_name" json:"project_head_name"`
	ManagerID                        string        `bson:"manager_id" json:"manager_id"`
	ManagerName                      string        `bson:"manager_name" json:"manager_name"`
	TotalClaimedProjectWorkingHours  float64       `bson:"total_claimed_project_working_hours" json:"total_claimed_project_working_hours"`
	TotalApprovedProjectWorkingHours float64       `bson:"total_approved_project_working_hours" json:"total_approved_project_working_hours"`
	Tasks                            []ProjectTask `bson:"tasks" json:"tasks"`
}

type TimeSheet struct {
	Day                  int                `bson:"day" json:"day"`
	Date                 string             `bson:"date" json:"date"`
	FirstCheckInTime     time.Time          `bson:"first_check_in_time" json:"first_check_in_time"`
	LastCheckOutTime     time.Time          `bson:"last_check_out_time" json:"last_check_out_time"`
	RequiredWorkingHours float64            `bson:"required_working_hours" json:"required_working_hours"`
	TotalWorkingHours    float64            `bson:"total_working_hours" json:"total_working_hours"`
	ClaimedWorkingHours  float64            `bson:"claimed_working_hours" json:"claimed_working_hours"`
	ApprovedWorkingHours float64            `bson:"approved_working_hours" json:"approved_working_hours"`
	IsWorkingDay         bool               `bson:"is_working_day" json:"is_working_day"`
	IsHoliday            bool               `bson:"is_holiday" json:"is_holiday"`
	IsOnLeave            bool               `bson:"is_on_leave" json:"is_on_leave"`
	IsWeeklyOffDay       bool               `bson:"is_weekly_off_day" json:"is_weekly_off_day"`
	Entries              []TimeSheetEntry   `bson:"entries" json:"entries"`
	Projects             []TimeSheetProject `bson:"projects" json:"projects"`
}

type TimeSheetEntry struct {
	ProjectId string  `bson:"project_id" json:"project_id"`
	TaskId    string  `bson:"task_id" json:"task_id"`
	Hours     float64 `bson:"hours" json:"hours"`
	Remark    string  `bson:"remark" json:"remark"`
}

type PendingApprovalEmployee struct {
	EmployeeID           string        `json:"employee_id" bson:"employee_id"`
	EmployeeName         string        `json:"employee_name" bson:"employee_name"`
	ClaimedWorkingHours  float64       `json:"claimed_working_hours" bson:"claimed_working_hours"`
	ApprovedWorkingHours float64       `json:"approved_working_hours" bson:"approved_working_hours"`
	Tasks                []ProjectTask `json:"tasks" bson:"tasks"`
}

type PendingApprovalProject struct {
	ProjectID            string                    `json:"project_id" bson:"project_id"`
	ProjectName          string                    `json:"project_name" bson:"project_name"`
	ClaimedProjectHours  float64                   `json:"claimed_project_hours" bson:"claimed_project_hours"`
	ApprovedProjectHours float64                   `json:"approved_project_hours" bson:"approved_project_hours"`
	Employees            []PendingApprovalEmployee `json:"employees" bson:"employees"`
}

type MonthlyPendingApproval struct {
	Id                    primitive.ObjectID     `json:"_id" bson:"_id"`
	Month                 int                    `json:"month" bson:"month"`
	Year                  int                    `json:"year" bson:"year"`
	ManagerID             string                 `json:"manager_id" bson:"manager_id"`
	ManagerName           string                 `json:"manager_name" bson:"manager_name"`
	IsProjectHead         bool                   `json:"is_project_head" bson:"is_project_head"`
	TotalPendingApprovals int                    `json:"total_pending_approvals" bson:"total_pending_approvals"`
	DailyPendingApproval  []DailyPendingApproval `json:"daily_pending_approval" bson:"daily_pending_approval"`
}

type DailyPendingApproval struct {
	Day                   int                      `json:"day" bson:"day"`
	TotalPendingApprovals int                      `json:"total_pending_approvals" bson:"total_pending_approvals"`
	Projects              []PendingApprovalProject `json:"projects" bson:"projects"`
}
