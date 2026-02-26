package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Query struct {
	Title     string `bson:"title"`
	SLAInDays int64  `bson:"sla_in_days"`
	Remarks   string `bson:"remarks"`
}

type EmployeeRef struct {
	EmployeeID string `bson:"employee_id"`
	Name       string `bson:"name"`
	Email      string `bson:"email"`
}

type PermissionData struct {
	Name             string `bson:"name"`
	PermissionStatus string `bson:"permission_status"`
}

type TicketConfig struct {
	ID                            primitive.ObjectID `bson:"_id"`
	DepartmentName                string             `bson:"department_name"`
	PrimaryEmployeeId             string             `bson:"primary_emp_id"`
	PrimaryEmployeeName           string             `bson:"primary_emp_name"`
	PrimaryEmployeeEmail          string             `bson:"primary_emp_email"`
	SecondaryEmployeeId           string             `bson:"secondary_emp_id"`
	SecondaryEmployeeName         string             `bson:"secondary_emp_name"`
	SecondaryEmployeeEmail        string             `bson:"secondary_emp_email"`
	IsSecondaryActive             bool               `bson:"is_secondary_active"`
	Resolvers                     []EmployeeRef      `bson:"resolvers"`
	Queries                       []Query            `bson:"queries"`
	TotalNumberOfTickets          int64              `bson:"total_number_of_tickets"`
	NumberOfOpenTickets           int64              `bson:"number_of_open_tickets"`
	NumberOfReopenTickets         int64              `bson:"number_of_reopen_tickets"`
	NumberOfClosedTickets         int64              `bson:"number_of_closed_tickets"`
	AverageTicketCloserTimeInDays int64              `bson:"average_ticket_closer_time_in_days"`
	NotifyNewTicketOnEmail        bool               `bson:"notify_new_ticket_on_email"`
	LockClosedTicketsAfterDays    int64              `bson:"lock_closed_tickets_after_days"`
	SendFeedbackReminders         bool               `bson:"send_feedback_reminders"`
	SendFeedbackReminderAfterDays int64              `bson:"send_feedback_reminder_after_days"`
	Country                       string             `bson:"country"`
	TimeZone                      string             `bson:"time_zone"`
	CreatedAt                     time.Time          `bson:"created_at"`
	UpdatedAt                     time.Time          `bson:"updated_at"`
}

type Ticket struct {
	ID                   primitive.ObjectID `bson:"_id"`
	TicketID             string             `bson:"ticket_id"`
	Title                string             `bson:"title"`
	QuerySLAInDays       int64              `bson:"query_sla_in_days"`
	QueryRemarks         string             `bson:"query_remarks"`
	Description          string             `bson:"description"`
	ReporteeId           string             `bson:"reportee_id"`
	ReporteeEmailId      string             `bson:"reportee_emailId"`
	ReporteeName         string             `bson:"reportee_name"`
	ReporteeDepartment   string             `bson:"reportee_department"`
	ReporteeDomain       string             `bson:"reportee_domain"`
	Department           string             `bson:"department"`
	MasterAssigneeId     string             `bson:"master_assignee_id"`
	MasterAssigneeName   string             `bson:"master_assignee_name"`
	MasterAssigneeEmail  string             `bson:"master_assignee_email"`
	CurrentAssigneeName  string             `bson:"current_assignee_name"`
	CurrentAssigneeId    string             `bson:"current_assignee_id"`
	CurrentAssigneeEmail string             `bson:"current_assignee_email"`
	CurrentResolvers     []EmployeeRef      `bson:"current_resolvers"`
	MasterResolvers      []EmployeeRef      `bson:"master_resolvers"`
	Status               string             `bson:"status"`
	PriorityLevel        string             `bson:"priority_level"`
	TicketType           string             `bson:"ticket_type"`
	Comments             []Comment          `bson:"comments"`
	DueDate              string             `bson:"due_date"`
	OpenDay              int                `bson:"open_day"`
	OpenMonth            int                `bson:"open_month"`
	OpenYear             int                `bson:"open_year"`
	CloseDay             int                `bson:"close_day"`
	CloseMonth           int                `bson:"close_month"`
	CloseYear            int                `bson:"close_year"`
	CloseDurationInDays  int                `bson:"close_duration_in_days"`
	NotificationActive   bool               `bson:"notification_active"`
	SendUpdatesOnEmail   bool               `bson:"send_updates_on_email"`
	FeedbackRemarks      string             `bson:"feedback_remarks"`
	Rating               float32            `bson:"rating"`
	RatingDate           string             `bson:"rating_date"`
	ReminderMailSent     bool               `bson:"reminder_mail_sent"`
	IsLocked             bool               `bson:"is_locked"`
	Country              string             `bson:"country"`
	TimeZone             string             `bson:"time_zone"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}

type Comment struct {
	TicketID       string   `bson:"ticket_id"`
	UserID         string   `bson:"user_id"`
	Username       string   `bson:"user_name"`
	Message        string   `bson:"message"`
	AttachmentURLs []string `bson:"attachment_urls"`
	CreatedAt      string   `bson:"created_at"`
}

type TicketIdCounter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Prefix  string             `bson:"prefix"`
	Counter int64              `bson:"counter"`
}

type Lead struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Name        string             `bson:"name"`
	EmailId     string             `bson:"email_Id"`
	PhoneNumber string             `bson:"reportee_phone_number"`
	CompanyName string             `bson:"company_name"`
	CompanySize string             `bson:"company_size"`
	Remarks     string             `bson:"remarks"`
	TicketId    string             `bson:"ticket_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type GetTrackingTicket struct {
	ID             primitive.ObjectID `bson:"_id"`
	TicketId       string             `bson:"ticket_id"`
	Date           time.Time          `bson:"date"`
	Name           string             `bson:"name"`
	EmployeeId     string             `bson:"employee_id"`
	Designation    string             `bson:"designation"`
	Phone          string             `bson:"phone"`
	Title          string             `bson:"title"`
	Permission     string             `bson:"permission"`
	Priority       string             `bson:"priority"`
	DuplicateOf    string             `bson:"duplicate_of"`
	Description    string             `bson:"description"`
	ImgURL         []string           `bson:"img_url"`
	ResolutionDesc string             `bson:"resolution_desc"`
	TicketStatus   TicketStatus       `bson:"ticket_status"`
	MoreData       []PermissionData   `bson:"more_data"`
	CreatedById    string             `bson:"created_by_id"`
	CreatedByName  string             `bson:"created_by_name"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

type TicketStatus struct {
	Current string   `bson:"current"`
	Allowed []string `bson:"allowed"`
}
