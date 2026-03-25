package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialURL struct {
	Platform string `bson:"platform"`
	URL      string `bson:"url"`
}

type LeadInfo struct {
	Id                      primitive.ObjectID `bson:"_id,omitempty"`
	LeadID                  string             `bson:"lead_id"`
	Title                   string             `bson:"title"`
	Description             string             `bson:"description"`
	Stage                   ClientMainStage    `bson:"stage"`
	DurationToCloseInMonths float32            `bson:"duration_to_close_in_months"`
	SubStage                string             `bson:"sub_stage,omitempty"`
	AssignedSPOCs           []string           `bson:"assigned_spocs,omitempty"`
	LeadRemarks             []string           `bson:"lead_remarks,omitempty"`
	Meetings                []MeetingInfo      `bson:"meetings,omitempty"`
	Remarks                 []LeadRemark       `bson:"remarks"`
	Notes                   []LeadRemark       `bson:"notes"`
	DealValue               float64            `bson:"deal_value"`
	Priority                string             `bson:"priority,omitempty"`
	Currency                string             `bson:"currency"`
	DocumentLinks           []Attachment       `bson:"document_links,omitempty"`
	AssociatedContacts      []string           `bson:"associated_contacts,omitempty"`
	ProbabilityToClose      int                `bson:"probability_to_close,omitempty"`
	LeadSource              string             `bson:"lead_source,omitempty"`
	ReferenceId             string             `bson:"reference_id,omitempty"`
	TenderId                string             `bson:"tender_id,omitempty"`
	TenderStartDate         string             `bson:"tender_start_date,omitempty"`
	TenderEndDate           string             `bson:"tender_end_date,omitempty"`
	Watchers                []string           `bson:"watchers,omitempty"`
	LeadStage               string             `bson:"lead_stage"`
	LeadSubStage            string             `bson:"lead_sub_stage"`
	OfferingType            string             `bson:"offering_type"`
	OfferingSubType         string             `bson:"offering_sub_type"`
	ActionItems             []ActionItem       `bson:"action_items"`
	CreatedBy               string             `bson:"created_by"`
	CreatedAt               time.Time          `bson:"created_at"`
	UpdatedAt               time.Time          `bson:"updated_at"`
}

type LeadRemark struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Note          string             `bson:"note"`
	CreatedAt     time.Time          `bson:"created_at"`
	CreatedByEid  string             `bson:"created_by_eid"`
	CreatedByName string             `bson:"created_by_name"`
}

type CompanyInfo struct {
	Id                   primitive.ObjectID `bson:"_id,omitempty"`
	CompanyId            string             `bson:"company_id"`
	CompanyName          string             `bson:"company_name"`
	CompanyLogo          string             `bson:"company_logo"`
	CurrentStage         ClientMainStage    `bson:"current_stage"`
	Website              string             `bson:"website"`
	Industry             string             `bson:"industry"`
	HeadquarterCountry   string             `bson:"headquarter_country"`
	BusinessActivities   string             `bson:"business_activities"`
	ClientDescription    string             `bson:"client_description"`
	YearFounded          int                `bson:"year_founded,omitempty"`
	Contacts             []string           `bson:"contacts,omitempty"`
	CompanySize          int                `bson:"company_size,omitempty"`
	AnnualRevenue        float64            `bson:"annual_revenue,omitempty"`
	RevenueCurrency      string             `bson:"revenue_currency,omitempty"`
	ProductsAndServices  []string           `bson:"products_and_services,omitempty"`
	OtherOfficeLocations []string           `bson:"other_office_locations,omitempty"`
	SourceOfIntroduction string             `bson:"source_of_introduction"`
	IntroducedBy         string             `bson:"introduced_by"`
	Tags                 []string           `bson:"tags,omitempty"`
	Priority             string             `bson:"priority,omitempty"`
	IsPublic             bool               `bson:"is_public,omitempty"`
	Attachments          []Attachment       `bson:"attachments,omitempty"`
	Remarks              []string           `bson:"remarks,omitempty"`
	Regex                string             `bson:"regex"`
	UpcomingReminderTime time.Time          `bson:"upcoming_reminder_time,omitempty"`
	Leads                []LeadInfo         `bson:"leads,omitempty"`
	EditHistory          []LeadRemark       `bson:"edit_history,omitempty"`
	IsActive             bool               `bson:"is_active"`
	CompanyStage         string             `bson:"company_stage"`
	CreatedBy            string             `bson:"created_by"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}

type ContactInfo struct {
	ContactId       string      `bson:"contact_id,omitempty"`
	EmpId           string      `bson:"emp_id,omitempty"`
	Name            string      `bson:"name"`
	Email           string      `bson:"email"`
	Mobile          string      `bson:"mobile"`
	SecondaryEmail  string      `bson:"secondary_email"`
	SecondaryMobile string      `bson:"secondary_mobile"`
	Designation     string      `bson:"designation"`
	Department      string      `bson:"department,omitempty"`
	DOB             string      `bson:"dob,omitempty"`
	Anniversary     string      `bson:"anniversary,omitempty"`
	Tags            []string    `bson:"tags,omitempty"`
	IsActive        bool        `bson:"is_active,omitempty"`
	Notes           string      `bson:"notes,omitempty"`
	CompanyName     string      `bson:"company_name,omitempty"`
	CompanyId       string      `bson:"company_id,omitempty"`
	RegexString     string      `bson:"regex_string,omitempty"`
	CompanyInfo     CompanyInfo `bson:"company_info,omitempty"`
	SocialURLs      []SocialURL `bson:"social_urls,omitempty"`
}

type Attendee struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Eid         string             `bson:"employee_id,omitempty"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email,omitempty"`
	Designation string             `bson:"designation,omitempty"`
	Mobile      string             `bson:"mobile,omitempty"`
}

