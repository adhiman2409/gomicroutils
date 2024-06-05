package hrms

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeTeamInfo struct {
	ID           primitive.ObjectID `bson:"_id"`
	EmployeeId   string             `bson:"employee_id"`
	FullName     string             `bson:"full_name"`
	EmailId      string             `bson:"email_id"`
	PhoneNumber  string             `bson:"phone_number"`
	Status       string             `bson:"status"`
	EmployeeType string             `bson:"employee_type"`
	Department   string             `bson:"department"`
	Designation  string             `bson:"designation"`
	Role         string             `bson:"role"`
}
