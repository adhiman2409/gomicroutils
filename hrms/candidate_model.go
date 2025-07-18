package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryComponentType string

const (
	FixedSalaryComponent    SalaryComponentType = "fixed"
	VariableSalaryComponent SalaryComponentType = "variable"
	BenefitSalaryComponent  SalaryComponentType = "benefit"
	OneTimeComponent        SalaryComponentType = "one_time"
)

type PaymentFrequency string

const (
	Monthly    PaymentFrequency = "monthly"
	Quarterly  PaymentFrequency = "quarterly"
	HalfYearly PaymentFrequency = "half_yearly"
	Yearly     PaymentFrequency = "yearly"
)

type CandidatePersonalInfo struct {
	FullName      string `bson:"full_name"`
	Gender        string `bson:"gender"`
	DOB           string `bson:"dob"`
	PhoneNumber   string `bson:"phone_number"`
	PersonalEmail string `bson:"personal_email"`
	AddressLine1  string `bson:"address_line_1"`
	AddressLine2  string `bson:"address_line_2"`
	City          string `bson:"city"`
	State         string `bson:"state"`
	Country       string `bson:"country"`
	Pincode       string `bson:"pincode"`
}

type CandidateProfessionalInfo struct {
	EmployeeType   string `bson:"employee_type"`
	Designation    string `bson:"designation"`
	Department     string `bson:"department"`
	EmailId        string `bson:"email_id"`
	OfficeLocation string `bson:"office_location"`
	WorkLocation   string `bson:"work_location"`
}

type SalaryComponent struct {
	ComponentName     string              `bson:"component_name"`
	ComponentType     SalaryComponentType `bson:"component_type"` // e.g., "fixed", "variable", "benefit"
	AnnualAmount      float32             `bson:"annual_amount"`
	PaymentFrequency  PaymentFrequency    `bson:"payment_frequency"`   // e.g., "monthly", "quarterly", "yearly"
	IsDrivedComponent bool                `bson:"is_drived_component"` // Indicates if the component is derived from CTC
	DrivedFrom        []string            `bson:"drived_from"`         // Name of the components from which this is derived, if applicable
	DrivedEquation    string              `bson:"drived_equation"`     // Factor to derive the component from CTC
}

type SalaryAnnexure struct {
	TemplateName       string            `bson:"template_name"`
	DocumentNumber     string            `bson:"document_number"`
	DocumentDate       string            `bson:"document_date"`
	TotalCostToCompany float32           `bson:"total_cost_to_company"`
	Components         []SalaryComponent `bson:"components"` // List of salary components
}

type OfferCoverLetter struct {
	TemplateName   string `bson:"template_name"`
	DocumentNumber string `bson:"document_number"`
	DocumentDate   string `bson:"document_date"`
}

type CandidateOfferLetter struct {
	DocumentNumber      string           `bson:"document_number"`
	DocumentDate        string           `bson:"document_date"`
	OfferCoverLetter    OfferCoverLetter `bson:"cover_letter"`
	SalaryAnnexure      SalaryAnnexure   `bson:"salary_annexure"`
	ProposedJoiningDate string           `bson:"proposed_joining_date"`
	AcceptedJoiningDate string           `bson:"accepted_joining_date"`
	ValidFrom           string           `bson:"valid_from"`
	ValidTo             string           `bson:"valid_to"`
	OfferStatus         string           `bson:"offer_status"`
	OfferStatusRemarks  string           `bson:"offer_status_remarks"`
	CreatedAt           string           `bson:"created_at"`
	UpdatedAt           string           `bson:"updated_at"`
	CreatedBy           string           `bson:"created_by"`
}

type CandidateDetails struct {
	Id               primitive.ObjectID        `bson:"_id"`
	CandidateId      string                    `bson:"candidate_id"`
	JobId            string                    `bson:"job_id"`
	Status           CandidateStatus           `bson:"status"`
	PersonalInfo     CandidatePersonalInfo     `bson:"candidate_personal_info"`
	ProfessionalInfo CandidateProfessionalInfo `bson:"candidate_professional_info"`
	OfferLetter      CandidateOfferLetter      `bson:"candidate_offer_letter"`
	DocumentList     DocumentList              `bson:"document_list"`
	ApproverId       string                    `bson:"approver_id"`
	ApproverName     string                    `bson:"approver_name"`
	ApproverEmail    string                    `bson:"approver_email"`
	Remarks          []string                  `bson:"remarks"`
	Country          string                    `bson:"country"`
	State            string                    `bson:"state"`
	TimeZone         string                    `bson:"time_zone"`
	CreatedBy        string                    `bson:"created_by"`
	CreatedAt        time.Time                 `bson:"created_at"`
}
