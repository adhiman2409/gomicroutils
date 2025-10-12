package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	EventTypeEmployeeCheckedIn      = "employee_checked_in"
	EventTypeEmployeeCheckedOut     = "employee_checked_out"
	EventTypeMonitoringControl      = "monitoring_control"
	EventTypeMonitoringSession      = "monitoring_session"
	EventTypeApplicationActivity    = "application_activity"
	EventTypeApplicationSubActivity = "application_sub_activity"
	EventTypeUserActivity           = "user_activity"
	EventTypeBrowserActivity        = "browser_activity"
	EventTypeBrowserSubActivity     = "browser_sub_activity"
	EventTypeDocumentActivity       = "document_activity"
	EventTypeUSBActivity            = "usb_activity"
	EventTypeLocationUpdate         = "location_update"
	EventTypeScreenshot             = "screenshot"
	EventTypeSystemStatus           = "system_status"
	EventTypeNetworkStatus          = "network_status"
	EventTypeApplicationUsage       = "application_usage"
)

type MonitoringConfig struct {
	Id                              primitive.ObjectID `bson:"_id"`
	StartMonitoringAfterCheckInOnly bool               `bson:"start_monitoring_after_check_in_only"`
	IsCheckedIn                     bool               `bson:"is_checked_in"`
	IsCheckedOut                    bool               `bson:"is_checked_out"`
	UserId                          string             `bson:"user_id"`
	MacAddress                      string             `bson:"mac_address"`
	Name                            string             `bson:"name"`
	Domain                          string             `bson:"domain"`
	Department                      string             `bson:"department"`
	Email                           string             `bson:"email"`
	MonitoringEnabled               bool               `bson:"monitoring_enabled"`
	IdleThreshold                   int                `bson:"idle_threshold"`
	ScreenshotInterval              int                `bson:"screenshot_interval"`
	TrackApplications               bool               `bson:"track_applications"`
	TrackLocation                   bool               `bson:"track_location"`
	TrackBrowser                    bool               `bson:"track_browser"`
	TrackDocuments                  bool               `bson:"track_documents"`
	TrackUSB                        bool               `bson:"track_usb"`
}

type ActivityLogEntry struct {
	Id                     primitive.ObjectID     `bson:"_id"`
	Day                    int                    `bson:"day"`
	Month                  int                    `bson:"month"`
	Year                   int                    `bson:"year"`
	EmployeeID             string                 `bson:"employee_id"`
	User                   string                 `bson:"user"`
	Timestamp              time.Time              `bson:"timestamp"`
	UserID                 string                 `bson:"user_id"`
	MacAddress             string                 `bson:"mac_address"`
	Message                string                 `bson:"message,omitempty"` // Kept for backward compatibility
	EventType              string                 `bson:"event_type"`
	Hostname               string                 `bson:"hostname"`
	IPAddress              string                 `bson:"ip_address,omitempty"`
	SessionID              string                 `bson:"session_id,omitempty"`
	Domain                 string                 `bson:"domain,omitempty"`
	ActivityInfo           ActivityInfo           `bson:"activity_info,omitempty"`
	UserActivityInfo       UserActivityInfo       `bson:"user_activity_info,omitempty"`
	SystemStatusInfo       SystemStatusInfo       `bson:"system_status_info,omitempty"`
	USBDeviceInfo          USBDeviceInfo          `bson:"usb_device_info,omitempty"`
	LocationInfo           LocationInfo           `bson:"location_info,omitempty"`
	NetworkStatus          NetworkStatus          `bson:"network_status,omitempty"`
	ScreenshotInfo         ScreenshotInfo         `bson:"screenshot_info,omitempty"`
	ApplicationUsageReport ApplicationUsageReport `bson:"application_usage_report,omitempty"`
}

