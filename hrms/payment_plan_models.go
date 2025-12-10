package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentDetails struct {
	Id              primitive.ObjectID `bson:"_id"`
	TenantId        string             `bson:"tenant_id"`
	PaymentId       string             `bson:"payment_id"`
	InvoiceId       string             `bson:"invoice_id"`
	SubscriptionId  string             `bson:"subscription_id"`
	PaymentDueDate  time.Time          `bson:"payment_due_date"`
	PaymentDate     time.Time          `bson:"payment_date"`
	IsOnTimePayment bool               `bson:"is_on_time_payment"`
	PaymentStatus   string             `bson:"payment_status"`
	PaymentCurrency string             `bson:"payment_currency"`
	PaymentMode     PaymentMethod      `bson:"payment_mode"`
	PaymentMeta     interface{}        `bson:"payment_meta"`
	ActualAmount    float64            `bson:"actual_amount"`
	PaymentAmount   float64            `bson:"payment_amount"`
	PendingAmount   float64            `bson:"pending_amount"`
	IsOverdue       bool               `bson:"is_overdue"`
	OverdueDays     int                `bson:"overdue_days"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type RazorpayMeta struct {
	RazorpayPaymentId string `bson:"razorpay_payment_id"`
	RazorpayOrderId   string `bson:"razorpay_order_id"`
	PaymentMode       string `bson:"payment_mode"`
}

type ChequeMeta struct {
	ChequeNumber string `bson:"cheque_number"`
	BankName     string `bson:"bank_name"`
	Branch       string `bson:"branch"`
	IssuerName   string `bson:"issuer_name"`
}

type BankTransferMeta struct {
	UTRNumber    string `bson:"utr_number"`
	BankName     string `bson:"bank_name"`
	AccountNo    string `bson:"account_no"`
	TransferMode string `bson:"transfer_mode"` // NEFT, RTGS, IMPS
}

type CashMeta struct {
	ReceiverName string `bson:"receiver_name"`
}

type SubscriptionDetails struct {
	Id                     primitive.ObjectID `bson:"_id"`
	SubscriptionID         string             `bson:"subscription_id"`
	TenantId               string             `bson:"tenant_id"`
	TenantName             string             `bson:"tenant_name"`
	Domain                 string             `bson:"domain"`
	PrimaryAdminId         string             `bson:"primary_admin_id"`
	PrimaryAdminName       string             `bson:"primary_admin_name"`
	PrimaryAdminEmail      string             `bson:"primary_admin_email"`
	SecondaryAdminId       string             `bson:"secondary_admin_id"`
	SecondaryAdminName     string             `bson:"secondary_admin_name"`
	SecondaryAdminEmail    string             `bson:"secondary_admin_email"`
	PlanId                 string             `bson:"plan_id"`
	Name                   string             `bson:"name"`
	BaseAmount             float64            `bson:"base_amount"`
	Currency               string             `bson:"currency"`
	Country                string             `bson:"country"`
	InitialLicenseCount    int                `bson:"initial_license_count"`
	LastInvoiceGeneratedAt time.Time          `bson:"last_invoice_generated_at"`
	IsDiscounted           bool               `bson:"is_discounted"`
	Discount               DiscountDetails    `bson:"discount"`
	AmountAfterDiscount    float64            `bson:"amount_after_discount"`
	PaymentFrequency       string             `bson:"payment_frequency"`
	StartDate              time.Time          `bson:"start_date"`
	EndDate                time.Time          `bson:"end_date"`
	Status                 string             `bson:"status"`
	CreatedAt              time.Time          `bson:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at"`
	NextBillingDate        time.Time          `bson:"next_billing_date"`
	LastBillingDate        time.Time          `bson:"last_billing_date"`
	PaymentHistory         []PaymentHistory   `bson:"payment_history"`
	TotalAmount            float64            `bson:"total_amount"`
	TotalPaidAmount        float64            `bson:"total_paid_amount"`
	TotalPendingAmount     float64            `bson:"total_pending_amount"`
	TotalRefundedAmount    float64            `bson:"total_refunded_amount"`
	ApplicableTaxIds       []string           `bson:"applicable_tax_ids"`
	HaveTIN                bool               `bson:"have_tin"`
	TaxIdentificationNum   string             `bson:"tax_identification_num"`
	BillingCycleDetails    BillingCycle       `bson:"billing_cycle"`
	NumberOfCycles         int                `bson:"number_of_cycles"`
	RemainingCycles        int                `bson:"remaining_cycles"`
	Notes                  string             `bson:"notes"`
}

type BillingCycle struct {
	ID                  primitive.ObjectID `bson:"_id"`
	BillingCycleID      string             `bson:"billing_cycle_id"`
	Frequency           string             `bson:"frequency"`
	UsageBufferDays     int                `bson:"usage_buffer_days"`
	CycleLengthInMonths int                `bson:"cycle_length_in_months"`
	PaymentBufferDays   int                `bson:"payment_buffer_days"`
}

