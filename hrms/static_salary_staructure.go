package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CountrySalaryStructure struct {
	SSID                                  string    `bson:"ssid"`
	FinancialYear                         string    `bson:"financial_year"`
	EffectiveFromDate                     string    `bson:"effective_from_date"`
	EffectiveToDate                       string    `bson:"effective_to_date"`
	IsUpcoming                            bool      `bson:"is_upcoming"`
	IsActive                              bool      `bson:"is_active"`
	IsExpired                             bool      `bson:"is_expired"`
	ExpiredOn                             string    `bson:"expired_on"`
	ExpiredSalaryStructureSSID            string    `bson:"expired_salary_structure_ssid"`
	EmployeeId                            string    `bson:"employee_id"`
	EmployeeName                          string    `bson:"employee_name"`
	EmployeeEmail                         string    `bson:"employee_email"`
	EmployeeJoiningDate                   string    `bson:"employee_joining_date"`
	Gender                                string    `bson:"gender"`
	Designation                           string    `bson:"designation"`
	Department                            string    `bson:"department"`
	EmployeeType                          string    `bson:"employee_type"`
	EmploymentStatus                      string    `bson:"employment_status"`
	OrganizationDefaultTaxRegime          string    `bson:"organization_default_tax_regime"`
	EmployeeAdoptedTaxRegime              string    `bson:"employee_adopted_tax_regime"`
	State                                 string    `bson:"state"`
	Country                               string    `bson:"country"`
	DeliveryLocation                      string    `bson:"delivery_location"`
	PAN                                   string    `bson:"pan"`
	UAN                                   string    `bson:"uan"`
	UID                                   string    `bson:"uid"`
	PFEstablishmentCode                   string    `bson:"pf_establishment_code"`
	BankName                              string    `bson:"bank_name"`
	BankBranch                            string    `bson:"bank_branch"`
	AccountNumber                         string    `bson:"account_number"`
	IFSC                                  string    `bson:"ifsc"`
	ESICApplicable                        bool      `bson:"esic_applicable"`
	LWFApplicable                         bool      `bson:"lwf_applicable"`
	AnnualFixedCTC                        float32   `bson:"annual_fixed_ctc"`
	MonthlyFixedCTC                       float32   `bson:"monthly_fixed_ctc"`
	AnnualEmployerContributionPF          float32   `bson:"annual_employer_contribution_pf"`
	MonthlyEmployerContributionPF         float32   `bson:"monthly_employer_contribution_pf"`
	AnnualEmployerContributionESIC        float32   `bson:"annual_employer_contribution_esic"`
	MonthlyEmployerContributionESIC       float32   `bson:"monthly_employer_contribution_esic"`
	AnnualEmployerContributionLWF         float32   `bson:"annual_employer_contribution_lwf"`
	MonthlyEmployerContributionLWF        float32   `bson:"monthly_employer_contribution_lwf"`
	AnnualMedicalInsurancePremium         float32   `bson:"annual_medical_insurance_premium"`
	MonthlyMedicalInsurancePremium        float32   `bson:"monthly_medical_insurance_premium"`
	AnnualEmployeeContributionPF          float32   `bson:"annual_employee_contribution_pf"`
	MonthlyEmployeeContributionPF         float32   `bson:"monthly_employee_contribution_pf"`
	AnnualEmployeeContributionESIC        float32   `bson:"annual_employee_contribution_esic"`
	MonthlyEmployeeContributionESIC       float32   `bson:"monthly_employee_contribution_esic"`
	AnnualEmployeeContributionLWF         float32   `bson:"annual_employee_contribution_lwf"`
	MonthlyEmployeeContributionLWF        float32   `bson:"monthly_employee_contribution_lwf"`
	AnnualGrossSalary                     float32   `bson:"annual_gross_salary"`
	MonthlyGrossSalary                    float32   `bson:"monthly_gross_salary"`
	AnnualBasicSalary                     float32   `bson:"annual_basic_salary"`
	MonthlyBasicSalary                    float32   `bson:"monthly_basic_salary"`
	AnnualHRA                             float32   `bson:"annual_hra"`
	MonthlyHRA                            float32   `bson:"monthly_hra"`
	AnnualLTA                             float32   `bson:"annual_lta"`
	MonthlyLTA                            float32   `bson:"monthly_lta"`
	AnnualBooksAndPeriodicals             float32   `bson:"annual_books_and_periodicals"`
	MonthlyBooksAndPeriodicals            float32   `bson:"monthly_books_and_periodicals"`
	AnnualBroadbandAndMobile              float32   `bson:"annual_broadband_and_mobile"`
	MonthlyBroadbandAndMobile             float32   `bson:"monthly_broadband_and_mobile"`
	AnnualFoodCoupons                     float32   `bson:"annual_food_coupons"`
	MonthlyFoodCoupons                    float32   `bson:"monthly_food_coupons"`
	AnnualUniformAndLaundry               float32   `bson:"annual_uniform_and_laundry"`
	MonthlyUniformAndLaundry              float32   `bson:"monthly_uniform_and_laundry"`
	AnnualConveyanceAllowance             float32   `bson:"annual_conveyance_allowance"`
	MonthlyConveyanceAllowance            float32   `bson:"monthly_conveyance_allowance"`
	AnnualStatutoryBonus                  float32   `bson:"annual_statutory_bonus"`
	MonthlyStatutoryBonus                 float32   `bson:"monthly_statutory_bonus"`
	AnnualGratuity                        float32   `bson:"annual_gratuity"`
	MonthlyGratuity                       float32   `bson:"monthly_gratuity"`
	IsOptedforCorporateNPS                bool      `bson:"is_opted_for_corporate_nps"`
	AnnualCorporateNPS                    float32   `bson:"annual_corporate_nps"`
	MonthlyCorporateNPS                   float32   `bson:"monthly_corporate_nps"`
	AnnualSpecialAllowance                float32   `bson:"annual_special_allowance"`
	MonthlySpecialAllowance               float32   `bson:"monthly_special_allowance"`
	AnnualPerformanceBonus                float32   `bson:"annual_performance_bonus"`
	JoiningBonus                          float32   `bson:"joining_bonus"`
	JoiningBonusPaymentFrequency          Frequency `bson:"joining_bonus_payment_frequency"`
	PerformanceBonusPaymentFrequency      Frequency `bson:"performance_bonus_payment_frequency"`
	AnnualRetentionBonus                  float32   `bson:"annual_retention_bonus"`
	RetentionBonusPaymentFrequency        Frequency `bson:"retention_bonus_payment_frequency"`
	AnnualVariableCTC                     float32   `bson:"annual_variable_ctc"`
	AnnualCTC                             float32   `bson:"annual_ctc"`
	MonthlyCTC                            float32   `bson:"monthly_ctc"`
	AnnualNetPay                          float32   `bson:"annual_net_pay"`
	MonthlyNetPay                         float32   `bson:"monthly_net_pay"`
	TotalGrossAmountPaidThisFinancialYear float32   `bson:"total_gross_amount_paid_this_financial_year"`
	TotalTaxPaidThisFinancialYear         float32   `bson:"total_tax_paid_this_financial_year"`
	TotalPaidMonthsThisFinancialYear      int       `bson:"total_paid_months_this_financial_year"`
	LastPaidMonthThisFinancialYear        string    `bson:"last_paid_month_this_financial_year"`
	Remarks                               string    `bson:"remarks"`
	CreatedAt                             time.Time `bson:"created_at"`
	UpdatedAt                             time.Time `bson:"updated_at"`
	CreatedBy                             string    `bson:"created_by"`
}

