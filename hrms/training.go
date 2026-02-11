package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TrainingCourse represents a training course in the system (course catalog)
type TrainingCourse struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TrainingId      string             `bson:"training_id" json:"training_id"`
	Title           string             `bson:"title" json:"title"`
	Description     string             `bson:"description" json:"description"`
	Category        string             `bson:"category" json:"category"` // Technical, Compliance, Soft Skills
	DurationInHours float32            `bson:"duration_in_hours" json:"duration_in_hours"`

	// Materials (GCS paths)
	Materials []TrainingMaterial `bson:"materials" json:"materials"`

	// Assignment Rules
	IsMandatory     bool           `bson:"is_mandatory" json:"is_mandatory"`
	AssignmentRules AssignmentRule `bson:"assignment_rules,omitempty" json:"assignment_rules,omitempty"`

	// Publishing
	Status      string    `bson:"status" json:"status"` // draft/published/archived
	PublishedBy string    `bson:"published_by,omitempty" json:"published_by,omitempty"`
	PublishedAt time.Time `bson:"published_at,omitempty" json:"published_at,omitempty"`

	// Metadata
	CreatedBy string    `bson:"created_by" json:"created_by"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// TrainingMaterial represents a single material (PPT, Doc, or Video) in a training
type TrainingMaterial struct {
	SectionIndex    int       `bson:"section_index" json:"section_index"` // 0, 1, 2...
	SectionTitle    string    `bson:"section_title" json:"section_title"`
	MaterialType    string    `bson:"material_type" json:"material_type"` // ppt/doc/video
	FileName        string    `bson:"file_name" json:"file_name"`
	FileURL         string    `bson:"file_url" json:"file_url"` // GCS path
	FileSize        int64     `bson:"file_size" json:"file_size"`
	DurationMinutes int       `bson:"duration_minutes,omitempty" json:"duration_minutes,omitempty"` // For videos
	UploadedAt      time.Time `bson:"uploaded_at" json:"uploaded_at"`
}

// AssignmentRule defines which employees should be auto-assigned a mandatory training
type AssignmentRule struct {
	Departments  []string `bson:"departments" json:"departments"`
	Designations []string `bson:"designations" json:"designations"`
}

// TrainingEnrollment represents an employee's enrollment in a training
type TrainingEnrollment struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// Training Reference (denormalized for performance)
	TrainingId       string  `bson:"training_id" json:"training_id"`
	TrainingTitle    string  `bson:"training_title" json:"training_title"`
	TrainingCategory string  `bson:"training_category" json:"training_category"`
	DurationInHours  float32 `bson:"duration_in_hours" json:"duration_in_hours"`

	// Employee Reference
	EmployeeId    string `bson:"employee_id" json:"employee_id"`
	EmployeeName  string `bson:"employee_name" json:"employee_name"`
	EmployeeEmail string `bson:"employee_email" json:"employee_email"`
	Department    string `bson:"department" json:"department"`
	Designation   string `bson:"designation" json:"designation"`
	ManagerId     string `bson:"manager_id" json:"manager_id"`
	ManagerName   string `bson:"manager_name" json:"manager_name"`
	ManagerEmail  string `bson:"manager_email" json:"manager_email"`

	// Enrollment Type
	EnrollmentType string `bson:"enrollment_type" json:"enrollment_type"` // mandatory/optional

	// Optional Training Approval (only for optional trainings)
	ApplicationDate time.Time `bson:"application_date,omitempty" json:"application_date,omitempty"`
	EmployeeRemarks string    `bson:"employee_remarks,omitempty" json:"employee_remarks,omitempty"`
	ApprovalStatus  string    `bson:"approval_status,omitempty" json:"approval_status,omitempty"` // pending/approved/rejected
	ManagerRemarks  string    `bson:"manager_remarks,omitempty" json:"manager_remarks,omitempty"`
	ApprovedAt      time.Time `bson:"approved_at,omitempty" json:"approved_at,omitempty"`

	// Progress Tracking
	Status      string    `bson:"status" json:"status"` // assigned/in_progress/completed
	StartedAt   time.Time `bson:"started_at,omitempty" json:"started_at,omitempty"`
	CompletedAt time.Time `bson:"completed_at,omitempty" json:"completed_at,omitempty"`

	// Video Progress (embedded array)
	VideoProgress   []VideoProgress `bson:"video_progress" json:"video_progress"`
	OverallProgress int             `bson:"overall_progress" json:"overall_progress"` // 0-100

	// Deadline Management
	AssignedDate time.Time `bson:"assigned_date" json:"assigned_date"`
	Deadline     time.Time `bson:"deadline" json:"deadline"`

	// Reminder Tracking
	RemindersSent []ReminderLog `bson:"reminders_sent" json:"reminders_sent"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// VideoProgress tracks an employee's progress on a specific video section
type VideoProgress struct {
	SectionIndex    int       `bson:"section_index" json:"section_index"`
	SectionTitle    string    `bson:"section_title" json:"section_title"`
	FileURL         string    `bson:"file_url" json:"file_url"`
	WatchPercentage int       `bson:"watch_percentage" json:"watch_percentage"` // 0-100
	LastWatchedAt   time.Time `bson:"last_watched_at" json:"last_watched_at"`
}

// ReminderLog tracks when reminders were sent for a training enrollment
type ReminderLog struct {
	ReminderType string    `bson:"reminder_type" json:"reminder_type"` // 7_days/3_days/1_day
	SentAt       time.Time `bson:"sent_at" json:"sent_at"`
}

// TrainingCounter is used for generating unique training IDs
type TrainingCounter struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Prefix  string             `bson:"prefix" json:"prefix"` // "TRAIN"
	Counter int64              `bson:"counter" json:"counter"`
}
