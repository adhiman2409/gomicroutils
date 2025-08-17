package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CandidatePersonalInfo struct {
	FullName      string `bson:"full_name"`
	Gender        string `bson:"gender"`
	DOB           string `bson:"dob"`
	PhoneNumber   string `bson:"phone_number"`
	PersonalEmail string `bson:"personal_email"`
	AddressLine1  string `bson:"address_line_1"`
	AddressLine2  string `bson:"address_line_2"`
	City          string `bson:"city"`
	State         string `bson:"state"`
	Country       string `bson:"country"`
	Pincode       string `bson:"pincode"`
}

type CandidateProfessionalInfo struct {
	EmployeeType   string `bson:"employee_type"`
	Designation    string `bson:"designation"`
	Department     string `bson:"department"`
	EmailId        string `bson:"email_id"`
	OfficeLocation string `bson:"office_location"`
	WorkLocation   string `bson:"work_location"`
}

type OfferCoverLetterTemplate struct {
	TemplateName string            `bson:"template_name"`
	DataMap      map[string]string `bson:"data_map"`
}

type CandidateOfferLetter struct {
	DocumentNumber           string                   `bson:"document_number"`
	DocumentDate             string                   `bson:"document_date"`
	OfferCoverLetterTemplate OfferCoverLetterTemplate `bson:"cover_letter_template"`
	RemunerationInfo         OrgSalaryStructure       `bson:"remuneration_info"`
	JoiningDate              string                   `bson:"joining_date"`
	ValidFrom                string                   `bson:"valid_from"`
	ValidTo                  string                   `bson:"valid_to"`
	OfferStatus              string                   `bson:"offer_status"`
	OfferStatusRemarks       string                   `bson:"offer_status_remarks"`
	CreatedAt                string                   `bson:"created_at"`
	UpdatedAt                string                   `bson:"updated_at"`
	CreatedBy                string                   `bson:"created_by"`
}

type CandidateDetails struct {
	Id               primitive.ObjectID        `bson:"_id"`
	CandidateId      string                    `bson:"candidate_id"`
	JobId            string                    `bson:"job_id"`
	Status           CandidateStatus           `bson:"status"`
	PersonalInfo     CandidatePersonalInfo     `bson:"candidate_personal_info"`
	ProfessionalInfo CandidateProfessionalInfo `bson:"candidate_professional_info"`
	OfferLetter      CandidateOfferLetter      `bson:"candidate_offer_letter"`
	DocumentList     DocumentList              `bson:"document_list"`
	ApproverId       string                    `bson:"approver_id"`
	ApproverName     string                    `bson:"approver_name"`
	ApproverEmail    string                    `bson:"approver_email"`
	Remarks          []string                  `bson:"remarks"`
	Country          string                    `bson:"country"`
	State            string                    `bson:"state"`
	TimeZone         string                    `bson:"time_zone"`
	CreatedBy        string                    `bson:"created_by"`
	CreatedAt        time.Time                 `bson:"created_at"`
}