type ArrearPaymentInfo struct {
	IsArrearNeedToBePaid   bool      `bson:"is_arrear_need_to_be_paid"`
	OldSalaryStructureSSID string    `bson:"old_salary_structure_ssid"`
	ArrearFromDate         time.Time `bson:"arrear_from_date"`
	ArrearToDate           time.Time `bson:"arrear_to_date"`
	NumberOfArrearDays     int       `bson:"number_of_arrear_days"`
	ArrearsAmount          float32   `bson:"arrears_amount"`
	IsArrearPaid           bool      `bson:"is_arrear_paid"`
	ArrearPaymentDate      time.Time `bson:"arrear_payment_date"`
	Remarks                string    `bson:"remarks"`
}

type SalaryStructureDetails struct {
	ID                      primitive.ObjectID       `bson:"_id"`
	EmployeeId              string                   `bson:"employee_id"`
	EmployeeName            string                   `bson:"employee_name"`
	EmploymentStatus        string                   `bson:"employment_status"`
	EmployeeType            string                   `bson:"employee_type"`
	UpcomingSalaryStructure CountrySalaryStructure   `bson:"upcoming_salary_structure"`
	ActiveSalaryStructure   CountrySalaryStructure   `bson:"active_salary_structure"`
	ExpiredSalaryStructures []CountrySalaryStructure `bson:"expired_salary_structures"`
	ArrearPaymentInfo       ArrearPaymentInfo        `bson:"arrear_payment_info"`
	PaidArrearPaymentInfo   []ArrearPaymentInfo      `bson:"paid_arrear_payment_info"`
	Remarks                 []string                 `bson:"remarks"`
}

