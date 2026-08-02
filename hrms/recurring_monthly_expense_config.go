package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RecurringMonthlyExpenseConfig is an org-level, HR/admin-configured recurring monthly
// expense (e.g. "Telephone", "Food") applied automatically to matching employees during
// InitiateSalaryProcessing, either as a one-time monthly earning or deduction.
type RecurringMonthlyExpenseConfig struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	// IsDeduction defaults to false, meaning the expense is added to earnings.
	IsDeduction   bool    `bson:"is_deduction"`
	MonthlyAmount float64 `bson:"monthly_amount"`
	// ValidUpToMonth is a "YYYYMM" integer (e.g. 202612 for December 2026) through which
	// (inclusive) the config stays active. A value <= 0 means no expiry.
	ValidUpToMonth int `bson:"valid_up_to_month"`
	// EmployeeTypes lists which employee types (Permanent, Trainee, Internship, ...) this
	// config applies to; matched case-insensitively against EmployeeSummary.EmployeeType.
	EmployeeTypes []string `bson:"employee_types"`
	// PayableToAllEmployeesOfSelectedTypes, when true, applies to every employee of the
	// selected types. When false, only employees listed in EmployeeIds are eligible.
	PayableToAllEmployeesOfSelectedTypes bool           `bson:"payable_to_all_employees_of_selected_types"`
	EmployeeIds                          []EmployeeData `bson:"employee_ids,omitempty"`
	IsActive                             bool           `bson:"is_active"`
	Remarks                              string         `bson:"remarks,omitempty"`
	CreatedBy                            string         `bson:"created_by"`
	CreatedAt                            time.Time      `bson:"created_at"`
	UpdatedAt                            time.Time      `bson:"updated_at"`
}

type EmployeeData struct {
	EmployeeId           string  `bson:"employee_id"`
	EmployeeName         string  `bson:"employee_name"`
	EmployeeType         string  `bson:"employee_type"`
	EmploymentStatus     string  `bson:"employment_status"`
	MonthlyAmount        float64 `bson:"monthly_amount"`
	IsPaymentOnHold      bool    `bson:"is_payment_on_hold"`
	IsManuallyOverridden bool    `bson:"is_manually_overridden"`
}
