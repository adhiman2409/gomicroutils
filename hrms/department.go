package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeInfo struct {
	EmployeeId       string   `bson:"employee_id"`
	FullName         string   `bson:"full_name"`
	Department       string   `bson:"department"`
	Designation      string   `bson:"designation"`
	EmailID          string   `bson:"email_id"`
	Status           string   `bson:"status"`
	ProfileImageURL  string   `bson:"profile_image_url"`
	CurrentStatus    string   `bson:"current_status"`
	CheckInSource    string   `bson:"check_in_source"`
	AttendanceStatus []string `bson:"attendance_status"`
}

type AttendanceStats struct {
	CheckedIn                   uint32  `bson:"checked_in"`
	CheckedOut                  uint32  `bson:"checked_out"`
	OnLeave                     uint32  `bson:"on_leave"`
	Absent                      uint32  `bson:"absent"`
	AverageLoginDurationToday   string  `bson:"average_login_duration_today"`
	AttendancePercentageToday   float64 `bson:"attendance_percentage_today"`
	AverageLoginDurationMonthly string  `bson:"average_login_duration_monthly"`
	AttendancePercentageMonthly float64 `bson:"attendance_percentage_monthly"`
	AverageLoginDurationYearly  string  `bson:"average_login_duration_yearly"`
	AttendancePercentageYearly  float64 `bson:"attendance_percentage_yearly"`
}

type Department struct {
	ID                              primitive.ObjectID `bson:"_id"`
	OrgId                           string             `bson:"org_id"`
	OrgName                         string             `bson:"org_name"`
	BudgetId                        string             `bson:"budget_id"`
	Name                            string             `bson:"name"`
	Alias                           string             `bson:"alias"`
	Abbreviation                    string             `bson:"abbreviation"`
	AdminId                         string             `bson:"admin_id"`
	AdminName                       string             `bson:"admin_name"`
	TotalEmployees                  uint32             `bson:"total_employees"`
	ActiveEmployees                 uint32             `bson:"active_employees"`
	NoticePeriodEmployees           uint32             `bson:"notice_period_employees"`
	ExitedEmployees                 uint32             `bson:"exited_employees"`
	EmployeeList                    []EmployeeInfo     `bson:"employee_list"`
	AttendanceStats                 AttendanceStats    `bson:"attendance_stats"`
	Description                     string             `bson:"description"`
	IsActive                        bool               `bson:"is_active"`
	CreatedBy                       string             `bson:"created_by"`
	CreatedAt                       time.Time          `bson:"created_at"`
	UpdatedAt                       time.Time          `bson:"updated_at"`
	AllocatedBudget                 float32            `bson:"allocated_budget"`
	UsedBudget                      float32            `bson:"used_budget"`
	RemainingBudget                 float32            `bson:"remaining_budget"`
	TotalPositions                  int                `bson:"total_positions"`
	ClosedPositions                 int                `bson:"closed_positions"`
	PendingPositions                int                `bson:"pending_positions"`
	AverageDurationToClosePositions float32            `bson:"average_duration_to_close_positions"`
}
