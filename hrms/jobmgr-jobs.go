package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobStats struct {
	State   string
	Counter int64
}

type EmpIdAndName struct {
	EID  string `bson:"eid"`
	Name string `bson:"name"`
}

type ClosedPositionInfo struct {
	ApplicationId string `bson:"application_id"`
	CandidateId   string `bson:"candidate_id"`
	CandidateName string `bson:"candidate_name"`
	JoiningDate   string `bson:"joining_date"`
	ClosedOn      string `bson:"closed_on"`
	ClosedById    string `bson:"closed_by_id"`
	ClosedBy      string `bson:"closed_by"`
	Remarks       string `bson:"remarks"`
}

type JobEntry struct {
	ID                       primitive.ObjectID   `bson:"_id"`
	Title                    string               `bson:"title"`
	JobId                    string               `bson:"job_id"`
	EmploymentType           string               `bson:"employment_type"`
	JobDomain                string               `bson:"job_domain"`
	Technology               string               `bson:"technology"`
	Designation              string               `bson:"designation"`
	Locations                []string             `bson:"locations"`
	ExperinceInYears         int64                `bson:"experince_in_years"`
	ExperinceRange           string               `bson:"experince_range"`
	NumberOfOpenings         int64                `bson:"number_of_openings"`
	NumberOfPositionsClosed  int64                `bson:"number_of_positions_closed"`
	NumberOfPositionsPending int64                `bson:"number_of_positions_pending"`
	ClosedPositionsInfo      []ClosedPositionInfo `bson:"closed_positions_info"`
	Description              string               `bson:"description"`
	MandatarySkills          []string             `bson:"mandatory_skills"`
	NiceToHaveSkills         []string             `bson:"nice_to_have_skills"`
	Qualification            []string             `bson:"qualification"`
	MaxNoticePeriodInMonths  int64                `bson:"max_notice_period_in_months"`
	BudgetRange              string               `bson:"budget_range"`
	BudgetInINR              int64                `bson:"budget_in_inr"`
	StartDate                string               `bson:"start_date"`
	EndDate                  string               `bson:"end_date"`
	Status                   JobPublishStatus     `bson:"status"`
	JobStatus                JobStatus            `bson:"job_status"`
	Client                   string               `bson:"client"`
	ClientId                 string               `bson:"client_id"`
	MetaTags                 string               `bson:"meta_tags"`
	JobAnalytics             []JobStats           `bson:"job_analytics"`
	AssignedToList           []EmpIdAndName       `bson:"assigned_to_list"`
	RequestedById            string               `bson:"requested_by_id"`
	RequestedByName          string               `bson:"requested_by_name"`
	Department               string               `bson:"department"`
	ProjectId                string               `bson:"project_id"`
	ProjectName              string               `bson:"project_name"`
	IsApproved               bool                 `bson:"is_approved"`
	ApprovedOn               string               `bson:"approved_on"`
	ManagerId                string               `bson:"manager_id"`
	ManagerName              string               `bson:"manager_name"`
	RecruiterId              string               `bson:"recruiter_id"`
	RecruiterName            string               `bson:"recruiter_name"`
	MarkedDeleted            bool                 `bson:"marked_deleted"`
	IsLive                   bool                 `bson:"is_live"`
	JDURL                    string               `bson:"jd_url"`
	ImageURL                 string               `bson:"image_url"`
	Remarks                  string               `bson:"remarks"`
	NaukriT                  string               `bson:"naukri_t"`
	CreatedAt                int64                `bson:"created_at"`
	UpdatedAt                int64                `bson:"updated_at"`
}
