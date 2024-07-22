package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketConfig struct {
	ID                     primitive.ObjectID `bson:"_id"`
	DepartmentName         string             `bson:"dept_name"`
	PrimaryEmployeeId      string             `bson:"primary_emp_id"`
	PrimaryEmployeeName    string             `bson:"primary_emp_name"`
	PrimaryEmployeeEmail   string             `bson:"primary_emp_email"`
	SecondaryEmployeeId    []string           `bson:"primary_emp_id"`
	SecondaryEmployeeName  []string           `bson:"primary_emp_name"`
	SecondaryEmployeeEmail []string           `bson:"primary_emp_email"`
	Queries                []string           `bson:"queries"`
}
