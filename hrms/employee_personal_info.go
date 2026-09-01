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
	District           string  `bson:"district"`
	Village            string  `bson:"village"`
	Tehsil             string  `bson:"tehsil"`
	Country            string  `bson:"country"`
	Zipcode            string  `bson:"zipcode"`
	Lat                float64 `bson:"lat"`
	Lng                float64 `bson:"lng"`
	GoogleLocationCode string  `bson:"google_location_code"`
}

type FamilyInfo struct {
	FamilyMemberId           string `bson:"family_member_id"`
	EmployeeId               string `bson:"employee_id"`
	EmployeeName             string `bson:"employee_name"`
	FamilyMemberName         string `bson:"family_member_name"`
	FamilyMemberDOB          string `bson:"family_member_dob"`
	FamilyMemberGender       string `bson:"family_member_gender"`
	FamilyMemberAadharNumber string `bson:"family_member_aadhar_number"`
	FamilyMemberRelation     string `bson:"family_member_relation"`
	CoveredInCorporateIns    bool   `bson:"covered_in_corporate_ins"`
	InsurancePolicyId        string `bson:"insurance_policy_id"`
	InsuranceCardURL         string `bson:"insurance_card_url"`
	InsuranceCardUpdateYear  string `bson:"insurance_card_update_year"`
	FamilyMemberPhoneNumber  string `bson:"family_member_phone_number"`
	IsEmergencyContact       bool   `bson:"is_emergency_contact"`
	UpdateAt                 string `bson:"updated_at"`
	UpdateBy                 string `bson:"updated_by"`
}

type EmergencyContact struct {
	ID          string `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	Relation    string `bson:"relation"`
	PhoneNumber string `bson:"phone_number"`
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
	EmergencyNumberTwo   string             `bson:"emergency_number_two"`
	PersonalEmail        string             `bson:"email_id"`
	PAN                  string             `bson:"pan"`
	UAN                  string             `bson:"uan"`
	AadharNumber         string             `bson:"aadhar_number"`
	PassportNumber       string             `bson:"passport_number"`
	EmiratesIdNumber     string             `bson:"emirates_id_number,omitempty"`
	UnifiedIdNumber      string             `bson:"unified_id_number,omitempty"`
	VisaFileNumber       string             `bson:"visa_file_number,omitempty"`
	VisaExpiryDate       string             `bson:"visa_expiry_date,omitempty"`
	PermanentAddress     Address            `bson:"permanent_address"`
	CommunicationAddress Address            `bson:"communication_address"`
	FamilyInfo           []FamilyInfo       `bson:"family_info"`
	EmergencyContacts    []EmergencyContact `bson:"emergency_contacts,omitempty"`
	BothAddressSame      bool               `bson:"both_address_same"`
	RegexText            string             `bson:"regex_text"`
	IsPersonalInfoLocked bool               `bson:"is_personal_info_locked"`
	IsFamilyInfoLocked   bool               `bson:"is_family_info_locked"`
	IsDocsLocked         bool               `bson:"is_docs_locked"`
	IsExperienceLocked   bool               `bson:"is_experience_locked"`
	IsPassportLocked     bool               `bson:"is_passport_locked"`
	CreatedBy            string             `bson:"created_by"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}
