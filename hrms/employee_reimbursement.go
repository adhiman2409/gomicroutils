package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseIdCounter struct {
	ID         primitive.ObjectID `bson:"_id"`
	EmployeeId string             `bson:"employee_id"`
	Prefix     string             `bson:"prefix"`
	Counter    int64              `bson:"counter"`
}

type RemarksInfo struct {
	EmployeeId   string `bson:"employee_id"`
	EmployeeName string `bson:"employee_name"`
	DateTime     string `bson:"date_time"`
	Remark       string `bson:"remark"`
}

type Expense struct {
	ID                      primitive.ObjectID `bson:"_id"`
	ExpenseId               string             `bson:"expense_id"`
	EmployeeId              string             `bson:"employee_id"`
	EmployeeName            string             `bson:"employee_name"`
	Department              string             `bson:"department"`
	DepartmentId            string             `bson:"department_id"`
	DesignationId           string             `bson:"designation_id"`
	ExpenseState            ExpenseState       `bson:"expense_state"`
	Category                string             `bson:"category"`
	SubCategory             string             `bson:"sub_category"`
	Day                     string             `bson:"day"`
	Month                   string             `bson:"month"`
	Year                    string             `bson:"year"`
	BillDate                string             `bson:"bill_date"`
	BillAmount              float32            `bson:"bill_amount"`
	BillAmountInINR         float32            `bson:"bill_amount_in_inr"`
	BillCurrency            string             `bson:"bill_currency"`
	BillNumber              string             `bson:"bill_number"`
	PaidTo                  string             `bson:"paid_to"`
	BillableProjectId       string             `bson:"billable_project_id"`
	BillableProject         string             `bson:"billable_project"`
	IsBillableToClient      bool               `bson:"is_billable_to_client"`
	ApprovedAmount          float32            `bson:"approved_amount"`
	ApprovedAmountInINR     float32            `bson:"approved_amount_in_inr"`
	RejectedAmount          float32            `bson:"rejected_amount"`
	RejectedAmountInINR     float32            `bson:"rejected_amount_in_inr"`
	ConversionRate          float32            `bson:"conversion_rate"`
	PrimaryApproverId       string             `bson:"primary_approver_id"`
	PrimaryApproverName     string             `bson:"primary_approver_name"`
	IsApprovedByPrimary     bool               `bson:"is_approved_by_primary"`
	SecondaryApproverId     string             `bson:"secondary_approver_id"`
	SecondaryApproverName   string             `bson:"secondary_approver_name"`
	NeedBothApproval        bool               `bson:"need_both_approval"`
	FinanceApproverId       string             `bson:"finance_approver_id"`
	FinanceApproverName     string             `bson:"finance_approver_name"`
	FinanceStatusUpdatedOn  string             `bson:"secondary_status_updated_on"`
	ReimbursementPaid       bool               `bson:"reimbursement_paid"`
	ReimbursementPaidAmount float32            `bson:"reimbursement_paid_amount"`
	ReimbursementPaidOn     time.Time          `bson:"reimbursement_paid_on"`
	Remarks                 []RemarksInfo      `bson:"remarks"`
	DocURL                  string             `bson:"doc_url"`
	Country                 string             `bson:"country"`
	TimeZone                string             `bson:"time_zone"`
	State                   string             `bson:"state"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}
