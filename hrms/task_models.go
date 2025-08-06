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
	Tasks      []Task             `bson:"tasks"`
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

type Task struct {
	StateId     string       `bson:"state_id"`
	Title       string       `bson:"title"`
	Description string       `bson:"description"`
	Assignees   string       `bson:"assignees"`
	AssignedBy  string       `bson:"assigned_by"`
	StartDate   time.Time    `bson:"start_date"`
	DueDate     time.Time    `bson:"due_date"`
	Labels      []string     `bson:"labels"`
	Priority    TaskPriority `bson:"priority"`
	Attachment  string       `bson:"attachment"`
	Remarks     string       `bson:"remarks"`
	UpdatedAt   time.Time    `bson:"updated_at"`
	CreatedAt   time.Time    `bson:"created_at"`
}

type TaskPriority int

const (
	Urgent = iota + 1
	High
	Medium
	Low
)

func (r TaskPriority) String() string {
	return [...]string{"Urgent", "High", "Medium", "Low"}[r-1]
}

func (r TaskPriority) TaskPriorityEnumIndex() int {
	return int(r)
}

func GetTaskPriorities() []string {
	return []string{"Urgent", "High", "Medium", "Low"}
}

func TaskPriorityFromString(s string) TaskPriority {
	switch s {
	case "Urgent":
		return Urgent
	case "High":
		return High
	case "Medium":
		return Medium
	default:
		return Low
	}
}
