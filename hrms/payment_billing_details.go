package hrms

type BankAccountDetails struct {
	AccountHolderName string `bson:"account_holder_name"`
	AccountNumber     string `bson:"account_number"`
	IFSCCode          string `bson:"ifsc_code"`
	IBANCode          string `bson:"iban_code"`
	BankName          string `bson:"bank_name"`
	BranchName        string `bson:"branch_name"`
	SwiftCode         string `bson:"swift_code"`
	State             string `bson:"state"`
}

type BankAccounts struct {
	Country     string               `bson:"country"`
	BankDetails []BankAccountDetails `bson:"bank_details"`
}

type BillingAddresses struct {
	Country   string           `bson:"country"`
	Addresses []BillingAddress `bson:"addresses"`
}

type BillingAddress struct {
	AddressLine1         string  `bson:"address_line_1"`
	AddressLine2         string  `bson:"address_line_2"`
	State                string  `bson:"state"`
	City                 string  `bson:"city"`
	District             string  `bson:"district"`
	Village              string  `bson:"village"`
	Tehsil               string  `bson:"tehsil"`
	Country              string  `bson:"country"`
	Zipcode              string  `bson:"zipcode"`
	Lat                  float64 `bson:"lat"`
	Lng                  float64 `bson:"lng"`
	TaxIdentificationNum TaxId   `bson:"tax_identification_num"`
	GoogleLocationCode   string  `bson:"google_location_code"`
}

type TaxId struct {
	TinType   string `bson:"tin_type"`
	TinNumber string `bson:"tin_number"`
}
