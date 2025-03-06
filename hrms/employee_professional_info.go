package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpProfessionalInfo struct {
	ID                    primitive.ObjectID `bson:"_id"`
	EmployeeId            string             `bson:"employee_id"`
	EmployeeType          string             `bson:"employee_type"`
	Department            string             `bson:"department"`
	Designation           string             `bson:"designation"`
	Role                  string             `bson:"role"`
	EmailId               string             `bson:"email_id"`
	JoiningDate           string             `bson:"joining_date"`
	JoiningDay            int64              `bson:"joining_day"`
	JoiningMonth          int64              `bson:"joining_month"`
	JoiningYear           int64              `bson:"joining_year"`
	ReleavingDate         string             `bson:"releaving_date"`
	ReleavingRemarks      string             `bson:"releaving_remarks"`
	EmploymentStatus      string             `bson:"employment_status"`
	IsRotationalShift     bool               `bson:"is_rotational_shift"`
	OfficeLocation        string             `bson:"office_location"`
	WorkLocation          string             `bson:"work_location"`
	CIBILScore            string             `bson:"cibil_score"`
	ClientInfo            Client             `bson:"client_info"`
	WorkAddress           Address            `bson:"work_address"`
	ReportingManagerName  string             `bson:"reporting_manager"`
	ReportingManagerId    string             `bson:"reporting_manager_id"`
	ReportingManagerEmail string             `bson:"reporting_manager_email"`
	RegexText             string             `bson:"regex_text"`
	CreatedBy             string             `bson:"created_by"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
}
