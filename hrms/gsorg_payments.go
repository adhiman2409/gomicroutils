package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderRequest struct {
	DonorId            string  `bson:"donor_id"`
	DonorName          string  `bson:"donor_name"`
	DonorEmail         string  `bson:"donor_email"`
	DonorPhoneNumber   string  `bson:"donor_mobile"`
	KeepIdentityHidden bool    `bson:"keep_identity_hidden"`
	DonationId         string  `bson:"donation_id"`
	DonationAmount     float64 `bson:"donation_amount"`
}

type PaymentEntity struct {
	ID             primitive.ObjectID `bson:"_id"`
	PaymentID      string             `bson:"payment_id"`
	Entity         string             `bson:"entity"`
	Amount         int                `bson:"amount"`
	Currency       string             `bson:"currency"`
	Status         string             `bson:"status"`
	OrderID        string             `bson:"order_id"`
	InvoiceID      string             `bson:"invoice_id"`
	International  bool               `bson:"international"`
	Method         string             `bson:"method"`
	AmountRefunded int                `bson:"amount_refunded"`
	RefundStatus   string             `bson:"refund_status"`
	Captured       bool               `bson:"captured"`
	Description    string             `bson:"description"`
	CardID         string             `bson:"card_id"`
	Bank           string             `bson:"bank"`
	Email          string             `bson:"email"`
	Wallet         string             `bson:"wallet"`
	VPA            string             `bson:"vpa"`
	Contact        string             `bson:"contact"`
	Notes          struct {
		DonorID    string `bson:"donor_id"`
		DonationID string `bson:"donation_id"`
	} `bson:"notes"`
	ErrorCode string `bson:"error_code"`
	ErrorDesc string `bson:"error_description"`
	Fee       int    `bson:"fee"`
	Tax       int    `bson:"tax"`
	CreatedAt int64  `bson:"created_at"`
}

type OrderEntity struct {
	ID         primitive.ObjectID `bson:"_id"`
	OrderId    string             `bson:"order_id"`
	Entity     string             `bson:"entity"`
	Amount     int                `bson:"amount"`
	AmountPaid int                `bson:"amount_paid"`
	AmountDue  int                `bson:"amount_due"`
	Currency   string             `bson:"currency"`
	Receipt    string             `bson:"receipt"`
	OfferID    string             `bson:"offer_id"`
	Status     string             `bson:"status"`
	Attempts   int                `bson:"attempts"`
	Notes      struct {
		DonorID    string `bson:"donor_id"`
		DonationID string `bson:"donation_id"`
	} `bson:"notes"`
	CreatedAt int64 `bson:"created_at"`
}
