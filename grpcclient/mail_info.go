package grpcclient

type MailRequest struct {
	From         string            `json:"from"`
	To           string            `json:"to"`
	Subject      string            `json:"subject"`
	DataMap      map[string]string `json:"data_map"`
	TemplateName string            `json:"template_name"`
	AckRequired  bool              `json:"ack_required"`
	Priority     string            `json:"priority"`
}

type MailResponse struct {
	Status        bool   `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlation_id"`
}
