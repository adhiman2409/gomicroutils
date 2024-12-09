package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmpAttendanceRegularization struct {
	ID                        primitive.ObjectID `bson:"_id"`
	EmployeeId                string             `bson:"employee_id"`
	EmployeeName              string             `bson:"employee_name"`
	EmployeeEmail             string             `bson:"employee_email"`
	Department                string             `bson:"department"`
	Domain                    string             `bson:"domain"`
	Day                       int32              `bson:"day"`
	Month                     int32              `bson:"month"`
	Year                      int32              `bson:"year"`
	CheckInTime               string             `bson:"check_in_time"`
	CheckOutTime              string             `bson:"check_out_time"`
	OldWorkingHours           float32            `bson:"old_working_hours"`
	NewWorkingHours           float32            `bson:"new_working_hours"`
	TotalWorkingHours         float32            `bson:"total_working_hours"`
	ApprovedNewWorkingHours   float32            `bson:"approved_new_working_hours"`
	ApprovedTotalWorkingHours float32            `bson:"approved_total_working_hours"`
	IsHalfDay                 bool               `bson:"is_half_day"`
	EmployeeRemarks           string             `bson:"employee_remarks"`
	Status                    string             `bson:"status"`
	RequestedOn               string             `bson:"requested_on"`
	RegularizedOn             string             `bson:"regularized_on"`
	ManagerRemarks            string             `bson:"manager_remarks"`
	IsHalfDayApproval         bool               `bson:"is_half_day_approval"`
	DocumentURL               string             `bson:"document_url"`
	ManagerId                 string             `bson:"manager_id"`
	ManagerName               string             `bson:"manager_name"`
	ManagerEmail              string             `bson:"manager_email"`
	CreatedAt                 time.Time          `bson:"created_at"`
	UpdatedAt                 time.Time          `bson:"updated_at"`
}