type SalaryStructureIdCounter struct {
	ID          primitive.ObjectID `bson:"_id"`
	EmployeeId  string             `bson:"employee_id"`
	CountryCode string             `bson:"country_code"`
	Prefix      string             `bson:"prefix"`
	Counter     int64              `bson:"counter"`
}

type PayrollSalaryConfig struct {
	ID                                       primitive.ObjectID    `bson:"_id"`
	EmployeeType                             string                `bson:"employee_type"`
	PayrollConfigName                        string                `bson:"payroll_config_name"`
	IsDefaultForEmployeeType                 bool                  `bson:"is_default_for_employee_type"`
	Description                              string                `bson:"description"`
	EmployeeId                               string                `bson:"employee_id"`
	EmployeeName                             string                `bson:"employee_name"`
	State                                    string                `bson:"state"`
	Country                                  string                `bson:"country"`
	DeliveryLocation                         string                `bson:"delivery_location"`
	IsFixedMonthlyPayout                     bool                  `bson:"is_fixed_monthly_payout"`
	FixedMonthlyPayoutAmount                 float32               `bson:"fixed_monthly_payout_amount"`
	IsFixedDailyPayout                       bool                  `bson:"is_fixed_daily_payout"`
	FixedDailyPayoutAmount                   float32               `bson:"fixed_daily_payout_amount"`
	DeductOnlyTDS                            bool                  `bson:"deduct_only_tds"`
	DefaultTDSPercentage                     float32               `bson:"default_tds_percentage"`
	OrganizationDefaultTaxRegime             string                `bson:"organization_default_tax_regime"` // e.g., "new", "old"
	EmployeeAdoptedTaxRegime                 string                `bson:"employee_adopted_tax_regime"`     // e.g., "new", "old"
	IsVariableSalaryPartOfTotalCTC           bool                  `bson:"is_variable_salary_part_of_total_ctc"`
	IsVariableSalaryPercentageOfTotalCTC     bool                  `bson:"is_variable_salary_percentage_of_total_ctc"`
	VariableSalaryPercentageOfTotalCTC       float32               `bson:"variable_salary_percentage_of_total_ctc"`
	FixedAnnualVariableSalary                float32               `bson:"fixed_annual_variable_salary"`
	IsOrgPFContributionIsBasedOnBasicSalary  bool                  `bson:"is_org_pf_contribution_is_based_on_basic_salary"`
	PFEmployerContributionPercentageOfBasic  float32               `bson:"pf_employer_contribution_percentage_of_basic"`
	IsEmpPFContributionIsBasedOnBasicSalary  bool                  `bson:"is_emp_pf_contribution_is_based_on_basic_salary"`
	PFEmployeeContributionPercentageOfBasic  float32               `bson:"pf_employee_contribution_percentage_of_basic"`
	PFMonthlyFixedEmployerContributionAmount float32               `bson:"pf_monthly_fixed_employer_contribution_amount"`
	PFMonthlyFixedEmployeeContributionAmount float32               `bson:"pf_monthly_fixed_employee_contribution_amount"`
	MonthlyBasicSalaryPercentageOfGross      float32               `bson:"monthly_basic_salary_percentage_odf_gross"`
	MinimumMonthlyBasicSalary                float32               `bson:"minimum_monthly_basic_salary"`
	MonthlyHRAPercentageOfBasic              float32               `bson:"monthly_hra_percentage_of_basic"`
	MonthlyStatutoryBonusPercentageOfBasic   float32               `bson:"monthly_statutory_bonus_percentage_of_basic"`
	MonthlyGratuityPercentageOfBasic         float32               `bson:"monthly_gratuity_percentage_of_basic"`
	AnnualLTA                                float32               `bson:"annual_lta"`
	MonthlyLTA                               float32               `bson:"monthly_lta"`
	AnnualBooksAndPeriodicals                float32               `bson:"annual_books_and_periodicals"`
	MonthlyBooksAndPeriodicals               float32               `bson:"monthly_books_and_periodicals"`
	AnnualBroadbandAndMobile                 float32               `bson:"annual_broadband_and_mobile"`
	MonthlyBroadbandAndMobile                float32               `bson:"monthly_broadband_and_mobile"`
	AnnualFoodCoupons                        float32               `bson:"annual_food_coupons"`
	MonthlyFoodCoupons                       float32               `bson:"monthly_food_coupons"`
	AnnualUniformAndLaundry                  float32               `bson:"annual_uniform_and_laundry"`
	MonthlyUniformAndLaundry                 float32               `bson:"monthly_uniform_and_laundry"`
	AnnualConveyanceAllowance                float32               `bson:"annual_conveyance_allowance"`
	MonthlyConveyanceAllowance               float32               `bson:"monthly_conveyance_allowance"`
	IsOptedforCorporateNPS                   bool                  `bson:"is_opted_for_corporate_nps"`
	AnnualCorporateNPSPercentageOfBasic      float32               `bson:"annual_corporate_nps_percentage_of_basic"`
	MonthlyCorporateNPSPercentageOfBasic     float32               `bson:"monthly_corporate_nps_percentage_of_basic"`
	AddRemainingAmountInSpecialAllowance     bool                  `bson:"add_remaining_amount_in_special_allowance"`
	ESICConfig                               []ESICConfig          `bson:"esi_config"`
	LWFConfig                                []LWFConfig           `bson:"lwf_config"`
	ProfessionalTax                          []ProfessionalTaxSlab `bson:"professional_tax"`
	AnnualMedicalInsurancePremium            float32               `bson:"annual_medical_insurance_premium"`
	UpdatedAt                                time.Time             `bson:"updated_at"`
	CreatedAt                                time.Time             `bson:"created_at"`
}

