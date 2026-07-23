package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TenantLeavesConfig is the Mongo representation of a tenant's leaves config.
// Like TenantResignationConfig, multiple docs can coexist per tenant in the
// "tenant-leaves-config" collection, one per configured scope (Level +
// Country/State/City/OfficeLabel); resolution picks the most specific level
// whose location fields match an employee (office > city > state > country > global).
type TenantLeavesConfig struct {
	ID                              primitive.ObjectID      `bson:"_id"`
	Level                           OrganizationConfigLevel `bson:"level"`
	Country                         string                  `bson:"country,omitempty"`
	State                           string                  `bson:"state,omitempty"`
	City                            string                  `bson:"city,omitempty"`
	OfficeLabel                     string                  `bson:"office_label,omitempty"`
	DefaultPrimaryApproverRole      string                  `bson:"default_approver_role"`
	ForceBothApprovalsForAllLeaves  bool                    `bson:"force_both_approvals_for_all_leaves"`
	DefaultSecondaryApproverRole    string                  `bson:"default_secondary_approver_role"`
	ForceBothApprovalsForBulkLeaves bool                    `bson:"force_both_approvals_for_bulk_leaves"`
	BulkLeaveCountThreshold         int                     `bson:"bulk_leave_count_threshold"`
	DefaultPrimaryApproverEID       string                  `bson:"default_primary_eid"`
	DefaultPrimaryApproverName      string                  `bson:"default_primary_name"`
	DefaultPrimaryApproverEmail     string                  `bson:"default_primary_email"`
	DefaultSecondaryApproverEID     string                  `bson:"default_secondary_eid"`
	DefaultSecondaryApproverName    string                  `bson:"default_secondary_name"`
	DefaultSecondaryApproverEmail   string                  `bson:"default_secondary_email"`
	BulkLeavePrimaryApproverEID     string                  `bson:"bulk_leave_primary_eid"`
	BulkLeavePrimaryApproverName    string                  `bson:"bulk_leave_primary_name"`
	BulkLeavePrimaryApproverEmail   string                  `bson:"bulk_leave_primary_email"`
	BulkLeaveSecondaryApproverEID   string                  `bson:"bulk_leave_secondary_eid"`
	BulkLeaveSecondaryApproverName  string                  `bson:"bulk_leave_secondary_name"`
	BulkLeaveSecondaryApproverEmail string                  `bson:"bulk_leave_secondary_email"`
	NotifyProjectHead               bool                    `bson:"notify_project_head"`
	HRFyiRecipients                 []ConfigRecipient       `bson:"hr_fyi_recipients"`
	UpdatedAt                       time.Time               `bson:"updated_at"`
}

// TenantWorkFromHomeConfig is the Mongo representation of a tenant's WFH
// config. Like TenantResignationConfig, multiple docs can coexist per tenant
// in the "tenant-wfh-config" collection, one per configured scope (Level +
// Country/State/City/OfficeLabel); resolution picks the most specific level
// whose location fields match an employee (office > city > state > country > global).
type TenantWorkFromHomeConfig struct {
	ID                            primitive.ObjectID      `bson:"_id"`
	Level                         OrganizationConfigLevel `bson:"level"`
	Country                       string                  `bson:"country,omitempty"`
	State                         string                  `bson:"state,omitempty"`
	City                          string                  `bson:"city,omitempty"`
	OfficeLabel                   string                  `bson:"office_label,omitempty"`
	DefaultPrimaryApproverRole    string                  `bson:"default_approver_role"`
	ForceBothApprovalsForAllWFH   bool                    `bson:"force_both_approvals_for_all_wfh"`
	DefaultSecondaryApproverRole  string                  `bson:"default_secondary_approver_role"`
	ForceBothApprovalsForBulkWFH  bool                    `bson:"force_both_approvals_for_bulk_wfh"`
	IsOfficeLocationBasedWFH      bool                    `bson:"is_office_location_based_wfh"`
	OfficeLocations               []string                `bson:"office_locations"`
	WFHAutoApprove                bool                    `bson:"wfh_auto_approve"`
	WFHAutoApproveLimit           int                     `bson:"wfh_auto_approve_limit"`
	WFHMonthlyLimit               int                     `bson:"wfh_monthly_limit"`
	WFHBothApprovalThreshold      int                     `bson:"wfh_both_approval_threshold"`
	DefaultPrimaryApproverEID     string                  `bson:"default_primary_eid"`
	DefaultPrimaryApproverName    string                  `bson:"default_primary_name"`
	DefaultPrimaryApproverEmail   string                  `bson:"default_primary_email"`
	DefaultSecondaryApproverEID   string                  `bson:"default_secondary_eid"`
	DefaultSecondaryApproverName  string                  `bson:"default_secondary_name"`
	DefaultSecondaryApproverEmail string                  `bson:"default_secondary_email"`
	NotifyProjectHead             bool                    `bson:"notify_project_head"`
	HRFyiRecipients               []ConfigRecipient       `bson:"hr_fyi_recipients"`
	UpdatedAt                     time.Time               `bson:"updated_at"`
}

type ConfigRecipient struct {
	Email string `bson:"email"`
	Name  string `bson:"name"`
}
