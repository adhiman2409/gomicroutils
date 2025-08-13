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

type TaskIdCounter struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectId string             `bson:"project_id"`
	Prefix    string             `bson:"prefix"`
	Counter   int64              `bson:"counter"`
}

type EmpSummary struct {
	ID                          primitive.ObjectID `bson:"_id"`
	EmployeeId                  string             `bson:"employee_id"`
	FullName                    string             `bson:"full_name"`
	EmailId                     string             `bson:"email_id"`
	EmploymentStatus            string             `bson:"employment_status"`
	Department                  string             `bson:"department"`
	Designation                 string             `bson:"designation"`
	ReportingManagerName        string             `bson:"reporting_manager"`
	ReportingManagerId          string             `bson:"reporting_manager_id"`
	ReportingManagerEmail       string             `bson:"reporting_manager_email"`
	ReportingManagerDesignation string             `bson:"reporting_manager_designation"`
	ReportingManagerDepartment  string             `bson:"reporting_manager_department"`
}

type Task struct {
	TaskId      string       `bson:"task_id"`
	StateId     string       `bson:"state_id"`
	Title       string       `bson:"title"`
	Description string       `bson:"description"`
	Assignees   []EmpSummary `bson:"assignees"`
	AssignedBy  EmpSummary   `bson:"assigned_by"`
	StartDate   time.Time    `bson:"start_date"`
	DueDate     time.Time    `bson:"due_date"`
	Labels      []string     `bson:"labels"`
	Priority    TaskPriority `bson:"priority"`
	Attachments []string     `bson:"attachments"`
	Remarks     []string     `bson:"remarks"`
	UpdatedAt   time.Time    `bson:"updated_at"`
	CreatedAt   time.Time    `bson:"created_at"`
}

type ProjectTasksResponse struct {
	StateId      string `bson:"state_id"`
	Name         string `bson:"name"`
	Abbreviation string `bson:"abbreviation"`
	Color        string `bson:"color"`
	Description  string `bson:"description"`
	Icon         string `bson:"icon"`
	Tasks        []Task `bson:"tasks"`
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
