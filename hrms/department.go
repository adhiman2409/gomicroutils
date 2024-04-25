package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrgId          string             `bson:"org_id"`
	OrgName        string             `bson:"org_name"`
	Name           string             `bson:"name"`
	Alias          string             `bson:"alias"`
	AdminId        string             `bson:"admin_id"`
	AdminName      string             `bson:"admin_name"`
	TotalEmployees uint32             `bson:"total_employees"`
	EmployeeList   []EmployeeSummary  `bson:"employee_list"`
	Description    string             `bson:"description"`
	IsActive       bool               `bson:"is_active"`
	CreatedBy      string             `bson:"created_by"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