type ActionItem struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	AssignedTo    []string           `bson:"assigned_to,omitempty"`
	DueDate       time.Time          `bson:"due_date"`
	Status        string             `bson:"status"`
	CreatedAt     time.Time          `bson:"created_at"`
	IsReminderSet bool               `bson:"is_reminder_set"`
	ReminderTime  time.Time          `bson:"reminder_time,omitempty"`
	MailSent      bool               `bson:"mail_sent,omitempty"`
	CreatedBy     string             `bson:"created_by"`
	Remarks       []LeadRemark       `bson:"remarks"`
	Attachments   []Attachment       `bson:"attachments,omitempty"`
}

type Attachment struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	FileName   string             `bson:"file_name"`
	FileURL    string             `bson:"file_url"`
	CreatedBy  string             `bson:"created_by"`
	UploadedAt time.Time          `bson:"uploaded_at"`
}

type Note struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	NoteContent string             `bson:"note_content"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type MeetingInfo struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	MeetingId         string             `bson:"meeting_id"`
	LeadId            string             `bson:"lead_id,omitempty"`
	Title             string             `bson:"title"`
	MeetingDate       time.Time          `bson:"meeting_date"`
	Mode              string             `bson:"mode"`
	InternalAttendees []string           `bson:"internal_attendees,omitempty"`
	ExternalAttendees []string           `bson:"external_attendees,omitempty"`
	MeetingNotes      []Note             `bson:"meeting_notes,omitempty"`
	ActionItems       []ActionItem       `bson:"action_items,omitempty"`
	Attachments       []Attachment       `bson:"attachments,omitempty"`
	Outcome           string             `bson:"outcome"`
	ReminderDate      time.Time          `bson:"reminder_date,omitempty"`
	IsReminderSet     bool               `bson:"is_reminder_set"`
	CreatedBy         string             `bson:"created_by"`
	CreatedByName     string             `bson:"created_by_name"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	MailSent          bool               `bson:"mail_sent,omitempty"`
}

type LeadOffering struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty"`
	Type      string                `bson:"type"` // Product / Service / Solution
	SubTypes  []LeadOfferingSubType `bson:"sub_types"`
	IsActive  bool                  `bson:"is_active"`
	CreatedAt time.Time             `bson:"created_at"`
	UpdatedAt time.Time             `bson:"updated_at"`
}

type LeadOfferingSubType struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description,omitempty"`
	IsActive    bool               `bson:"is_active"`
}

type LeadStage struct {
	Id           primitive.ObjectID `bson:"_id"`
	StageName    string             `bson:"stage_name"`
	Index        int                `bson:"index"`
	IsStartState bool               `bson:"is_start_state"`
	IsEndState   bool               `bson:"is_end_state"`
	SubStages    []string           `bson:"sub_stages"`
}

type ClientMainStage int

const (
	LLead ClientMainStage = iota + 1
	Suspect
	Prospect
	ProposalSent
	Negotiation
	ContractSigned
	CClient
	NoGo
	InProcess
	SSubmitted
	Technical
	WWon
	LLost
	CClosed
)

func (s ClientMainStage) String() string {
	stages := []string{
		"Lead",
		"Suspect",
		"Prospect",
		"Proposal Sent",
		"Negotiation",
		"Contract Signed",
		"Client",
		"NO-GO",
		"In-Process",
		"Submitted",
		"Technical",
		"Won",
		"Lost",
		"Closed",
	}

	i := int(s) - 1
	if i < 0 || i >= len(stages) {
		return "Unknown"
	}

	return stages[i]
}

func GetClientMainStages(domain string) []string {
	if domain == "qbitlabs.unirms.com" {
		return QbitLabsStages
	}
	return DefaultStages
}

func ClientMainStageFromString(s string) ClientMainStage {
	switch s {
	case "Lead":
		return LLead
	case "Suspect":
		return Suspect
	case "Prospect":
		return Prospect
	case "Proposal Sent":
		return ProposalSent
	case "Negotiation":
		return Negotiation
	case "Contract Signed":
		return ContractSigned
	case "Client":
		return CClient
	case "NO-GO":
		return NoGo
	case "In-Process":
		return InProcess
	case "Submitted":
		return SSubmitted
	case "Technical":
		return Technical
	case "Won":
		return WWon
	case "Lost":
		return LLost
	case "Closed":
		return CClosed
	default:
		return LLead // fallback
	}
}

var DefaultStages = []string{
	"Lead",
	"Suspect",
	"Prospect",
	"Proposal Sent",
	"Negotiation",
	"Contract Signed",
	"Client",
}

var QbitLabsStages = []string{
	"NO-GO",
	"In-Process",
	"Submitted",
	"Technical",
	"Won",
	"Lost",
	"Closed",
}