type ESICConfig struct {
	State                  string  `bson:"state"`
	ESIEmployerPercentage  float32 `bson:"esi_employer_percentage"`
	ESIEmployeePercentage  float32 `bson:"esi_employee_percentage"`
	ESIMaximumWage         float32 `bson:"esi_maximum_wage"`
	ESIMinimumWage         float32 `bson:"esi_minimum_wage"`
	ESIApplicableWageLimit float32 `bson:"esi_applicable_wage_limit"`
}

type LWFConfig struct {
	CityName                       string  `bson:"city_name"`
	LWFEmployerMonthlyContribution float32 `bson:"lwf_employer_monthly_contribution"`
	LWFEmployeeMonthlyContribution float32 `bson:"lwf_employee_monthly_contribution"`
}

type EmployeeTaxSavingDeclaration struct {
	EmployeeID                           string  `bson:"employee_id"`
	EmployeeName                         string  `bson:"employee_name"`
	FinancialYear                        string  `bson:"financial_year"`
	OptedTaxRegime                       string  `bson:"opted_tax_regime"`
	IncomeOtherThanSalary                float64 `bson:"income_other_than_salary"`
	HomeLoanInterest                     float64 `bson:"home_loan_interest"`
	HRAExemption                         float64 `bson:"hra_exemption"`
	Section80CDeduction                  float64 `bson:"section_80c_deduction"`
	Section80CCD2EmployerNPSContribution float64 `bson:"section_80ccd2_employer_nps_contribution"`
	Section80DHealthInsuranceForSelf     float64 `bson:"section_80d_health_insurance_for_self"`
	Section80DHealthInsuranceForFamily   float64 `bson:"section_80d_health_insurance_for_family"`
	OtherDeductions                      float64 `bson:"other_deductions"`
}

