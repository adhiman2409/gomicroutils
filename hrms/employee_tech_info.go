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
	ProjectName             string    `bson:"project_name"`
	ProjectType             string    `bson:"project_type"`
	IsCurrentProject        bool      `bson:"is_current_project"`
	IsBillable              bool      `bson:"is_billable"`
	StartDate               time.Time `bson:"start_date"`
	EndDate                 time.Time `bson:"end_date"`
	Description             string    `bson:"description"`
	RoleAndResponsibilities string    `bson:"role_and_responsibilities"`
	CompanyName             string    `bson:"company_name"`
	ClientName              string    `bson:"client_name"`
	ClientManagerName       string    `bson:"client_manager_name"`
	ClientLeadName          string    `bson:"client_lead_name"`
	Skills                  []string  `bson:"s_skills"`
	URL                     string    `bson:"url"`
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
	URL                 string `bson:"url"`
	CreatedBy           string `bson:"created_by"`
	CreatorDesignation  string `bson:"creator_designation"`
}

type SeparationDocument struct {
	DocumentName string `bson:"document_name"`
	DocumentURL  string `bson:"document_url"`
}

type SeparationInfo struct {
	ResignationDate                   string               `bson:"resignation_date"`
	ResignationStatus                 string               `bson:"resignation_status"`
	EmployeeRemarks                   string               `bson:"employee_remarks"`
	ReportingManagerRemarks           string               `bson:"reporting_manager_remarks"`
	HRRemarks                         string               `bson:"hr_remarks"`
	PrimaryApproverId                 string               `bson:"primary_approver_id"`
	PrimaryApproverName               string               `bson:"primary_approver_name"`
	IsAcceptedByPrimaryApprover       bool                 `bson:"is_accepted_by_primary_approver"`
	PrimaryApproverAcceptanceDate     string               `bson:"primary_approver_acceptance_date"`
	SecondaryApproverId               string               `bson:"secondary_approver_id"`
	SecondaryApproverName             string               `bson:"secondary_approver_name"`
	SecondaryApproverAcceptanceDate   string               `bson:"secondary_approver_acceptance_date"`
	IsAcceptedBySecondaryApprover     bool                 `bson:"is_accepted_by_secondary_approver"`
	ResignationAcceptanceDate         string               `bson:"resignation_acceptance_date"`
	KnowledgeTransferStartDate        string               `bson:"knowledge_transfer_start_date"`
	KnowledgeTransferEndDate          string               `bson:"knowledge_transfer_end_date"`
	IsKnowledgeTransferCompleted      bool                 `bson:"is_knowledge_transfer_completed"`
	KnowledgeTransferStatus           string               `bson:"knowledge_transfer_status"`
	KnowledgeTransferOwnerId          string               `bson:"knowledge_transfer_owner_id"`
	KnowledgeTransferOwnerName        string               `bson:"knowledge_transfer_owner_name"`
	KnowledgeTransferOwnerDesignation string               `bson:"knowledge_transfer_owner_designation"`
	KnowledgeTransferOwnerRemarks     string               `bson:"knowledge_transfer_owner_remarks"`
	IsNoticePeriodServed              bool                 `bson:"is_notice_period_served"`
	NoticePeriodStartDate             string               `bson:"notice_period_start_date"`
	NoticePeriodEndDate               string               `bson:"notice_period_end_date"`
	IsNoticePeriodExtended            bool                 `bson:"is_notice_period_extended"`
	NoticePeriodExtensionDate         string               `bson:"notice_period_extension_date"`
	IsNoticePeriodBuyoutRequested     bool                 `bson:"is_notice_period_buyout_requested"`
	NoticePeriodBuyoutDays            string               `bson:"notice_period_buyout_days"`
	NoticePeriodBuyoutStatus          string               `bson:"notice_period_buyout_staus"`
	NoticePeriodBuyoutRemarks         string               `bson:"notice_period_buyout_remarks"`
	NoticePeriodBuyoutAcceptanceDate  string               `bson:"notice_period_buyout_acceptance_date"`
	NoticePeriodBuyoutRejectionDate   string               `bson:"notice_period_buyout_rejection_date"`
	NoticePeriodBuyoutAcceptedBy      string               `bson:"notice_period_buyout_accepted_by"`
	NoticePeriodBuyoutRejectedBy      string               `bson:"notice_period_buyout_rejected_by"`
	IsExitInterviewScheduled          bool                 `bson:"is_exit_interview_scheduled"`
	ExitInterviewDate                 string               `bson:"exit_interview_date"`
	ExitInterviewerId                 string               `bson:"exit_interviewer_id"`
	ExitInterviewerName               string               `bson:"exit_interviewer_name"`
	ExitInterviewEmployeeRemarks      string               `bson:"exit_interview_employee_remarks"`
	ExitInterviewerRemarks            string               `bson:"exit_interviewer_remarks"`
	SeparationDocuments               []SeparationDocument `bson:"separation_documents"`
	RetentionRequestDate              string               `bson:"resignation_request_date"`
	EmployeeRetentionRemarks          string               `bson:"employee_retention_remarks"`
	ReportingManagerRetentionRemarks  string               `bson:"reporting_manager_retention_remarks"`
	HRRetentionRemarks                string               `bson:"hr_retention_remarks"`
	IsRetentionAcceptedByHR           bool                 `bson:"is_retention_accepted_by_hr"`
	RetentionAcceptanceDate           string               `bson:"retention_acceptance_date"`
	ExitDate                          string               `bson:"exit_date"`
}

