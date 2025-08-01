package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectIdCounter struct {
	ID       primitive.ObjectID `bson:"_id"`
	ClientId string             `bson:"client_id"`
	Prefix   string             `bson:"prefix"`
	Counter  int64              `bson:"counter"`
}

type ProjectShifts struct {
	ID                primitive.ObjectID `bson:"_id"`
	ProjectId         string             `bson:"project_id"`
	ShiftType         ShiftType          `bson:"shift_type"`
	Abbrivation       string             `bson:"abbrivation"`
	IsActive          bool               `bson:"is_active"`
	IsOverlappedShift bool               `bson:"is_overlapped_shift"`
	OrgCheckInTime    string             `bson:"org_check_in_time"`
	OrgCheckOutTime   string             `bson:"org_check_out_time"`
	Color             string             `bson:"color"`
	TextColor         string             `bson:"text_color"`
}

type ProjectInfo struct {
	ID                      primitive.ObjectID  `bson:"_id"`
	ClientId                string              `bson:"client_id"`
	ProjectId               string              `bson:"project_id"`
	ProjectName             string              `bson:"project_name"`
	ProjectType             string              `bson:"project_type"`
	ProjectStatus           string              `bson:"project_status"`
	IsExternalProject       bool                `bson:"is_external_project"`
	IsDefault               bool                `bson:"is_default"`
	IsBillable              bool                `bson:"is_billable"`
	IsRemoteWorking         bool                `bson:"is_remote_working"`
	StartDate               string              `bson:"start_date"`
	EndDate                 string              `bson:"end_date"`
	ProjectHeadId           string              `bson:"project_head_id"`
	ProjectHeadName         string              `bson:"project_head_name"`
	ProjectHeadEmail        string              `bson:"project_head_email"`
	ProjectCoordinatorName  string              `bson:"project_coordinator_name"`
	ProjectCoordinatorId    string              `bson:"project_coordinator_id"`
	ProjectCoordinatorEmail string              `bson:"project_coordinator_email"`
	ProjectWorkLocation     string              `bson:"project_work_location"`
	ProjectWorkAddresses    []Address           `bson:"project_work_addresses"`
	ProjectDescription      string              `bson:"project_description"`
	ProjectRemarks          string              `bson:"project_remarks"`
	ProjectHolidayList      []string            `bson:"project_holiday_list"`
	ProjectDailyWorkingHour string              `bson:"project_daily_working_hour"`
	ProjectWeeklyOffDay     string              `bson:"project_weekly_off_day"`
	TeamMembers             []ProjectTeamMember `bson:"team_members"`
	ReleasedTeamMembers     []ProjectTeamMember `bson:"released_team_members"`
	ShiftTypes              []ShiftType         `bson:"shift_types"`
	CreatedBy               string              `bson:"created_by"`
	CreatedAt               time.Time           `bson:"created_at"`
	UpdatedAt               time.Time           `bson:"updated_at"`
}

type ProjectTeamMember struct {
	ID                         string    `bson:"_id"`
	ProjectId                  string    `bson:"project_id"`
	EmployeeId                 string    `bson:"employee_id"`
	EmployeeName               string    `bson:"employee_name"`
	EmployeeStatus             string    `bson:"employee_status"`
	EmployeeJoiningDate        string    `bson:"employee_joining_date"`
	EmployeeReleavingDate      string    `bson:"employee_releaving_date"`
	IsReleasedFromProject      bool      `bson:"is_released_from_project"`
	EmployeeProjectEmail       string    `bson:"employee_project_email"`
	EmployeeProjectDesignation string    `bson:"employee_project_designation"`
	EmployeeProjectDepartment  string    `bson:"employee_project_department"`
	EmployeeWorkLocation       string    `bson:"employee_work_location"`
	EmployeeWorkAddress        Address   `bson:"employee_work_address"`
	ProjectManagerId           string    `bson:"project_manager_id"`
	ProjectManagerName         string    `bson:"project_manager_name"`
	ProjectManagerEmail        string    `bson:"project_manager_email"`
	ReportingManagerId         string    `bson:"reporting_manager_id"`
	ReportingManagerName       string    `bson:"reporting_manager_name"`
	ReportingManagerEmail      string    `bson:"reporting_manager_email"`
	IsBillable                 bool      `bson:"is_billable"`
	IsRotationalShift          bool      `bson:"is_rotational_shift"`
	IsRemoteWorking            bool      `bson:"is_remote_working"`
	ShiftType                  ShiftType `bson:"shift_type"`
	CreatedBy                  string    `bson:"created_by"`
	CreatedAt                  time.Time `bson:"created_at"`
	UpdatedAt                  time.Time `bson:"updated_at"`
}