type YearlyPayrollAndTaxDetails struct {
	GrossIncome                          float64   `bson:"gross_income"`
	StandardDeduction                    float64   `bson:"standard_deduction"`
	NetIncome                            float64   `bson:"net_income"`
	IncomeOtherThanSalary                float64   `bson:"income_other_than_salary"`
	HomeLoanInterest                     float64   `bson:"home_loan_interest"`
	HRAExemption                         float64   `bson:"hra_exemption"`
	Section80CDeduction                  float64   `bson:"section_80c_deduction"`
	Section80CCD1EmployeeNPSContribution float64   `bson:"section_80ccd1_employee_nps_contribution"`
	Section80CCD2EmployerNPSContribution float64   `bson:"section_80ccd2_employer_nps_contribution"`
	Section80DHealthInsuranceForSelf     float64   `bson:"section_80d_health_insurance_for_self"`
	Section80DHealthInsuranceForParents  float64   `bson:"section_80d_health_insurance_for_parents"`
	Section80TTAInterestOnSavings        float64   `bson:"section_80tta_interest_on_savings"`
	Section80TTBInterestOnDeposits       float64   `bson:"section_80ttb_interest_on_deposits"`
	Section80EInterestOnEducationLoan    float64   `bson:"section_80e_interest_on_education_loan"`
	Section80GDonations                  float64   `bson:"section_80g_donations"`
	OtherDeductions                      float64   `bson:"other_deductions"`
	NetTaxableIncome                     float64   `bson:"net_taxable_income"`
	TaxSlabTable                         []TaxSlab `bson:"tax_slab_table"`
	TaxOnTotalIncome                     float64   `bson:"tax_on_total_income"`
	RebateUnderSection87A                float64   `bson:"rebate_under_section_87a"`
	TaxAfterRebate                       float64   `bson:"tax_after_rebate"`
	SurchargeOnTax                       float64   `bson:"surcharge_on_tax"`
	TaxWithSurcharge                     float64   `bson:"tax_with_surcharge"`
	CessOnTaxWithSurcharge               float64   `bson:"cess_on_tax_with_surcharge"`
	TotalTaxPayable                      float64   `bson:"total_tax_payable"`
}

type MonthlyEarnings struct {
	BasicSalary        float64 `bson:"basic_salary"`
	HouseRentAllowance float64 `bson:"house_rent_allowance"`
	StatutoryBonus     float64 `bson:"statutory_bonus"`
	SpecialAllowance   float64 `bson:"special_allowance"`
	OtherEarnings      float64 `bson:"other_earnings"`
	TotalGrossSalary   float64 `bson:"total_gross_salary"`
}

type MonthlyDeductions struct {
	LabourWelfareFund float64 `bson:"labour_welfare_fund"`
	ESI               float64 `bson:"esi"`
	ProvidentFund     float64 `bson:"provident_fund"`
	IncomeTax         float64 `bson:"income_tax"`
	OtherDeductions   float64 `bson:"other_deductions"`
	TotalDeductions   float64 `bson:"total_deductions"`
}

type MonthlyEmployerContributions struct {
	LabourWelfareFund          float64 `bson:"labour_welfare_fund"`
	EmployeeStateInsurance     float64 `bson:"employee_state_insurance"`
	ProvidentFund              float64 `bson:"provident_fund"`
	MedicalInsurancePremium    float64 `bson:"medical_insurance_premium"`
	Gratuity                   float64 `bson:"gratuity"`
	TotalEmployerContributions float64 `bson:"total_employer_contributions"`
}

type OneTimeMonthlyEarning struct {
	Name      string  `bson:"name"`
	Type      string  `bson:"type"` // e.g., "bonus", "arrears", "incentive", "reimbursement"
	Amount    float64 `bson:"amount"`
	Id        string  `bson:"id"`
	IsTaxable bool    `bson:"is_taxable"`
	Remarks   string  `bson:"remarks"`
}

type OneTimeMonthlyDeduction struct {
	Name      string  `bson:"name"`
	Type      string  `bson:"type"` // e.g., "loan", "advance", "penalty"
	Id        string  `bson:"id"`
	Amount    float64 `bson:"amount"`
	IsTaxable bool    `bson:"is_taxable"`
	Remarks   string  `bson:"remarks"`
}

