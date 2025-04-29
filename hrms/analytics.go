package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenderWise struct {
	Gender    string    `bson:"gender"`
	Count     int       `bson:"count"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

type DepartmentWise struct {
	Department string    `bson:"department"`
	Count      int       `bson:"count"`
	UpdatedAt  time.Time `bson:"updated_at"`
	CreatedAt  time.Time `bson:"created_at"`
}

type ExperienceWise struct {
	ExperienceGroup string    `bson:"experience_group"`
	Count           int       `bson:"count"`
	UpdatedAt       time.Time `bson:"updated_at"`
	CreatedAt       time.Time `bson:"created_at"`
}

type AgeWise struct {
	AgeGroup  string    `bson:"age_group"`
	Count     int       `bson:"count"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}
type TenureWise struct {
	TenureGroup string    `bson:"tenure_group"`
	Count       int       `bson:"count"`
	UpdatedAt   time.Time `bson:"updated_at"`
	CreatedAt   time.Time `bson:"created_at"`
}

type ClientWise struct {
	ClientName string    `bson:"client_name"`
	Count      int       `bson:"count"`
	UpdatedAt  time.Time `bson:"updated_at"`
	CreatedAt  time.Time `bson:"created_at"`
}

type SkillWise struct {
	SkillName string    `bson:"skill_name"`
	Count     int       `bson:"count"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

type EventWise struct {
	EventName string    `bson:"event_name"`
	Count     int       `bson:"count"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

type LocationWise struct {
	City      string    `bson:"city"`
	State     string    `bson:"state"`
	Country   string    `bson:"country"`
	Count     int       `bson:"count"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

type HiringAnalytics struct {
	Day            int              `bson:"day"`
	Month          int              `bson:"month"`
	Year           int              `bson:"year"`
	JoiningCount   int              `bson:"joining_count"`
	GenderWise     []GenderWise     `bson:"gender_wise"`
	DepartmentWise []DepartmentWise `bson:"department_wise"`
	ExperienceWise []ExperienceWise `bson:"experience_wise"`
	UpdatedAt      time.Time        `bson:"updated_at"`
	CreatedAt      time.Time        `bson:"created_at"`
}

type ExitAnalytics struct {
	Day            int              `bson:"day"`
	Month          int              `bson:"month"`
	Year           int              `bson:"year"`
	ExitCount      int              `bson:"exit_count"`
	GenderWise     []GenderWise     `bson:"gender_wise"`
	DepartmentWise []DepartmentWise `bson:"department_wise"`
	ExperienceWise []ExperienceWise `bson:"experience_wise"`
	UpdatedAt      time.Time        `bson:"updated_at"`
	CreatedAt      time.Time        `bson:"created_at"`
}

type EmploymentAnalytics struct {
	Day                   int       `bson:"day"`
	Month                 int       `bson:"month"`
	Year                  int       `bson:"year"`
	TotalEmployees        int       `bson:"total_employees"`
	ActiveEmployees       int       `bson:"active_employees"`
	InactiveEmployees     int       `bson:"inactive_employees"`
	NoticePeriodEmployees int       `bson:"notice_period_employees"`
	OnPIPEmployees        int       `bson:"on_pip_employees"`
	SuspendedEmployees    int       `bson:"suspended_employees"`
	UpdatedAt             time.Time `bson:"updated_at"`
	CreatedAt             time.Time `bson:"created_at"`
}

type EmployeeAnalytics struct {
	Day                 int              `bson:"day"`
	Month               int              `bson:"month"`
	Year                int              `bson:"year"`
	EmployeeCount       int              `bson:"employee_count"`
	BillableCount       int              `bson:"billable_count"`
	NonBillableCount    int              `bson:"non_billable_count"`
	WorkFromHomeCount   int              `bson:"work_from_home_count"`
	WorkFromOfficeCount int              `bson:"work_from_office_count"`
	ClientWise          []ClientWise     `bson:"client_wise"`
	LocationWise        []LocationWise   `bson:"location_wise"`
	TenureWise          []TenureWise     `bson:"tenure_wise"`
	SkillWise           []SkillWise      `bson:"skill_wise"`
	DepartmentWise      []DepartmentWise `bson:"department_wise"`
	ExperienceWise      []ExperienceWise `bson:"experience_wise"`
	AgeWise             []AgeWise        `bson:"age_wise"`
	UpdatedAt           time.Time        `bson:"updated_at"`
	CreatedAt           time.Time        `bson:"created_at"`
}

type EventsAnalytics struct {
	Day         int         `bson:"day"`
	Month       int         `bson:"month"`
	Year        int         `bson:"year"`
	EventsCount int         `bson:"events_count"`
	EventWise   []EventWise `bson:"event_wise"`
	UpdatedAt   time.Time   `bson:"updated_at"`
	CreatedAt   time.Time   `bson:"created_at"`
}

type EAnanlytics struct {
	ID                  primitive.ObjectID    `bson:"_id"`
	NextUpdateType      string                `bson:"next_update_type"`
	HiringAnalytics     []HiringAnalytics     `bson:"hiring_analytics"`
	ExitAnalytics       []ExitAnalytics       `bson:"exit_analytics"`
	EmploymentAnalytics []EmploymentAnalytics `bson:"employment_analytics"`
	EventsAnalytics     []EventsAnalytics     `bson:"events_analytics"`
	EmployeeAnalytics   []EmployeeAnalytics   `bson:"employee_analytics"`
	ImageURL            CompletionStats       `bson:"image_url"`
	About               CompletionStats       `bson:"about"`
	PersonalDetails     CompletionStats       `bson:"personal_details"`
	ProfessionalDetails CompletionStats       `bson:"professional_details"`
	FamilyDetails       CompletionStats       `bson:"family_details"`
	BankDetails         CompletionStats       `bson:"bank_details"`
	KeySkills           CompletionStats       `bson:"key_skills"`
	Education           CompletionStats       `bson:"education"`
	Experience          CompletionStats       `bson:"experince"`
	Projects            CompletionStats       `bson:"projects"`
	Documents           CompletionStats       `bson:"documents"`
	UpdatedAt           time.Time             `bson:"updated_at"`
	CreatedAt           time.Time             `bson:"created_at"`
}

type HourlyStats struct {
	Hour          int `bson:"hour"`
	CheckInCount  int `bson:"check_in_count"`
	CheckOutCount int `bson:"check_out_count"`
}

type LiveAttendanceStats struct {
	Day                 int             `bson:"day"`
	Month               int             `bson:"month"`
	Year                int             `bson:"year"`
	Weekday             string          `bson:"weekday"`
	IsWorkingDay        bool            `bson:"is_working_day"`
	IsWeeklyOff         bool            `bson:"is_weekly_off"`
	IsHoliday           bool            `bson:"is_holiday"`
	TotalEmployees      int             `bson:"total_employees"`
	PresentCount        int             `bson:"present_count"`
	AbsentCount         int             `bson:"absent_count"`
	LeaveCount          int             `bson:"leave_count"`
	CheckInCount        int             `bson:"check_in_count"`
	CheckOutCount       int             `bson:"check_out_count"`
	AverageWorkingHours int             `bson:"average_working_hours"`
	HourlyStats         [24]HourlyStats `bson:"hourly_stats"`
	UpdatedAt           time.Time       `bson:"updated_at"`
	CreatedAt           time.Time       `bson:"created_at"`
}

type LiveLeaveStats struct {
	Day           int       `bson:"day"`
	Month         int       `bson:"month"`
	Year          int       `bson:"year"`
	AppliedCount  int       `bson:"applied_count"`
	ApprovedCount int       `bson:"approved_count"`
	PendingCount  int       `bson:"pending_count"`
	RejectedCount int       `bson:"rejected_count"`
	ApprovalRate  int       `bson:"approval_rate"`
	RejectionRate int       `bson:"rejection_rate"`
	UpdatedAt     time.Time `bson:"updated_at"`
	CreatedAt     time.Time `bson:"created_at"`
}

type AttendanceAnalytics struct {
	Day            int `bson:"day"`
	Month          int `bson:"month"`
	Year           int `bson:"year"`
	TotalEmployees int `bson:"total_employees"`
	TenureDays     int `bson:"tenure_days"`
	WorkingDays    int `bson:"working_days"`
	LeaveDays      int `bson:"leave_days"`
	LOPDays        int `bson:"lop_days"`
}

type LAnalytics struct {
	ID                  primitive.ObjectID    `bson:"_id"`
	NextUpdateType      string                `bson:"next_update_type"`
	LiveAttendanceStats LiveAttendanceStats   `bson:"live_attendance_stats"`
	LiveLeaveStats      LiveLeaveStats        `bson:"live_leave_stats"`
	AttendanceAnalytics []AttendanceAnalytics `bson:"attendance_analytics"`
}

type PendingEmployee struct {
	EmployeeName  string `bson:"employee_name"`
	EmployeeId    string `bson:"employee_id"`
	EmployeeEmail string `bson:"employee_email"`
}

type CompletionStats struct {
	TotalAvtiveEmployees int               `bson:"total_active_employees"`
	CompletedCount       int               `bson:"completed_count"`
	PendingCount         int               `bson:"pending_count"`
	PendingEmployeeList  []PendingEmployee `bson:"pending_employee_list"`
	UpdatedAt            time.Time         `bson:"updated_at"`
}
