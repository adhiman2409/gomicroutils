package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryProcessingRequest struct {
	Id                             primitive.ObjectID            `bson:"_id"`
	RequestId                      string                        `bson:"request_id"`
	Year                           string                        `bson:"year"`
	Month                          string                        `bson:"month"`
	NumberOfDaysInMonth            int                           `bson:"number_of_days_in_month"`
	NumberOfPayableDays            int                           `bson:"number_of_payable_days"`
	OrganizationDefaultTaxRegime   string                        `bson:"organization_default_tax_regime"`
	EmployeeMonthlySalaryDetails   []EmployeeMonthlySalaryDetail `bson:"employee_monthly_salary_details"`
	TotalNumberOfSalaries          int                           `bson:"total_number_of_salaries"`
	TotalNumberOfOnHoldSalaries    int                           `bson:"total_number_of_on_hold_salaries"`
	TotalNumberOfProcessedSalaries int                           `bson:"total_number_of_processed_salaries"`
	TotalSalaryAmount              float64                       `bson:"total_salary_amount"`
	RequesterId                    string                        `bson:"requested_by_id"`
	RequesterName                  string                        `bson:"requested_by_name"`
	RequesterRemarks               string                        `bson:"requester_remarks"`
	RequestedAt                    time.Time                     `bson:"requested_at"`
	ApproverId                     string                        `bson:"approver_id"`
	RequestStatus                  string                        `bson:"request_status"`
	ApproverName                   string                        `bson:"approver_name"`
	ApproverRemarks                string                        `bson:"approver_remarks"`
	ApprovedAt                     time.Time                     `bson:"approved_at"`
	CreatedAt                      time.Time                     `bson:"created_at"`
	CreatedBy                      string                        `bson:"created_by"`
	UpdatedAt                      time.Time                     `bson:"updated_at"`
}

type AttendanceSheet struct {
	Day                  string  `bson:"day"`
	Month                string  `bson:"month"`
	Year                 string  `bson:"year"`
	Weekday              string  `bson:"weekday"`
	IsHoliday            bool    `bson:"is_holiday"`
	IsWeeklyOffDay       bool    `bson:"is_weekly_off_day"`
	IsWorkingDay         bool    `bson:"is_working_day"`
	IsOnLeave            bool    `bson:"is_on_leave"`
	IsHalfDayLeave       bool    `bson:"is_half_day_leave"`
	IsFullDayLOP         bool    `bson:"is_full_day_lop"`
	IsHalfDayLOP         bool    `bson:"is_half_day_lop"`
	RequiredWorkingHours float32 `bson:"required_working_hours"`
	ActualWorkingHours   float32 `bson:"actual_working_hours"`
	Status               string  `bson:"status"`
	Remarks              string  `bson:"remarks"`
}

type ReimbursementDetails struct {
	ReimbursementID   string  `bson:"reimbursement_id"`
	ReimbursementType string  `bson:"reimbursement_type"`
	ReimbursementDate string  `bson:"reimbursement_date"`
	Amount            float64 `bson:"amount"`
	Status            string  `bson:"status"`
	ApprovedBy        string  `bson:"approved_by"`
	ApprovalDate      string  `bson:"approval_date"`
	Remarks           string  `bson:"remarks"`
}

type EmployeeMonthlySalaryDetail struct {
	EmployeeID                   string                       `bson:"employee_id"`
	EmployeeName                 string                       `bson:"employee_name"`
	Designation                  string                       `bson:"designation"`
	Department                   string                       `bson:"department"`
	DeliveryLocation             string                       `bson:"delivery_location"`
	JoiningDate                  string                       `bson:"joining_date"`
	EmploymentStatus             string                       `bson:"employment_status"`
	EmployeeType                 string                       `bson:"employee_type"`
	Gender                       string                       `bson:"gender"`
	EmployeeAdoptedTaxRegime     string                       `bson:"employee_adopted_tax_regime"`
	State                        string                       `bson:"state"`
	Country                      string                       `bson:"country"`
	PAN                          string                       `bson:"pan"`
	UAN                          string                       `bson:"uan"`
	UID                          string                       `bson:"uid"`
	PFEstablishmentCode          string                       `bson:"pf_establishment_code"`
	BankName                     string                       `bson:"bank_name"`
	BankBranch                   string                       `bson:"bank_branch"`
	AccountNumber                string                       `bson:"account_number"`
	IFSC                         string                       `bson:"ifsc"`
	IsShortWorkingDaysOrHours    bool                         `bson:"is_short_working_days_or_hours"`
	TotalWorkingDays             int                          `bson:"total_working_days"`
	ActualWorkingDays            int                          `bson:"actual_working_days"`
	RequiredWorkingHours         float32                      `bson:"required_working_hours"`
	ActualWorkingHours           float32                      `bson:"actual_working_hours"`
	AttendanceSheet              []AttendanceSheet            `bson:"attendance"`
	Reimbursements               []ReimbursementDetails       `bson:"reimbursements"`
	SalaryProcessingStatus       SalaryStatus                 `bson:"salary_processing_status"`
	SalaryProcessingDate         string                       `bson:"salary_processing_date"`
	IsManuallyUpdated            bool                         `bson:"is_manually_updated"`
	MonthlyEarnings              MonthlyEarnings              `bson:"monthly_earnings"`
	TotalGrossEarnings           float64                      `bson:"total_gross_earnings"`
	MonthlyDeductions            MonthlyDeductions            `bson:"monthly_deductions"`
	TotalDeductions              float64                      `bson:"total_deductions"`
	NetSalary                    float64                      `bson:"net_salary"`
	MonthlyEmployerContributions MonthlyEmployerContributions `bson:"monthly_employer_contributions"`
	TotalEmployerContributions   float64                      `bson:"total_employer_contributions"`
	OneTimeMonthlyEarning        []OneTimeMonthlyEarning      `bson:"one_time_monthly_earnings"`
	OneTimeMonthlyDeduction      []OneTimeMonthlyDeduction    `bson:"one_time_monthly_deductions"`
	Remarks                      string                       `bson:"remarks"`
}