type QuestionAnswer struct {
	Domain              string    `bson:"domain"`
	Department          string    `bson:"department"`
	Designation         string    `bson:"designation"`
	EmployeeId          string    `bson:"employee_id"`
	Title               string    `bson:"title"`
	Description         string    `bson:"description"`
	IsMandatory         bool      `bson:"is_mandatory"`
	InputType           InputType `bson:"input_type"`
	MinimumScore        string    `bson:"minimum_score"`
	MaximumScore        string    `bson:"maximum_score"`
	EmployeeScore       string    `bson:"employee_score"`
	ManagerScore        string    `bson:"manager_score"`
	MultiSelectOptions  []string  `bson:"multi_select_options"`
	SingleSelectOptions []string  `bson:"single_select_options"`
	BinaryOptions       []string  `bson:"binary_options"`
	SelectedOptions     []string  `bson:"selected_options"`
	EmployeeRemarks     string    `bson:"employee_remarks"`
	ManagerRemarks      string    `bson:"manager_remarks"`
}

type Rating struct {
	Title           string `bson:"title"`
	Description     string `bson:"description"`
	MinimumScore    string `bson:"minimum_score"`
	MaximumScore    string `bson:"maximum_score"`
	EmployeeScore   string `bson:"employee_score"`
	ManagerScore    string `bson:"manager_score"`
	EmployeeRemarks string `bson:"employee_remarks"`
	ManagerRemarks  string `bson:"manager_remarks"`
}

type HighLight struct {
	Title           string `bson:"title"`
	Description     string `bson:"description"`
	EmployeeRemarks string `bson:"employee_remarks"`
	ManagerRemarks  string `bson:"manager_remarks"`
}

type HRFeedback struct {
	IsAppericiationEmails bool   `bson:"is_appericiation_emails"`
	IsEscalationEmails    bool   `bson:"is_escalation_emails"`
	URL                   string `bson:"url"`
	Remarks               string `bson:"remarks"`
}

