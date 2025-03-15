package hrms

import (
	"time"
)

type DocumentInfo struct {
	DocName             string    `json:"doc_name"`
	DocType             string    `json:"doc_type"`
	DocURL              string    `json:"doc_url"`
	IsMandotory         bool      `json:"is_mandatory"`
	DocStatus           string    `json:"doc_status"`
	DocRemarks          string    `json:"doc_remarks"`
	IsVerified          bool      `json:"is_verified"`
	VerifiedBy          string    `json:"verified_by"`
	VerifiedOn          time.Time `json:"verified_at"`
	VerificationMode    string    `json:"verification_mode"`
	VerificationRemarks string    `json:"verification_remarks"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type StudentSchoolInfo struct {
	SchoolId            string    `json:"school_id"`
	StudentId           string    `json:"student_id"`
	Name                string    `json:"name"`
	Type                string    `json:"type"`
	ContactName         string    `json:"contact_name"`
	ContactDesignation  string    `json:"contact_designation"`
	ContactPhone        string    `json:"contact_phone"`
	ContactEmail        string    `json:"contact_email"`
	Address             Address   `json:"address"`
	PrincipalName       string    `json:"principal_name"`
	PrincipalPhone      string    `json:"principal_phone"`
	PrincipalEmail      string    `json:"principal_email"`
	PrincipalRemarks    string    `json:"principal_remarks"`
	Remarks             string    `json:"remarks"`
	Website             string    `json:"website"`
	IsVerified          bool      `json:"is_verified"`
	VerifiedBy          string    `json:"verified_by"`
	VerifiedOn          time.Time `json:"verified_at"`
	VerificationMode    string    `json:"verification_mode"`
	VerificationRemarks string    `json:"verification_remarks"`
}

type StudentFeeInfo struct {
	SchoolId                    string    `json:"school_id"`
	MonthlyTotalFee             float64   `json:"monthly_total_fee"`
	FeePaymentCycle             string    `json:"fee_payment_cycle"`
	FeePaymentLastDate          string    `json:"fee_payment_last_date"`
	SchoolBankAccountHolderName string    `json:"bank_account_holder_name"`
	BankName                    string    `json:"bank_name"`
	BankAccountNumber           string    `json:"bank_account_number"`
	BankIFSC                    string    `json:"ifsc"`
	CancelledChequeURL          string    `json:"cancelled_cheque_url"`
	IsVerified                  bool      `json:"is_verified"`
	VerifiedBy                  string    `json:"verified_by"`
	VerifiedOn                  time.Time `json:"verified_at"`
	VerificationMode            string    `json:"verification_mode"`
	VerificationRemarks         string    `json:"verification_remarks"`
}

type StudentAcademicInfo struct {
	StudentId                  string         `json:"student_id"`
	AdmissionNumber            string         `json:"admission_number"`
	Class                      string         `json:"academic_class"`
	Section                    string         `json:"academic_section"`
	RollNumber                 string         `json:"academic_roll_number"`
	Stream                     string         `json:"academic_stream"`
	Medium                     string         `json:"academic_medium"`
	Board                      string         `json:"academic_board"`
	AidRequiredForAcademicYear string         `json:"aid_required_for_acadmic_year"`
	SchoolName                 string         `json:"academic_school_name"`
	SchoolType                 string         `json:"academic_school_type"`
	Documents                  []DocumentInfo `json:"academic_documents"`
	DocumentCompletionStatus   string         `json:"academic_document_completion_status"`
	DocumentCompletionDate     time.Time      `json:"academic_document_completion_date"`
	IsVerified                 bool           `json:"is_verified"`
	VerifiedBy                 string         `json:"verified_by"`
	VerifiedOn                 time.Time      `json:"verified_at"`
	VerificationMode           string         `json:"verification_mode"`
	VerificationRemarks        string         `json:"verification_remarks"`
	CreatedAt                  time.Time      `json:"created_at"`
	UpdatedAt                  time.Time      `json:"updated_at"`
}

type StudentPersonalInfo struct {
	StudentId                        string         `json:"student_id"`
	StudentName                      string         `json:"student_name"`
	Gender                           string         `json:"gender"`
	DOB                              string         `json:"dob"`
	IsOrphan                         bool           `json:"is_orphan"`
	IsSemiOrphan                     bool           `json:"is_semi_orphan"`
	IsParentDisabled                 bool           `json:"is_parent_disabled"`
	ParentDisability                 string         `json:"parent_disability"`
	GuardianName                     string         `json:"guardian_name"`
	GuardianPhoneNumber              string         `json:"guardian_phone_number"`
	Address                          Address        `json:"address"`
	GuardianRelation                 string         `json:"guardian_relation"`
	PersonalDocuments                []DocumentInfo `json:"personal_documents"`
	PersonalDocumentCompletionStatus string         `json:"personal_document_completion_status"`
	PersonalDocumentCompletionDate   time.Time      `json:"personal_document_completion_date"`
	IsSocialMediaConsentGiven        bool           `json:"is_social_media_consent_given"`
	IsVerified                       bool           `json:"is_verified"`
	VerifiedBy                       string         `json:"verified_by"`
	VerifiedOn                       time.Time      `json:"verified_at"`
	VerificationMode                 string         `json:"verification_mode"`
	VerificationRemarks              string         `json:"verification_remarks"`
}

type ScholarshipReferrerInfo struct {
	StudentId               string `json:"student_id"`
	ReferrerId              string `json:"referrer_id"`
	Name                    string `json:"name"`
	PhoneNumber             string `json:"phone_number"`
	Email                   string `json:"email"`
	HowReferrerKnowsStudent string `json:"how_referrer_knows_student"`
	Remarks                 string `json:"remarks"`
}

type RequestVerification struct {
	StudentId                  string    `json:"student_id"`
	IsVerificationInitiated    bool      `json:"is_verification_initiated"`
	VerificationInitiationDate string    `json:"verification_initiation_date"`
	VerificationMode           string    `json:"verification_mode"`
	VeerifierId                string    `json:"veerifier_id"`
	VerifierName               string    `json:"verifier_name"`
	IsVerificationCompleted    bool      `json:"is_verification_completed"`
	VerificationCompletionDate time.Time `json:"verification_completion_date"`
	VerificationRemarks        string    `json:"verification_remarks"`
}

type StudentScolershipRequest struct {
	ID                  string                  `json:"id"`
	StudentId           string                  `json:"student_id"`
	RequestId           string                  `json:"request_id"`
	RequestStatus       string                  `json:"request_status"`
	PersonalInfo        StudentPersonalInfo     `json:"personal_info"`
	ReferrerInfo        ScholarshipReferrerInfo `json:"referrer_info"`
	StudentAcademicInfo StudentAcademicInfo     `json:"student_academic_info"`
	StudentFeeInfo      StudentFeeInfo          `json:"student_fee_info"`
	StudentSchoolInfo   StudentSchoolInfo       `json:"student_school_info"`
	VerificationInfo    RequestVerification     `json:"verification_info"`
	CaseOfficerId       string                  `json:"case_officer_id"`
	CaseOfficerName     string                  `json:"case_officer_name"`
	CaseOfficerRemarks  string                  `json:"case_officer_remarks"`
	IsRequestActive     bool                    `json:"is_request_active"`
	RequestClosureDate  time.Time               `json:"request_closure_date"`
	Domain              string                  `json:"domain"`
	RegexText           string                  `json:"regex_text"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
}
