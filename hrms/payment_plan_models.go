package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentDetails struct {
	Id                     primitive.ObjectID `bson:"_id"`
	TenantId               string             `bson:"tenant_id"`
	PrimaryAdminId         string             `bson:"primary_admin_id"`
	PrimaryAdminName       string             `bson:"primary_admin_name"`
	PrimaryAdminEmail      string             `bson:"primary_admin_email"`
	SecondaryAdminId       string             `bson:"secondary_admin_id"`
	SecondaryAdminName     string             `bson:"secondary_admin_name"`
	SecondaryAdminEmail    string             `bson:"secondary_admin_email"`
	SubscriptionId         string             `bson:"subscription_id"`
	PlanId                 string             `bson:"plan_id"`
	Day                    int                `bson:"day"`
	Month                  int                `bson:"month"`
	Year                   int                `bson:"year"`
	PaymentDueDate         time.Time          `bson:"payment_due_date"`
	PaymentDate            time.Time          `bson:"payment_date"`
	IsOnTimePayment        bool               `bson:"is_on_time_payment"`
	PaymentId              string             `bson:"payment_id"`
	PaymentStatus          string             `bson:"payment_status"`
	PaymentCurrency        string             `bson:"payment_currency"`
	PaymentMethod          string             `bson:"payment_method"`
	ActualAmount           float32            `bson:"actual_amount"`
	PaymentAmount          float32            `bson:"payment_amount"`
	PendingAmount          float32            `bson:"pending_amount"`
	IsOverdue              bool               `bson:"is_overdue"`
	OverdueDays            int                `bson:"overdue_days"`
	InterestOnOverdue      float32            `bson:"interest_on_overdue"`
	OverdueInterestRate    float64            `bson:"overdue_interest_rate"`
	OverdueInterestAmount  float32            `bson:"overdue_interest_amount"`
	OverdueAmount          float32            `bson:"overdue_amount"`
	FirstReminderMailSent  bool               `bson:"first_reminder_mail_sent"`
	SecondReminderMailSent bool               `bson:"second_reminder_mail_sent"`
	ThirdReminderMailSent  bool               `bson:"third_reminder_mail_sent"`
	CreatedAt              time.Time          `bson:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

type SubscriptionDetails struct {
	Id                   primitive.ObjectID `bson:"_id"`
	SubscriptionID       string             `bson:"subscription_id"`
	TenantId             string             `bson:"tenant_id"`
	PrimaryAdminId       string             `bson:"primary_admin_id"`
	PrimaryAdminName     string             `bson:"primary_admin_name"`
	PrimaryAdminEmail    string             `bson:"primary_admin_email"`
	SecondaryAdminId     string             `bson:"secondary_admin_id"`
	SecondaryAdminName   string             `bson:"secondary_admin_name"`
	SecondaryAdminEmail  string             `bson:"secondary_admin_email"`
	PlanId               string             `bson:"plan_id"`
	Name                 string             `bson:"name"`
	Amount               float32            `bson:"amount"`
	Currency             string             `bson:"currency"`
	Country              string             `bson:"country"`
	Discount             DiscountDetails    `bson:"discount"`
	DiscountedAmount     float32            `bson:"discounted_amount"`
	PaymentFrequency     string             `bson:"payment_frequency"`
	StartDate            time.Time          `bson:"start_date"`
	EndDate              time.Time          `bson:"end_date"`
	Status               string             `bson:"status"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
	NextBillingDate      time.Time          `bson:"next_billing_date"`
	LastBillingDate      time.Time          `bson:"last_billing_date"`
	PaymentHistory       []PaymentDetails   `bson:"payment_history"`
	TotalAmount          float32            `bson:"total_amount"`
	TotalPaidAmount      float32            `bson:"total_paid_amount"`
	TotalPendingAmount   float32            `bson:"total_pending_amount"`
	TotalRefundedAmount  float32            `bson:"total_refunded_amount"`
	HaveTIN              bool               `bson:"have_tin"`
	TaxIdentificationNum string             `bson:"tax_identification_num"`
}

type TaxDetails struct {
	Id         primitive.ObjectID `bson:"_id"`
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
	Percentage  float64            `bson:"percentage"` // applicable if DiscountType is percentage
	Amount      int                `bson:"amount"`     // aplicable if DiscountType is fixed
	Currency    string             `bson:"currency"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	StartDate   time.Time          `bson:"start_date"`
	EndDate     time.Time          `bson:"end_date"`
}
