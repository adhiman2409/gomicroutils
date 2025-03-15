package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentInfo struct {
	DocName              string    `bson:"doc_name"`
	DocType              string    `bson:"doc_type"`
	DocURL               string    `bson:"doc_url"`
	IsMandotory          bool      `bson:"is_mandatory"`
	DocStatus            string    `bson:"doc_status"`
	DocRemarks           string    `bson:"doc_remarks"`
	IsVerified           bool      `bson:"is_verified"`
	VerifiedBy           string    `bson:"verified_by"`
	VerifiedOn           time.Time `bson:"verified_at"`
	VerificationMode     string    `bson:"verification_mode"`
	VerificationStatus   string    `bson:"verification_status"`
	IsVerificationPassed bool      `bson:"is_verification_passed"`
	VerificationRemarks  string    `bson:"verification_remarks"`
	CreatedAt            time.Time `bson:"created_at"`
	UpdatedAt            time.Time `bson:"updated_at"`
}

type StudentSchoolInfo struct {
	Name                 string    `bson:"name"`
	Type                 string    `bson:"type"`
	ContactName          string    `bson:"contact_name"`
	ContactDesignation   string    `bson:"contact_designation"`
	ContactPhone         string    `bson:"contact_phone"`
	ContactEmail         string    `bson:"contact_email"`
	Address              Address   `bson:"address"`
	PrincipalName        string    `bson:"principal_name"`
	PrincipalPhone       string    `bson:"principal_phone"`
	PrincipalEmail       string    `bson:"principal_email"`
	PrincipalRemarks     string    `bson:"principal_remarks"`
	Remarks              string    `bson:"remarks"`
	Website              string    `bson:"website"`
	IsVerified           bool      `bson:"is_verified"`
	VerifiedBy           string    `bson:"verified_by"`
	VerifiedOn           time.Time `bson:"verified_at"`
	VerificationMode     string    `bson:"verification_mode"`
	VerificationStatus   string    `bson:"verification_status"`
	IsVerificationPassed bool      `bson:"is_verification_passed"`
	VerificationRemarks  string    `bson:"verification_remarks"`
	Regex                string    `bson:"regex"`
}

type StudentFeeInfo struct {
	MonthlyTotalFee       float64   `bson:"monthly_total_fee"`
	FeePaymentCycle       string    `bson:"fee_payment_cycle"`
	FeePaymentLastDate    string    `bson:"fee_payment_last_date"`
	BankAccountHolderName string    `bson:"bank_account_holder_name"`
	BankName              string    `bson:"bank_name"`
	BankAccountNumber     string    `bson:"bank_account_number"`
	BankIFSC              string    `bson:"ifsc"`
	CancelledChequeURL    string    `bson:"cancelled_cheque_url"`
	IsVerified            bool      `bson:"is_verified"`
	VerifiedBy            string    `bson:"verified_by"`
	VerifiedOn            time.Time `bson:"verified_at"`
	VerificationMode      string    `bson:"verification_mode"`
	VerificationStatus    string    `bson:"verification_status"`
	IsVerificationPassed  bool      `bson:"is_verification_passed"`
	VerificationRemarks   string    `bson:"verification_remarks"`
}

type StudentAcademicInfo struct {
	AdmissionNumber            string         `bson:"admission_number"`
	Class                      string         `bson:"academic_class"`
	Section                    string         `bson:"academic_section"`
	RollNumber                 string         `bson:"academic_roll_number"`
	Stream                     string         `bson:"academic_stream"`
	Medium                     string         `bson:"academic_medium"`
	Board                      string         `bson:"academic_board"`
	AidRequiredForAcademicYear string         `bson:"aid_required_for_acadmic_year"`
	SchoolName                 string         `bson:"academic_school_name"`
	SchoolType                 string         `bson:"academic_school_type"`
	Documents                  []DocumentInfo `bson:"academic_documents"`
	DocumentCompletionStatus   string         `bson:"academic_document_completion_status"`
	DocumentCompletionDate     time.Time      `bson:"academic_document_completion_date"`
	IsVerified                 bool           `bson:"is_verified"`
	VerifiedBy                 string         `bson:"verified_by"`
	VerifiedOn                 time.Time      `bson:"verified_at"`
	VerificationMode           string         `bson:"verification_mode"`
	VerificationStatus         string         `bson:"verification_status"`
	IsVerificationPassed       bool           `bson:"is_verification_passed"`
	VerificationRemarks        string         `bson:"verification_remarks"`
	CreatedAt                  time.Time      `bson:"created_at"`
	UpdatedAt                  time.Time      `bson:"updated_at"`
}