type Appraisal struct {
	Year                           string           `bson:"year"`
	AppraisalStartDate             string           `bson:"appraisal_start_date"`
	AppraisalEndDate               string           `bson:"appraisal_end_date"`
	AppraisalStatus                string           `bson:"appraisal_status"`
	IsSubmittedByEmployee          bool             `bson:"is_submitted_by_employee"`
	IsSubmittedByManager           bool             `bson:"is_submitted_by_manager"`
	IsSubmittedByHR                bool             `bson:"is_submitted_by_hr"`
	IsSubmittedByClient            bool             `bson:"is_submitted_by_client"`
	EmployeeId                     string           `bson:"employee_id"`
	EmployeeName                   string           `bson:"employee_name"`
	EmployeeDesignation            string           `bson:"employee_designation"`
	EmployeeDepartment             string           `bson:"employee_department"`
	EmployeeLocation               string           `bson:"employee_location"`
	EmployeeDateOfJoining          string           `bson:"employee_date_of_joining"`
	EmployeeTenureInDays           string           `bson:"employee_tenure_in_days"`
	EmployeeAttendancePercentage   string           `bson:"employee_attendance_percentage"`
	EmployeeEngineeringManagerName string           `bson:"employee_engineering_manager_name"`
	EmployeeEngineeringManagerId   string           `bson:"employee_engineering_manager_id"`
	EmployeeClientManagerName      string           `bson:"employee_client_manager_name"`
	EmployeeClientManagerId        string           `bson:"employee_client_manager_id"`
	EmployeeHRManagerName          string           `bson:"employee_hr_manager_name"`
	EmployeeHRManagerId            string           `bson:"employee_hr_manager_id"`
	Projects                       []Project        `bson:"projects"`
	KeySkills                      []Rating         `bson:"key_skills"`
	AssesmentsmentRating           []Rating         `bson:"assessment_rating"`
	HighLights                     []HighLight      `bson:"highlights"`
	QuestionAnswers                []QuestionAnswer `bson:"question_answers"`
	Achivements                    []Achivement     `bson:"achivements"`
	Awards                         []Award          `bson:"awards"`
	HRFeedbacks                    []HRFeedback     `bson:"hr_feedbacks"`
	OverallEmployeeRating          string           `bson:"overall_employee_rating"`
	OverallManagerRating           string           `bson:"overall_manager_rating"`
	OverallEmployeeRemarks         string           `bson:"overall_employee_remarks"`
	OverallManagerRemarks          string           `bson:"overall_manager_remarks"`
	GoalsAndExpectionsForNextYear  string           `bson:"goals_and_expections_for_next_year"`
	CurrentCTC                     float32          `bson:"current_ctc"`
	CurrentPLBonusAmount           float32          `bson:"current_pl_bonus_amount"`
	NewCTC                         float32          `bson:"new_ctc"`
	IncrementPercentage            float32          `bson:"increment_percentage"`
	IncrementAmount                float32          `bson:"increment_amount"`
	PromotionDesignation           string           `bson:"promotion_designation"`
	ApprovedPLBonusAmount          float32          `bson:"approved_pl_bonus_amount"`
	ApprovedPLBonusPercentage      float32          `bson:"approved_pl_bonus_percentage"`
	DateOfCurrentDesignation       string           `bson:"date_of_current_designation"`
	TotalLeavesAvailed             float32          `bson:"total_leaves_availed"`
	TotalYearOfExperience          float32          `bson:"total_year_of_experience"`
	TotalYearOfRelevantExperience  float32          `bson:"total_year_of_relevant_experience"`
	DesignationAtTheTimeOfJoining  string           `bson:"designation_at_the_time_of_joining"`
}

type EmployeeTechInfo struct {
	ID                     primitive.ObjectID `bson:"_id"`
	EmployeeId             string             `bson:"employee_id"`
	About                  string             `bson:"about"`
	TotalExperinceInMonths int                `bson:"total_experince_in_months"`
	ReleaseDate            time.Time          `bson:"release_date"`
	PSkill                 string             `bson:"p_skill"`
	IsAppraisalActive      bool               `bson:"is_appraisal_active"`
	SSkills                []string           `bson:"s_skills"`
	Experiences            []Experience       `bson:"experiences"`
	Projects               []Project          `bson:"projects"`
	Achivements            []Achivement       `bson:"achivements"`
	Awards                 []Award            `bson:"awards"`
	Feedbacks              []Feedback         `bson:"feedbacks,omitempty"`
	OldAppraisals          []Appraisal        `bson:"old_appraisals,omitempty"`
	ActiveAppraisal        Appraisal          `bson:"active_appraisal,omitempty"`
	Separations            []SeparationInfo   `bson:"separations"`
	IsSeparationInfoLocked bool               `bson:"is_separation_info_locked"`
	IsProfileEditingLocked bool               `bson:"is_profile_editing_locked"`
}

type InputType int

const (
	Text InputType = iota + 1
	Binary
	StarRating
	SingleSelect
	MultiSelect
)

func (r InputType) String() string {
	return [...]string{"Text", "Binary", "StarRating", "SingleSelect", "MultiSelect"}[r-1]
}

func (r InputType) EnumIndex() int {
	return int(r)
}

func GetAllInputTypes() []string {
	return []string{"Text", "Binary", "StarRating", "SingleSelect", "MultiSelect"}
}

func InputTypeFromString(s string) InputType {
	if s == "Text" {
		return Text
	} else if s == "Binary" {
		return Binary
	} else if s == "StarRating" {
		return StarRating
	} else if s == "SingleSelect" {
		return SingleSelect
	} else {
		return MultiSelect
	}
}
