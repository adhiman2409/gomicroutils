package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrgId       string             `bson:"org_id"`
	DepId       string             `bson:"dep_id"`
	Title       string             `bson:"title"`
	IsActive    bool               `bson:"is_active"`
	Description string             `bson:"description"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