type StudentPersonalInfo struct {
	StudentName                      string         `bson:"student_name"`
	Gender                           string         `bson:"gender"`
	DOB                              string         `bson:"dob"`
	IsOrphan                         bool           `bson:"is_orphan"`
	IsSemiOrphan                     bool           `bson:"is_semi_orphan"`
	IsParentDisabled                 bool           `bson:"is_parent_disabled"`
	ParentDisability                 string         `bson:"parent_disability"`
	GuardianName                     string         `bson:"guardian_name"`
	GuardianPhoneNumber              string         `bson:"guardian_phone_number"`
	Address                          Address        `bson:"address"`
	GuardianRelation                 string         `bson:"guardian_relation"`
	PersonalDocuments                []DocumentInfo `bson:"personal_documents"`
	PersonalDocumentCompletionStatus string         `bson:"personal_document_completion_status"`
	PersonalDocumentCompletionDate   time.Time      `bson:"personal_document_completion_date"`
	IsSocialMediaConsentGiven        bool           `bson:"is_social_media_consent_given"`
	IsVerified                       bool           `bson:"is_verified"`
	VerifiedBy                       string         `bson:"verified_by"`
	VerifiedOn                       time.Time      `bson:"verified_at"`
	VerificationMode                 string         `bson:"verification_mode"`
	VerificationStatus               string         `bson:"verification_status"`
	IsVerificationPassed             bool           `bson:"is_verification_passed"`
	VerificationRemarks              string         `bson:"verification_remarks"`
}

type ScholarshipReferrerInfo struct {
	ReferrerId              string `bson:"referrer_id"`
	Name                    string `bson:"name"`
	PhoneNumber             string `bson:"phone_number"`
	Email                   string `bson:"email"`
	HowReferrerKnowsStudent string `bson:"how_referrer_knows_student"`
	KeepIdentityHidden      bool   `bson:"keep_identity_hidden"`
	Remarks                 string `bson:"remarks"`
}

type RequestVerification struct {
	VerificationMode      string    `bson:"verification_mode"`
	VeerifierId           string    `bson:"veerifier_id"`
	VerifierName          string    `bson:"verifier_name"`
	VerifierPhoneNumber   string    `bson:"verifier_phone_number"`
	VerificationStartDate time.Time `bson:"verification_start_date"`
	VerificationEndDate   time.Time `bson:"verification_end_date"`
	VerificationStatus    string    `bson:"verification_status"`
	IsVerificationPassed  bool      `bson:"is_verification_passed"`
	VerificationRemarks   string    `bson:"verification_remarks"`
}

type StudentScholarshipRequest struct {
	ID                  primitive.ObjectID      `bson:"_id"`
	StudentId           string                  `bson:"student_id"`
	RequestId           string                  `bson:"request_id"`
	RequestStatus       string                  `bson:"request_status"`
	StudentPersonalInfo StudentPersonalInfo     `bson:"student_personal_info"`
	ReferrerInfo        ScholarshipReferrerInfo `bson:"referrer_info"`
	StudentAcademicInfo StudentAcademicInfo     `bson:"student_academic_info"`
	StudentFeeInfo      StudentFeeInfo          `bson:"student_fee_info"`
	StudentSchoolInfo   StudentSchoolInfo       `bson:"student_school_info"`
	VerificationInfo    RequestVerification     `bson:"verification_info"`
	CaseOfficerId       string                  `bson:"case_officer_id"`
	CaseOfficerName     string                  `bson:"case_officer_name"`
	CaseOfficerRemarks  string                  `bson:"case_officer_remarks"`
	IsRequestActive     bool                    `bson:"is_request_active"`
	RequestClosureDate  time.Time               `bson:"request_closure_date"`
	Domain              string                  `bson:"domain"`
	RegexText           string                  `bson:"regex_text"`
	CreatedAt           time.Time               `bson:"created_at"`
	UpdatedAt           time.Time               `bson:"updated_at"`
}