type ActivityReport struct {
	Id                         primitive.ObjectID `bson:"_id"`
	Day                        int                `bson:"day"`
	Month                      int                `bson:"month"`
	Year                       int                `bson:"year"`
	EmployeeID                 string             `bson:"employee_id"`
	EmployeeName               string             `bson:"employee_name"`
	IsCurrentlyCheckedIn       bool               `bson:"is_currently_checked_in"`
	FirstCheckInTimeStamp      time.Time          `bson:"first_checkin_timestamp"`
	FirstCheckInSource         string             `bson:"first_checkin_source,omitempty"`
	LastCheckOutTimeStamp      time.Time          `bson:"last_checkout_timestamp"`
	LastCheckOutSource         string             `bson:"last_checkout_source,omitempty"`
	LastKnownLocation          LocationInfo       `bson:"last_known_location,omitempty"`
	FirstActivityTimeStamp     time.Time          `bson:"first_activity_timestamp"`
	FirstActivityApplication   string             `bson:"first_activity_application,omitempty"`
	LastActivityTimeStamp      time.Time          `bson:"last_activity_timestamp"`
	LastActivityApplication    string             `bson:"last_activity_application,omitempty"`
	LastUserStatusInfo         UserActivityInfo   `bson:"last_user_status_info,omitempty"`
	LastScreenshotInfo         ScreenshotInfo     `bson:"last_screenshot_info,omitempty"`
	LastNetworkStatusInfo      NetworkStatus      `bson:"last_network_status_info,omitempty"`
	LastApplicationInfo        ApplicationUsage   `bson:"last_activity_info,omitempty"`
	TotalEvents                int                `bson:"total_events"`
	TotalActiveTime            float64            `bson:"total_active_time_seconds"`
	TotalIdleTime              float64            `bson:"total_idle_time_seconds"`
	MonitoringWindows          []MonitoringWindow `bson:"monitoring_windows,omitempty"`
	MonitoringWindowSizeInMins int                `bson:"monitoring_window_size_in_mins"`
	ApplicationUsages          []ApplicationUsage `bson:"application_usages,omitempty"`
}

type MonitoringWindow struct {
	Start                           time.Time `bson:"start"`
	End                             time.Time `bson:"end"`
	ActiveTime                      float64   `bson:"active_time_seconds"`
	IdleTime                        float64   `bson:"idle_time_seconds"`
	IsCheckinInThisWindow           bool      `bson:"is_checkin_in_this_window"`
	CheckinTimeStamp                time.Time `bson:"checkin_timestamp,omitempty"`
	IsCheckoutInThisWindow          bool      `bson:"is_checkout_in_this_window"`
	CheckoutTimeStamp               time.Time `bson:"checkout_timestamp,omitempty"`
	IsMonitoringStartedInThisWindow bool      `bson:"is_monitoring_started_in_this_window"`
	StartMonitoringTimeStamp        time.Time `bson:"start_monitoring_timestamp,omitempty"`
	IsMonitoringEndedInThisWindow   bool      `bson:"is_monitoring_ended_in_this_window"`
	EndMonitoringTimeStamp          time.Time `bson:"end_monitoring_timestamp,omitempty"`
	City                            string    `bson:"city,omitempty"`
	Country                         string    `bson:"country,omitempty"`
	SSID                            string    `bson:"ssid,omitempty"`
	IPAddress                       string    `bson:"ip_address,omitempty"`
	ScreenshotURL                   string    `bson:"screenshot_url,omitempty"`
}

// ApplicationUsage represents time spent on an application
type ApplicationUsage struct {
	AppName        string         `bson:"app_name"`
	TimeStamp      time.Time      `bson:"timestamp"`
	TimeSpentToday float64        `bson:"time_spent_today_seconds"` // Total time today
	LastActiveTime time.Time      `bson:"last_active_time"`
	SessionCount   int            `bson:"session_count"` // Number of times app was activated
	ActivityInfo   []ActivityInfo `bson:"activity_info,omitempty"`
}

// ActivityInfo represents current activity information
type ActivityInfo struct {
	TimeStamp    time.Time `bson:"timestamp"`
	TimeSpent    float64   `bson:"time_spent_seconds"` // Total time today
	AppName      string    `bson:"app_name,omitempty"`
	WindowTitle  string    `bson:"window_title,omitempty"`
	DocumentName string    `bson:"document_name,omitempty"`
	BrowserURL   string    `bson:"browser_url,omitempty"`
	IsIncognito  bool      `bson:"is_incognito,omitempty"`
	ProcessID    int       `bson:"process_id,omitempty"`
}

