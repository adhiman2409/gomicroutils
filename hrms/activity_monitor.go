package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityType string

const (
	ApplicationUsage ActivityType = "application_usage"
	WebBrowsing      ActivityType = "web_browsing"
	Screenshot       ActivityType = "screenshot"
)

type Activity struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Day            int                `bson:"day"`   // DD
	Month          int                `bson:"month"` // MM
	Year           int                `bson:"year"`  // YYYY
	EmployeeID     string             `bson:"employee_id"`
	UnixTimestamp  int64              `bson:"timestamp"`
	ActivityType   ActivityType       `bson:"activity_type"` // "app_usage", "web_history", "screenshot"
	AppUsage       []AppUsage         `bson:"app_usage,omitempty"`
	WebHistory     []WebHistory       `bson:"web_history,omitempty"`
	ScreenshotURLs []string           `bson:"screenshot_urls,omitempty"`
}

type AppUsage struct {
	AppName      string    `bson:"app_name"`
	Window       string    `bson:"window"`
	StartTime    time.Time `bson:"start_time"`
	EndTime      time.Time `bson:"end_time"`
	Keystrokes   int       `bson:"keystrokes"`
	MouseClicks  int       `bson:"mouse_clicks"`
	Scrolls      int       `bson:"scrolls"`
	IdleTime     int       `bson:"idle_time"`
	TotalTime    int       `bson:"total_time"`
	Intensity    int       `bson:"intensity"`
	ProcessName  string    `bson:"process_name"`
	ProcessPath  string    `bson:"process_path"`
	ProcessId    int       `bson:"process_id"`
	Url          string    `bson:"url"`
	Domain       string    `bson:"domain"`
	Title        string    `bson:"title"`
	Productive   string    `bson:"productive"`
	Productivity int       `bson:"productivity"`
}

type WebHistory struct {
	URL       string    `bson:"url"`
	Title     string    `bson:"title"`
	StartTime time.Time `bson:"start_time"`
	EndTime   time.Time `bson:"end_time"`
}
