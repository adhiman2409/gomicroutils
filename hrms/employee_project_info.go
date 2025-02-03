package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientInfo struct {
	ID                primitive.ObjectID `bson:"_id"`
	ClientId          string             `bson:"client_id"`
	ClientName        string             `bson:"client_name"`
	ClientType        string             `bson:"client_type"`
	ClientStatus      string             `bson:"client_status"`
	ClientDescription string             `bson:"client_description"`
	ClientWebsite     string             `bson:"client_website"`
	ClientLocation    string             `bson:"client_location"`
	ReferencedBy      string             `bson:"referenced_by"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}

type KYCInfo struct {
	ID             primitive.ObjectID  `bson:"_id"`
	ClientId       string              `bson:"client_id"`
	Date           string              `bson:"date"`
	ClientContacts []ClientContactInfo `bson:"client_contacts"`
	PSContacts     []string            `bson:"ps_contacts"`
	Remarks        string              `bson:"remarks"`
	ActionItems    []string            `bson:"action_items"`
	ReminderDate   string              `bson:"reminder_date"`
	NextDate       string              `bson:"next_date"`
	Documents      []string            `bson:"documents"`
	CreatedBy      string              `bson:"created_by"`
	CreatedAt      time.Time           `bson:"created_at"`
	UpdatedAt      time.Time           `bson:"updated_at"`
}

type ClientContactInfo struct {
	ID                 primitive.ObjectID `bson:"_id"`
	ClientId           string             `bson:"client_id"`
	ContactName        string             `bson:"contact_name"`
	ContactEmail       string             `bson:"contact_email"`
	ContactPhone       string             `bson:"contact_phone"`
	ContactDesignation string             `bson:"contact_designation"`
	ContactRole        string             `bson:"contact_role"`
	ContactDescription string             `bson:"contact_description"`
	CreatedBy          string             `bson:"created_by"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
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
	ID                      primitive.ObjectID `bson:"_id"`
	ClientId                string             `bson:"client_id"`
	ProjectId               string             `bson:"project_id"`
	ProjectName             string             `bson:"project_name"`
	ProjectType             string             `bson:"project_type"`
	ProjectStatus           string             `bson:"project_status"`
	IsExternalProject       bool               `bson:"is_external_project"`
	IsBillable              bool               `bson:"is_billable"`
	IsRemoteWorking         bool               `bson:"is_remote_working"`
	StartDate               string             `bson:"start_date"`
	EndDate                 string             `bson:"end_date"`
	ProjectHead             string             `bson:"project_head"`
	ProjectHeadEmail        string             `bson:"project_head_email"`
	ProjectCoordinator      string             `bson:"project_coordinator"`
	ProjectCoordinatorEmail string             `bson:"project_coordinator_email"`
	ProjectWorkLocation     string             `bson:"project_work_location"`
	ProjectWorkAddresses    []Address          `bson:"project_work_addresses"`
	ProjectDescription      string             `bson:"project_description"`
	ProjectRemarks          string             `bson:"project_remarks"`
	ProjectHolidayList      []string           `bson:"project_holiday_list"`
	ProjectDailyWorkingHour string             `bson:"project_daily_working_hour"`
	ProjectWeeklyOffDay     string             `bson:"project_weekly_off_day"`
	ShiftTypes              []ShiftType        `bson:"shift_types"`
	CreatedBy               string             `bson:"created_by"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}

type ProjectTeamMember struct {
	ID                         primitive.ObjectID `bson:"_id"`
	ProjectId                  string             `bson:"project_id"`
	EmployeeId                 string             `bson:"employee_id"`
	EmployeeName               string             `bson:"employee_name"`
	EmployeeStatus             string             `bson:"employee_status"`
	EmployeeJoiningDate        string             `bson:"employee_joining_date"`
	EmployeeReleavingDate      string             `bson:"employee_releaving_date"`
	EmployeeProjectEmail       string             `bson:"employee_project_email"`
	EmployeeProjectDesignation string             `bson:"employee_project_designation"`
	EmployeeProjectDepartment  string             `bson:"employee_project_department"`
	EmployeeWorkLocation       string             `bson:"employee_work_location"`
	EmployeeWorkAddress        Address            `bson:"employee_work_address"`
	ProjectManager             string             `bson:"project_manager"`
	ProjectManagerEmail        string             `bson:"project_manager_email"`
	ReportingManager           string             `bson:"reporting_manager"`
	ReportingManagerEmail      string             `bson:"reporting_manager_email"`
	IsBillable                 bool               `bson:"is_billable"`
	IsRotationalShift          bool               `bson:"is_rotational_shift"`
	IsRemoteWorking            bool               `bson:"is_remote_working"`
	ShiftType                  ShiftType          `bson:"shift_type"`
	CreatedBy                  string             `bson:"created_by"`
	CreatedAt                  time.Time          `bson:"created_at"`
	UpdatedAt                  time.Time          `bson:"updated_at"`
}
