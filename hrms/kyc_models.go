package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeadInfo struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	LeadID        string             `bson:"lead_id"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	Stage         string             `bson:"stage"`
	SubStage      string             `bson:"sub_stage,omitempty"`
	AssignedSPOCs []ContactInfo      `bson:"assigned_spocs,omitempty"`
	LeadRemarks   []string           `bson:"lead_remarks,omitempty"`
	Meetings      []MeetingInfo      `bson:"meetings,omitempty"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

type CompanyInfo struct {
	ID                 string `bson:"_id,omitempty"`
	CompanyId          string `bson:"company_id"`
	CompanyName        string `bson:"company_name"`
	CompanyLogo        string `bson:"company_logo"`
	Website            string `bson:"website"`
	Industry           string `bson:"industry"`
	HeadquarterCountry string `bson:"headquarter_country"`
	BusinessActivities string `bson:"business_activities"`
	ClientDescription  string `bson:"client_description"`

	YearFounded          int           `bson:"year_founded,omitempty"`
	Contacts             []ContactInfo `bson:"contacts,omitempty"`
	CompanySize          int           `bson:"company_size,omitempty"`
	AnnualRevenue        float64       `bson:"annual_revenue,omitempty"`
	RevenueCurrency      string        `bson:"revenue_currency,omitempty"`
	ProductsAndServices  []string      `bson:"products_and_services,omitempty"`
	OtherOfficeLocations []string      `bson:"other_office_locations,omitempty"`

	SourceOfIntroduction string       `bson:"source_of_introduction"`
	IntroducedBy         string       `bson:"introduced_by"`
	Tags                 []string     `bson:"tags,omitempty"`
	Priority             string       `bson:"priority,omitempty"`
	IsPublic             bool         `bson:"is_public,omitempty"`
	Attachments          []Attachment `bson:"attachments,omitempty"`
	Remarks              []string     `bson:"remarks,omitempty"`
	Regex                string       `bson:"regex"`
	UpcomingReminderTime time.Time    `bson:"upcoming_reminder_time,omitempty"`

	Leads []LeadInfo `bson:"leads,omitempty"`

	CreatedBy string    `bson:"created_by"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type ContactInfo struct {
	ContactId   string      `bson:"contact_id,omitempty"`
	EmpId       string      `bson:"emp_id,omitempty"`
	Name        string      `bson:"name"`
	Email       string      `bson:"email"`
	Mobile      string      `bson:"mobile"`
	Designation string      `bson:"designation"`
	Department  string      `bson:"department,omitempty"`
	DOB         string      `bson:"dob,omitempty"`
	Anniversary string      `bson:"anniversary,omitempty"`
	Tags        []string    `bson:"tags,omitempty"`
	IsActive    bool        `bson:"is_active,omitempty"`
	Notes       string      `bson:"notes,omitempty"`
	CompanyName string      `bson:"company_name,omitempty"`
	CompanyId   string      `bson:"company_id,omitempty"`
	RegexString string      `bson:"regex_string,omitempty"`
	CompanyInfo CompanyInfo `bson:"company_info,omitempty"`
}

type Attendee struct {
	ID          string `bson:"_id,omitempty"`
	Eid         string `bson:"employee_id,omitempty"`
	Name        string `bson:"name"`
	Email       string `bson:"email,omitempty"`
	Designation string `bson:"designation,omitempty"`
	Mobile      string `bson:"mobile,omitempty"`
}

type ActionItem struct {
	ID            string        `bson:"_id,omitempty"`
	Title         string        `bson:"title"`
	Description   string        `bson:"description"`
	AssignedTo    []ContactInfo `bson:"assigned_to,omitempty"`
	DueDate       time.Time     `bson:"due_date"`
	Status        string        `bson:"status"`
	CreatedAt     time.Time     `bson:"created_at"`
	IsReminderSet bool          `bson:"is_reminder_set"`
	ReminderTime  time.Time     `bson:"reminder_time,omitempty"`
	MailSent      bool          `bson:"mail_sent,omitempty"`
	CreatedBy     string        `bson:"created_by"`
}

type Attachment struct {
	ID         string    `bson:"_id,omitempty"`
	FileName   string    `bson:"file_name"`
	FileURL    string    `bson:"file_url"`
	UploadedAt time.Time `bson:"uploaded_at"`
}

type Note struct {
	ID          string    `bson:"_id,omitempty"`
	NoteContent string    `bson:"note_content"`
	CreatedBy   string    `bson:"created_by"`
	CreatedAt   time.Time `bson:"created_at"`
}

type MeetingInfo struct {
	ID                string       `bson:"_id,omitempty"`
	MeetingId         string       `bson:"meeting_id"`
	LeadId            string       `bson:"lead_id,omitempty"`
	Title             string       `bson:"title"`
	MeetingDate       time.Time    `bson:"meeting_date"`
	Mode              string       `bson:"mode"`
	InternalAttendees []Attendee   `bson:"internal_attendees,omitempty"`
	ExternalAttendees []Attendee   `bson:"external_attendees,omitempty"`
	MeetingNotes      []Note       `bson:"meeting_notes,omitempty"`
	ActionItems       []ActionItem `bson:"action_items,omitempty"`
	Attachments       []Attachment `bson:"attachments,omitempty"`
	Outcome           string       `bson:"outcome"`
	ReminderDate      time.Time    `bson:"reminder_date,omitempty"`
	IsReminderSet     bool         `bson:"is_reminder_set"`
	CreatedBy         string       `bson:"created_by"`
	CreatedByName     string       `bson:"created_by_name"`
	CreatedAt         time.Time    `bson:"created_at"`
	UpdatedAt         time.Time    `bson:"updated_at"`
	MailSent          bool         `bson:"mail_sent,omitempty"`
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
	}
	i := int(s) - 1
	if i < 0 || i >= len(stages) {
		return "Unknown"
	}
	return stages[i]
}

func GetClientMainStages() []string {
	return []string{
		"Lead",
		"Suspect",
		"Prospect",
		"Proposal Sent",
		"Negotiation",
		"Contract Signed",
		"Client",
	}
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
	default:
		return LLead // fallback
	}
}
