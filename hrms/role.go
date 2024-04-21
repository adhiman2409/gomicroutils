package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id"`
	DepId       primitive.ObjectID `bson:"dep_id"`
	Title       string             `bson:"title"`
	IsActive    bool               `bson:"is_active"`
	Description string             `bson:"description"`
	CreatedBy   string             `bson:"created_by"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
