package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeSummary struct {
	ID                                primitive.ObjectID `bson:"_id"`
	EmployeeId                        string             `bson:"employee_id"`
	FullName                          string             `bson:"full_name"`
	DOB                               string             `bson:"dob"`
	PhoneNumber                       string             `bson:"phone_number"`
	EmailId                           string             `bson:"email_id"`
	EmployeeType                      string             `bson:"employee_type"`
	Department                        string             `bson:"department"`
	Designation                       string             `bson:"designation"`
	Role                              string             `bson:"role"`
	JoiningDate                       string             `bson:"joining_date"`
	JoiningDay                        int64              `bson:"joining_day"`
	JoiningMonth                      int64              `bson:"joining_month"`
	JoiningYear                       int64              `bson:"joining_year"`
	EmploymentStatus                  string             `bson:"employment_status"`
	ProbationOrTrainingPeriodInMonths int64              `bson:"probation_or_training_period_in_months"`
	ConfirmationDate                  string             `bson:"confirmation_date"`
	IsRotationalShift                 bool               `bson:"is_rotational_shift"`
	CurrentProject                    string             `bson:"current_project"`
	CurrentClient                     string             `bson:"current_client"`
	IsBillable                        bool               `bson:"is_billable"`
	OfficeLocation                    string             `bson:"office_location"`
	WorkLocation                      string             `bson:"work_location"`
	ReportingManagerName              string             `bson:"reporting_manager"`
	ReportingManagerId                string             `bson:"reporting_manager_id"`
	ReportingManagerEmail             string             `bson:"reporting_manager_email"`
	NewJoineeRewardId                 string             `bson:"new_joinee_reward_id"`
	ImgURL                            string             `bson:"img_url"`
	ReleavingDate                     string             `bson:"releaving_date"`
	IsProfileEditingLocked            bool               `bson:"is_profile_editing_locked"`
	Country                           string             `bson:"country"`
	TimeZone                          string             `bson:"time_zone"`
	RegexText                         string             `bson:"regex_text"`
	CreatedBy                         string             `bson:"created_by"`
	CreatedAt                         time.Time          `bson:"created_at"`
	UpdatedAt                         time.Time          `bson:"updated_at"`
}
