package hrms

// TenantLeavesConfig is the Mongo representation of a tenant's leaves config,
// stored as a single document in the tenant DB's "tenant-leaves-config" collection.
type TenantLeavesConfig struct {
	FromEmail        string            `bson:"from_email"`
	HRFyiRecipients  []ConfigRecipient `bson:"hr_fyi_recipients"`
	WFHFyiRecipients []ConfigRecipient `bson:"wfh_fyi_recipients"`
	// nil means the doc predates this field; treated as true (default on).
	NotifyProjectHead *bool         `bson:"notify_project_head,omitempty"`
	ApproverRules     ApproverRules `bson:"approver_rules"`
	Policy            LeavePolicy   `bson:"policy"`
}

type ConfigRecipient struct {
	Email string `bson:"email"`
	Name  string `bson:"name"`
}

type ApproverRules struct {
	UseTeamLeadAsPrimary bool            `bson:"use_team_lead_as_primary"`
	FixedSecondaryEID    string          `bson:"fixed_secondary_eid"`
	WFHSecondaryEID      string          `bson:"wfh_secondary_eid"`
	BulkEscalation       *BulkEscalation `bson:"bulk_escalation,omitempty"`
}

type BulkEscalation struct {
	Threshold     float64 `bson:"threshold"`
	PrimaryEID    string  `bson:"primary_eid"`
	EscalationEID string  `bson:"escalation_eid"`
}

type LeavePolicy struct {
	ForceBothApprovals    bool `bson:"force_both_approvals"`
	WFHForceBothApprovals bool `bson:"wfh_force_both_approvals"`
	// nil means the doc predates this field; treated as true (auto-approve on).
	WFHAutoApprove           *bool `bson:"wfh_auto_approve,omitempty"`
	WFHAutoApproveLimit      int   `bson:"wfh_auto_approve_limit"`
	WFHMonthlyLimit          int   `bson:"wfh_monthly_limit"`
	WFHBothApprovalThreshold int   `bson:"wfh_both_approval_threshold"`
}
