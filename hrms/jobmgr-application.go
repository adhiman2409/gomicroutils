package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EducationDetail struct {
	Degree      string `bson:"degree"`
	Institution string `bson:"institution"`
	Board       string `bson:"board"`
	Stream      string `bson:"stream"`
	Percentage  string `bson:"percentage"`
	YearOfPass  int64  `bson:"year_of_pass"`
}

type ProjectDetail struct {
	Title        string   `bson:"title"`
	Duration     string   `bson:"duration"`
	Technologies []string `bson:"technologies"`
	Description  string   `bson:"description"`
}

type InternshipDetail struct {
	Title        string   `bson:"title"`
	Duration     string   `bson:"duration"`
	Technologies []string `bson:"technologies"`
	Description  string   `bson:"description"`
}

type Skill struct {
	Name  string  `bson:"name"`
	Years float32 `bson:"years"`
}

type JobApplicationEntry struct {
	ID                         primitive.ObjectID     `bson:"_id"`
	JobTitle                   string                 `bson:"job_title"`
	JobId                      string                 `bson:"job_id"`
	ManagerId                  string                 `bson:"manager_id"`
	ManagerName                string                 `bson:"manager_name"`
	RecruiterId                string                 `bson:"recruiter_id"`
	RecruiterName              string                 `bson:"recruiter_name"`
	ApplicationId              string                 `bson:"application_id"`
	EmailId                    string                 `bson:"email_id"`
	PhoneNumber                string                 `bson:"phone_number"`
	Name                       string                 `bson:"name"`
	ExperinceInYears           float64                `bson:"experince_in_years"`
	NoticePeriodInDays         int64                  `bson:"notice_period_in_days"`
	CurrentCompany             string                 `bson:"current_company"`
	CurrentCTC                 float64                `bson:"current_ctc"`
	ExpectedCTC                float64                `bson:"expected_ctc"`
	CurrentWorkLocation        string                 `bson:"current_work_location"`
	ReadyToRelocate            bool                   `bson:"ready_to_relocate"`
	LastWorkingDay             string                 `bson:"last_working_day"`
	JoiningInWeeks             int64                  `bson:"joining_in_weeks"`
	CurrentlyServingNotice     bool                   `bson:"currently_serving_notice"`
	OfferInHand                bool                   `bson:"offer_in_hand"`
	OfferedCTC                 float64                `bson:"offered_ctc"`
	ResumeURL                  string                 `bson:"resume_url"`
	IsViewed                   bool                   `bson:"is_viewed"`
	Remarks                    string                 `bson:"remarks"`
	TimeLine                   []TimeLineInfo         `bson:"time_line"`
	ApplicationDate            string                 `bson:"application_date"`
	Status                     JobApplicationStatus   `bson:"status"`
	IsContacted                bool                   `bson:"is_contacted"`
	IsVerified                 bool                   `bson:"is_verified"`
	RegexText                  string                 `bson:"regex_text"`
	Source                     string                 `bson:"source"`
	CurrentDesignation         string                 `bson:"current_designation"`
	ProjectDetails             []ProjectDetail        `bson:"project_details"`
	InternshipDetails          []InternshipDetail     `bson:"internship_details"`
	EducationDetails           []EducationDetail      `bson:"education_details"`
	Skills                     []Skill                `bson:"skills"`
	AreaOfInterest             string                 `bson:"area_of_interest"`
	WillingToOfferCommitment   bool                   `bson:"willing_to_offer_commitment"`
	CommitmentDurationInMonths int64                  `bson:"commitment_duration_in_months"`
	AiAnalysisReport           map[string]interface{} `bson:"ai_analysis_report"`
	UpdatedAt                  int64                  `bson:"updated_at"`
}

type TimeLineInfo struct {
	Date    string `bson:"date"`
	Time    string `bson:"time"`
	Remarks string `bson:"remarks"`
	Name    string `bson:"name"`
}