type MonthlyPayrollAndTaxDetails struct {
	Year                         string                       `bson:"year"`
	Month                        string                       `bson:"month"`
	NumberOfDaysInMonth          int                          `bson:"number_of_days_in_month"`
	NumberOfPayableDays          int                          `bson:"number_of_payable_days"`
	SalaryProcessingStatus       SalaryStatus                 `bson:"salary_processing_status"`
	SalaryProcessingDate         string                       `bson:"salary_processing_date"`
	SalaryStructureSSID          string                       `bson:"salary_structure_ssid"`
	IsHybridSSIDMonth            bool                         `bson:"is_hybrid_ssid_month"`
	NewSSIDStartDate             string                       `bson:"new_ssid_start_date"`
	OldSSIDEndDate               string                       `bson:"old_ssid_end_date"`
	OldSalaryStructureSSID       string                       `bson:"old_salary_structure_ssid"`
	IsManuallyUpdated            bool                         `bson:"is_manually_updated"`
	MonthlyEarnings              MonthlyEarnings              `bson:"monthly_earnings"`
	TotalGrossEarnings           float64                      `bson:"total_gross_earnings"`
	MonthlyDeductions            MonthlyDeductions            `bson:"monthly_deductions"`
	TotalDeductions              float64                      `bson:"total_deductions"`
	NetSalary                    float64                      `bson:"net_salary"`
	TaxableIncome                float64                      `bson:"taxable_income"`
	MonthlyEmployerContributions MonthlyEmployerContributions `bson:"monthly_employer_contributions"`
	TotalEmployerContributions   float64                      `bson:"total_employer_contributions"`
	OneTimeMonthlyEarning        []OneTimeMonthlyEarning      `bson:"one_time_monthly_earnings"`
	OneTimeMonthlyDeduction      []OneTimeMonthlyDeduction    `bson:"one_time_monthly_deductions"`
	Remarks                      []string                     `bson:"remarks"`
	CreatedAt                    time.Time                    `bson:"created_at"`
	UpdatedAt                    time.Time                    `bson:"updated_at"`
}

type EmployeePayrollMaster struct {
	ID                                          primitive.ObjectID            `bson:"_id"`
	EmployeeID                                  string                        `bson:"employee_id"`
	EmployeeName                                string                        `bson:"employee_name"`
	JoiningDate                                 string                        `bson:"joining_date"`
	EmploymentStatus                            string                        `bson:"employment_status"`
	OptedTaxRegime                              string                        `bson:"opted_tax_regime"`
	FinancialYear                               string                        `bson:"financial_year"`
	AssessmentYear                              string                        `bson:"assessment_year"`
	CurrentFinancialYearCTC                     float64                       `bson:"current_financial_year_ctc"`
	CurrentFinancialYearGrossSalary             float64                       `bson:"current_financial_year_gross_salary"`
	TotalGrossAmountPaidThisFinancialYear       float32                       `bson:"total_gross_amount_paid_this_financial_year"`
	TotalTaxPaidThisFinancialYear               float32                       `bson:"total_tax_paid_this_financial_year"`
	TotalPaidMonthsThisFinancialYear            int                           `bson:"total_paid_months_this_financial_year"`
	LastPaidMonthThisFinancialYear              string                        `bson:"last_paid_month_this_financial_year"`
	YearlyPayrollAndTaxDetails                  YearlyPayrollAndTaxDetails    `bson:"yearly_payroll_and_tax_details"`
	MonthlyPayrollAndTaxDetails                 []MonthlyPayrollAndTaxDetails `bson:"monthly_payroll_and_tax_details"`
	TotalGrossAmountPendingForThisFinancialYear float32                       `bson:"total_gross_amount_pending_for_this_financial_year"`
	TotalTaxPendingForThisFinancialYear         float32                       `bson:"total_tax_pending_for_this_financial_year"`
	TotalPendingMonthsForThisFinancialYear      int                           `bson:"total_pending_months_for_this_financial_year"`
	CreatedAt                                   time.Time                     `bson:"created_at"`
	UpdatedAt                                   time.Time                     `bson:"updated_at"`
}
