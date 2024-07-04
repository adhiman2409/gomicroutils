package gcs

type InventoryFileUploadResponse struct {
	BucketId    string   `json:"bucketId,omitempty"`
	StoreId     string   `json:"storeId,omitempty"`
	Domain      string   `json:"domain,omitempty"`
	Department  string   `json:"department,omitempty"`
	DocName     []string `json:"doc_name,omitempty"`
	DocCategory []string `json:"doc_category,omitempty"`
	DocType     string   `json:"doc_type,omitempty"`
	DocPath     []string `json:"doc_path,omitempty"`
}
