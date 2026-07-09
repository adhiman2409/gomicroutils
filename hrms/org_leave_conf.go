package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrgLeaveConf struct {
	ID                                 primitive.ObjectID `bson:"_id"`
	OrgName                            string             `bson:"org_name"`
	Title                              string             `bson:"title"`
	IsFixed                            bool               `bson:"is_fixed"`
	Frequency                          string             `bson:"frequency"`
	LeaveCount                         float32            `bson:"leave_count"`
	LeaveUnit                          string             `bson:"leave_unit"`
	MaxLeaveCount                      float32            `bson:"max_leave_count"`
	IsCarryForwardAllowed              bool               `bson:"is_carry_forward_allowed"`
	MaxCarryForwardCount               float32            `bson:"max_carry_forward_count"`
	BulkLeaveCount                     int                `bson:"bulk_leave_count"`
	BulkLeaveNoticeInDays              int                `bson:"bulk_leave_notice_in_days"`
	IsEncashmentAllowed                bool               `bson:"encashment_allowed"`
	IsManualAllocationAllowed          bool               `bson:"is_manual_allocation_allowed"`
	MaxManualAllocationCount           float32            `bson:"max_manual_allocation_count"`
	TotalManualAllocationAllowedInYear float32            `bson:"total_manual_allocation_allowed_in_year"`
	ApplicableAfterWorkingDays         int                `bson:"applicable_after_working_days"`
	DocumentRequired                   bool               `bson:"document_required"`
	WeeklyOffAndHolidayIncluded        bool               `bson:"weekly_off_and_holiday_included"`
	IncludeWeeklyOffDays               bool               `bson:"include_weekly_off_days"`
	IncludeHolidays                    bool               `bson:"include_holidays"`
	MaxLeaveAllowedCount               float32            `bson:"max_leave_allowed_count"`
	IsActive                           bool               `bson:"is_active"`
	Description                        string             `bson:"description"`
	YearStartDate                      string             `bson:"year_start_date"`
	YearEndDate                        string             `bson:"year_end_date"`
	State                              string             `bson:"state"`
	Country                            string             `bson:"country"`
	CreatedBy                          string             `bson:"created_by"`
	CreatedAt                          time.Time          `bson:"created_at"`
	UpdatedAt                          time.Time          `bson:"updated_at"`
}

type LeaveRevertRequest struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	LeaveObjId      string             `bson:"leave_obj_id"`
	LeaveHumanId    string             `bson:"leave_human_id"`
	EmployeeId      string             `bson:"employee_id"`
	EmployeeName    string             `bson:"employee_name"`
	EmployeeEmail   string             `bson:"employee_email"`
	LeaveType       string             `bson:"leave_type"`
	LeaveStartDate  string             `bson:"leave_start_date"`
	LeaveEndDate    string             `bson:"leave_end_date"`
	AmountOfLeaves  float32            `bson:"amount_of_leaves"`
	InitiatorId     string             `bson:"initiator_id"`
	InitiatorName   string             `bson:"initiator_name"`
	InitiatorRole   string             `bson:"initiator_role"`
	ApproverId      string             `bson:"approver_id"`
	ApproverName    string             `bson:"approver_name"`
	ApproverEmail   string             `bson:"approver_email"`
	Reason          string             `bson:"reason"`
	Status          string             `bson:"status"`
	DecisionRemarks string             `bson:"decision_remarks"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
	DecidedAt       time.Time          `bson:"decided_at"`
}
