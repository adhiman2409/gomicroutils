package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentDetails struct {
	Id                     primitive.ObjectID `json:"_id"`
	TenantId               string             `json:"tenant_id"`
	PrimaryAdminId         string             `json:"primary_admin_id"`
	PrimaryAdminName       string             `json:"primary_admin_name"`
	PrimaryAdminEmail      string             `json:"primary_admin_email"`
	SecondaryAdminId       string             `json:"secondary_admin_id"`
	SecondaryAdminName     string             `json:"secondary_admin_name"`
	SecondaryAdminEmail    string             `json:"secondary_admin_email"`
	SubscriptionId         string             `json:"subscription_id"`
	PlanId                 string             `json:"plan_id"`
	Day                    int                `json:"day"`
	Month                  int                `json:"month"`
	Year                   int                `json:"year"`
	PaymentDueDate         time.Time          `json:"payment_due_date"`
	PaymentDate            time.Time          `json:"payment_date"`
	IsOnTimePayment        bool               `json:"is_on_time_payment"`
	PaymentId              string             `json:"payment_id"`
	PaymentStatus          string             `json:"payment_status"`
	PaymentCurrency        string             `json:"payment_currency"`
	PaymentMethod          string             `json:"payment_method"`
	ActualAmount           float32            `json:"actual_amount"`
	PaymentAmount          float32            `json:"payment_amount"`
	PendingAmount          float32            `json:"pending_amount"`
	IsOverdue              bool               `json:"is_overdue"`
	OverdueDays            int                `json:"overdue_days"`
	InterestOnOverdue      float32            `json:"interest_on_overdue"`
	OverdueInterestRate    float64            `json:"overdue_interest_rate"`
	OverdueInterestAmount  float32            `json:"overdue_interest_amount"`
	OverdueAmount          float32            `json:"overdue_amount"`
	FirstReminderMailSent  bool               `json:"first_reminder_mail_sent"`
	SecondReminderMailSent bool               `json:"second_reminder_mail_sent"`
	ThirdReminderMailSent  bool               `json:"third_reminder_mail_sent"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
}

type SubscriptionDetails struct {
	Id                   primitive.ObjectID `json:"_id"`
	SubscriptionID       string             `json:"subscription_id"`
	TenantId             string             `json:"tenant_id"`
	PrimaryAdminId       string             `json:"primary_admin_id"`
	PrimaryAdminName     string             `json:"primary_admin_name"`
	PrimaryAdminEmail    string             `json:"primary_admin_email"`
	SecondaryAdminId     string             `json:"secondary_admin_id"`
	SecondaryAdminName   string             `json:"secondary_admin_name"`
	SecondaryAdminEmail  string             `json:"secondary_admin_email"`
	PlanId               string             `json:"plan_id"`
	Name                 string             `json:"name"`
	Amount               float32            `json:"amount"`
	Currency             string             `json:"currency"`
	Country              string             `json:"country"`
	Discount             DiscountDetails    `json:"discount"`
	DiscountedAmount     float32            `json:"discounted_amount"`
	PaymentFrequency     Frequency          `json:"payment_frequency"`
	StartDate            int64              `json:"start_date"`
	EndDate              int64              `json:"end_date"`
	Status               string             `json:"status"`
	CreatedAt            int64              `json:"created_at"`
	UpdatedAt            int64              `json:"updated_at"`
	NextBillingDate      int64              `json:"next_billing_date"`
	LastBillingDate      int64              `json:"last_billing_date"`
	PaymentHistory       []PaymentDetails   `json:"payment_history"`
	TotalAmount          float32            `json:"total_amount"`
	TotalPaidAmount      float32            `json:"total_paid_amount"`
	TotalPendingAmount   float32            `json:"total_pending_amount"`
	TotalRefundedAmount  float32            `json:"total_refunded_amount"`
	HaveTIN              bool               `json:"have_tin"`
	TaxIdentificationNum string             `json:"tax_identification_num"`
}

type TaxDetails struct {
	Id         primitive.ObjectID `json:"_id"`
	Year       int                `json:"year"`
	Type       string             `json:"type"` // e.g., "GST", "VAT"
	Name       string             `json:"name"`
	Country    string             `json:"country"`
	Percentage float64            `json:"percentage"`
}

type DiscountDetails struct {
	Id          primitive.ObjectID `json:"_id"`
	DiscountID  string             `json:"discount_id"`
	Code        string             `json:"code"`
	Type        string             `json:"type"`       // percentage or fixed
	Percentage  float64            `json:"percentage"` // applicable if DiscountType is percentage
	Amount      int                `json:"amount"`     // aplicable if DiscountType is fixed
	Currency    string             `json:"currency"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     time.Time          `json:"end_date"`
}
