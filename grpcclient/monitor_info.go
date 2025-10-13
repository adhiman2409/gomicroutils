package grpcclient

type EmployeeCheckInCheckOutRequest struct {
	Day          int32  `json:"day"`
	Month        int32  `json:"month"`
	Year         int32  `json:"year"`
	EmployeeId   string `json:"employee_id"`
	IsCheckedIn  bool   `json:"is_checked_in"`
	IsCheckedOut bool   `json:"is_checked_out"`
	Domain       string `json:"domain"`
}

type EmployeeCheckInCheckOutResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
