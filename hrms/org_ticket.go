package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketConfig struct {
	ID                     primitive.ObjectID `bson:"_id"`
	DepartmentName         string             `bson:"dept_name"`
	PrimaryEmployeeId      string             `bson:"primary_emp_id"`
	PrimaryEmployeeName    string             `bson:"primary_emp_name"`
	PrimaryEmployeeEmail   string             `bson:"primary_emp_email"`
	SecondaryEmployeeId    string             `bson:"secondary_emp_id"`
	SecondaryEmployeeName  string             `bson:"secondary_emp_name"`
	SecondaryEmployeeEmail string             `bson:"secondary_emp_email"`
	Queries                []string           `bson:"queries"`
}

type Ticket struct {
	ID                   primitive.ObjectID `bson:"_id"`
	TicketID             string             `bson:"ticket_id"`
	Title                string             `bson:"title"`
	Description          string             `bson:"description"`
	ReporteeId           string             `bson:"reportee_id"`
	ReporteeEmailId      string             `bson:"reportee_emailId"`
	ReporteeName         string             `bson:"reportee_name"`
	ReporteeDepartment   string             `bson:"reportee_department"`
	AssigneeDepartment   string             `bson:"assignee_department"`
	MasterAssigneeId     string             `bson:"master_assignee_id"`
	MasterAssigneeName   string             `bson:"master_assignee_name"`
	MasterAssigneeEmail  string             `bson:"master_assignee_email"`
	CurrentAssigneeName  string             `bson:"current_assignee_name"`
	CurrentAssigneeId    string             `bson:"current_assignee_id"`
	CurrentAssigneeEmail string             `bson:"current_assignee_email"`
	IsViewed             bool               `bson:"is_viewed"`
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
	Rating               float32            `bson:"rating"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}

type Comment struct {
	TicketID       string    `bson:"ticket_id"`
	UserID         string    `bson:"user_id"`
	Username       string    `bson:"user_name"`
	Message        string    `bson:"message"`
	AttachmentURLs []string  `bson:"attachment_urls"`
	CreatedAt      time.Time `bson:"created_at"`
}
