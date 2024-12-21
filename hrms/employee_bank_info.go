package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type BankInfo struct {
	ID            primitive.ObjectID `bson:"_id"`
	EmployeeId    string             `bson:"employee_id"`
	FullName      string             `bson:"full_name"`
	BankName      string             `bson:"bank_name"`
	AccountNUmber string             `bson:"account_number"`
	IFSC          string             `bson:"ifsc"`
	ChequeURL     string             `bson:"cheque_url"`
	CreatedAt     string             `bson:"created_at"`
	UpdatedAt     string             `bson:"updated_at"`
}
