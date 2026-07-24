package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpProfessionalInfo struct {
	ID                                primitive.ObjectID   `bson:"_id"`
	EmployeeId                        string               `bson:"employee_id"`
	EmployeeType                      string               `bson:"employee_type"`
	Department                        string               `bson:"department"`
	Designation                       string               `bson:"designation"`
	Role                              string               `bson:"role"`
	EmailId                           string               `bson:"email_id"`
	JoiningDate                       string               `bson:"joining_date"`
	JoiningDay                        int64                `bson:"joining_day"`
	JoiningMonth                      int64                `bson:"joining_month"`
	JoiningYear                       int64                `bson:"joining_year"`
	ProbationOrTrainingPeriodInMonths int64                `bson:"probation_or_training_period_in_months"`
	ConfirmationDate                  string               `bson:"confirmation_date"`
	ReleavingDate                     string               `bson:"releaving_date"`
	ReleavingRemarks                  string               `bson:"releaving_remarks"`
	MarkedReleavedOn                  time.Time            `bson:"marked_releaved_on"`
	EmploymentStatus                  string               `bson:"employment_status"`
	IsRotationalShift                 bool                 `bson:"is_rotational_shift"`
	ShiftName                         string               `bson:"shift_name"`
	OfficeLocation                    string               `bson:"office_location"`
	WorkLocation                      string               `bson:"work_location"`
	CIBILScore                        string               `bson:"cibil_score"`
	ClientInfo                        Client               `bson:"client_info"`
	WorkAddress                       Address              `bson:"work_address"`
	CurrentProduct                    string               `bson:"current_product"`
	CurrentProject                    string               `bson:"current_project"`
	CurrentClient                     string               `bson:"current_client"`
	IsBillable                        bool                 `bson:"is_billable"`
	ReportingManagerName              string               `bson:"reporting_manager"`
	ReportingManagerId                string               `bson:"reporting_manager_id"`
	ReportingManagerEmail             string               `bson:"reporting_manager_email"`
	TeamLeadName                      string               `bson:"team_lead"`
	TeamLeadId                        string               `bson:"team_lead_id"`
	TeamLeadEmail                     string               `bson:"team_lead_email"`
	SupervisorName                    string               `bson:"supervisor_name"`
	SupervisorId                      string               `bson:"supervisor_id"`
	SupervisorEmail                   string               `bson:"supervisor_email"`
	IsProfessionalInfoLocked          bool                 `bson:"is_professional_info_locked"`
	Country                           string               `bson:"country"`
	State                             string               `bson:"state"`
	City                              string               `bson:"city"`
	OfficeLabel                       string               `bson:"office_label"`
	TimeZone                          string               `bson:"time_zone"`
	RegexText                         string               `bson:"regex_text"`
	PIPOn                             time.Time            `bson:"pip_on"`
	NoticePeriodOn                    time.Time            `bson:"notice_period_on"`
	InactiveOn                        time.Time            `bson:"inactive_on"`
	DepartmentTimeline                []DepartmentRecord   `bson:"department_timeline"`
	DesignationTimeline               []DesignationRecord  `bson:"designation_timeline"`
	OfficeTimeline                    []OfficeRecord       `bson:"office_timeline"`
	WorkLocationTimeline              []WorkLocationRecord `bson:"work_location_timeline"`
	PromotionDueDate                  time.Time            `bson:"promotion_due_date"`
	TransferDueDate                   time.Time            `bson:"transfer_due_date"`
	IsOnDeputation                    bool                 `bson:"is_on_deputation"`
	ActiveDeputation                  DeputationRecord     `bson:"active_deputation"`
	DeputationTimeline                []DeputationRecord   `bson:"deputation_timeline"`
	EmploymentTimeline                []EmploymentRecord   `bson:"employment_timeline"`
	ProjectTimeline                   []ProjectRecord      `bson:"project_timeline"`
	MohrePersonCode                   string               `bson:"mohre_person_code"`
	CreatedBy                         string               `bson:"created_by"`
	CreatedAt                         time.Time            `bson:"created_at"`
	UpdatedAt                         time.Time            `bson:"updated_at"`
}

type EmpStatusHistory struct {
	ID         primitive.ObjectID `bson:"_id"`
	EmployeeId string             `bson:"employee_id"`
	FromStatus string             `bson:"from_status"`
	ToStatus   string             `bson:"to_status"`
	ChangedBy  string             `bson:"changed_by"`
	Remarks    string             `bson:"remarks"`
	ChangedAt  time.Time          `bson:"changed_at"`
}
