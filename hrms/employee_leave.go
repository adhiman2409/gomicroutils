package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeLeave struct {
	ID          primitive.ObjectID `bson:"_id"`
	EmployeeId  string             `bson:"employee_id"`
	FullName    string             `bson:"full_name"`
	PhoneNumber string             `bson:"phone_number"`
	EmailId     string             `bson:"email_id"`
	CasualLeave string             `bson:"casual_leave"`
	SickLeave   string             `bson:"sick_leave"`
	EarnedLeave string             `bson:"earned_leave"`
}

type LeaveStats struct {
	IsActive                          bool    `bson:"is_active"`
	LeaveType                         string  `bson:"leave_type"`
	Year                              string  `bson:"year"`
	CarryForwardLeaves                float32 `bson:"carry_forward_leaves"`
	EarnedLeavesThisYear              float32 `bson:"earned_leaves_this_year"`
	TotalLeaves                       float32 `bson:"total_leaves"`
	AvailableLeaves                   float32 `bson:"available_leaves"`
	ConsumedLeaves                    float32 `bson:"consumed_leaves"`
	AppliedLeaves                     float32 `bson:"applied_leaves"`
	ApplicableAfterWorkingDays        int32   `bson:"applicable_after_working_days"`
	ActivationDate                    string  `bson:"activation_date"`
	DeactivationDate                  string  `bson:"deactivation_date"`
	IncludeWeeklyOffDays              bool    `bson:"include_weekly_off_days"`
	IncludeHolidays                   bool    `bson:"include_holidays"`
	MaxLeaveAllowedCount              float32 `bson:"max_leave_allowed_count"`
	BulkLeaveNoticeInDays             int     `bson:"bulk_leave_notice_in_days"`
	WorkingDaysRemainingForActivation int32   `bson:"working_days_remaining_for_activation"`
	WorkingDaysRemainingForNextBatch  int32   `bson:"working_days_remaining_for_next_batch"`
	LeavesEligibleForEncashment       float32 `bson:"leaves_eligible_for_encashment"`
	LeavesEligibleForCarryForward     float32 `bson:"leaves_eligible_for_carry_forward"`
	LeavesEncashedLastYear            float32 `bson:"leaves_encashed_last_year"`
	LeavesCarryForwardLastYear        float32 `bson:"leaves_carry_forward_last_year"`
	LastYearEarnedPrivilegeLeave      float32 `bson:"last_year_earned_privilege_leave"`
	LastLeavesCreditedOn              string  `bson:"last_leaves_credited_on"`
	Country                           string  `bson:"country"`
	TimeZone                          string  `bson:"time_zone"`
}

type EmployeeLeaveStats struct {
	ID                               primitive.ObjectID `bson:"_id"`
	EmployeeId                       string             `bson:"employee_id"`
	EmployeeType                     string             `bson:"employee_type"`
	FullName                         string             `bson:"full_name"`
	PhoneNumber                      string             `bson:"phone_number"`
	IsMale                           bool               `bson:"is_male"`
	IsMarried                        bool               `bson:"is_married"`
	EmailId                          string             `bson:"email_id"`
	EmpJoiningDate                   string             `bson:"emp_joining_date"`
	EmpConfirmationDate              string             `bson:"emp_confirmation_date"`
	TenureInDays                     float32            `bson:"tenure_in_days"`
	TotalWorkingDays                 float32            `bson:"total_working_days"`
	TotalAbsentDays                  float32            `bson:"total_absent_days"`
	TotalPresentDays                 float32            `bson:"total_present_days"`
	AttendanceStatsUpdatedOn         string             `bson:"attendance_stats_updated_on"`
	LeaveStatsUpdatedOn              string             `bson:"leave_stats_updated_on"`
	EarnedLeaveCycleInDays           int32              `bson:"earned_leave_cycle_in_days"`
	RegularizationReminderMailSentOn string             `bson:"regularization_reminder_mail_sent_on"`
	Year                             string             `bson:"year"`
	Country                          string             `bson:"country"`
	TimeZone                         string             `bson:"time_zone"`
	LeavesStats                      []LeaveStats       `bson:"leaves_stats"`
}

