package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TimeSheetTask struct {
	ID          string    `bson:"_id"`
	TaskID      string    `bson:"task_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	StartDate   time.Time `bson:"start_date"`
	EndDate     time.Time `bson:"end_date"`
	EmployeeID  string    `bson:"employee_id"`
	ManagerID   string    `bson:"manager_id"`
	CreatedBy   string    `bson:"created_by"`
}

type MonthlyTimeSheet struct {
	ID             primitive.ObjectID `bson:"_id"`
	Month          int                `bson:"month"`
	Year           int                `bson:"year"`
	EmployeeID     string             `bson:"employee_id"`
	ManagerID      string             `bson:"manager_id"`
	DailyTimesheet []TimeSheet        `bson:"daily_timesheet"`
}

type TimeSheet struct {
	Day               int              `bson:"day"`
	FirstCheckInTime  time.Time        `bson:"first_check_in_time"`
	LastCheckOutTime  time.Time        `bson:"last_check_out_time"`
	TotalWorkingHours float64          `bson:"total_working_hours"`
	IsWorkingDay      bool             `bson:"is_working_day"`
	IsHoliday         bool             `bson:"is_holiday"`
	IsOnLeave         bool             `bson:"is_on_leave"`
	IsWeeklyOffDay    bool             `bson:"is_weekly_off_day"`
	Entries           []TimeSheetEntry `bson:"entries"`
}

type TimeSheetEntry struct {
	TaskID string  `bson:"task_id"`
	Hours  float64 `bson:"hours"`
	Remark string  `bson:"remark"`
}
