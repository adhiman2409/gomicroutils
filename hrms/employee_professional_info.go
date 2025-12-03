package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpProfessionalInfo struct {
	ID                                primitive.ObjectID `bson:"_id"`
	EmployeeId                        string             `bson:"employee_id"`
	EmployeeType                      string             `bson:"employee_type"`
	Department                        string             `bson:"department"`
	Designation                       string             `bson:"designation"`
	Role                              string             `bson:"role"`
	EmailId                           string             `bson:"email_id"`
	JoiningDate                       string             `bson:"joining_date"`
	JoiningDay                        int64              `bson:"joining_day"`
	JoiningMonth                      int64              `bson:"joining_month"`
	JoiningYear                       int64              `bson:"joining_year"`
	ProbationOrTrainingPeriodInMonths int64              `bson:"probation_or_training_period_in_months"`
	ConfirmationDate                  string             `bson:"confirmation_date"`
	ReleavingDate                     string             `bson:"releaving_date"`
	ReleavingRemarks                  string             `bson:"releaving_remarks"`
	MarkedReleavedOn                  time.Time          `bson:"marked_releaved_on"`
	EmploymentStatus                  string             `bson:"employment_status"`
	IsRotationalShift                 bool               `bson:"is_rotational_shift"`
	OfficeLocation                    string             `bson:"office_location"`
	WorkLocation                      string             `bson:"work_location"`
	CIBILScore                        string             `bson:"cibil_score"`
	ClientInfo                        Client             `bson:"client_info"`
	WorkAddress                       Address            `bson:"work_address"`
	CurrentProject                    string             `bson:"current_project"`
	CurrentClient                     string             `bson:"current_client"`
	IsBillable                        bool               `bson:"is_billable"`
	ReportingManagerName              string             `bson:"reporting_manager"`
	ReportingManagerId                string             `bson:"reporting_manager_id"`
	ReportingManagerEmail             string             `bson:"reporting_manager_email"`
	IsProfessionalInfoLocked          bool               `bson:"is_professional_info_locked"`
	Country                           string             `bson:"country"`
	TimeZone                          string             `bson:"time_zone"`
	RegexText                         string             `bson:"regex_text"`
	CreatedBy                         string             `bson:"created_by"`
	CreatedAt                         time.Time          `bson:"created_at"`
	UpdatedAt                         time.Time          `bson:"updated_at"`
}
