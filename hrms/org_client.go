package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type ClientHoliday struct {
// 	ID          primitive.ObjectID `bson:"_id"`
// 	ClientId    string             `bson:"client_id"`
// 	ClientName  string             `bson:"client_name"`
// 	Day         string             `bson:"day"`
// 	Month       string             `bson:"month"`
// 	Year        string             `bson:"year"`
// 	Type        string             `bson:"type"`
// 	WeekDay     string             `bson:"week_day"`
// 	Description string             `bson:"description"`
// }

type ClientInfo struct {
	ID                         primitive.ObjectID `bson:"_id"`
	OrganizationId             string             `bson:"organization_id"`
	OrganizationName           string             `bson:"organization_name"`
	ClientId                   string             `bson:"client_id"`
	ClientName                 string             `bson:"client_name"`
	ClientType                 string             `bson:"client_type"`
	ClientStatus               string             `bson:"client_status"`
	ClientDescription          string             `bson:"client_description"`
	IsDefault                  bool               `bson:"is_default"`
	IsExternalProject          bool               `bson:"is_external_project"`
	IsNDASigned                bool               `bson:"is_nda_signed"`
	NDASignDate                time.Time          `bson:"nda_sign_date"`
	NDADocument                string             `bson:"nda_document"`
	ClientWebsite              string             `bson:"client_website"`
	PrimaryContactName         string             `bson:"primary_contact_name"`
	PrimaryContactEmail        string             `bson:"primary_contact_email"`
	PrimaryContactIDesignation string             `bson:"primary_contact_designation"`
	ProjectHeadId              string             `bson:"project_head_id"`
	ProjectHeadName            string             `bson:"project_head_name"`
	ProjectHeadEmail           string             `bson:"project_head_email"`
	ExternalProjectHeadId      string             `bson:"external_project_head_id"`
	ExternalProjectHeadName    string             `bson:"external_project_head_name"`
	ExternalProjectHeadEmail   string             `bson:"external_project_head_email"`
	FullAddress                Address            `bson:"full_address"`
	City                       string             `bson:"city"`
	State                      string             `bson:"state"`
	Country                    string             `bson:"country"`
	Contacts                   []string           `bson:"contacts"`
	CreatedBy                  string             `bson:"created_by"`
	CreatedAt                  time.Time          `bson:"created_at"`
	UpdatedAt                  time.Time          `bson:"updated_at"`
}

// type Participant struct {
// 	IsInternal  bool   `bson:"is_internal"`
// 	Id          string `bson:"id"`
// 	Name        string `bson:"name"`
// 	EmailId     string `bson:"email_id"`
// 	PhoneNumber string `bson:"phone_number"`
// 	Designation string `bson:"designation"`
// 	Remarks     string `bson:"remarks"`
// }

// type ActionItem struct {
// 	Items                  []string  `bson:"items"`
// 	Status                 string    `bson:"status"`
// 	IsClientActionItem     string    `bson:"client_action_item"`
// 	OwnerId                string    `bson:"owner_id"`
// 	OwnerName              string    `bson:"owner_name"`
// 	OwnerEmailId           string    `bson:"owner_email_id"`
// 	SendActionItemsToOwner bool      `bson:"send_action_items_to_owner"`
// 	CompletionDateAndTime  time.Time `bson:"completion_date_and_time"`
// 	ReminderDateAndTime    time.Time `bson:"reminder_date_and_time"`
// 	IsReminderMailSent     bool      `bson:"is_reminder_mail_sent"`
// 	Remarks                string    `bson:"remarks"`
// 	Documents              []string  `bson:"documents"`
// 	CreatedById            string    `bson:"created_by_id"`
// 	CreatedByName          string    `bson:"created_by_name"`
// }

// type MeetingInfo struct {
// 	ID                             primitive.ObjectID `bson:"_id"`
// 	MeetingId                      string             `bson:"meeting_id"`
// 	ClientId                       string             `bson:"client_id"`
// 	ClientName                     string             `bson:"client_name"`
// 	MeetingDateAndTime             time.Time          `bson:"meeting_date_and_time"`
// 	MeetingType                    string             `bson:"meeting_type"`
// 	MeetingLocation                string             `bson:"meeting_location"`
// 	Participants                   []Participant      `bson:"participants"`
// 	Title                          string             `bson:"title"`
// 	Description                    string             `bson:"description"`
// 	Documents                      []string           `bson:"documents"`
// 	MinutsOfMeeting                []string           `bson:"minuts_of_meeting"`
// 	ActionItems                    []ActionItem       `bson:"action_items"`
// 	TotalActionItems               int                `bson:"total_action_items"`
// 	PendingActionItems             int                `bson:"pending_action_items"`
// 	CompletedActionItems           int                `bson:"completed_action_items"`
// 	NextMeetingDate                time.Time          `bson:"next_meeting_date"`
// 	NextMeetingRemarks             time.Time          `bson:"next_meeting_remarks"`
// 	NextMeetingReminderDateAndTime time.Time          `bson:"next_meeting_reminder_date_and_time"`
// 	IsReminderMailSent             bool               `bson:"is_reminder_mail_sent"`
// 	NextMeetingId                  string             `bson:"next_meeting_id"`
// 	PreviousMeetingId              string             `bson:"previous_meeting_id"`
// 	CreatedBy                      string             `bson:"created_by"`
// 	CreatedAt                      time.Time          `bson:"created_at"`
// 	UpdatedAt                      time.Time          `bson:"updated_at"`
// }

type ClientContactInfo struct {
	ID                 primitive.ObjectID `bson:"_id"`
	ClientId           string             `bson:"client_id"`
	ClientName         string             `bson:"client_name"`
	ContactName        string             `bson:"contact_name"`
	ContactEmail       string             `bson:"contact_email"`
	ContactPhone       string             `bson:"contact_phone"`
	ContactDesignation string             `bson:"contact_designation"`
	ContactRole        string             `bson:"contact_role"`
	ContactDescription string             `bson:"contact_description"`
	CreatedBy          string             `bson:"created_by"`
	ReferencedBy       string             `bson:"referenced_by"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
}

type ClientIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}
