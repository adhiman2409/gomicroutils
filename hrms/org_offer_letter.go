package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bonus struct {
	BonusAmount       float64 `bson:"bonus_amount"`
	BonusTitle        string  `bson:"bonus_title"`
	BonusType         string  `bson:"type"`
	BonusPaymentCycle string  `bson:"bonus_payment_cycle"`
}

type OfferLetter struct {
	ID                          primitive.ObjectID  `bson:"_id"`
	OrgName                     string              `bson:"org_name"`
	Department                  string              `bson:"department"`
	OfferLetterDate             string              `bson:"offer_letter_date"`
	OfferExpiryDate             string              `bson:"offer_expiry_date"`
	OfferedJoiningDate          string              `bson:"offered_joining_date"`
	DocumentNumber              string              `bson:"document_number"`
	CandidateId                 string              `bson:"candidate_id"`
	CandidateName               string              `bson:"candidate_name"`
	CandidateEmail              string              `bson:"candidate_email"`
	CandidatePhone              string              `bson:"candidate_phone"`
	Designation                 string              `bson:"designation"`
	CTC                         float64             `bson:"ctc"`
	CTCInWords                  string              `bson:"ctc_in_words"`
	CTCUnits                    string              `bson:"ctc_units"`
	IsWorkFromHome              bool                `bson:"is_work_from_home"`
	WorkLocation                string              `bson:"work_location"`
	WorkAddress                 string              `bson:"work_address"`
	IncludeBonus                bool                `bson:"include_bonus"`
	Bonus                       []Bonus             `bson:"bonus"`
	IncludeTravelAllowance      bool                `bson:"include_travel_allowance"`
	TravelAllowanceAmount       float64             `bson:"travel_allowance_amount"`
	TravelAllowanceTitle        string              `bson:"travel_allowance_title"`
	Status                      string              `bson:"status"`
	AcceptanceDate              string              `bson:"acceptance_date"`
	CandidateJoiningDate        string              `bson:"candidate_joining_date"`
	CandidateJoinedDate         string              `bson:"candidate_joined_date"`
	ApprovedDate                string              `bson:"approved_date"`
	ApproverRemarks             []string            `bson:"approver_remarks"`
	ApproverId                  string              `bson:"approver_id"`
	ApproverName                string              `bson:"approver_name"`
	IsApproved                  bool                `bson:"is_approved"`
	SignatoryId                 string              `bson:"signatory_id"`
	SignatoryName               string              `bson:"signatory_name"`
	SignatoryDesignation        string              `bson:"signatory_designation"`
	RecruiterId                 string              `bson:"recruiter_id"`
	RecruiterName               string              `bson:"recruiter_name"`
	RefererId                   string              `bson:"referer_id"`
	RefererName                 string              `bson:"referer_name"`
	ReferalAmount               float64             `bson:"referal_amount"`
	ReferalEligibilityAfterDays int                 `bson:"referal_eligibility_after_days"`
	DocumentPublished           bool                `bson:"document_published"`
	DocumentPublishedDate       bool                `bson:"document_published_date"`
	DocumentLocked              bool                `bson:"document_locked"`
	IsRevisedDocument           bool                `bson:"is_revised_document"`
	OrignalDocumentNumber       string              `bson:"orignal_document_number"`
	DiscardedDocument           bool                `bson:"discarded_document"`
	NewRevisedDocumentNumber    string              `bson:"new_revised_document_number"`
	DraftDocumentPath           string              `bson:"draft_document_path"`
	PublishedDocumentPath       string              `bson:"published_document_path"`
	AcceptedDocumentPath        string              `bson:"accepted_document_path"`
	OfferLetterAnnexure         OfferLetterAnnexure `bson:"offer_letter_annexure"`
}

type OfferLetterAnnexure struct {
	ID                            primitive.ObjectID `bson:"_id"`
	DocumentNumber                string             `bson:"document_number"`
	IsCandidate                   bool               `bson:"is_candidate"`
	CandidateId                   string             `bson:"candidate_id"`
	CandidateName                 string             `bson:"candidate_name"`
	Designation                   string             `bson:"designation"`
	Department                    string             `bson:"department"`
	EffectiveDate                 string             `bson:"effective_date"`
	AnnualFixedCTC                float32            `bson:"annual_fixed_salary"`
	AnnualBasicSalary             float32            `bson:"annual_basic_salary"`
	AnnualHRA                     float32            `bson:"annual_hra"`
	AnnualSpecialPay              float32            `bson:"annual_special_pay"`
	AnnualGrossSalary             float32            `bson:"annual_gross_salary"`
	AnnualEmployeeContributionPF  float32            `bson:"annual_employee_contribution_pf"`
	AnnualNetPay                  float32            `bson:"annual_net_pay"`
	MonthlyFixedCTC               float32            `bson:"monthly_fixed_salary"`
	MonthlyBasicSalary            float32            `bson:"monthly_basic_salary"`
	MonthlyHRA                    float32            `bson:"monthly_hra"`
	MonthlySpecialPay             float32            `bson:"monthly_special_pay"`
	MonthlyGrossSalary            float32            `bson:"monthly_gross_salary"`
	MonthlyEmployeeContributionPF float32            `bson:"monthly_employee_contribution_pf"`
	MonthlyNetPay                 float32            `bson:"monthly_net_pay"`
	AnnualPerformanceBonus        float32            `bson:"annual_performance_bonus"`
	AnnualRetentionBonus          float32            `bson:"annual_retention_bonus"`
	TotalVariableCTC              float32            `bson:"total_variable_ctc"`
	AnnualEmployerContributionPF  float32            `bson:"annual_employer_contribution_pf"`
	AnnualGratuity                float32            `bson:"annual_gratuity"`
	AnnualLeaveEncashment         float32            `bson:"annual_leave_encashment"`
	AnnualMedicalInsurance        float32            `bson:"annual_medical_insurance"`
	TotalCompanyBenefits          float32            `bson:"total_company_benefits"`
	TotalAnnualCostToCompany      float32            `bson:"total_annual_cost_to_company"`
	OfferLetterAnnexurePath       string             `bson:"offer_letter_annexure_path"`
	CreatedAt                     string             `bson:"created_at"`
	UpdatedAt                     string             `bson:"updated_at"`
	CreatedBy                     string             `bson:"created_by"`
}