// ApplicationUsageReport represents a summary of application usage
type ApplicationUsageReport struct {
	Applications []ApplicationUsage `bson:"applications"`
	TotalTime    float64            `bson:"total_time_seconds"`
	TopApps      []ApplicationUsage `bson:"top_apps,omitempty"` // Top 5 apps by time
}

// SystemStatusInfo represents system status information
type SystemStatusInfo struct {
	Message string `bson:"message"`
	Details string `bson:"details,omitempty"`
}

// USBDeviceInfo represents USB device information
type USBDeviceInfo struct {
	VendorID    int    `bson:"vendor_id"`
	ProductID   int    `bson:"product_id"`
	VendorName  string `bson:"vendor_name"`
	ProductName string `bson:"product_name"`
	Action      string `bson:"action"` // "connected" or "disconnected"
}

// LocationInfo represents device location
type LocationInfo struct {
	City      string    `bson:"city"`
	Country   string    `bson:"country"`
	Latitude  float64   `bson:"latitude"`
	Longitude float64   `bson:"longitude"`
	Accuracy  float64   `bson:"accuracy"`
	Timestamp time.Time `bson:"timestamp"`
	Available bool      `bson:"available"`
}

// UserActivityInfo represents user activity status
type UserActivityInfo struct {
	Status               string    `bson:"status"` // "Active" or "Idle"
	TimeStamp            time.Time `bson:"timestamp,omitempty"`
	TotalIdleTimeToday   float64   `bson:"total_idle_time_today,omitempty"`
	TotalActiveTimeToday float64   `bson:"total_active_time_today,omitempty"`
}

// NetworkStatus represents network connectivity status
type NetworkStatus struct {
	TimeStamp          time.Time `bson:"timestamp,omitempty"`
	IsOnline           bool      `bson:"is_online"`
	SSID               string    `bson:"ssid,omitempty"`
	TotalUpTimeToday   float64   `bson:"total_uptime_today,omitempty"`
	TotalDownTimeToday float64   `bson:"total_downtime_today,omitempty"`
	IPAddress          string    `bson:"ip_address"`
}

// ScreenshotInfo represents screenshot metadata
type ScreenshotInfo struct {
	Filename   string    `bson:"filename"`
	Timestamp  time.Time `bson:"timestamp"`
	MacAddress string    `bson:"mac_address"`
	UserID     string    `bson:"user_id"`
	FilePath   string    `bson:"file_path"`
	URL        string    `bson:"url,omitempty"`
	Domain     string    `bson:"domain,omitempty"`
}

// ApplicationTimeBreakdown represents the percentage of time spent on each application
type ApplicationTimeBreakdown struct {
	AppName           string  `bson:"app_name"`
	TimeSpent         float64 `bson:"time_spent_seconds"`
	PercentageOfTotal float64 `bson:"percentage_of_total"`
}

// EmployeeAnalyticsReport represents a comprehensive analytics report for an employee
type EmployeeAnalyticsReport struct {
	Id                   primitive.ObjectID         `bson:"_id"`
	Day                  int                        `bson:"day"`
	Month                int                        `bson:"month"`
	Year                 int                        `bson:"year"`
	EmployeeID           string                     `bson:"employee_id"`
	EmployeeName         string                     `bson:"employee_name"`
	Domain               string                     `bson:"domain"`
	ReportTimestamp      time.Time                  `bson:"report_timestamp"`
	TotalTime            float64                    `bson:"total_time_seconds"` // Total time from first to last event
	ActiveTime           float64                    `bson:"active_time_seconds"`
	IdleTime             float64                    `bson:"idle_time_seconds"`
	ApplicationBreakdown []ApplicationTimeBreakdown `bson:"application_breakdown"`
	TotalApplicationTime float64                    `bson:"total_application_time_seconds"` // Sum of all app time
}
