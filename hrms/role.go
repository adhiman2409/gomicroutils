package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ORG_COLLECTION_ROLES = "roles"

type Role struct {
	ID          primitive.ObjectID `bson:"_id"`
	DepId       string             `bson:"dep_id"`
	DepName     string             `bson:"dep_name"`
	Title       string             `bson:"title"`
	IsActive    bool               `bson:"is_active"`
	Description string             `bson:"description"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
