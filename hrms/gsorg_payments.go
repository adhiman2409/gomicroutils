package hrms

type OrderRequest struct {
	DonorId            string  `bson:"donor_id"`
	DonorName          string  `bson:"donor_name"`
	DonorEmail         string  `bson:"donor_email"`
	DonorPhoneNumber   string  `bson:"donor_mobile"`
	KeepIdentityHidden bool    `bson:"keep_identity_hidden"`
	DonationId         string  `bson:"donation_id"`
	DonationAmount     float64 `bson:"donation_amount"`
}
