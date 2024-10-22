package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgSalaryComponent struct {
	ID                   primitive.ObjectID `bson:"_id"`
	ComponentName        string             `bson:"component_name"`
	Abbrivation          string             `bson:"abbrivation"`
	Description          string             `bson:"description"`
	IsTaxableComponent   bool               `bson:"is_taxable_component"`
	IsCreditComponent    bool               `bson:"is_credit_component"`
	Formula              string             `bson:"formula"`
	Cycle                string             `bson:"cycle"` //once, monthly, Biannual, Annual
	IsPaidDayDepandant   bool               `bson:"is_paid_day_depandant"`
	AmountBasedOnFormula bool               `bson:"amount_based_on_formula"`
	CalculatedAmount     float32            `bson:"calculated_amount"`
	StaticAmount         float32            `bson:"static_amount"`
	FinalAmount          float32            `bson:"final_amount"`
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

type OrgSalaryStrucutre struct {
	ID                    primitive.ObjectID        `bson:"_id"`
	OrgName               string                    `bson:"org_name"`
	SalaryComponentGroups []OrgSalaryComponentGroup `bson:"salary_component_groups"`
}