type TaxDetails struct {
	Id         primitive.ObjectID `bson:"_id"`
	TaxId      string             `bson:"tax_id"`
	Year       int                `bson:"year"`
	Type       string             `bson:"type"` // e.g., "GST", "VAT"
	Name       string             `bson:"name"`
	Country    string             `bson:"country"`
	Percentage float64            `bson:"percentage"`
}

type DiscountDetails struct {
	Id          primitive.ObjectID `bson:"_id"`
	DiscountID  string             `bson:"discount_id"`
	Code        string             `bson:"code"`
	Type        string             `bson:"type"`       // percentage or fixed
	Percentage  float64            `bson:"percentage"` // if percentage type
	Amount      int                `bson:"amount"`     // if fixed discount
	Currency    string             `bson:"currency"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	StartDate   time.Time          `bson:"start_date"`
	EndDate     time.Time          `bson:"end_date"`
}

type Invoice struct {
	Id                             primitive.ObjectID `bson:"_id"`
	InvoiceID                      string             `bson:"invoice_id"`
	InvoiceNumber                  string             `bson:"invoice_number"`
	SubscriptionID                 string             `bson:"subscription_id"`
	TenantId                       string             `bson:"tenant_id"`
	TenantName                     string             `bson:"tenant_name"`
	Domain                         string             `bson:"domain"`
	PrimaryAdminEmail              string             `bson:"primary_admin_email"`
	SecondaryAdminEmail            string             `bson:"secondary_admin_email"`
	BillingCycleType               string             `bson:"billing_cycle_type"`
	BillingPeriodStart             time.Time          `bson:"billing_period_start"`
	BillingPeriodEnd               time.Time          `bson:"billing_period_end"`
	IsFirstInvoice                 bool               `bson:"is_first_invoice"`
	InitialLicenseCount            int                `bson:"initial_license_count"`
	UsersOnLastDay                 int                `bson:"users_on_last_day"`
	NewUsersInPeriod               int                `bson:"new_users_in_period"`
	BufferDays                     int                `bson:"buffer_days"`
	LicenseBaseCost                float64            `bson:"license_base_cost"`
	LicenseCostAfterDiscount       float64            `bson:"license_cost_after_discount"`
	InitialLicenseCost             float64            `bson:"initial_license_cost"`
	LastDayUsersCostBeforeDiscount float64            `bson:"last_day_users_cost_before_discount"`
	NewUsersCostBeforeDiscount     float64            `bson:"new_users_cost_before_discount"`
	GrossTotalBeforeDiscount       float64            `bson:"gross_total_before_discount"`
	LastDayUsersCostAfterDiscount  float64            `bson:"last_day_users_cost_after_discount"`
	NewUsersCostAfterDiscount      float64            `bson:"new_users_cost_after_discount"`
	GrossTotalAfterDiscount        float64            `bson:"gross_total_after_discount"`
	TaxBreakdown                   TaxBreakdown       `bson:"tax_breakdown"`
	TaxAmount                      float64            `bson:"tax_amount"`
	DiscountAmount                 float64            `bson:"discount_amount"`
	NetTotal                       float64            `bson:"net_total"`
	Currency                       string             `bson:"currency"`

	PaymentStatus string    `bson:"payment_status"`
	PaymentMethod string    `bson:"payment_method"`
	PaymentDate   time.Time `bson:"payment_date,omitempty"`
	DueDate       time.Time `bson:"due_date"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
	Notes         string    `bson:"notes"`
}

type TaxBreakdown struct {
	TotalTaxAmount float64   `bson:"total_tax_amount"`
	TaxLines       []TaxLine `bson:"tax_lines"`
}

type TaxLine struct {
	TaxType   string  `bson:"tax_type"`
	TaxRate   float64 `bson:"tax_rate"`
	TaxAmount float64 `bson:"tax_amount"`
}

type PaymentHistory struct {
	InvoiceId        string    `bson:"invoice_id"`
	PaymentId        string    `bson:"payment_id"`
	Amount           float64   `bson:"amount"`
	PaymentStatus    string    `bson:"payment_status"`
	PaymentDate      time.Time `bson:"payment_date"`
	Reminders        Reminders `bson:"reminders"`
	BillingStartDate time.Time `bson:"billing_start_date"`
	BillingEndDate   time.Time `bson:"billing_end_date"`
}

type CountryTax struct {
	ID         primitive.ObjectID `bson:"_id"`
	Country    string             `bson:"country" json:"country"`
	Components []TaxDetails       `bson:"components" json:"components"`
}

type Reminders struct {
	ReminderCount    int         `bson:"reminder_count"`
	LastReminderDate time.Time   `bson:"last_reminder_date"`
	NextReminderDate time.Time   `bson:"next_reminder_date"`
	ReminderDates    []time.Time `bson:"reminder_dates"`
}
