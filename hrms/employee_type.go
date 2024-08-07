package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
