package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type IncrementLetterStatus string

const (
	IncrementLetterStatusDraft    IncrementLetterStatus = "draft"
	IncrementLetterStatusApproved IncrementLetterStatus = "approved"
	IncrementLetterStatusActive   IncrementLetterStatus = "active"
	IncrementLetterStatusInactive IncrementLetterStatus = "inactive"
	IncrementLetterStatusExpired  IncrementLetterStatus = "expired"
)

type IncrementLetterCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Year    string             `bson:"year"`
	Month   string             `bson:"month"`
	Counter int64              `bson:"counter"`
}

type IncrementLetter struct {
	ID                            primitive.ObjectID    `bson:"_id"`
	DocumentNumber                string                `bson:"document_number"`
	FinancialYear                 string                `bson:"financial_year"`
	Status                        IncrementLetterStatus `bson:"status"`
	EmployeeId                    string                `bson:"employee_id"`
	EmployeeName                  string                `bson:"employee_name"`
	Designation                   string                `bson:"designation"`
	Department                    string                `bson:"department"`
	EffectiveDate                 string                `bson:"effective_date"`
	AnnualFixedCTC                float32               `bson:"annual_fixed_salary"`
	AnnualBasicSalary             float32               `bson:"annual_basic_salary"`
	AnnualHRA                     float32               `bson:"annual_hra"`
	AnnualSpecialPay              float32               `bson:"annual_special_pay"`
	AnnualLTA                     float32               `bson:"annual_lta"`
	AnnualBooksAndPeriodicals     float32               `bson:"annual_books_and_periodicals"`
	AnnualBroadbandAndMobile      float32               `bson:"annual_broadband_and_mobile"`
	AnnualFoodCoupons             float32               `bson:"annual_food_coupons"`
	AnnualGrossSalary             float32               `bson:"annual_gross_salary"`
	AnnualEmployeeContributionPF  float32               `bson:"annual_employee_contribution_pf"`
	AnnualNetPay                  float32               `bson:"annual_net_pay"`
	MonthlyFixedCTC               float32               `bson:"monthly_fixed_salary"`
	MonthlyBasicSalary            float32               `bson:"monthly_basic_salary"`
	MonthlyHRA                    float32               `bson:"monthly_hra"`
	MonthlySpecialPay             float32               `bson:"monthly_special_pay"`
	MonthlyLTA                    float32               `bson:"monthly_lta"`
	MonthlyBooksAndPeriodicals    float32               `bson:"monthly_books_and_periodicals"`
	MonthlyBroadbandAndMobile     float32               `bson:"monthly_broadband_and_mobile"`
	MonthlyFoodCoupons            float32               `bson:"monthly_food_coupons"`
	MonthlyGrossSalary            float32               `bson:"monthly_gross_salary"`
	MonthlyEmployeeContributionPF float32               `bson:"monthly_employee_contribution_pf"`
	MonthlyCorporateNPS           float32               `bson:"monthly_corporate_nps"`
	MonthlyNetPay                 float32               `bson:"monthly_net_pay"`
	AnnualPerformanceBonus        float32               `bson:"annual_performance_bonus"`
	AnnualRetentionBonus          float32               `bson:"annual_retention_bonus"`
	TotalVariableCTC              float32               `bson:"total_variable_ctc"`
	AnnualEmployerContributionPF  float32               `bson:"annual_employer_contribution_pf"`
	IsOptedforCorporateNPS        bool                  `bson:"is_opted_for_corporate_nps"`
	AnnualCorporateNPS            float32               `bson:"annual_corporate_nps"`
	AnnualGratuity                float32               `bson:"annual_gratuity"`
	AnnualLeaveEncashment         float32               `bson:"annual_leave_encashment"`
	AnnualMedicalInsurance        float32               `bson:"annual_medical_insurance"`
	TotalCompanyBenefits          float32               `bson:"total_company_benefits"`
	TotalAnnualCostToCompany      float32               `bson:"total_annual_cost_to_company"`
	OfferLetterAnnexurePath       string                `bson:"offer_letter_annexure_path"`
	CreatedAt                     string                `bson:"created_at"`
	UpdatedAt                     string                `bson:"updated_at"`
	CreatedBy                     string                `bson:"created_by"`
}
