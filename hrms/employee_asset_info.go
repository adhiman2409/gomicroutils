package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetTransitInfo struct {
	TransitType  string `bson:"asset_handover_type"` // handover, return
	TransitMode  string `bson:"asset_handover_mode"` // Pickup ByEmployee, Courier, Runner, Deleviry By Hand
	Remarks      string `bson:"asset_handover_remarks"`
	StartDate    string `bson:"asset_handover_start_date"`
	DeliveryDate string `bson:"asset_handover_delivery_date"`
	TrackingId   string `bson:"asset_handover_tracking_id"`
	TrackingURL  string `bson:"asset_handover_tracking_url"`
	InvoiceURL   string `bson:"asset_handover_invoice_url"`
}

type EmployeeAssetInfo struct {
	ID                        primitive.ObjectID `bson:"_id"`
	EmployeeID                string             `bson:"employee_id"`
	EmployeeEmail             string             `bson:"employee_email"`
	EmployeeName              string             `bson:"employee_name"`
	EmployeeLocation          string             `bson:"employee_location"`
	Department                string             `bson:"department"`
	AssetPSId                 string             `bson:"asset_ps_id"`
	AssetSerialNumber         string             `bson:"asset_serial_number"`
	AssetName                 string             `bson:"asset_name"`
	AssetType                 string             `bson:"asset_type"`
	AssetDescription          string             `bson:"asset_description"`
	IsAllocated               bool               `bson:"is_allocated"`
	AllocationDate            string             `bson:"allocation_date"`
	AllocationType            string             `bson:"allocation_type"`             //Permanent/Temporary
	AllocationDurationInDays  string             `bson:"allocation_duration_in_days"` //if Temporary
	InTransit                 bool               `bson:"in_transit"`
	AssetTransitInfoList      []AssetTransitInfo `bson:"asset_transit_info_list"`
	ReturnDate                string             `bson:"return_date"`
	AssetImages               []string           `bson:"asset_images"`
	AssetDocuments            []string           `bson:"asset_documents"`
	ImagesUpdatedOn           string             `bson:"image_updated_on"`
	AssetAggrementStatus      string             `bson:"asset_aggrement_status"`
	TermsAndConditions        string             `bson:"terms_and_conditions"`
	EmployeeConsentStatus     string             `bson:"employee_consent_status"`
	EmployeeConsentReceivedAt string             `bson:"employee_consent_received_at"`
	Remarks                   string             `bson:"remarks"`
	CreatedAt                 time.Time          `bson:"created_at"`
	UpdatedAt                 time.Time          `bson:"updated_at"`
}
