package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	AddressLine1       string  `bson:"address_line_1"`
	AddressLine2       string  `bson:"address_line_2"`
	State              string  `bson:"state"`
	City               string  `bson:"city"`
	Zipcode            string  `bson:"zipcode"`
	Lat                float64 `bson:"lat"`
	Lng                float64 `bson:"lng"`
	GoogleLocationCode string  `bson:"google_location_code"`
}

type EmpPersonalInfo struct {
	ID                   primitive.ObjectID `bson:"_id"`
	EmployeeId           string             `bson:"employee_id"`
	FirstName            string             `bson:"first_name"`
	MiddleName           string             `bson:"middle_name"`
	LastName             string             `bson:"last_name"`
	MotherName           string             `bson:"mother_name"`
	FatherName           string             `bson:"father_name"`
	Gender               string             `bson:"gender"`
	DOB                  string             `bson:"dob"`
	MaritalStatus        string             `bson:"marital_status"`
	SpouseName           string             `bson:"spouse_name"`
	Nationality          string             `bson:"nationality"`
	BloodGroup           string             `bson:"blood_group"`
	PhoneNumber          string             `bson:"phone_number"`
	EmergencyNumber      string             `bson:"emergency_number"`
	PersonalEmail        string             `bson:"email_id"`
	PAN                  string             `bson:"pan"`
	AadharNumber         string             `bson:"aadhar_number"`
	PassportNumber       string             `bson:"passport_number"`
	PermanentAddress     Address            `bson:"permanent_address"`
	CommunicationAddress Address            `bson:"communication_address"`
	BothAddressSame      bool               `bson:"both_address_same"`
	RegexText            string             `bson:"regex_text"`
	CreatedBy            string             `bson:"created_by"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}
