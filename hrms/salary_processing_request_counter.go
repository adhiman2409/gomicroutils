package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

// SalaryProcessingRequestCounter tracks the next RequestCount to assign for a given
// (month, financial_year) pair within a tenant domain, allowing multiple
// SalaryProcessingRequest documents to coexist for the same month+FY.
type SalaryProcessingRequestCounter struct {
	Id            primitive.ObjectID `bson:"_id"`
	Month         string             `bson:"month"`
	FinancialYear string             `bson:"financial_year"`
	Counter       int                `bson:"counter"`
}
