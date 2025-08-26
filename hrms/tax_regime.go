package hrms

type TaxRegime struct {
	ID                     string          `bson:"_id"`
	RegimeName             string          `bson:"regime_name"`    // e.g., "Old", "New"
	FinancialYear          string          `bson:"financial_year"` // e.g., "2023-2024"
	Description            string          `bson:"description"`
	StandardDeduction      float64         `bson:"standard_deduction"` // e.g., 50000
	Slabs                  []TaxSlab       `bson:"slabs"`
	RebateUnderSection87A  Rebate          `bson:"rebate"`
	Surcharges             []SurchargeSlab `bson:"surcharges"`                 // Percentage, e.g., 10 for 10%
	CessOnTaxWithSurcharge float64         `bson:"cess_on_tax_with_surcharge"` // Percentage, e.g., 4 for 4%
}

type Rebate struct {
	MaxTaxableIncome float64 `bson:"max_income"` // Inclusive
	RebateAmount     float64 `bson:"rebate_amount"`
}

type TaxSlab struct {
	MinIncome     float64 `bson:"min_income"`     // Inclusive
	MaxIncome     float64 `bson:"max_income"`     // Inclusive
	RateOnIncome  float64 `bson:"rate_on_income"` // Percentage, e.g., 5 for 5%
	CalculatedTax float64 `bson:"calculated_tax"` // Pre-calculated tax for the slab
}

type SurchargeSlab struct {
	MinIncome float64 `bson:"min_income"`  // Inclusive
	MaxIncome float64 `bson:"max_income"`  // Inclusive
	RateOnTax float64 `bson:"rate_on_tax"` // Percentage, e.g., 5 for 5%
}
