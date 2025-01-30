package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientInfo struct {
	ClientId     string    `bson:"client_id"`
	ClientName   string    `bson:"client_name"`
	ClientType   string    `bson:"client_type"`
	ClientStatus string    `bson:"client_status"`
	CreatedBy    string    `bson:"created_by"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

type ProjectInfo struct {
	ID                    primitive.ObjectID `bson:"_id"`
	ClientId              string             `bson:"client_id"`
	ProjectId             string             `bson:"project_id"`
	ProjectName           string             `bson:"project_name"`
	ProjectType           string             `bson:"project_type"`
	ProjectStatus         string             `bson:"project_status"`
	StartDate             string             `bson:"start_date"`
	EndDate               string             `bson:"end_date"`
	ProjectManager        string             `bson:"project_manager"`
	ProjectManagerEmail   string             `bson:"project_manager_email"`
	ReportingManager      string             `bson:"reporting_manager"`
	ReportingManagerEmail string             `bson:"reporting_manager_email"`
	ProjectWorkLocation   string             `bson:"project_work_location"`
	ProjectWorkAddresses  []Address          `bson:"project_work_addresses"`
	CreatedBy             string             `bson:"created_by"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
}

type ProjectTeam struct {
	ID                    primitive.ObjectID `bson:"_id"`
	ProjectId             string             `bson:"project_id"`
	EmployeeId            string             `bson:"employee_id"`
	EmployeeName          string             `bson:"employee_name"`
	EmployeeEmail         string             `bson:"employee_email"`
	EmployeeRole          string             `bson:"employee_role"`
	EmployeeDesignation   string             `bson:"employee_designation"`
	EmployeeDepartment    string             `bson:"employee_department"`
	EmployeeType          string             `bson:"employee_type"`
	EmployeeStatus        string             `bson:"employee_status"`
	EmployeeJoiningDate   string             `bson:"employee_joining_date"`
	EmployeeReleavingDate string             `bson:"employee_releaving_date"`
	EmployeeWorkLocation  string             `bson:"employee_work_location"`
	EmployeeWorkAddress   Address            `bson:"employee_work_address"`
	CreatedBy             string             `bson:"created_by"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
}
