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
	SecondaryEmployeeId    string             `bson:"primary_emp_id"`
	SecondaryEmployeeName  string             `bson:"primary_emp_name"`
	SecondaryEmployeeEmail string             `bson:"primary_emp_email"`
	Queries                []string           `bson:"queries"`
}

type Ticket struct {
	ID                     primitive.ObjectID `bson:"_id"`
	TicketID               string             `bson:"ticket_id"`
	Title                  string             `bson:"title"`
	Description            string             `bson:"description"`
	ReporterEmailId        string             `bson:"reporter_emailId"`
	ReporterName           string             `bson:"reporter_name"`
	ReporterDepartment     string             `bson:"reporter_department"`
	AssigneeDepartment     string             `bson:"assignee_department"`
	Status                 string             `bson:"status"`
	PriorityLevel          string             `bson:"priority_level"`
	TicketType             string             `bson:"ticket_type"`
	Comments               []Comment          `bson:"comments"`
	DueDate                string             `bson:"due_date"`
	CreatedAt              time.Time          `bson:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at"`
}

type Comment struct {
	TicketID  string    `bson:"ticket_id"`
	UserID    string    `bson:"user_id"`
	Username  string    `bson:"user_name"`
	Message   string    `bson:"message"`
	CreatedAt time.Time `bson:"created_at"`
}
