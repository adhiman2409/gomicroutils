package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ORG_COLLECTION_DEPARTMENTS = "departments"

type Department struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrgId       string             `bson:"org_id"`
	Name        string             `bson:"name"`
	AdminId     string             `bson:"admin_id"`
	Description string             `bson:"description"`
	IsActive    bool               `bson:"is_active"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
