package grpcclient

type MailRequest struct {
	From         string            `json:"from"`
	To           string            `json:"to"`
	Subject      string            `json:"subject"`
	DataMap      map[string]string `json:"data_map"`
	TemplateName string            `json:"template_name"`
	AckRequired  bool              `json:"ack_required"`
	Priority     string            `json:"priority"`
	Cc           []string          `json:"cc"`     // Added CC field for carbon copy recipients
	Domain       string            `json:"domain"` // Added Domain field for specifying the domain
}

type NewMailRequest struct {
	From         string         `json:"from"`
	To           string         `json:"to"`
	Subject      string         `json:"subject"`
	DataMap      map[string]any `json:"data_map"`
	TemplateName string         `json:"template_name"`
	AckRequired  bool           `json:"ack_required"`
	Priority     string         `json:"priority"`
	Cc           []string       `json:"cc"`     // Added CC field for carbon copy recipients
	Domain       string         `json:"domain"` // Added Domain field for specifying the domain
}

type MailResponse struct {
	Status        bool   `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlation_id"`
}
