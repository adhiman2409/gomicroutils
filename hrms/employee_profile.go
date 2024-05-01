package hrms

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeProfile struct {
	ID                   primitive.ObjectID `bson:"_id"`
	EmployeeId           string             `bson:"employee_id"`
	FullName             string             `bson:"full_name"`
	PhoneNumber          string             `bson:"phone_number"`
	EmailId              string             `bson:"email_id"`
	EmployeeType         string             `bson:"employee_type"`
	Department           string             `bson:"department"`
	Designation          string             `bson:"designation"`
	Role                 string             `bson:"role"`
	EmploymentStatus     string             `bson:"employment_status"`
	ReportingManagerName string             `bson:"reporting_manager"`
	ImgURL               string             `bson:"img_url"`
}
