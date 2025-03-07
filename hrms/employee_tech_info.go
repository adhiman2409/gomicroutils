package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	StartYearAndMonth  string `bson:"start_year_and_month"`
	EndYearAndMonth    string `bson:"end_year_and_month"`
	ExperienceInMonths int    `bson:"experience_in_months"`
	CompanyName        string `bson:"company_name"`
	Designation        string `bson:"designation"`
}

type Project struct {
	ProjectName      string    `bson:"project_name"`
	ProjectType      string    `bson:"project_type"`
	IsCurrentProject bool      `bson:"is_current_project"`
	IsBillable       bool      `bson:"is_billable"`
	StartDate        time.Time `bson:"start_date"`
	EndDate          time.Time `bson:"end_date"`
	Description      string    `bson:"description"`
	CompanyName      string    `bson:"company_name"`
	ClientName       string    `bson:"client_name"`
	Skills           []string  `bson:"s_skills"`
	URL              string    `bson:"url"`
}

type Achivement struct {
	Year        string `bson:"year"`
	Month       string `bson:"month"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	CompanyName string `bson:"company_name"`
	ClientName  string `bson:"client_name"`
	URL         string `bson:"url"`
}

type Award struct {
	Year        string `bson:"year"`
	Month       string `bson:"month"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	CompanyName string `bson:"company_name"`
	ClientName  string `bson:"client_name"`
	URL         string `bson:"url"`
}

type Feedback struct {
	EmployeeDesignation string `bson:"employee_designation"`
	Day                 string `bson:"day"`
	Year                string `bson:"year"`
	Month               string `bson:"month"`
	Title               string `bson:"title"`
	Description         string `bson:"description"`
	CreatedBy           string `bson:"created_by"`
	CreatorDesignation  string `bson:"creator_designation"`
}

// type SeparationDocument struct {
// 	DocumentName string `bson:"document_name"`
// 	DocumentURL  string `bson:"document_url"`
// }

// type SeparationInfo struct {
// 	ResignationDate                   string               `bson:"resignation_date"`
// 	EmployeeRemarks                   string               `bson:"resignation_remarks"`
// 	ReportingManagerRemarks           string               `bson:"resignation_remarks"`
// 	HumanResourceRemarks              string               `bson:"human_resource_remarks"`
// 	ResignationAcceptanceDate         string               `bson:"resignation_acceptance_date"`
// 	KnowledgeTransferStartDate        string               `bson:"knowledge_transfer_start_date"`
// 	KnowledgeTransferEndDate          string               `bson:"knowledge_transfer_end_date"`
// 	IsKnowledgeTransferCompleted      bool                 `bson:"is_knowledge_transfer_completed"`
// 	KnowledgeTransferStatus           string               `bson:"knowledge_transfer_status"`
// 	KnowledgeTransferOwnerId          string               `bson:"knowledge_transfer_owner_id"`
// 	KnowledgeTransferOwnerName        string               `bson:"knowledge_transfer_owner_name"`
// 	KnowledgeTransferOwnerDesignation string               `bson:"knowledge_transfer_owner_designation"`
// 	KnowledgeTransferOwnerRemarks     string               `bson:"knowledge_transfer_owner_remarks"`
// 	IsNoticePeriodServed              bool                 `bson:"is_notice_period_served"`
// 	NoticePeriodStartDate             string               `bson:"notice_period_start_date"`
// 	NoticePeriodEndDate               string               `bson:"notice_period_end_date"`
// 	IsNoticePeriodExtended            bool                 `bson:"is_notice_period_extended"`
// 	NoticePeriodExtensionDate         string               `bson:"notice_period_extension_date"`
// 	IsNoticePeriodBuyoutRequested     bool                 `bson:"is_notice_period_buyout_requested"`
// 	NoticePeriodBuyoutDays            string               `bson:"notice_period_buyout_days"`
// 	NoticePeriodBuyoutStatus          string               `bson:"notice_period_buyout_staus"`
// 	NoticePeriodBuyoutRemarks         string               `bson:"notice_period_buyout_remarks"`
// 	NoticePeriodBuyoutAcceptanceDate  string               `bson:"notice_period_buyout_acceptance_date"`
// 	NoticePeriodBuyoutRejectionDate   string               `bson:"notice_period_buyout_rejection_date"`
// 	NoticePeriodBuyoutAcceptedBy      string               `bson:"notice_period_buyout_accepted_by"`
// 	NoticePeriodBuyoutRejectedBy      string               `bson:"notice_period_buyout_rejected_by"`
// 	IsExitInterviewScheduled          bool                 `bson:"is_exit_interview_scheduled"`
// 	ExitInterviewDate                 string               `bson:"exit_interview_date"`
// 	ExitInterviewerId                 string               `bson:"exit_interviewer_id"`
// 	ExitInterviewerName               string               `bson:"exit_interviewer_name"`
// 	ExitInterviewEmployeeRemarks      string               `bson:"exit_interview_employee_remarks"`
// 	ExitInterviewerRemarks            string               `bson:"exit_interviewer_remarks"`

// 	SeparationDocuments               []SeparationDocument `bson:"separation_documents"`

// 	ExitDate                          string               `bson:"exit_date"`
// }

type EmployeeTechInfo struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeId             string             `bson:"employee_id"`
	About                  string             `bson:"about"`
	TotalExperinceInMonths int                `bson:"total_experince_in_months"`
	ReleaseDate            time.Time          `bson:"release_date"`
	PSkill                 string             `bson:"p_skill"`
	SSkills                []string           `bson:"s_skills"`
	Experiences            []Experience       `bson:"experiences"`
	Projects               []Project          `bson:"projects"`
	Achivements            []Achivement       `bson:"achivements"`
	Awards                 []Award            `bson:"awards"`
	Feedbacks              []Feedback         `json:"feedbacks,omitempty"`
	IsProfileEditingLocked bool               `bson:"is_profile_editing_locked"`
}
