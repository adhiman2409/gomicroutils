package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConfigFields struct {
	Field string `bson:"field"`
	Value string `bson:"value"`
}

type Category struct {
	Category      string   `bson:"category"`
	SubCategories []string `bson:"sub_categories"`
}

type NewAssetConfig struct {
	Category    string    `bson:"category"`
	SubCategory string    `bson:"sub_category"`
	ConfigField []string  `bson:"config_field"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type VendorInfo struct {
	ID          primitive.ObjectID `bson:"_id"`
	VendorName  string             `bson:"vendor_name"`
	ContactName string             `bson:"contact_name"`
	Email       string             `bson:"email"`
	Phone       string             `bson:"phone"`
	Address     Address            `bson:"address"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Tags        []string           `bson:"tags"`
}

type AssetReturnInfo struct {
	AssetId              string          `bson:"asset_id"`
	ReturnEntryUpdatedBy EmployeeSummary `bson:"return_entry_updated_by"`
	ReturnedTo           EmployeeSummary `bson:"returned_to"`
	Remarks              []string        `bson:"remarks"`
	ReturnDocuments      []string        `bson:"return_documents"`
	ReturnedAt           time.Time       `bson:"returned_at"`
}

type AssetInfo struct {
	ID                               primitive.ObjectID          `bson:"_id"`
	AssetId                          string                      `bson:"asset_id"`
	UniqueId                         string                      `bson:"unique_id"`
	DepartmentId                     string                      `bson:"department_id"`
	DepartmentName                   string                      `bson:"department_name"`
	ProjectId                        string                      `bson:"project_id"`
	ProjectName                      string                      `bson:"project_name"`
	ProjectHeadId                    string                      `bson:"project_head_id"`
	ProjectHeadName                  string                      `bson:"project_head_name"`
	ProjectHeadEmailId               string                      `bson:"project_head_email_id"`
	AssetSerialNumber                string                      `bson:"asset_serial_number"`
	AssetCategory                    string                      `bson:"asset_category"` //Hardware, Software
	AssetSubCategory                 string                      `bson:"asset_sub_category"`
	AssetType                        string                      `bson:"asset_type"` //laptop, screen, keyboard, mouse etc
	AssetBrandName                   string                      `bson:"asset_brand_name"`
	AssetDescription                 string                      `bson:"asset_description"`
	AssetModel                       string                      `bson:"asset_model"`
	AssetConfig                      AssetConfig                 `bson:"asset_config"`
	AssetLocation                    string                      `bson:"asset_location"`
	InvoiceNumber                    string                      `bson:"invoice_number"`
	AssetNewConfig                   []ConfigFields              `bson:"asset_new_config"`
	Accessories                      []string                    `bson:"accessories"`
	ProcurementType                  string                      `bson:"procurement_type"` // Purchased, Rental, Subscription
	RenewalFrequency                 Frequency                   `bson:"renewal_frequency"`
	SubscriptionEndDate              time.Time                   `bson:"subscription_end_date"`
	PriceInINR                       float64                     `bson:"price_in_inr"`
	Price                            float64                     `bson:"price"`
	Currency                         string                      `bson:"currency"` // INR, USD, EUR etc
	ProcurementDate                  time.Time                   `bson:"procurement_date"`
	ProcurementVendor                string                      `bson:"procurement_vendor"`
	ProcurementInvoiceURL            string                      `bson:"procurement_invoice_url"`
	AssetImages                      []string                    `bson:"asset_images"`
	AssetDocuments                   []string                    `bson:"asset_documents"`
	ImagesUpdatedOn                  time.Time                   `bson:"image_updated_on"`
	Status                           AssetStatus                 `bson:"status"` // Available, Allocated, Under Maintenance, Retired
	IsInsured                        bool                        `bson:"is_insured"`
	IsPhysicallyVerified             bool                        `bson:"is_physically_verified"`
	AssetInsuranceInfo               AssetInsuranceInfo          `bson:"asset_insurance_info"`
	AssetPhysicalVerificationInfo    AssetPhysicalVerification   `bson:"asset_physical_verification_info"`
	AssetPhysicalVerificationHistory []AssetPhysicalVerification `bson:"asset_physical_verification_history"`
	AssetDisposalInfo                AssetDisposalInfo           `bson:"asset_disposal_info"`
	Remarks                          []string                    `bson:"remarks"`
	WarrantyEndDate                  time.Time                   `bson:"warranty_end_date"`
	WarrantyProvider                 string                      `bson:"warranty_provider"`
	AllocationInfo                   []AssetAllocationInfo       `bson:"allocation_info"`
	AssetReturnInfo                  AssetReturnInfo             `bson:"asset_return_info"`
	AssetReturnHistory               []AssetReturnInfo           `bson:"asset_return_history"`
	AllocationHistory                []AllocationHistory         `bson:"allocation_history"`
	VendorInfo                       VendorInfo                  `bson:"vendor_info"`
	CreatedAt                        time.Time                   `bson:"created_at"`
	UpdatedAt                        time.Time                   `bson:"updated_at"`
}