type Leave struct {
	Day                 string `bson:"day"`
	Month               string `bson:"month"`
	Year                string `bson:"year"`
	LeaveType           string `bson:"leave_type"`
	LeaveStatus         string `bson:"leave_status"`
	LeaveDuration       string `bson:"leave_duration"`
	IsWithoutPayLeave   bool   `bson:"is_without_pay_leave"`
	LeaveWithdrawalDate string `bson:"leave_withdrawal_date"`
	LeaveApprovalDate   string `bson:"leave_approval_date"`
	Remarks             string `bson:"remarks"`
}

type EmployeeLeaveObj struct {
	ID                                   primitive.ObjectID `bson:"_id"`
	LeaveId                              string             `bson:"leave_id"`
	EmployeeId                           string             `bson:"employee_id"`
	FullName                             string             `bson:"full_name"`
	EmailId                              string             `bson:"email_id"`
	LeaveType                            string             `bson:"leave_type"`
	LeaveStatus                          string             `bson:"leave_status"`
	ISBulkLeave                          bool               `bson:"is_bulk_leave"`
	RequiredNoticeDays                   float32            `bson:"required_notice_days"`
	ActualNoticeDays                     float32            `bson:"actual_notice_days"`
	LeaveStartDate                       string             `bson:"leave_start_date"`
	LeaveEndDate                         string             `bson:"leave_end_date"`
	LeaveDuration                        string             `bson:"leave_duration"`
	AmountOfLeaves                       float32            `bson:"amount_of_leaves"`
	IncludeNonWorkingDays                bool               `bson:"include_non_working_days"`
	TotalAppliedLeaves                   float32            `bson:"total_applied_leaves"`
	TotalOutOfOfficeDays                 float32            `bson:"total_out_of_office_days"`
	IncludeWithoutPayLeave               bool               `bson:"include_without_pay_leave"`
	TotalWithouPayLeaves                 float32            `bson:"total_without_pay_leaves"`
	PrimaryApproverId                    string             `bson:"primary_approver_id"`
	PrimaryApproverName                  string             `bson:"primary_approver_name"`
	PrimaryApproverDesignation           string             `bson:"primary_approver_designation"`
	IsApprovedByPrimary                  bool               `bson:"is_approved_by_primary"`
	SecondaryApproverId                  string             `bson:"secondary_approver_id"`
	SecondaryApproverName                string             `bson:"secondary_approver_name"`
	SecondaryApproverDesignation         string             `bson:"secondary_approver_designation"`
	NeedBothApproval                     bool               `bson:"need_both_approval"`
	LeveApplicationDate                  string             `bson:"leave_application_date"`
	PrimaryLeaveApprovalDate             string             `bson:"primary_leave_approval_date"`
	SecondaryLeaveApprovalDate           string             `bson:"secondary_leave_approval_date"`
	RejectionDate                        string             `bson:"rejection_date"`
	IncludeWithdrawalLeaves              bool               `bson:"include_withdrawal_leaves"`
	TotalApprovedLeaves                  float32            `bson:"total_approved_leaves"`
	IsExternalApprovalInitiated          bool               `bson:"is_external_approval_initiated"`
	ExternalApprovalInitiaedForProjectId string             `bson:"external_approval_initiated_for_project_id"`
	ExternalApproverEmail                string             `bson:"external_approver_email"`
	ExternalApproverName                 string             `bson:"external_approver_name"`
	ExternalApprovalStatus               string             `bson:"external_approval_status"`
	ExternalApprovalRemarks              string             `bson:"external_approval_remarks"`
	ExternalApprovalInitiatedAt          time.Time          `bson:"external_approval_initiated_at"`
	ExternalApprovalStatusReceivedAt     time.Time          `bson:"external_approval_status_received_at"`
	Leaves                               []Leave            `bson:"leave"`
	Remarks                              string             `bson:"remarks"`
	DocURL                               string             `bson:"doc_url"`
	Country                              string             `bson:"country"`
	TimeZone                             string             `bson:"time_zone"`
	CreatedAt                            time.Time          `bson:"created_at"`
	UpdatedAt                            time.Time          `bson:"updated_at"`
}

type LeaveCarryForward struct {
	ID                primitive.ObjectID `bson:"_id"`
	EmployeeId        string             `bson:"employee_id"`
	CarryForwardCount float64            `bson:"carry_forward_count"`
	EncashmentCount   float64            `bson:"encashment_count"`
	Year              int                `bson:"year"`
}
