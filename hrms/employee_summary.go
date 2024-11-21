package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeSummary struct {
	ID                    primitive.ObjectID `bson:"_id"`
	EmployeeId            string             `bson:"employee_id"`
	FullName              string             `bson:"full_name"`
	DOB                   string             `bson:"dob"`
	PhoneNumber           string             `bson:"phone_number"`
	EmailId               string             `bson:"email_id"`
	EmployeeType          string             `bson:"employee_type"`
	Department            string             `bson:"department"`
	Designation           string             `bson:"designation"`
	Role                  string             `bson:"role"`
	JoiningDate           string             `bson:"joining_date"`
	JoiningDay            string             `bson:"joining_day"`
	JoiningMonth          string             `bson:"joining_month"`
	JoiningYear           string             `bson:"joining_year"`
	EmploymentStatus      string             `bson:"employment_status"`
	OfficeLocation        string             `bson:"office_location"`
	WorkLocation          string             `bson:"work_location"`
	ReportingManagerName  string             `bson:"reporting_manager"`
	ReportingManagerId    string             `bson:"reporting_manager_id"`
	ReportingManagerEmail string             `bson:"reporting_manager_email"`
	ImgURL                string             `bson:"img_url"`
	RegexText             string             `bson:"regex_text"`
	CreatedBy             string             `bson:"created_by"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
}
