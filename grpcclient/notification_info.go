package grpcclient

type NotificationRequest struct {
	EmployeeId     string            `json:"employee_id"`
	Title          string            `json:"title"`
	Body           string            `json:"body"`
	ImageUrl       string            `json:"image_url"`
	DataMap        map[string]string `json:"data_map"`
	Domain         string            `json:"domain"`
	Group          string            `json:"group"`
	Topic          string            `json:"topic"`
	BroadcastToAll bool              `json:"broadcast_to_all"`
}

type NotificationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
