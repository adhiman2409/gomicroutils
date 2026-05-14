package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductLocationIdCounter struct {
	ProductId string `bson:"product_id"`
	Counter   int64  `bson:"counter"`
}

type LeadOfferingIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type LeadOffering struct {
	Id                   primitive.ObjectID    `bson:"_id"`
	ProductId            string                `bson:"product_id"`
	ProductName          string                `bson:"product_name"`
	ProductDescription   string                `bson:"product_description"`
	Type                 string                `bson:"type"` // Product / Service / Solution
	SubTypes             []LeadOfferingSubType `bson:"sub_types"`
	IsActive             bool                  `bson:"is_active"`
	CreatedAt            time.Time             `bson:"created_at"`
	UpdatedAt            time.Time             `bson:"updated_at"`
	ProductLocation      []ProductLocation     `bson:"product_location"`
	InvoiceRequestFields map[string]any        `bson:"invoice_request_fields"`
}

type ProductLocation struct {
	ProductLocationId   string          `bson:"product_location_id"`
	TeamLeader          string          `bson:"team_leader"`
	TeamMembers         []LOCTeamMember `bson:"team_members"`
	LocationCountry     string          `bson:"location_country"`
	LocationState       string          `bson:"location_state"`
	LocationCity        string          `bson:"location_city"`
	LocationLat         float64         `bson:"location_lat"`
	LocationLng         float64         `bson:"location_lng"`
	BasePrice           float64         `bson:"base_price"`
	MinPrice            float64         `bson:"min_price"`
	Currency            string          `bson:"currency,"`
	AutomaticAssignment bool            `bson:"automatic_assignment,"`
	LicenseType         string          `bson:"license_type"` // One Time Purchase / Monthly / Quarterly / Yearly
	CreateAt            time.Time       `bson:"created_at"`
	UpdatedAt           time.Time       `bson:"updated_at"`
	LastAssignedIndex   int             `bson:"last_assigned_index"`
}

type LOCTeamMember struct {
	MemberId              string  `bson:"member_id"`
	MemberLocationCountry string  `bson:"location_country"`
	MemberLocationState   string  `bson:"location_state"`
	MemberLocationCity    string  `bson:"location_city"`
	MemberLocationLat     float64 `bson:"location_lat"`
	MemberLocationLng     float64 `bson:"location_lng"`
}

type LeadOfferingSubType struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description,omitempty"`
	IsActive    bool               `bson:"is_active"`
}

type LeadWithCompany struct {
	CompanyId   string   `bson:"company_id"`
	CompanyName string   `bson:"company_name"`
	Lead        LeadInfo `bson:"lead"`
}

type LeadStage struct {
	Id             primitive.ObjectID `bson:"_id"`
	StageName      string             `bson:"stage_name"`
	Index          int                `bson:"index"`
	IsStartState   bool               `bson:"is_start_state"`
	IsEndState     bool               `bson:"is_end_state"`
	IsDroppedState bool               `bson:"is_dropped_state"`
	IsDNDState     bool               `bson:"is_dnd_state"`
	SubStages      []string           `bson:"sub_stages"`
}

type MeetingStatusConfig struct {
	Id         primitive.ObjectID `bson:"_id"`
	StatusName string             `bson:"status_name"`
	IsActive   bool               `bson:"is_active"`
}

type Vendor struct {
	Id          primitive.ObjectID  `bson:"_id"`
	VendorId    string              `bson:"vendor_id"`
	VendorName  string              `bson:"vendor_name"`
	Address     VendorAddressDB     `bson:"address"`
	BankDetails VendorBankDetailsDB `bson:"bank_details"`
	IsActive    bool                `bson:"is_active"`
	CreatedAt   time.Time           `bson:"created_at"`
	UpdatedAt   time.Time           `bson:"updated_at"`
}

type VendorAddressDB struct {
	Country string `bson:"country"`
	State   string `bson:"state"`
	City    string `bson:"city"`
}

type VendorBankDetailsDB struct {
	BankName      string `bson:"bank_name"`
	AccountNumber string `bson:"account_number"`
	IFSECode      string `bson:"ifsc_code"`
	AccountHolder string `bson:"account_holder"`
	AccountType   string `bson:"account_type"`
}

// InvoiceRequestInfo is the DB model for a sales-to-finance invoice request.
// ProductFields holds product-specific data; its schema is extensible without
// requiring a model change when new product types are introduced.
type InvoiceRequestInfo struct {
	ID              primitive.ObjectID           `bson:"_id,omitempty"`
	RequestId       string                       `bson:"request_id"`
	ProductId       string                       `bson:"product_id"`
	ProductType     string                       `bson:"product_type"`
	InvoiceType     string                       `bson:"invoice_type"`
	BillingType     string                       `bson:"billing_type"`
	CompanyId       string                       `bson:"company_id"`
	CompanyName     string                       `bson:"company_name"`
	ProductFields   map[string]any               `bson:"product_fields"`
	Attachments     []InvoiceRequestAttachmentDB `bson:"attachments"`
	PaymentTerms    string                       `bson:"payment_terms"`
	PaymentStatus   string                       `bson:"payment_status"`
	RequestStatus   string                       `bson:"request_status"`
	InvoiceId       string                       `bson:"invoice_id"`
	RequestedBy     string                       `bson:"requested_by"`
	RequestedByName string                       `bson:"requested_by_name"`
	Remarks         string                       `bson:"remarks"`
	CreatedAt       time.Time                    `bson:"created_at"`
	UpdatedAt       time.Time                    `bson:"updated_at"`
}

type InvoiceRequestAttachmentDB struct {
	AttachmentType string `bson:"attachment_type"`
	FileURL        string `bson:"file_url"`
	FileName       string `bson:"file_name"`
}

type InvoiceRequestIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}
