package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TrainingModule is an ordered group of lessons inside a TrainingCourse.
type TrainingModule struct {
	ModuleId    string           `bson:"module_id" json:"module_id"` // MOD-<12 hex>
	Title       string           `bson:"title" json:"title"`
	Description string           `bson:"description,omitempty" json:"description,omitempty"`
	Order       int              `bson:"order" json:"order"`
	Lessons     []TrainingLesson `bson:"lessons" json:"lessons"`
}

// TrainingLesson is one unit of content. Exactly one content field is used,
// depending on Type: video/slides/document -> AssetId, embed -> Embed,
// text -> Body, quiz -> QuizId.
type TrainingLesson struct {
	LessonId        string         `bson:"lesson_id" json:"lesson_id"` // LES-<12 hex>, stable progress key
	Title           string         `bson:"title" json:"title"`
	Type            string         `bson:"type" json:"type"` // video/slides/document/embed/text/quiz
	Order           int            `bson:"order" json:"order"`
	Required        bool           `bson:"required" json:"required"`
	DurationMinutes int            `bson:"duration_minutes,omitempty" json:"duration_minutes,omitempty"`
	AssetId         string         `bson:"asset_id,omitempty" json:"asset_id,omitempty"`
	Embed           *TrainingEmbed `bson:"embed,omitempty" json:"embed,omitempty"`
	Body            string         `bson:"body,omitempty" json:"body,omitempty"` // sanitized HTML
	QuizId          string         `bson:"quiz_id,omitempty" json:"quiz_id,omitempty"`
	SlideCount      int            `bson:"slide_count,omitempty" json:"slide_count,omitempty"`
	DripDays        int            `bson:"drip_days,omitempty" json:"drip_days,omitempty"`
	ArchivedAt      time.Time      `bson:"archived_at,omitempty" json:"archived_at,omitempty"` // soft delete
}

// TrainingEmbed is a validated, normalized third-party video link.
type TrainingEmbed struct {
	Provider string `bson:"provider" json:"provider"` // youtube/vimeo/loom/gdrive
	URL      string `bson:"url" json:"url"`           // normalized embed URL
	VideoId  string `bson:"video_id" json:"video_id"`
}

// LessonProgress is an employee's progress on one lesson (keyed by LessonId).
type LessonProgress struct {
	LessonId            string    `bson:"lesson_id" json:"lesson_id"`
	ModuleId            string    `bson:"module_id" json:"module_id"`
	State               string    `bson:"state" json:"state"` // not_started/in_progress/completed
	Percentage          int       `bson:"percentage" json:"percentage"`
	LastPositionSeconds int       `bson:"last_position_seconds,omitempty" json:"last_position_seconds,omitempty"`
	CompletedAt         time.Time `bson:"completed_at,omitempty" json:"completed_at,omitempty"`
	UpdatedAt           time.Time `bson:"updated_at" json:"updated_at"`
}

// TrainingAsset records a file stored through the existing employee upload API
// (POST /employee/upload); DocPath is the doc_path it returned.
type TrainingAsset struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AssetId         string             `bson:"asset_id" json:"asset_id"` // AST-<12 hex>
	TrainingId      string             `bson:"training_id" json:"training_id"`
	Kind            string             `bson:"kind" json:"kind"` // video/slides/document/image
	FileName        string             `bson:"file_name" json:"file_name"`
	ContentType     string             `bson:"content_type" json:"content_type"`
	SizeBytes       int64              `bson:"size_bytes" json:"size_bytes"`
	DurationSeconds int                `bson:"duration_seconds,omitempty" json:"duration_seconds,omitempty"` // video only
	DocPath         string             `bson:"doc_path" json:"doc_path"`
	UploadedBy      string             `bson:"uploaded_by" json:"uploaded_by"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
}

// AudienceRules selects employees by profile facets. Values within a facet are
// OR'ed, facets are AND'ed, an empty facet does not filter.
type AudienceRules struct {
	Departments     []string `bson:"departments" json:"departments"`
	Designations    []string `bson:"designations" json:"designations"`
	WorkLocations   []string `bson:"work_locations" json:"work_locations"`
	OfficeLocations []string `bson:"office_locations" json:"office_locations"`
	EmployeeTypes   []string `bson:"employee_types" json:"employee_types"`
}

// InviteCounts summarizes one invite batch.
type InviteCounts struct {
	Matched         int `bson:"matched" json:"matched"`
	Assigned        int `bson:"assigned" json:"assigned"`
	AlreadyEnrolled int `bson:"already_enrolled" json:"already_enrolled"`
	SkippedInactive int `bson:"skipped_inactive" json:"skipped_inactive"`
	Failed          int `bson:"failed" json:"failed"`
}

// InviteResult is the outcome for one employee in a batch.
type InviteResult struct {
	EmployeeId string `bson:"employee_id" json:"employee_id"`
	Name       string `bson:"name" json:"name"`
	Status     string `bson:"status" json:"status"` // assigned/already_enrolled/inactive/not_found/error
}

// TrainingInviteBatch is the audit record of one admin invite.
type TrainingInviteBatch struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BatchId        string             `bson:"batch_id" json:"batch_id"` // INV-<12 hex>
	TrainingId     string             `bson:"training_id" json:"training_id"`
	Mode           string             `bson:"mode" json:"mode"` // rules/ids/all
	Rules          AudienceRules      `bson:"rules" json:"rules"`
	EmployeeIds    []string           `bson:"employee_ids" json:"employee_ids"`
	ExcludeIds     []string           `bson:"exclude_ids" json:"exclude_ids"`
	EnrollmentType string             `bson:"enrollment_type" json:"enrollment_type"`
	DeadlineDays   int                `bson:"deadline_days" json:"deadline_days"`
	Notify         bool               `bson:"notify" json:"notify"`
	Message        string             `bson:"message,omitempty" json:"message,omitempty"`
	Counts         InviteCounts       `bson:"counts" json:"counts"`
	Failures       []InviteResult     `bson:"failures" json:"failures"` // only people NOT enrolled; enrolled ones are found via enrollments.invite_batch_id
	Status         string             `bson:"status" json:"status"`     // queued/running/completed/partial
	CreatedBy      string             `bson:"created_by" json:"created_by"`
	CreatedByName  string             `bson:"created_by_name" json:"created_by_name"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	CompletedAt    time.Time          `bson:"completed_at,omitempty" json:"completed_at,omitempty"`
}

