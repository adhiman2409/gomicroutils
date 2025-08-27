package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyInfo struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	CompanyId          string             `bson:"company_id"`
	CompanyName        string             `bson:"company_name"`
	CompanyLogo        string             `bson:"company_logo"`
	Website            string             `bson:"website"`
	Industry           string             `bson:"industry"`
	HeadquarterCountry string             `bson:"headquarter_country"`
	BusinessActivities string             `bson:"business_activities"`
	ClientDescription  string             `bson:"client_description"`

	YearFounded          int           `bson:"year_founded,omitempty"`
	Contacts             []ContactInfo `bson:"contacts,omitempty"`
	CompanySize          int           `bson:"company_size,omitempty"`
	AnnualRevenue        float64       `bson:"annual_revenue,omitempty"`
	RevenueCurrency      string        `bson:"revenue_currency,omitempty"`
	ProductsAndServices  []string      `bson:"products_and_services,omitempty"`
	OtherOfficeLocations []string      `bson:"other_office_locations,omitempty"`

	SourceOfIntroduction string        `bson:"source_of_introduction"`
	IntroducedBy         string        `bson:"introduced_by"`
	CurrentStage         string        `bson:"current_stage"` // Lead, Prospect, Client, etc.
	AssignedSPOCs        []ContactInfo `bson:"assigned_spocs,omitempty"`
	Tags                 []string      `bson:"tags,omitempty"`
	Priority             string        `bson:"priority,omitempty"`
	IsPublic             bool          `bson:"is_public,omitempty"`
	Attachments          []Attachment  `bson:"attachments,omitempty"`

	CreatedBy string    `bson:"created_by"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type ContactInfo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	ContactId   string             `bson:"contact_id"`
	EmpId       string             `bson:"emp_id"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Mobile      string             `bson:"mobile"`
	Designation string             `bson:"designation"`
	Department  string             `bson:"department"`
	DOB         string             `bson:"dob,omitempty"`
	Anniversary string             `bson:"anniversary "`
	Tags        []string           `bson:"tags"`
	IsActive    bool               `bson:"is_active"`
	Notes       string             `bson:"notes"`
	CompanyName string             `bson:"company_name,omitempty"`
	CompanyId   string             `bson:"company_id,omitempty"`
	RegexString string             `bson:"regex_string,omitempty"`
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
	Id            primitive.ObjectID `bson:"_id"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	AssignedTo    []ContactInfo      `bson:"assigned_to"`
	DueDate       time.Time          `bson:"due_date"`
	Status        string             `bson:"status"` // Open, In Progress, Completed
	CreatedAt     time.Time          `bson:"created_at"`
	IsReminderSet bool               `bson:"is_reminder_set"`
	ReminderTime  time.Time          `bson:"reminder_time,omitempty"`
	MailSent      bool               `bson:"mail_sent,omitempty"`
	CreatedBy     string             `bson:"created_by"`
}

type Attachment struct {
	Id       primitive.ObjectID `bson:"_id"`
	FileName string             `bson:"file_name"`
	FileURL  string             `bson:"file_url"`
}

type Note struct {
	Id          primitive.ObjectID `bson:"_id"`
	NoteContent string             `bson:"note_content"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type MeetingInfo struct {
	Id                primitive.ObjectID `bson:"_id"`
	MeetingId         string             `bson:"meeting_id"` // Link meeting to company/client
	CompanyID         string             `bson:"company_id"` // Link meeting to company/client
	Title             string             `bson:"title"`
	MeetingDate       time.Time          `bson:"meeting_date"`
	Mode              string             `bson:"mode"` // Virtual, Phone, In-person
	InternalAttendees []Attendee         `bson:"internal_attendees"`
	ExternalAttendees []Attendee         `bson:"external_attendees"`
	MeetingNotes      []Note             `bson:"meeting_notes"`
	ActionItems       []ActionItem       `bson:"action_items"`
	Attachments       []Attachment       `bson:"attachments"`
	Outcome           string             `bson:"outcome"` // e.g., Demo Done, Follow-up Needed
	ReminderDate      time.Time          `bson:"reminder_date,omitempty"`
	MailSent          bool               `bson:"mail_sent,omitempty"`
	IsReminderSet     bool               `bson:"is_reminder_set"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}
