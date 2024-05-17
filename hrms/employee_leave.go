package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	LeaveType            string  `bson:"leave_type"`
	Year                 string  `bson:"year"`
	CarryForwardLeaves   float32 `bson:"carry_forward_leaves"`
	EarnedLeavesThisYear float32 `bson:"earned_leaves_this_year"`
	TotalLeaves          float32 `bson:"total_leaves"`
	AvailableLeaves      float32 `bson:"available_leaves"`
	ConsumedLeaves       float32 `bson:"consumed_leaves"`
	AppliedLeaves        float32 `bson:"applied_leaves"`
}

type EmployeeLeaveStatus struct {
	ID             primitive.ObjectID `bson:"_id"`
	EmployeeId     string             `bson:"employee_id"`
	FullName       string             `bson:"full_name"`
	PhoneNumber    string             `bson:"phone_number"`
	EmailId        string             `bson:"email_id"`
	EmpJoiningDate string             `bson:"emp_joining_date"`
	Year           string             `bson:"year"`
	LeavesStats    []LeaveStats       `bson:"leaves_stats"`
}

type Leave struct {
	Day                 string `bson:"day"`
	Month               string `bson:"month"`
	Year                string `bson:"year"`
	LeaveType           string `bson:"leave_type"`
	LeaveStatus         string `bson:"leave_status"`
	IsWithoutPayLeave   bool   `bson:"is_without_pay_leave"`
	LeaveWithdrawalDate string `bson:"leave_withdrawal_date"`
	LeaveApprovalDate   string `bson:"leave_approval_date"`
	Remarks             string `bson:"remarks"`
}

type EmployeeLeaveObj struct {
	ID                           primitive.ObjectID `bson:"_id"`
	EmployeeId                   string             `bson:"employee_id"`
	FullName                     string             `bson:"full_name"`
	PhoneNumber                  string             `bson:"phone_number"`
	EmailId                      string             `bson:"email_id"`
	LeaveType                    string             `bson:"leave_type"`
	LeaveStatus                  string             `bson:"leave_status"`
	ISBulkLeave                  bool               `bson:"is_bulk_leave"`
	RequiredNoticeDays           float32            `bson:"required_notice_days"`
	ActualNoticeDays             float32            `bson:"actual_notice_days"`
	LeaveStartDate               string             `bson:"leave_start_date"`
	LeaveEndDate                 string             `bson:"leave_end_date"`
	IncludeNonWorkingDays        bool               `bson:"include_non_working_days"`
	TotalAppliedLeaves           float32            `bson:"total_applied_leaves"`
	TotalOutOfOfficeDays         float32            `bson:"total_out_of_office_days"`
	IncludeWithoutPayLeave       bool               `bson:"include_without_pay_leave"`
	TotalWithouPayLeaves         float32            `bson:"total_without_pay_leaves"`
	PrimaryApproverId            string             `bson:"primary_approver_id"`
	PrimaryApproverName          string             `bson:"primary_approver_name"`
	PrimaryApproverDesignation   string             `bson:"primary_approver_designation"`
	SecondaryApproverId          string             `bson:"secondary_approver_id"`
	SecondaryApproverName        string             `bson:"secondary_approver_name"`
	SecondaryApproverDesignation string             `bson:"secondary_approver_designation"`
	NeedBothApproval             bool               `bson:"need_both_approval"`
	LeveApplicationDate          string             `bson:"leave_application_date"`
	PrimaryLeaveApprovalDate     string             `bson:"primary_leave_approval_date"`
	SecondaryLeaveApprovalDate   string             `bson:"secondary_leave_approval_date"`
	IncludeWithdrawalLeaves      bool               `bson:"include_withdrawal_leaves"`
	TotalApprovedLeaves          float32            `bson:"total_approved_leaves"`
	Leaves                       []Leave            `bson:"leave"`
	Remarks                      string             `bson:"remarks"`
}
