package hrms

type NewTaxRegime struct {
	ID                     string          `bson:"_id"`
	FinancialYear          string          `bson:"financial_year"`
	StandardDeduction      float64         `bson:"standard_deduction"` // e.g., 50000
	Slabs                  []TaxSlab       `bson:"slabs"`
	Rebate                 Rebate          `bson:"rebate"`
	Surcharges             []SurchargeSlab `bson:"surcharges"`                 // Percentage, e.g., 10 for 10%
	CessOnTaxWithSurcharge float64         `bson:"cess_on_tax_with_surcharge"` // Percentage, e.g., 4 for 4%
}

type Rebate struct {
	MaxTaxableIncome float64 `bson:"max_income"` // Inclusive
	RebateAmount     float64 `bson:"rebate_amount"`
}

type TaxSlab struct {
	MinIncome    float64 `bson:"min_income"`     // Inclusive
	MaxIncome    float64 `bson:"max_income"`     // Inclusive
	RateOnIncome float64 `bson:"rate_on_income"` // Percentage, e.g., 5 for 5%
}

type SurchargeSlab struct {
	MinIncome float64 `bson:"min_income"`  // Inclusive
	MaxIncome float64 `bson:"max_income"`  // Inclusive
	RateOnTax float64 `bson:"rate_on_tax"` // Percentage, e.g., 5 for 5%
}
