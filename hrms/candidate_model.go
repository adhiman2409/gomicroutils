package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	JoiningDate    string `bson:"joining_date"`
	OfficeLocation string `bson:"office_location"`
	WorkLocation   string `bson:"work_location"`
}

type AnnualComponents struct {
	FixedSalary            float32 `bson:"fixed_salary"`
	BasicSalary            float32 `bson:"basic_salary"`
	HRA                    float32 `bson:"hra"`
	SpecialPay             float32 `bson:"special_pay"`
	LTA                    float32 `bson:"lta"`
	BooksAndPeriodicals    float32 `bson:"books_and_periodicals"`
	BroadbandAndMobile     float32 `bson:"broadband_and_mobile"`
	FoodCoupons            float32 `bson:"food_coupons"`
	GrossSalary            float32 `bson:"gross_salary"`
	EmployeeContributionPF float32 `bson:"employee_contribution_pf"`
	EmployerContributionPF float32 `bson:"employer_contribution_pf"`
	IsOptedforCorporateNPS bool    `bson:"is_opted_for_corporate_nps"`
	CorporateNPS           float32 `bson:"corporate_nps"`
	Gratuity               float32 `bson:"gratuity"`
	LeaveEncashment        float32 `bson:"leave_encashment"`
	MedicalInsurance       float32 `bson:"medical_insurance"`
	PerformanceBonus       float32 `bson:"performance_bonus"`
	RetentionBonus         float32 `bson:"retention_bonus"`
	NetPay                 float32 `bson:"net_pay"`
	TotalVariableCTC       float32 `bson:"total_variable_ctc"`
	TotalCompanyBenefits   float32 `bson:"total_company_benefits"`
	TotalCostToCompany     float32 `bson:"total_cost_to_company"`
}

type MonthlyComponents struct {
	FixedCTC               float32 `bson:"fixed_salary"`
	BasicSalary            float32 `bson:"basic_salary"`
	HRA                    float32 `bson:"hra"`
	SpecialPay             float32 `bson:"special_pay"`
	LTA                    float32 `bson:"lta"`
	BooksAndPeriodicals    float32 `bson:"books_and_periodicals"`
	BroadbandAndMobile     float32 `bson:"broadband_and_mobile"`
	FoodCoupons            float32 `bson:"food_coupons"`
	GrossSalary            float32 `bson:"gross_salary"`
	EmployeeContributionPF float32 `bson:"employee_contribution_pf"`
	CorporateNPS           float32 `bson:"corporate_nps"`
	NetPay                 float32 `bson:"net_pay"`
}

type OneTimeComponents struct {
	JoiningBonus float32 `bson:"joining_bonus"`
	Relocation   float32 `bson:"relocation"`
}

type AnnexureTemplate struct {
	TemplateName      string            `bson:"template_name"`
	DocumentNumber    string            `bson:"document_number"`
	DocumentDate      string            `bson:"document_date"`
	AnnualComponents  AnnualComponents  `bson:"annual_components"`
	MonthlyComponents MonthlyComponents `bson:"monthly_components"`
	OneTimeComponents OneTimeComponents `bson:"one_time_components"`
}

type OfferCoverLetterTemplate struct {
	TemplateName   string `bson:"template_name"`
	DocumentNumber string `bson:"document_number"`
	DocumentDate   string `bson:"document_date"`
	ValidFrom      string `bson:"valid_from"`
	ValidTo        string `bson:"valid_to"`
}

type CandidateOfferLetter struct {
	DocumentNumber           string                   `bson:"document_number"`
	DocumentDate             string                   `bson:"document_date"`
	OfferCoverLetterTemplate OfferCoverLetterTemplate `bson:"cover_letter_template"`
	AnnexureTemplate         AnnexureTemplate         `bson:"salary_annexure_template"`
	ValidFrom                string                   `bson:"valid_from"`
	ValidTo                  string                   `bson:"valid_to"`
	OfferStatus              string                   `bson:"offer_status"`
	OfferStatusRemarks       string                   `bson:"offer_status_remarks"`
	CreatedAt                string                   `bson:"created_at"`
	UpdatedAt                string                   `bson:"updated_at"`
	CreatedBy                string                   `bson:"created_by"`
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
	CreatedBy        string                    `bson:"created_by"`
	CreatedAt        time.Time                 `bson:"created_at"`
}
