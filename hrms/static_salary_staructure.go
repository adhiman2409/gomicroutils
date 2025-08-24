package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IndianSalaryStructure struct {
	SSID                                    string    `bson:"ssid"`
	FinancialYearStartDate                  string    `bson:"financial_year_start_date"`
	FinancialYearEndDate                    string    `bson:"financial_year_end_date"`
	EffectiveFromDate                       string    `bson:"effective_from_date"`
	IsExpired                               bool      `bson:"is_expired"`
	ExpiredOn                               string    `bson:"expired_on"`
	EmployeeId                              string    `bson:"employee_id"`
	EmployeeName                            string    `bson:"employee_name"`
	EmployeeEmail                           string    `bson:"employee_email"`
	EmployeeJoiningDate                     string    `bson:"employee_joining_date"`
	Gender                                  string    `bson:"gender"`
	Designation                             string    `bson:"designation"`
	Department                              string    `bson:"department"`
	EmploymentStatus                        string    `bson:"employment_status"`
	OrganizationDefaultTaxRegime            string    `bson:"organization_default_tax_regime"`
	EmployeeAdoptedTaxRegime                string    `bson:"employee_adopted_tax_regime"`
	State                                   string    `bson:"state"`
	DeliveryLocation                        string    `bson:"delivery_location"`
	PAN                                     string    `bson:"pan"`
	UAN                                     string    `bson:"uan"`
	UID                                     string    `bson:"uid"`
	PFEstablishmentCode                     string    `bson:"pf_establishment_code"`
	BankName                                string    `bson:"bank_name"`
	BankBranch                              string    `bson:"bank_branch"`
	AccountNumber                           string    `bson:"account_number"`
	IFSC                                    string    `bson:"ifsc"`
	ESICApplicable                          bool      `bson:"esic_applicable"`
	LWFApplicable                           bool      `bson:"lwf_applicable"`
	AnnualFixedCTC                          float32   `bson:"annual_fixed_ctc"`
	MonthlyFixedCTC                         float32   `bson:"monthly_fixed_ctc"`
	AnnualEmployerContributionPF            float32   `bson:"annual_employer_contribution_pf"`
	MonthlyEmployerContributionPF           float32   `bson:"monthly_employer_contribution_pf"`
	AnnualEmployerContributionESIC          float32   `bson:"annual_employer_contribution_esic"`
	MonthlyEmployerContributionESIC         float32   `bson:"monthly_employer_contribution_esic"`
	AnnualEmployerContributionLWF           float32   `bson:"annual_employer_contribution_lwf"`
	MonthlyEmployerContributionLWF          float32   `bson:"monthly_employer_contribution_lwf"`
	AnnualMedicalInsurancePremium           float32   `bson:"annual_medical_insurance_premium"`
	MonthlyMedicalInsurancePremium          float32   `bson:"monthly_medical_insurance_premium"`
	AnnualEmployeeContributionPF            float32   `bson:"annual_employee_contribution_pf"`
	MonthlyEmployeeContributionPF           float32   `bson:"monthly_employee_contribution_pf"`
	AnnualEmployeeContributionESIC          float32   `bson:"annual_employee_contribution_esic"`
	MonthlyEmployeeContributionESIC         float32   `bson:"monthly_employee_contribution_esic"`
	AnnualEmployeeContributionLWF           float32   `bson:"annual_employee_contribution_lwf"`
	MonthlyEmployeeContributionLWF          float32   `bson:"monthly_employee_contribution_lwf"`
	AnnualGrossSalary                       float32   `bson:"annual_gross_salary"`
	MonthlyGrossSalary                      float32   `bson:"monthly_gross_salary"`
	AnnualBasicSalary                       float32   `bson:"annual_basic_salary"`
	MonthlyBasicSalary                      float32   `bson:"monthly_basic_salary"`
	AnnualHRA                               float32   `bson:"annual_hra"`
	MonthlyHRA                              float32   `bson:"monthly_hra"`
	AnnualLTA                               float32   `bson:"annual_lta"`
	MonthlyLTA                              float32   `bson:"monthly_lta"`
	AnnualBooksAndPeriodicals               float32   `bson:"annual_books_and_periodicals"`
	MonthlyBooksAndPeriodicals              float32   `bson:"monthly_books_and_periodicals"`
	AnnualBroadbandAndMobile                float32   `bson:"annual_broadband_and_mobile"`
	MonthlyBroadbandAndMobile               float32   `bson:"monthly_broadband_and_mobile"`
	AnnualFoodCoupons                       float32   `bson:"annual_food_coupons"`
	MonthlyFoodCoupons                      float32   `bson:"monthly_food_coupons"`
	AnnualStatutoryBonus                    float32   `bson:"annual_statutory_bonus"`
	MonthlyStatutoryBonus                   float32   `bson:"monthly_statutory_bonus"`
	AnnualGratuity                          float32   `bson:"annual_gratuity"`
	MonthlyGratuity                         float32   `bson:"monthly_gratuity"`
	IsOptedforCorporateNPS                  bool      `bson:"is_opted_for_corporate_nps"`
	AnnualCorporateNPS                      float32   `bson:"annual_corporate_nps"`
	MonthlyCorporateNPS                     float32   `bson:"monthly_corporate_nps"`
	AnnualSpecialAllowance                  float32   `bson:"annual_special_allowance"`
	MonthlySpecialAllowance                 float32   `bson:"monthly_special_allowance"`
	AnnualPerformanceBonus                  float32   `bson:"annual_performance_bonus"`
	JoiningBonus                            float32   `bson:"joining_bonus"`
	JoiningBonusPaymentFrequency            Frequency `bson:"joining_bonus_payment_frequency"`
	PerformanceBonusPaymentFrequency        Frequency `bson:"performance_bonus_payment_frequency"`
	AnnualRetentionBonus                    float32   `bson:"annual_retention_bonus"`
	RetentionBonusPaymentFrequency          Frequency `bson:"retention_bonus_payment_frequency"`
	AnnualVariableCTC                       float32   `bson:"annual_variable_ctc"`
	AnnualCTC                               float32   `bson:"annual_ctc"`
	MonthlyCTC                              float32   `bson:"monthly_ctc"`
	AnnualNetPay                            float32   `bson:"annual_net_pay"`
	MonthlyNetPay                           float32   `bson:"monthly_net_pay"`
	UseOrgDefaultEmployeePFContributions    bool      `bson:"use_org_default_employee_pf_contributions"`
	EmployeePFContributionPercentage        float32   `bson:"employee_pf_contribution_percentage"`
	IsOldSalaryStructureAvailable           bool      `bson:"is_old_salary_structure_available"`
	OldSalaryStructureSSID                  string    `bson:"old_salary_structure_ssid"`
	OldAnnualFixedCTC                       float32   `bson:"old_annual_fixed_ctc"`
	OldAnnualPerformanceBonus               float32   `bson:"old_annual_performance_bonus"`
	OldAnnualRetentionBonus                 float32   `bson:"old_annual_retention_bonus"`
	OldAnnualVariableCTC                    float32   `bson:"old_annual_variable_ctc"`
	OldUseOrgDefaultEmployeePFContributions bool      `bson:"old_use_org_default_employee_pf_contributions"`
	OldEmployeePFContributionPercentage     float32   `bson:"old_employee_pf_contribution_percentage"`
	IsArrearNeedToBePaid                    bool      `bson:"is_arrear_need_to_be_paid"`
	ArrearFromDate                          time.Time `bson:"arrear_from_date"`
	ArrearToDate                            time.Time `bson:"arrear_to_date"`
	NumberOfArrearDays                      int       `bson:"number_of_arrear_days"`
	ArrearsAmount                           float32   `bson:"arrears_amount"`
	IsArrearPaid                            bool      `bson:"is_arrear_paid"`
	ArrearPaymentDate                       time.Time `bson:"arrear_payment_date"`
	Remarks                                 []string  `bson:"remarks"`
	CreatedAt                               time.Time `bson:"created_at"`
	UpdatedAt                               time.Time `bson:"updated_at"`
	CreatedBy                               string    `bson:"created_by"`
}

type IndianSalaryStructureDetails struct {
	ID                      primitive.ObjectID      `bson:"_id"`
	EmployeeId              string                  `bson:"employee_id"`
	EmployeeName            string                  `bson:"employee_name"`
	EmploymentStatus        string                  `bson:"employment_status"`
	ActiveSalaryStructure   IndianSalaryStructure   `bson:"active_salary_structure"`
	ExpiredSalaryStructures []IndianSalaryStructure `bson:"expired_salary_structures"`
	Remarks                 []string                `bson:"remarks"`
}

type SalaryStructureIdCounter struct {
	ID          primitive.ObjectID `bson:"_id"`
	EmployeeId  string             `bson:"employee_id"`
	CountryCode string             `bson:"country_code"`
	Prefix      string             `bson:"prefix"`
	Counter     int64              `bson:"counter"`
}
