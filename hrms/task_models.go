package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectStateList struct {
	ID         primitive.ObjectID `bson:"_id"`
	ProjectId  string             `bson:"project_id"`
	TaskStates []TaskState        `bson:"task_states"`
	CreatedAt  time.Time          `bson:"created_at"`
}

type TaskState struct {
	StateId      string `bson:"state_id"`
	Name         string `bson:"name"`
	Abbreviation string `bson:"abbreviation"`
	Color        string `bson:"color"`
	Description  string `bson:"description"`
	Icon         string `bson:"icon"`
}

type StateIdCounter struct {
	ID        primitive.ObjectID `bson:"_id"`
	ProjectId string             `bson:"project_id"`
	Prefix    string             `bson:"prefix"`
	Counter   int64              `bson:"counter"`
}
