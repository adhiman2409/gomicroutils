package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reimbursment struct {
	EmployeeId      string    `bson:"employee_id"`
	Category        string    `bson:"category"`
	Day             string    `bson:"day"`
	Month           string    `bson:"month"`
	Year            string    `bson:"year"`
	BillDate        string    `bson:"bill_date"`
	BillAmount      float32   `bson:"bill_amount"`
	BillNumber      string    `bson:"bill_number"`
	PaidTo          string    `bson:"paid_to"`
	IsBillable      bool      `bson:"is_billable"`
	BillableCompany string    `bson:"billable_company"`
	ApprovedAmount  float32   `bson:"approved_amount"`
	Remarks         string    `bson:"remarks"`
	DocURL          string    `bson:"doc_url"`
	Approved        bool      `bson:"approved"`
	ApprovedBy      string    `bson:"approved_by"`
	ApprovedOn      time.Time `bson:"approved_on"`
	Rejected        bool      `bson:"rejected"`
	RejectedBy      string    `bson:"rejected_by"`
	RejectedOn      time.Time `bson:"rejected_on"`
	RejectionReason string    `bson:"rejection_reason"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

type ReimbursmentObj struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EmployeeId                   string             `bson:"employee_id"`
	FullName                     string             `bson:"full_name"`
	EmailId                      string             `bson:"email_id"`
	Designation                  string             `bson:"designation"`
	ReimbursmentId               string             `bson:"reimbursment_id"`
	ReimbursmentType             string             `bson:"reimbursment_type"`
	Title                        string             `bson:"title"`
	ReimbursmentStatus           string             `bson:"reimbursment_status"`
	RequiredNoticeDays           float32            `bson:"required_notice_days"`
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
	Reimbursments                []Reimbursment     `bson:"reimbursments"`
	Remarks                      string             `bson:"remarks"`
	DocURL                       string             `bson:"doc_url"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
}