type AssetConfig struct {
	ScreenSize    string `bson:"screen_size"`
	Processor     string `bson:"processor"`
	SystemRam     string `bson:"system_ram"`
	SystemStorage string `bson:"system_storage"`
}

type AssetInsuranceInfo struct {
	AssetId            string    `bson:"asset_id"`
	AssetSerialNumber  string    `bson:"asset_serial_number"`
	InsuranceId        string    `bson:"insurance_id"`
	ProviderName       string    `bson:"provider_name"`
	PolicyNumber       string    `bson:"policy_number"`
	CoverageAmount     float64   `bson:"coverage_amount"`
	PremiumAmount      float64   `bson:"premium_amount"`
	StartDate          time.Time `bson:"start_date"`
	EndDate            time.Time `bson:"end_date"`
	CoverageType       string    `bson:"coverage_type"`        // Full, Partial, Theft-only, acidental
	ContactNumber      string    `bson:"contact_number"`       // Insurance provider's contact
	TermsAndConditions string    `bson:"terms_and_conditions"` // Policy DOcuments url []
	Remarks            []string  `bson:"remarks"`
	Attachments        []string  `bson:"attachments"`
	IsActive           bool      `bson:"is_active"` // Whether insurance is currently valid
	CreatedAt          time.Time `bson:"created_at"`
	UpdatedAt          time.Time `bson:"updated_at"`
}

type AssetDisposalInfo struct {
	DisposalType        string    `bson:"disposal_type"` // Sell, Recycle, Donate, Trash
	DisposedById        string    `bson:"disposed_by_id"`
	DisposedByName      string    `bson:"disposed_by_name"`
	DisposalDate        time.Time `bson:"disposal_date"`
	DisposalCost        float64   `bson:"disposal_cost"`
	AssetDisposalAmount float64   `bson:"asset_disposal_amount"`
	Remarks             []string  `bson:"remarks"`
	Attachments         []string  `bson:"attachments"`
	CreatedAt           time.Time `bson:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at"`
}

type AssetPhysicalVerification struct {
	VerificationDate time.Time `bson:"verification_date"`
	VerifiedById     string    `bson:"verified_by_id"`
	VerifiedByName   string    `bson:"verified_by_name"`
	Remarks          []string  `bson:"remarks"`
	Attachments      []string  `bson:"attachments"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
}

type AssetAllocationInfo struct {
	AssetId             string      `bson:"asset_id"`
	AssetSerialNumber   string      `bson:"asset_serial_number"`
	EmployeeID          string      `bson:"employee_id"`
	EmployeeEmail       string      `bson:"employee_email"`
	EmployeeName        string      `bson:"employee_name"`
	EmployeeLocation    string      `bson:"employee_location"`
	EmployementStatus   string      `bson:"employement_status"`
	AllocatedAt         time.Time   `bson:"allocated_at"`
	AllocationDocuments []string    `bson:"allocation_documents"`
	AllocatedTill       time.Time   `bson:"allocated_till"`
	AllocationType      string      `bson:"allocation_type"` //permanent, temporary
	Remarks             []string    `bson:"remarks"`
	Status              AssetStatus `bson:"status"`
	ReturnedAt          time.Time   `bson:"returned_at"`
}

type AllocationHistory struct {
	AssetId             string      `bson:"asset_id"`
	AssetSerialNumber   string      `bson:"asset_serial_number"`
	EmployeeID          string      `bson:"employee_id"`
	EmployeeEmail       string      `bson:"employee_email"`
	EmployeeName        string      `bson:"employee_name"`
	EmployeeLocation    string      `bson:"employee_location"`
	AllocatedAt         time.Time   `bson:"allocated_at"`
	AllocationDocuments []string    `bson:"allocation_documents"`
	AllocatedTill       time.Time   `bson:"allocated_till"`
	AllocationType      string      `bson:"allocation_type"`
	Remarks             []string    `bson:"remarks"`
	Status              AssetStatus `bson:"status"`
	ReturnedAt          time.Time   `bson:"returned_at"`
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
