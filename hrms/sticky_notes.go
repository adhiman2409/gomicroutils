package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StickyNote struct {
	ID               primitive.ObjectID `bson:"_id"`
	ProjectId        string             `bson:"project_id"`
	Content          string             `bson:"content"`
	SharedWith       []string           `bson:"shared_with"`
	SharedWithDetail []SharedWith       `bson:"shared_with_info"`
	CreatedBy        string             `bson:"created_by"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
	UpdatedBy        string             `bson:"updated_by"`
	UpdateHistory    []UpdateStickyNote `bson:"update_history"`
	Position         Position           `bson:"position"`
}

type SharedWith struct {
	ID    string `bson:"id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type Position struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
}

type UpdateStickyNote struct {
	UpdatedAt time.Time `bson:"updated_at"`
	UpdatedBy string    `bson:"updated_by"`
}
