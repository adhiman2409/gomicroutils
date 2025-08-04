package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskState struct {
	ID          primitive.ObjectID `bson:"_id"`
	ProjectId   string             `bson:"project_id"`
	StateId     string             `bson:"state_id"`
	Name        string             `bson:"name"`
	Color       string             `bson:"color"`
	Description string             `bson:"description"`
	Icon        string             `bson:"icon"`
	Order       string             `bson:"order"`
	Default     bool               `bson:"default"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	CreatedAt   time.Time          `bson:"created_at"`
}
