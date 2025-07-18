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
	CreatedAt        string             `bson:"created_at"`
	UpdatedAt        string             `bson:"updated_at"`
	UpdatedBy        string             `bson:"updated_by"`
}
