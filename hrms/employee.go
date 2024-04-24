package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const EMPLOYEE_COLLECTION_PERSONAL = "employee-personal"
const EMPLOYEE_COLLECTION_PROFESSIONAL = "employee-professional"
const EMPLOYEE_COLLECTION_DOCUMENTS = "employee-documents"
const EMPLOYEE_COLLECTION_SUMMARY = "employee-summary"
const ORG_COLLECTION_EMPLOYEE_TYPE = "employee-types"

type EmployeeType struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrgId       string             `bson:"org_id"`
	OrgName     string             `bson:"org_name"`
	DepId       string             `bson:"dep_id"`
	DepName     string             `bson:"dep_name"`
	Title       string             `bson:"title"`
	IsActive    bool               `bson:"is_active"`
	Description string             `bson:"description"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type EmployeeSummary struct {
	EmpId  string `bson:"emp_id"`
	Name   string `bson:"name"`
	Role   string `bson:"role"`
	ImgURL string `bson:"img_url"`
}