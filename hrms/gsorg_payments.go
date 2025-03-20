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

type OrderResponse struct {
	ID         primitive.ObjectID `bson:"_id"`
	OrderId    string             `bson:"order_id"`
	DonationId string             `bson:"donation_id"`
	DonorId    string             `bson:"donor_id"`
	Entity     string             `bson:"entity"`
	Amount     int64              `bson:"amonth"`
	AmountPaid int64              `bson:"amount_paid"`
	AmountDue  int64              `bson:"amount_due"`
	Currency   string             `bson:"currency"`
	Receipt    string             `bson:"receipt"`
	OfferId    string             `bson:"offer_id"`
	Status     string             `bson:"status"`
	Attempts   int                `bson:"attempts"`
	CreatedAt  int64              `bson:"created_at"`
}
