package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DonationCollection struct {
	Month             string  `bson:"month"`
	Year              string  `bson:"year"`
	TotalAmount       float64 `bson:"total_amount"`
	TotalContributors int     `bson:"total_contributors"`
}

type Subscription struct {
	ID                   primitive.ObjectID       `bson:"_id"`
	MasterDonationId     string                   `bson:"master_donation_id"`
	SubscriptionId       string                   `bson:"subscription_id"`
	SubscriptionType     string                   `bson:"subscription_type"`
	StartDate            time.Time                `bson:"start_date"`
	EndDate              time.Time                `bson:"end_date"`
	MonthlyAmount        float64                  `bson:"monthly_amount"`
	TotalAmount          float64                  `bson:"total_amount"`
	TotalMonths          int                      `bson:"total_months"`
	TotalAmountPaid      float64                  `bson:"total_amount_paid"`
	TotalMonthsPaid      int                      `bson:"total_months_paid"`
	TotalAmountLeft      float64                  `bson:"total_amount_left"`
	TotalMonthsLeft      int                      `bson:"total_months_left"`
	TotalAmountUsed      float64                  `bson:"total_amount_used"`
	TotalMonthsUsed      int                      `bson:"total_months_used"`
	NumberOfStudents     int                      `bson:"number_of_students"`
	PrefferedPaymentDate string                   `bson:"preffered_payment_date"`
	IsAutoRenewal        bool                     `bson:"is_auto_renewal"`
	IsActive             bool                     `bson:"is_active"`
	PaymentMode          string                   `bson:"payment_mode"`
	PaymentHistory       []DonationPaymentHistory `bson:"payment_history"`
	SubscriptionStatus   string                   `bson:"subscription_status"`
	Remarks              string                   `bson:"remarks"`
	CreatedAt            time.Time                `bson:"created_at"`
	UpdatedAt            time.Time                `bson:"updated_at"`
}

type DonationPaymentHistory struct {
	DonationId          string    `bson:"donation_id"`
	DonationAmount      float64   `bson:"donation_amount"`
	DonationDate        time.Time `bson:"donation_date"`
	DonationRemarks     string    `bson:"donation_remarks"`
	DonationStatus      string    `bson:"donation_status"`
	DonationType        string    `bson:"donation_type"`
	DonationMode        string    `bson:"donation_mode"`
	DonationReceiptId   string    `bson:"donation_receipt_id"`
	DonationReceiptURL  string    `bson:"donation_receipt_url"`
	DonationReceiptDate time.Time `bson:"donation_receipt_date"`
}

type Donor struct {
	ID                     primitive.ObjectID       `bson:"_id"`
	DonorId                string                   `bson:"donor_id"`
	DonorName              string                   `bson:"donor_name"`
	DonorEmail             string                   `bson:"donor_email"`
	DonorPhone             string                   `bson:"donor_phone"`
	TotalDonation          float64                  `bson:"total_donation"`
	TotalDonationCount     int                      `bson:"total_donation_count"`
	KeepIdentityHidden     bool                     `bson:"keep_identity_hidden"`
	DonationPaymentHistory []DonationPaymentHistory `bson:"donation_payment_history"`
	Subscriptions          []Subscription           `bson:"subscriptions"`
	CreatedAt              time.Time                `bson:"created_at"`
	UpdatedAt              time.Time                `bson:"updated_at"`
}

type DonationPayment struct {
	ID                  primitive.ObjectID `bson:"_id"`
	DonationId          string             `bson:"donation_id"`
	DonorId             string             `bson:"donor_id"`
	DonorName           string             `bson:"donor_name"`
	DonationAmount      float64            `bson:"donation_amount"`
	Day                 string             `bson:"day"`
	Month               string             `bson:"month"`
	Year                string             `bson:"year"`
	DonationDate        time.Time          `bson:"donation_date"`
	DonationRemarks     string             `bson:"donation_remarks"`
	DonationStatus      string             `bson:"donation_status"`
	DonationType        string             `bson:"donation_type"`
	DonationMode        string             `bson:"donation_mode"`
	DonationReceiptId   string             `bson:"donation_receipt_id"`
	DonationReceiptURL  string             `bson:"donation_receipt_url"`
	DonationReceiptDate time.Time          `bson:"donation_receipt_date"`
	CreatedAt           time.Time          `bson:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at"`
}

type MasterDonation struct {
	ID                    primitive.ObjectID   `bson:"_id"`
	DonationId            string               `bson:"donation_id"`
	Title                 string               `bson:"title"`
	Description           string               `bson:"description"`
	DonationMinAmount     float64              `bson:"donation_min_amount"`
	DonationMaxAmount     float64              `bson:"donation_max_amount"`
	IsFixedAmountDonation bool                 `bson:"is_fixed_amount_donation"`
	FixedDonationAmount   float64              `bson:"fixed_donation_amount"`
	IsTargetBasedDonation bool                 `bson:"is_target_based_donation"`
	TargetDonationAmount  float64              `bson:"target_donation_amount"`
	MonthlyTargetAmount   float64              `bson:"monthly_target_amount"`
	TotalDonationAmount   float64              `bson:"total_donation_amount"`
	TotalDonors           int                  `bson:"total_donors"`
	TotalAmountPaid       float64              `bson:"total_scholarship_amount_paid"`
	TotalAmountLeft       float64              `bson:"total_scholarship_amount_left"`
	TotalAmountUsed       float64              `bson:"total_scholarship_amount_used"`
	DonationCategory      string               `bson:"donation_category"`
	DonationSubCategory   string               `bson:"donation_sub_category"`
	DonationStatus        string               `bson:"donation_status"`
	AssociatedDoneeId     string               `bson:"associated_donee_id"`
	AssociatedDoneeName   string               `bson:"associated_donee_name"`
	AutoStopDonation      bool                 `bson:"auto_stop_donation"`
	Donations             []DonationPayment    `bson:"donations"`
	MontlyCollection      []DonationCollection `bson:"monthly_collection"`
	IsPublished           bool                 `bson:"is_published"`
	PublishDate           time.Time            `bson:"publish_date"`
	UnpublishDate         time.Time            `bson:"unpublish_date"`
	HeroImageURL          string               `bson:"hero_image_url"`
	ImageURLs             []string             `bson:"image_urls"`
	Rejex                 string               `bson:"rejex"`
	Remarks               string               `bson:"remarks"`
	CreatedAt             time.Time            `bson:"created_at"`
	UpdatedAt             time.Time            `bson:"updated_at"`
}
