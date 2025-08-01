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

type OrgSalaryStructure struct {
	ID                    primitive.ObjectID        `bson:"_id"`
	OrgName               string                    `bson:"org_name"`
	Name                  string                    `bson:"name"`
	SalaryComponentGroups []OrgSalaryComponentGroup `bson:"salary_component_groups"`
}
