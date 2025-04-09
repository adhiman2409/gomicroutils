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
	Month          string           `bson:"month"`
	Year           string           `bson:"year"`
	JoiningCount   int              `bson:"joining_count"`
	GenderWise     []GenderWise     `bson:"gender_wise"`
	DepartmentWise []DepartmentWise `bson:"department_wise"`
	ExperienceWise []ExperienceWise `bson:"experience_wise"`
	UpdatedAt      time.Time        `bson:"updated_at"`
	CreatedAt      time.Time        `bson:"created_at"`
}

type ExitAnalytics struct {
	Month          string           `bson:"month"`
	Year           string           `bson:"year"`
	ExitCount      int              `bson:"exit_count"`
	GenderWise     []GenderWise     `bson:"gender_wise"`
	DepartmentWise []DepartmentWise `bson:"department_wise"`
	ExperienceWise []ExperienceWise `bson:"experience_wise"`
	UpdatedAt      time.Time        `bson:"updated_at"`
	CreatedAt      time.Time        `bson:"created_at"`
}

type EmploymentAnalytics struct {
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
	Month            string           `bson:"month"`
	Year             string           `bson:"year"`
	EmployeeCount    int              `bson:"employee_count"`
	BillableCount    int              `bson:"billable_count"`
	NonBillableCount int              `bson:"non_billable_count"`
	GenderWise       []GenderWise     `bson:"gender_wise"`
	DepartmentWise   []DepartmentWise `bson:"department_wise"`
	ExperienceWise   []ExperienceWise `bson:"experience_wise"`
	AgeWise          []AgeWise        `bson:"age_wise"`
	ClientWise       []ClientWise     `bson:"client_wise"`
	LocationWise     []LocationWise   `bson:"location_wise"`
	TenureWise       []TenureWise     `bson:"tenure_wise"`
	SkillWise        []SkillWise      `bson:"skill_wise"`
	UpdatedAt        time.Time        `bson:"updated_at"`
	CreatedAt        time.Time        `bson:"created_at"`
}

type EventsAnalytics struct {
	Month       string      `bson:"month"`
	Year        string      `bson:"year"`
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
	EmployeeAnalytics   []EmployeeAnalytics   `bson:"employee_analytics"`
	EventsAnalytics     []EventsAnalytics     `bson:"events_analytics"`
	UpdatedAt           time.Time             `bson:"updated_at"`
	CreatedAt           time.Time             `bson:"created_at"`
}
