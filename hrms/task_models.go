package hrms

import "time"

type ProjectStateList struct {
	ProjectId  string      `bson:"project_id"`
	TaskStates []TaskState `bson:"task_states"`
	CreatedAt  time.Time   `bson:"created_at"`
}

type TaskState struct {
	StateId      string `bson:"state_id"`
	Name         string `bson:"name"`
	Abbreviation string `bson:"abbreviation"`
	Color        string `bson:"color"`
	Description  string `bson:"description"`
	Icon         string `bson:"icon"`
}
