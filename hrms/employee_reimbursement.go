package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReimbursementIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type ExpenseIdCounter struct {
	ID         primitive.ObjectID `bson:"_id"`
	EmployeeId string             `bson:"employee_id"`
	Prefix     string             `bson:"prefix"`
	Counter    int64              `bson:"counter"`
}

type Expense struct {
	ID              primitive.ObjectID `bson:"_id"`
	ExpenseId       string             `json:"expense_id"`
	EmployeeId      string             `json:"employee_id"`
	ExpenseState    ExpenseState       `json:"expense_state"`
	Category        string             `json:"category"`
	SubCategory     string             `json:"sub_category"`
	Day             string             `json:"day"`
	Month           string             `json:"month"`
	Year            string             `json:"year"`
	BillDate        string             `json:"bill_date"`
	BillAmount      float32            `json:"bill_amount"`
	BillNumber      string             `json:"bill_number"`
	PaidTo          string             `json:"paid_to"`
	IsBillable      bool               `json:"is_billable"`
	BillableProject string             `json:"billable_project"`
	ApprovedAmount  float32            `json:"approved_amount"`
	EmployeeRemarks string             `json:"employee_remarks"`
	ApproverRemarks string             `json:"approver_remarks"`
	ReimbursementId string             `json:"reimbursement_id"`
	DocURL          string             `json:"doc_url"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type Reimbursement struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EmployeeId                   string             `json:"employee_id"`
	FullName                     string             `json:"full_name"`
	EmailId                      string             `json:"email_id"`
	Designation                  string             `json:"designation"`
	ReimbursementId              string             `json:"reimbursement_id"`
	Title                        string             `json:"title"`
	ReimbursementStatus          string             `json:"reimbursement_status"`
	StartDate                    string             `json:"start_date"`
	EndDate                      string             `json:"end_date"`
	TotalAmount                  float32            `json:"total_amount"`
	PrimaryApproverId            string             `json:"primary_approver_id"`
	PrimaryApproverName          string             `json:"primary_approver_name"`
	PrimaryApproverDesignation   string             `json:"primary_approver_designation"`
	IsApprovedByPrimary          bool               `json:"is_approved_by_primary"`
	SecondaryApproverId          string             `json:"secondary_approver_id"`
	SecondaryApproverName        string             `json:"secondary_approver_name"`
	SecondaryApproverDesignation string             `json:"secondary_approver_designation"`
	NeedBothApproval             bool               `json:"need_both_approval"`
	ApplicationDate              string             `json:"application_date"`
	PrimaryApprovalDate          string             `json:"primary_approval_date"`
	SecondaryApprovalDate        string             `json:"secondary_approval_date"`
	RejectionDate                string             `json:"rejection_date"`
	TotalApprovedAmount          float32            `json:"total_approved_amount"`
	Expenses                     []Expense          `json:"expenses"`
	ReimbursementPaid            bool               `json:"reimbursement_paid"`
	ReimbursementPaidAmount      float32            `json:"reimbursement_paid_amount"`
	ReimbursementPaidOn          time.Time          `json:"reimbursement_paid_on"`
	IsAcceptedByFinance          bool               `json:"is_accepted_by_finance"`
	MailSentToFinance            bool               `json:"mail_sent_to_finance"`
	MailSentToFinanceOn          time.Time          `json:"mail_sent_to_finance_on"`
	Remarks                      string             `json:"remarks"`
	CreatedAt                    time.Time          `json:"created_at"`
	UpdatedAt                    time.Time          `json:"updated_at"`
}
