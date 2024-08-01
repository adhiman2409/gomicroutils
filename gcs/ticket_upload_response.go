package gcs

type TicketFileUploadResponse struct {
	TicketId      string   `json:"ticketId,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	Department    string   `json:"department,omitempty"`
	DocName       []string `json:"doc_name,omitempty"`
	DocType       string   `json:"doc_type,omitempty"`
	TicketPathArr []string `json:"ticket_path_arr,omitempty"`
}
