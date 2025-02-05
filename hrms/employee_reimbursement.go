package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reimbursement struct {
	EmployeeId          string    `bson:"employee_id"`
	Category            string    `bson:"category"`
	Day                 string    `bson:"day"`
	Month               string    `bson:"month"`
	Year                string    `bson:"year"`
	NeedBothApproval    bool      `bson:"need_both_approval"`
	BillDate            string    `bson:"bill_date"`
	BillAmount          float32   `bson:"bill_amount"`
	BillNumber          string    `bson:"bill_number"`
	PaidTo              string    `bson:"paid_to"`
	IsBillable          bool      `bson:"is_billable"`
	BillableCompany     string    `bson:"billable_company"`
	ApprovedAmount      float32   `bson:"approved_amount"`
	Remarks             string    `bson:"remarks"`
	DocURL              string    `bson:"doc_url"`
	PrimaryApproved     bool      `bson:"primary_approved"`
	PrimaryApprovedOn   time.Time `bson:"primary_approved_on"`
	SecondaryApproved   bool      `bson:"secondary_approved"`
	SecondaryApprovedOn time.Time `bson:"secondary_approved_on"`
	PrimaryRejected     bool      `bson:"primary_rejected"`
	PrimaryRejectedOn   time.Time `bson:"primary_rejected_on"`
	SecondaryRejected   bool      `bson:"secondary_rejected"`
	SecondaryRejectedOn time.Time `bson:"secondary_rejected_on"`
	RejectionReason     string    `bson:"rejection_reason"`
	Status              string    `bson:"status"`
	CreatedAt           time.Time `bson:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at"`
}

type ReimbursementObj struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EmployeeId                   string             `bson:"employee_id"`
	FullName                     string             `bson:"full_name"`
	EmailId                      string             `bson:"email_id"`
	Designation                  string             `bson:"designation"`
	ReimbursementId              string             `bson:"reimbursement_id"`
	ReimbursementType            string             `bson:"reimbursement_type"`
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
	Reimbursements               []Reimbursement    `bson:"reimbursements"`
	ReimbursementPaid            bool               `bson:"reimbursement_paid"`
	ReimbursementPaidOn          time.Time          `bson:"reimbursement_paid_on"`
	MailSentToFinance            bool               `bson:"mail_sent_to_finance"`
	MailSentToFinanceOn          time.Time          `bson:"mail_sent_to_finance_on"`
	Remarks                      string             `bson:"remarks"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}