// ---------- M2: quizzes ----------

type QuizOption struct {
	OptionId string `bson:"option_id" json:"option_id"`
	Text     string `bson:"text" json:"text"`
}

type QuizQuestion struct {
	QuestionId       string       `bson:"question_id" json:"question_id"`
	Type             string       `bson:"type" json:"type"` // single/multiple/true_false
	Text             string       `bson:"text" json:"text"`
	Options          []QuizOption `bson:"options" json:"options"`
	CorrectOptionIds []string     `bson:"correct_option_ids" json:"correct_option_ids"` // never sent to learners pre-review
	Points           int          `bson:"points" json:"points"`
	Explanation      string       `bson:"explanation,omitempty" json:"explanation,omitempty"`
}

type TrainingQuiz struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	QuizId           string             `bson:"quiz_id" json:"quiz_id"` // QUZ-<12 hex>
	TrainingId       string             `bson:"training_id" json:"training_id"`
	LessonId         string             `bson:"lesson_id" json:"lesson_id"`
	PassMarkPct      int                `bson:"pass_mark_pct" json:"pass_mark_pct"`
	MaxAttempts      int                `bson:"max_attempts" json:"max_attempts"` // 0 = unlimited
	TimeLimitMinutes int                `bson:"time_limit_minutes" json:"time_limit_minutes"`
	ShuffleQuestions bool               `bson:"shuffle_questions" json:"shuffle_questions"`
	ShuffleOptions   bool               `bson:"shuffle_options" json:"shuffle_options"`
	ShowAnswers      string             `bson:"show_answers" json:"show_answers"` // never/after_submit/after_pass
	Questions        []QuizQuestion     `bson:"questions" json:"questions"`
	Version          int                `bson:"version" json:"version"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

type QuizAnswer struct {
	QuestionId string   `bson:"question_id" json:"question_id"`
	OptionIds  []string `bson:"option_ids" json:"option_ids"`
}

type TrainingQuizAttempt struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AttemptId    string             `bson:"attempt_id" json:"attempt_id"` // ATT-<12 hex>
	QuizId       string             `bson:"quiz_id" json:"quiz_id"`
	QuizVersion  int                `bson:"quiz_version" json:"quiz_version"`
	TrainingId   string             `bson:"training_id" json:"training_id"`
	EnrollmentId string             `bson:"enrollment_id" json:"enrollment_id"`
	EmployeeId   string             `bson:"employee_id" json:"employee_id"`
	QuestionIds  []string           `bson:"question_ids" json:"question_ids"` // order served
	Answers      []QuizAnswer       `bson:"answers" json:"answers"`
	ScorePoints  int                `bson:"score_points" json:"score_points"`
	MaxPoints    int                `bson:"max_points" json:"max_points"`
	ScorePct     int                `bson:"score_pct" json:"score_pct"`
	Passed       bool               `bson:"passed" json:"passed"`
	Status       string             `bson:"status" json:"status"` // in_progress/submitted
	StartedAt    time.Time          `bson:"started_at" json:"started_at"`
	ExpiresAt    time.Time          `bson:"expires_at,omitempty" json:"expires_at,omitempty"`
	SubmittedAt  time.Time          `bson:"submitted_at,omitempty" json:"submitted_at,omitempty"`
}

// ---------- M2: certificates ----------

type TrainingCertificate struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CertificateId    string             `bson:"certificate_id" json:"certificate_id"`       // CRT-<12 hex>
	VerificationCode string             `bson:"verification_code" json:"verification_code"` // 16-char base32
	TrainingId       string             `bson:"training_id" json:"training_id"`
	TrainingTitle    string             `bson:"training_title" json:"training_title"`
	EmployeeId       string             `bson:"employee_id" json:"employee_id"`
	EmployeeName     string             `bson:"employee_name" json:"employee_name"`
	ScorePct         int                `bson:"score_pct,omitempty" json:"score_pct,omitempty"`
	TemplateId       string             `bson:"template_id,omitempty" json:"template_id,omitempty"`
	IssuedAt         time.Time          `bson:"issued_at" json:"issued_at"`
	RevokedAt        time.Time          `bson:"revoked_at,omitempty" json:"revoked_at,omitempty"`
}

type TrainingCertificateTemplate struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TemplateId       string             `bson:"template_id" json:"template_id"`
	Name             string             `bson:"name" json:"name"`
	OrgName          string             `bson:"org_name" json:"org_name"`
	SignatoryName    string             `bson:"signatory_name" json:"signatory_name"`
	SignatoryTitle   string             `bson:"signatory_title" json:"signatory_title"`
	SignatureAssetId string             `bson:"signature_asset_id,omitempty" json:"signature_asset_id,omitempty"`
	LogoAssetId      string             `bson:"logo_asset_id,omitempty" json:"logo_asset_id,omitempty"`
	AccentColor      string             `bson:"accent_color,omitempty" json:"accent_color,omitempty"`
	IsDefault        bool               `bson:"is_default" json:"is_default"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}
