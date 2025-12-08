package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID              primitive.ObjectID `bson:"_id"`
	BudgetId        string             `bson:"budget_id"`
	DepartmentID    string             `bson:"department_id"`
	DepartmentName  string             `bson:"department_name"`
	AmountAllocated float32            `bson:"amount_allocated"`
	Currency        string             `bson:"currency"`
	Status          string             `bson:"status"`
	FiscalYear      string             `bson:"fiscal_year"`
	CreatedBy       EmployeeSummary    `bson:"created_by"`
	AmountSpent     float32            `bson:"amount_spent"`
	RemainingBudget float32            `bson:"remaining_budget"`
	Remarks         string             `bson:"remarks"`
	UpdatedAt       time.Time          `bson:"updated_at"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type BudgetHistory struct {
	ID              primitive.ObjectID `bson:"_id"`
	BudgetId        string             `bson:"budget_id"`
	DepartmentID    string             `bson:"department_id"`
	FiscalYear      string             `bson:"fiscal_year"`
	TransactionType string             `bson:"transaction_type"`
	Amount          float32            `bson:"amount"`
	Currency        string             `bson:"currency"`
	AllocatedBy     string             `bson:"allocated_by"`
	Remarks         string             `bson:"remarks"`
	TransactionDate time.Time          `bson:"transaction_date"`
}

type BudgetIdCounter struct {
	ID             primitive.ObjectID `bson:"_id"`
	DepartmentName string             `bson:"department_name"`
	Counter        int64              `bson:"counter"`
}

type DepartmentDailyStat struct {
	ID                primitive.ObjectID `bson:"_id"`
	Date              string             `bson:"date" json:"date"`
	Department        string             `bson:"department" json:"department"`
	TotalLoginMinutes float64            `bson:"total_login_minutes" json:"total_login_minutes"`
	PresentCount      int                `bson:"present_count" json:"present_count"`
	WorkingDayCount   int                `bson:"working_day_count" json:"working_day_count"`
	Headcount         int                `bson:"headcount" json:"headcount"`
}
