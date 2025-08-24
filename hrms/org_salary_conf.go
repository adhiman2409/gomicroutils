package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgSalaryComponent struct {
	ID                        primitive.ObjectID `bson:"_id"`
	ComponentName             string             `bson:"component_name"`
	Abbrivation               string             `bson:"abbrivation"`
	Description               string             `bson:"description"`
	IsBaseComponent           bool               `bson:"is_base_component"` // Indicates if this is a base component
	IsTaxableComponent        bool               `bson:"is_taxable_component"`
	IsCreditComponent         bool               `bson:"is_credit_component"`
	Formula                   string             `bson:"formula"`
	PaymentFrequency          Frequency          `bson:"payment_frequency"` //once, monthly, Biannual, Annual
	IsPresentDayDepandant     bool               `bson:"is_present_day_depandant"`
	IsDrivedComponent         bool               `bson:"is_drived_component"`
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

type OrgStockOptionComponents struct {
	StockCount                   float32 `bson:"stock_count"`
	TotalStockValue              float32 `bson:"total_stock_value"`
	TotalVestingDurationInMonths string  `bson:"total_vesting_duration_in_months"`
	VestingFrequencyInMonths     string  `bson:"vesting_frequency_in_months"`
	StockOptionType              string  `bson:"stock_option_type"` // e.g., ESOP, RSU
}

type OrgSalaryStructure struct {
	ID                    primitive.ObjectID        `bson:"_id"`
	Name                  string                    `bson:"name"`
	Abbrivation           string                    `bson:"abbrivation"`
	Description           string                    `bson:"description"`
	BaseComponentValue    float32                   `bson:"base_component_value"`
	SalaryComponentGroups []OrgSalaryComponentGroup `bson:"salary_component_groups"`
	IsStockOptionEnabled  bool                      `bson:"is_stock_option_enabled"`
	StockOptionComponents OrgStockOptionComponents  `bson:"stock_option_components"`
	TotalCTC              float32                   `bson:"total_ctc"`
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

type OrgSalaryConfig struct {
	ID                                       primitive.ObjectID `bson:"_id"`
	State                                    int                `bson:"state"`
	Country                                  string             `bson:"country"`
	DeliveryLocation                         string             `bson:"delivery_location"`
	OrganizationDefaultTaxRegime             string             `bson:"organization_default_tax_regime"` // e.g., "new", "old"
	IsPFContributionIsBasedOnBasicSalary     bool               `bson:"is_pf_contribution_is_based_on_basic_salary"`
	PFEmployerContributionPercentageOfBasic  float32            `bson:"pf_employer_contribution_percentage_of_basic"`
	PFEmployeeContributionPercentageOfBasic  float32            `bson:"pf_employee_contribution_percentage_of_basic"`
	PFMonthlyFixedEmployerContributionAmount float32            `bson:"pf_monthly_fixed_employer_contribution_amount"`
	PFMonthlyFixedEmployeeContributionAmount float32            `bson:"pf_monthly_fixed_employee_contribution_amount"`
	MonthlyBasicSalaryPercentageOfGross      float32            `bson:"monthly_basic_salary_percentage_odf_gross"`
	MinimumMonthlyBasicSalary                float32            `bson:"minimum_monthly_basic_salary"`
	MonthlyHRAPercentageOfBasic              float32            `bson:"monthly_hra_percentage_of_basic"`
	MonthlyStatutoryBonusPercentageOfBasic   float32            `bson:"monthly_statutory_bonus_percentage_of_basic"`
	MonthlyGratuityPercentageOfBasic         float32            `bson:"monthly_gratuity_percentage_of_basic"`
	AddRemainingAmountInSpecialAllowance     bool               `bson:"add_remaining_amount_in_special_allowance"`
	ESICConfig                               []ESICConfig       `bson:"esi_config"`
	LWFConfig                                []LWFConfig        `bson:"lwf_config"`
	AnnualMedicalInsurancePremium            float32            `bson:"annual_medical_insurance_premium"`
}
