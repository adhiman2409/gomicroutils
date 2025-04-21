package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientInfo struct {
	ID                primitive.ObjectID `bson:"_id"`
	ClientId          string             `bson:"client_id"`
	ClientName        string             `bson:"client_name"`
	ClientType        string             `bson:"client_type"`
	ClientStatus      string             `bson:"client_status"`
	ClientDescription string             `bson:"client_description"`
	ClientWebsite     string             `bson:"client_website"`
	ClientAddress     []Address          `bson:"client_address"`
	CreatedBy         string             `bson:"created_by"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
}

type ClientHoliday struct {
	ID          primitive.ObjectID `bson:"_id"`
	ClientId    string             `bson:"client_id"`
	ClientName  string             `bson:"client_name"`
	Day         string             `bson:"day"`
	Month       string             `bson:"month"`
	Year        string             `bson:"year"`
	Type        string             `bson:"type"`
	WeekDay     string             `bson:"week_day"`
	Description string             `bson:"description"`
}

type KYCInfo struct {
	ID             primitive.ObjectID  `bson:"_id"`
	ClientId       string              `bson:"client_id"`
	Date           string              `bson:"date"`
	ClientContacts []ClientContactInfo `bson:"client_contacts"`
	PSContacts     []string            `bson:"ps_contacts"`
	Remarks        string              `bson:"remarks"`
	ActionItems    []string            `bson:"action_items"`
	ReminderDate   string              `bson:"reminder_date"`
	NextDate       string              `bson:"next_date"`
	Documents      []string            `bson:"documents"`
	CreatedBy      string              `bson:"created_by"`
	CreatedAt      time.Time           `bson:"created_at"`
	UpdatedAt      time.Time           `bson:"updated_at"`
}

type ClientContactInfo struct {
	ID                 primitive.ObjectID `bson:"_id"`
	ClientId           string             `bson:"client_id"`
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
