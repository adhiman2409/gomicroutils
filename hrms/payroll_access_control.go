package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type PayrollAccessControl struct {
	Id                 primitive.ObjectID `json:"_id"`
	EmployeeID         string             `json:"employee_id"`
	EmployeeName       string             `json:"employee_name"`
	EmploymentStatus   string             `json:"employment_status"`
	IsAccessController bool               `json:"is_access_controller"`
	IsInitiator        bool               `json:"is_initiator"`
	IsApprover         bool               `json:"is_approver"`
	IsProcessor        bool               `json:"is_processor"`
}
