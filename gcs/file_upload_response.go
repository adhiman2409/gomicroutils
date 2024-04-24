package gcs

type FileUploadResponse struct {
	Name         string `json:"name,omitempty"`
	DocumentType string `json:"document_type,omitempty"`
	Path         string `json:"path,omitempty"`
}
