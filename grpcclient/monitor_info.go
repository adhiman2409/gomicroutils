package grpcclient

type EmployeeCheckInCheckOutRequest struct {
	EmployeeId   string `json:"employee_id"`
	IsCheckedIn  bool   `json:"is_checked_in"`
	IsCheckedOut bool   `json:"is_checked_out"`
	Domain       string `json:"domain"`
}

type EmployeeCheckInCheckOutResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
