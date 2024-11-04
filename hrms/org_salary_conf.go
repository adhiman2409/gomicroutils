package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgSalaryComponent struct {
	ID                        primitive.ObjectID `bson:"_id"`
	ComponentName             string             `bson:"component_name"`
	Abbrivation               string             `bson:"abbrivation"`
	Description               string             `bson:"description"`
	IsTaxableComponent        bool               `bson:"is_taxable_component"`
	IsCreditComponent         bool               `bson:"is_credit_component"`
	Formula                   string             `bson:"formula"`
	Cycle                     string             `bson:"cycle"` //once, monthly, Biannual, Annual
	IsPaidDayDepandant        bool               `bson:"is_paid_day_depandant"`
	AmountBasedOnFormula      bool               `bson:"amount_based_on_formula"`
	AbbrivationsUsedInFormula []string           `bson:"abbrivations_used_in_formula"`
	CalculatedAmount          float32            `bson:"calculated_amount"`
	StaticAmount              float32            `bson:"static_amount"`
	FinalAmount               float32            `bson:"final_amount"`
}

type OrgSalaryComponentGroup struct {
	ID               primitive.ObjectID   `bson:"_id"`
	Name             string               `bson:"name"`
	SerialNumber     int                  `bson:"serial_number"`
	SalaryComponents []OrgSalaryComponent `bson:"salary_components"`
	Formula          string               `bson:"formula"`
	CalculatedAmount float32              `bson:"calculated_amount"`
	TotalAmount      float32              `bson:"total_amount"`
}

type OrgSalaryStructure struct {
	ID                    primitive.ObjectID        `bson:"_id"`
	OrgName               string                    `bson:"org_name"`
	Name                  string                    `bson:"name"`
	SalaryComponentGroups []OrgSalaryComponentGroup `bson:"salary_component_groups"`
}

type OfferLetter struct {
	ID                          primitive.ObjectID `bson:"_id"`
	OrgName                     string             `bson:"org_name"`
	OfferLetterDate             string             `bson:"offer_letter_date"`
	OfferExpiryDate             string             `bson:"offer_expiry_date"`
	OfferedJoiningDate          string             `bson:"offered_joining_date"`
	DocumentNumber              string             `bson:"document_number"`
	CandidateId                 string             `bson:"candidate_id"`
	CandidateName               string             `bson:"candidate_name"`
	CandidateEmail              string             `bson:"candidate_email"`
	CandidatePhone              string             `bson:"candidate_phone"`
	Designation                 string             `bson:"designation"`
	CTC                         float64            `bson:"ctc"`
	CTCInWords                  string             `bson:"ctc_in_words"`
	CTCUnits                    string             `bson:"ctc_units"`
	IsWorkFromHome              bool               `bson:"is_work_from_home"`
	WorkLocation                string             `bson:"work_location"`
	WorkAddress                 string             `bson:"work_address"`
	IncludeBonus                bool               `bson:"include_bonus"`
	BonusAmount                 float64            `bson:"bonus_amount"`
	BonusTitle                  string             `bson:"bonus_title"`
	BonusType                   string             `bson:"type"`
	BonusPaymentCycle           string             `bson:"bonus_payment_cycle"`
	IncludeTravelAllowance      bool               `bson:"include_travel_allowance"`
	TravelAllowanceAmount       float64            `bson:"travel_allowance_amount"`
	TravelAllowanceTitle        string             `bson:"travel_allowance_title"`
	SalaryStructureName         string             `bson:"salary_structure_name"`
	OfferTemplateName           string             `bson:"offer_template_path"`
	CurrentStatus               string             `bson:"current_status"`
	AcceptanceDate              string             `bson:"acceptance_date"`
	CandidateJoiningDate        string             `bson:"candidate_joining_date"`
	CandidateJoinedDate         string             `bson:"candidate_joined_date"`
	ReviewerId                  string             `bson:"reviewer_id"`
	ReviewerName                string             `bson:"reviewer_name"`
	IsReviewed                  bool               `bson:"is_reviewed"`
	ReviewedDate                string             `bson:"reviewed_date"`
	ReviewRemarks               []string           `bson:"review_remarks"`
	RecruiterId                 string             `bson:"recruiter_id"`
	RecruiterName               string             `bson:"recruiter_name"`
	RefererId                   string             `bson:"referer_id"`
	RefererName                 string             `bson:"referer_name"`
	ReferalAmount               float64            `bson:"referal_amount"`
	ReferalEligibilityAfterDays int                `bson:"referal_eligibility_after_days"`
	ReferalAmountPaid           bool               `bson:"referal_amount_paid"`
	DocumentPublished           bool               `bson:"document_published"`
	DocumentPublishedDate       bool               `bson:"document_published_date"`
	DocumentLocked              bool               `bson:"document_locked"`
	IsRevisedDocument           bool               `bson:"is_revised_document"`
	OrignalDocumentNumber       string             `bson:"orignal_document_number"`
	DiscardedDocument           bool               `bson:"discarded_document"`
	NewRevisedDocumentNumber    string             `bson:"new_revised_document_number"`
	DraftDocumentPath           string             `bson:"draft_document_path"`
	PublishedDocumentPath       string             `bson:"published_document_path"`
	AcceptedDocumentPath        string             `bson:"accepted_document_path"`
}
