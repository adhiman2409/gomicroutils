package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetInfo struct {
	ID                  primitive.ObjectID `bson:"_id"`
	AssetId             string             `bson:"asset_id"`
	AssetSerialNumber   string             `bson:"asset_serial_number"`
	AssetCategory       string             `bson:"asset_category"` //Hardware, Software
	AssetType           string             `bson:"asset_type"`     //laptop, screen, keyboard, mouse etc
	AssetBrandName      string             `bson:"asset_brand_name"`
	AssetDescription    string             `bson:"asset_description"`
	AssetModel          string             `bson:"asset_model"`
	AssetConfig         string             `bson:"asset_config"`
	Accessories         []string           `bson:"accessories"`
	ProcurementType     string             `bson:"procurement_type"` // Purchased, Rental, Subscription
	RenewalFrequency    string             `bson:"renewal_frequency"`
	SubscriptionEndDate time.Time          `bson:"subscription_end_date"`
	PriceInINR          float64            `bson:"price_in_inr"`
	AssetImages         []string           `bson:"asset_images"`
	AssetDocuments      []string           `bson:"asset_documents"`
	ImagesUpdatedOn     time.Time          `bson:"image_updated_on"`
	Status              string             `bson:"status"` // Available, Allocated, Under Maintenance, Retired
	IsInsured           bool               `bson:"is_insured"`
	Remarks             []string           `bson:"remarks"`
	WarrantyEndDate     time.Time          `bson:"warranty_end_date"`
	WarrantyProvider    string             `bson:"warranty_provider"`
	CreatedAt           time.Time          `bson:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at"`
}

type AssetInsuranceInfo struct {
	AssetId             string                `bson:"asset_id"`
	AssetSerialNumber   string                `bson:"asset_serial_number"`
	InsuranceId         string                `bson:"insurance_id"`
	ProviderName        string                `bson:"provider_name"`
	PolicyNumber        string                `bson:"policy_number"`
	CoverageAmount      float64               `bson:"coverage_amount"`
	PremiumAmount       float64               `bson:"premium_amount"`
	StartDate           time.Time             `bson:"start_date"`
	EndDate             time.Time             `bson:"end_date"`
	CoverageType        string                `bson:"coverage_type"`  // Full, Partial, Theft-only
	ContactNumber       string                `bson:"contact_number"` // Insurance provider's contact
	TermsAndConditions  string                `bson:"terms_and_conditions"`
	IsActive            bool                  `bson:"is_active"` // Whether insurance is currently valid
	AssetAllocationInfo []AssetAllocationInfo `bson:"asset_allocation_info"`
	AssetTransitInfo    []AssetTransitInfo    `bson:"asset_transit_info"`
	CreatedAt           time.Time             `bson:"created_at"`
	UpdatedAt           time.Time             `bson:"updated_at"`
}

type AssetAllocationInfo struct {
	AssetId             string    `bson:"asset_id"`
	AssetSerialNumber   string    `bson:"asset_serial_number"`
	EmployeeID          string    `bson:"employee_id"`
	EmployeeEmail       string    `bson:"employee_email"`
	EmployeeName        string    `bson:"employee_name"`
	EmployementStatus   string    `bson:"employement_status"`
	AllocatedAt         time.Time `bson:"allocated_at"`
	AllocationDocuments []string  `bson:"allocation_documents"`
	AllocatedTill       time.Time `bson:"allocated_till"`
	AllocationType      string    `bson:"allocation_type"` //permanent, temporary
	Status              string    `bson:"status"`
	ReturnedAt          time.Time `bson:"returned_at"`
}

type AllocationHistory struct {
	AssetId             string    `bson:"asset_id"`
	AssetSerialNumber   string    `bson:"asset_serial_number"`
	EmployeeID          string    `bson:"employee_id"`
	EmployeeEmail       string    `bson:"employee_email"`
	EmployeeName        string    `bson:"employee_name"`
	AllocatedAt         time.Time `bson:"allocated_at"`
	AllocationDocuments []string  `bson:"allocation_documents"`
	AllocatedTill       time.Time `bson:"allocated_till"`
	AllocationType      string    `bson:"allocation_type"`
	Status              string    `bson:"status"`
	ReturnedAt          time.Time `bson:"returned_at"`
}

type AssetTransitInfo struct {
	AssetId           string    `bson:"asset_id"`
	AssetSerialNumber string    `bson:"asset_serial_number"`
	TransitType       string    `bson:"transit_type"` // handover, return
	TransitMode       string    `bson:"transit_mode"` // Pickup ByEmployee, Courier, Runner, Deleviry By Hand
	Remarks           string    `bson:"remark"`
	DispatchDate      time.Time `bson:"dispatch_date"`
	DeliveryDate      time.Time `bson:"delivery_date"`
	TrackingId        string    `bson:"tracking_id"`
	TrackingURL       string    `bson:"tracking_url"`
	InvoiceURL        string    `bson:"invoice_url"`
	CourierCompany    string    `bson:"courier_company"`
	CreatedAt         time.Time `bson:"created_at"`
}
