package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type BankInfo struct {
	ID               primitive.ObjectID `bson:"_id"`
	EmployeeId       string             `bson:"employee_id"`
	FullName         string             `bson:"full_name"`
	BankName         string             `bson:"bank_name"`
	AccountNumber    string             `bson:"account_number"`
	IsSalaryAccount  bool               `bson:"is_salary_account"`
	IFSC             string             `bson:"ifsc"`
	ChequeURL        string             `bson:"cheque_url"`
	IsBankInfoLocked bool               `bson:"is_bank_info_locked"`
	Country          string             `bson:"country"`
	State            string             `bson:"state"`
	ZipCode          string             `bson:"zip_code"`
	SWIFT            string             `bson:"swift"`          // SWIFT/BIC code
	RoutingNumber    string             `bson:"routing_number"` // ABA routing number (US)
	SortCode         string             `bson:"sort_code"`      // UK/Ireland sort code
	BSB              string             `bson:"bsb"`            // Australia/New Zealand BSB code
	BankAddress      string             `bson:"bank_address"`   // Bank branch address
	Currency         string             `bson:"currency"`       // ISO 4217 currency code (e.g. USD, EUR)
	CreatedAt        string             `bson:"created_at"`
	UpdatedAt        string             `bson:"updated_at"`
	UpdatedBy        string             `bson:"updated_by"`
}
