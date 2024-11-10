package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobApplication struct {
	ID                      primitive.ObjectID `bson:"_id"`
	JobTitle                string             `bson:"job_title"`
	JobId                   string             `bson:"job_id"`
	EmploymentType          string             `bson:"employment_type"`
	JobDomain               string             `bson:"job_domain"`
	Technology              string             `bson:"technology"`
	Designation             string             `bson:"designation"`
	Location                string             `bson:"location"`
	ExperinceInYears        int64              `bson:"experince_in_years"`
	ExperinceRange          string             `bson:"experince_range"`
	NumberOfOpenings        int64              `bson:"number_of_openings"`
	Description             string             `bson:"description"`
	MandatarySkills         []string           `bson:"mandatory_skills"`
	NiceToHaveSkills        []string           `bson:"nice_to_have_skills"`
	Qualification           []string           `bson:"qualification"`
	MaxNoticePeriodInMonths int64              `bson:"max_notice_period_in_months"`
	StartDate               string             `bson:"start_date"`
	EndDate                 string             `bson:"end_date"`
	CurrentStatus           string             `bson:"current_status"`
	Client                  string             `bson:"client"`
	MetaTags                string             `bson:"meta_tags"`
	ApplicationList         []string           `bson:"application_list"`
	ManagerId               string             `bson:"manager_id"`
	ManagerName             string             `bson:"manager_nmae"`
	RecruiterId             string             `bson:"recruiter_id"`
	RecruiterName           string             `bson:"recruiter_name"`
	MarkedDeleted           bool               `bson:"marked_deleted"`
	IsLive                  bool               `bson:"is_live"`
	JDURL                   string             `bson:"jd_url"`
	ImageURL                string             `bson:"image_url"`
	CreatedAt               int64              `bson:"created_at"`
	UpdatedAt               int64              `bson:"updated_at"`
}

type JobApplicationEntry struct {
	ID                     primitive.ObjectID `bson:"_id"`
	JobTitle               string             `bson:"job_title"`
	JobId                  string             `bson:"job_id"`
	ManagerId              string             `bson:"manager_id"`
	ManagerName            string             `bson:"manager_name"`
	RecruiterId            string             `bson:"recruiter_id"`
	RecruiterName          string             `bson:"recruiter_name"`
	ApplicationId          string             `bson:"application_id"`
	EmailId                string             `bson:"email_id"`
	PhoneNumber            string             `bson:"phone_number"`
	Name                   string             `bson:"name"`
	ExperinceInYears       float64            `bson:"experince_in_years"`
	NoticePeriodInDays     int64              `bson:"notice_period_in_days"`
	CurrentCompany         string             `bson:"current_company"`
	CurrentCTC             float64            `bson:"current_ctc"`
	ExpectedCTC            float64            `bson:"expected_ctc"`
	CurrentWorkLocation    string             `bson:"current_work_location"`
	ReadyToRelocate        bool               `bson:"ready_to_relocate"`
	LastWorkingDay         string             `bson:"last_working_day"`
	JoiningInWeeks         int64              `bson:"joining_in_weeks"`
	CurrentlyServingNotice bool               `bson:"currently_serving_notice"`
	OfferInHand            bool               `bson:"offer_in_hand"`
	OfferedCTC             float64            `bson:"offered_ctc"`
	ResumeURL              string             `bson:"resume_url"`
	IsViewed               bool               `bson:"is_viewed"`
	Remarks                string             `bson:"remarks"`
	TimeLine               []TimeLineInfo     `bson:"time_line"`
	ApplicationDate        string             `bson:"application_date"`
	Status                 JobAppStatus       `bson:"status"`
	IsContacted            bool               `bson:"is_contacted"`
	IsVerified             bool               `bson:"is_verified"`
	RegexText              string             `bson:"regex_text"`
	Source                 string             `bson:"source"`
	UpdatedAt              int64              `bson:"updated_at"`
}

type TimeLineInfo struct {
	Date    string `bson:"date"`
	Time    string `bson:"time"`
	Remarks string `bson:"remarks"`
	Name    string `bson:"name"`
}
