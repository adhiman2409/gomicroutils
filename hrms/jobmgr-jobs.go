package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobStats struct {
	State   string
	Counter int64
}

type JobEntry struct {
	ID                      primitive.ObjectID `bson:"_id"`
	Title                   string             `bson:"title"`
	JobId                   string             `bson:"job_id"`
	EmploymentType          string             `bson:"employment_type"`
	JobDomain               string             `bson:"job_domain"`
	Technology              string             `bson:"technology"`
	Designation             string             `bson:"designation"`
	Locations               []string           `bson:"locations"`
	ExperinceInYears        int64              `bson:"experince_in_years"`
	ExperinceRange          string             `bson:"experince_range"`
	NumberOfOpenings        int64              `bson:"number_of_openings"`
	Description             string             `bson:"description"`
	MandatarySkills         []string           `bson:"mandatory_skills"`
	NiceToHaveSkills        []string           `bson:"nice_to_have_skills"`
	Qualification           []string           `bson:"qualification"`
	MaxNoticePeriodInMonths int64              `bson:"max_notice_period_in_months"`
	BudgetRange             string             `bson:"budget_range"`
	BudgetInINR             int64              `bson:"budget_in_inr"`
	StartDate               string             `bson:"start_date"`
	EndDate                 string             `bson:"end_date"`
	Status                  JobStatus          `bson:"status"`
	Client                  string             `bson:"client"`
	MetaTags                string             `bson:"meta_tags"`
	JobAnalytics            []JobStats         `bson:"job_analytics"`
	ManagerId               string             `bson:"manager_id"`
	ManagerName             string             `bson:"manager_name"`
	RecruiterId             string             `bson:"recruiter_id"`
	RecruiterName           string             `bson:"recruiter_name"`
	MarkedDeleted           bool               `bson:"marked_deleted"`
	IsLive                  bool               `bson:"is_live"`
	JDURL                   string             `bson:"jd_url"`
	ImageURL                string             `bson:"image_url"`
	NaukriT                 string             `bson:"naukri_t"`
	Remarks                 string             `bson:"remarks"`
	CreatedAt               int64              `bson:"created_at"`
	UpdatedAt               int64              `bson:"updated_at"`
}
