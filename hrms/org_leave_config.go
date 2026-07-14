package hrms

// TenantLeavesConfig is the Mongo representation of a tenant's leaves config,
// stored as a single document in the tenant DB's "tenant-leaves-config" collection.
type TenantLeavesConfig struct {
	DefaultPrimaryApproverRole      string            `bson:"default_approver_role"`
	ForceBothApprovalsForAllLeaves  bool              `bson:"force_both_approvals_for_all_leaves"`
	DefaultSecondaryApproverRole    string            `bson:"default_secondary_approver_role"`
	ForceBothApprovalsForBulkLeaves bool              `bson:"force_both_approvals_for_bulk_leaves"`
	BulkLeaveCountThreshold         int               `bson:"bulk_leave_count_threshold"`
	DefaultPrimaryApproverEID       string            `bson:"default_primary_eid"`
	DefaultPrimaryApproverName      string            `bson:"default_primary_name"`
	DefaultPrimaryApproverEmail     string            `bson:"default_primary_email"`
	DefaultSecondaryApproverEID     string            `bson:"default_secondary_eid"`
	DefaultSecondaryApproverName    string            `bson:"default_secondary_name"`
	DefaultSecondaryApproverEmail   string            `bson:"default_secondary_email"`
	BulkLeavePrimaryApproverEID     string            `bson:"bulk_leave_primary_eid"`
	BulkLeavePrimaryApproverName    string            `bson:"bulk_leave_primary_name"`
	BulkLeavePrimaryApproverEmail   string            `bson:"bulk_leave_primary_email"`
	BulkLeaveSecondaryApproverEID   string            `bson:"bulk_leave_secondary_eid"`
	BulkLeaveSecondaryApproverName  string            `bson:"bulk_leave_secondary_name"`
	BulkLeaveSecondaryApproverEmail string            `bson:"bulk_leave_secondary_email"`
	NotifyProjectHead               bool              `bson:"notify_project_head"`
	HRFyiRecipients                 []ConfigRecipient `bson:"hr_fyi_recipients"`
}

// TenantWorkFromHomeConfig is the Mongo representation of a tenant's WFH
// config, stored as a single document in the tenant DB's
// "tenant-wfh-config" collection.
type TenantWorkFromHomeConfig struct {
	DefaultPrimaryApproverRole    string            `bson:"default_approver_role"`
	ForceBothApprovalsForAllWFH   bool              `bson:"force_both_approvals_for_all_wfh"`
	DefaultSecondaryApproverRole  string            `bson:"default_secondary_approver_role"`
	ForceBothApprovalsForBulkWFH  bool              `bson:"force_both_approvals_for_bulk_wfh"`
	IsOfficeLocationBasedWFH      bool              `bson:"is_office_location_based_wfh"`
	OfficeLocations               []string          `bson:"office_locations"`
	WFHAutoApprove                bool              `bson:"wfh_auto_approve"`
	WFHAutoApproveLimit           int               `bson:"wfh_auto_approve_limit"`
	WFHMonthlyLimit               int               `bson:"wfh_monthly_limit"`
	WFHBothApprovalThreshold      int               `bson:"wfh_both_approval_threshold"`
	DefaultPrimaryApproverEID     string            `bson:"default_primary_eid"`
	DefaultPrimaryApproverName    string            `bson:"default_primary_name"`
	DefaultPrimaryApproverEmail   string            `bson:"default_primary_email"`
	DefaultSecondaryApproverEID   string            `bson:"default_secondary_eid"`
	DefaultSecondaryApproverName  string            `bson:"default_secondary_name"`
	DefaultSecondaryApproverEmail string            `bson:"default_secondary_email"`
	NotifyProjectHead             bool              `bson:"notify_project_head"`
	HRFyiRecipients               []ConfigRecipient `bson:"hr_fyi_recipients"`
}

type ConfigRecipient struct {
	Email string `bson:"email"`
	Name  string `bson:"name"`
}
