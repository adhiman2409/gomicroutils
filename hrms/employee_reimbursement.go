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
	ExpenseId       string             `bson:"expense_id"`
	EmployeeId      string             `bson:"employee_id"`
	ExpenseState    ExpenseState       `bson:"expense_state"`
	Category        string             `bson:"category"`
	SubCategory     string             `bson:"sub_category"`
	Day             string             `bson:"day"`
	Month           string             `bson:"month"`
	Year            string             `bson:"year"`
	BillDate        string             `bson:"bill_date"`
	BillAmount      float32            `bson:"bill_amount"`
	BillNumber      string             `bson:"bill_number"`
	PaidTo          string             `bson:"paid_to"`
	IsBillable      bool               `bson:"is_billable"`
	BillableProject string             `bson:"billable_project"`
	ApprovedAmount  float32            `bson:"approved_amount"`
	EmployeeRemarks string             `bson:"employee_remarks"`
	ApproverRemarks string             `bson:"approver_remarks"`
	ReimbursementId string             `bson:"reimbursement_id"`
	DocURL          string             `bson:"doc_url"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type Reimbursement struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EmployeeId                   string             `bson:"employee_id"`
	FullName                     string             `bson:"full_name"`
	EmailId                      string             `bson:"email_id"`
	Designation                  string             `bson:"designation"`
	ReimbursementId              string             `bson:"reimbursement_id"`
	Title                        string             `bson:"title"`
	ReimbursementStatus          string             `bson:"reimbursement_status"`
	StartDate                    string             `bson:"start_date"`
	EndDate                      string             `bson:"end_date"`
	TotalAmount                  float32            `bson:"total_amount"`
	PrimaryApproverId            string             `bson:"primary_approver_id"`
	PrimaryApproverName          string             `bson:"primary_approver_name"`
	PrimaryApproverDesignation   string             `bson:"primary_approver_designation"`
	IsApprovedByPrimary          bool               `bson:"is_approved_by_primary"`
	SecondaryApproverId          string             `bson:"secondary_approver_id"`
	SecondaryApproverName        string             `bson:"secondary_approver_name"`
	SecondaryApproverDesignation string             `bson:"secondary_approver_designation"`
	NeedBothApproval             bool               `bson:"need_both_approval"`
	ApplicationDate              string             `bson:"application_date"`
	PrimaryApprovalDate          string             `bson:"primary_approval_date"`
	SecondaryApprovalDate        string             `bson:"secondary_approval_date"`
	RejectionDate                string             `bson:"rejection_date"`
	TotalApprovedAmount          float32            `bson:"total_approved_amount"`
	Expenses                     []Expense          `bson:"expenses"`
	ReimbursementPaid            bool               `bson:"reimbursement_paid"`
	ReimbursementPaidAmount      float32            `bson:"reimbursement_paid_amount"`
	ReimbursementPaidOn          time.Time          `bson:"reimbursement_paid_on"`
	IsAcceptedByFinance          bool               `bson:"is_accepted_by_finance"`
	MailSentToFinance            bool               `bson:"mail_sent_to_finance"`
	MailSentToFinanceOn          time.Time          `bson:"mail_sent_to_finance_on"`
	Remarks                      string             `bson:"remarks"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}
