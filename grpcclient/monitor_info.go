package grpcclient

type EmployeeCheckInCheckOutRequest struct {
	Day               int32  `json:"day"`
	Month             int32  `json:"month"`
	Year              int32  `json:"year"`
	EmployeeId        string `json:"employee_id"`
	IsCheckedIn       bool   `json:"is_checked_in"`
	CheckInTimestamp  string `json:"check_in_timestamp"`
	IsCheckedOut      bool   `json:"is_checked_out"`
	CheckOutTimestamp string `json:"check_out_timestamp"`
	Domain            string `json:"domain"`
}

type EmployeeCheckInCheckOutResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
