package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type UnirmsOrderRequest struct {
	PayeeId          string  `bson:"payee_id"`
	PayeeName        string  `bson:"payee_name"`
	PayeeEmail       string  `bson:"payee_email"`
	PayeePhoneNumber string  `bson:"payee_phone_number"`
	InvoiceId        string  `bson:"invoice_id"`
	Amount           float64 `bson:"amount"`
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
		TenantID       string `bson:"tenant_id"`
		SubscriptionId string `bson:"subscription_id"`
		InvoiceID      string `bson:"invoice_id"`
	} `bson:"notes"`
	Fee          int    `bson:"fee"`
	Tax          int    `bson:"tax"`
	ErrorCode    string `bson:"error_code"`
	ErrorDesc    string `bson:"error_description"`
	ErrorSrc     string `bson:"error_source"`
	ErrorStep    string `bson:"error_step"`
	ErrorReason  string `bson:"error_reason"`
	AcquirerData struct {
		Code    string `bson:"code"`
		Message string `bson:"message"`
	} `bson:"acquirer_data"`
	CreatedAt int64 `bson:"created_at"`
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
		TenantID       string `bson:"tenant_id"`
		SubscriptionId string `bson:"subscription_id"`
		InvoiceID      string `bson:"invoice_id"`
	} `bson:"notes"`
	CreatedAt int64 `bson:"created_at"`
}
