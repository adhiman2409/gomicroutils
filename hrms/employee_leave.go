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

type EmployeeLeaveStatus struct {
	EmployeeId  string `bson:"employee_id"`
	LeaveType   string `bson:"leave_type"`
	LeaveStatus string `bson:"leave_status"`
	Day         string `bson:"day"`
	Month       string `bson:"month"`
	Year        string `bson:"year"`
	Remarks     string `bson:"remarks"`
}
