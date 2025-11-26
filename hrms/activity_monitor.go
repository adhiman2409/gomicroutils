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
	EventTypeInstantScreenshot      = "instant_screenshot"
	EventTypeInstantVideo           = "instant_video"
)

type MonitoringConfig struct {
	Id                              primitive.ObjectID `bson:"_id"`
	StartMonitoringAfterCheckInOnly bool               `bson:"start_monitoring_after_check_in_only"`
	IsCheckedIn                     bool               `bson:"is_checked_in"`
	IsCheckedOut                    bool               `bson:"is_checked_out"`
	UserId                          string             `bson:"user_id"`
	MacAddress                      string             `bson:"mac_address"`
	StripMacAddress                 string             `bson:"strip_mac_address"`
	Name                            string             `bson:"name"`
	Domain                          string             `bson:"domain"`
	FrontendURL                     string             `bson:"frontend_url"`
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
	VideoInfo              VideoInfo              `bson:"video_info,omitempty"`
	ApplicationUsageReport ApplicationUsageReport `bson:"application_usage_report,omitempty"`
}

type ActivityReport struct {
	Id                           primitive.ObjectID `bson:"_id"`
	Day                          int                `bson:"day"`
	Month                        int                `bson:"month"`
	Year                         int                `bson:"year"`
	EmployeeID                   string             `bson:"employee_id"`
	EmployeeName                 string             `bson:"employee_name"`
	IsCurrentlyCheckedIn         bool               `bson:"is_currently_checked_in"`
	FirstCheckInTimeStamp        time.Time          `bson:"first_checkin_timestamp"`
	FirstCheckInSource           string             `bson:"first_checkin_source,omitempty"`
	TotalCheckInCheckoutSessions int                `bson:"total_checkin_checkout_sessions"`
	LastCheckOutTimeStamp        time.Time          `bson:"last_checkout_timestamp"`
	LastCheckOutSource           string             `bson:"last_checkout_source,omitempty"`
	TotalEvents                  int                `bson:"total_events"`
	TotalActiveTime              float64            `bson:"total_active_time_seconds"`
	TotalIdleTime                float64            `bson:"total_idle_time_seconds"`
	TotalSleepTime               float64            `bson:"total_sleep_time_seconds"`
	TotalOfflineTime             float64            `bson:"total_offline_time_seconds"`
	TotalCheckInTime             float64            `bson:"total_checkin_time_seconds"`
	TotalSessionTime             float64            `bson:"total_session_time_seconds"`
	TotalNetworkUpTime           float64            `bson:"total_network_uptime_seconds"`
	TotalNetworkDownTime         float64            `bson:"total_network_downtime_seconds"`
	MonitoringWindows            []MonitoringWindow `bson:"monitoring_windows,omitempty"`
	MonitoringWindowSizeInMins   int                `bson:"monitoring_window_size_in_mins"`
	ActivityInfo                 []ActivityInfo     `bson:"activity_info,omitempty"`
	ApplicationUsages            []ApplicationUsage `bson:"application_usages,omitempty"`
	CheckInOutHistory            []CheckInOutInfo   `bson:"checkin_checkout_history,omitempty"`
	ScreenshotInfos              []ScreenshotInfo   `bson:"screenshot_infos,omitempty"`
	VideoInfos                   []VideoInfo        `bson:"video_infos,omitempty"`
	LocationInfos                []LocationInfo     `bson:"location_infos,omitempty"`
}

type CheckInOutInfo struct {
	CheckInTimeStamp  time.Time `bson:"checkin_timestamp"`
	CheckInSource     string    `bson:"checkin_source,omitempty"`
	CheckOutTimeStamp time.Time `bson:"checkout_timestamp"`
	CheckOutSource    string    `bson:"checkout_source,omitempty"`
	TotalDuration     float64   `bson:"total_duration_seconds"`
}

type MonitoringWindow struct {
	Start                           time.Time `bson:"start"`
	End                             time.Time `bson:"end"`
	ActiveTime                      float64   `bson:"active_time_seconds"`
	IdleTime                        float64   `bson:"idle_time_seconds"`
	SleepTime                       float64   `bson:"sleep_time_seconds"`
	OfflineTime                     float64   `bson:"offline_time_seconds"`
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
	Status                 string    `bson:"status"` // "Active" or "Idle"
	TimeStamp              time.Time `bson:"timestamp,omitempty"`
	IdleDurationSeconds    float64   `bson:"idle_duration_seconds,omitempty"`
	ActiveDurationSeconds  float64   `bson:"active_duration_seconds,omitempty"`
	OfflineDurationSeconds float64   `bson:"offline_duration_seconds,omitempty"`
	SessionDurationSeconds float64   `bson:"session_duration_seconds,omitempty"`
	SleepDurationSeconds   float64   `bson:"sleep_duration_seconds,omitempty"`
}

// NetworkStatus represents network connectivity status
// UpTimeSeconds: Network uptime in the last minute (not total for the day)
// DownTimeSeconds: Network downtime in the last minute (not total for the day)
// SessionDurationSeconds: Time between consecutive NetworkStatus events
type NetworkStatus struct {
	TimeStamp              time.Time `bson:"timestamp,omitempty"`
	SSID                   string    `bson:"ssid,omitempty"`
	UpTimeSeconds          float64   `bson:"uptime,omitempty"`
	DownTimeSeconds        float64   `bson:"downtime,omitempty"`
	SessionDurationSeconds float64   `bson:"session_duration_seconds,omitempty"`
	IPAddress              string    `bson:"ip_address"`
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

// VideoInfo represents video recording metadata
type VideoInfo struct {
	Filename        string    `bson:"filename"`
	Timestamp       time.Time `bson:"timestamp"`
	MacAddress      string    `bson:"mac_address"`
	UserID          string    `bson:"user_id"`
	FilePath        string    `bson:"file_path"`
	URL             string    `bson:"url,omitempty"`
	Domain          string    `bson:"domain,omitempty"`
	DurationSeconds int32     `bson:"duration_seconds,omitempty"`
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

type UpdateInfo struct {
	Version              string    `json:"version"`
	DownloadURL          string    `json:"download_url"`
	Checksum             string    `json:"checksum"`
	Criticality          string    `json:"criticality"`
	ReleaseNotes         string    `json:"release_notes"`
	MinCompatibleVersion string    `json:"min_compatible_version"`
	SizeBytes            int       `json:"size_bytes"`
	Platform             string    `json:"platform"`
	Architecture         string    `json:"architecture"`
	ReleasedAt           time.Time `json:"released_at"`
	Domain               string    `json:"domain"`
}

type ApplicationUpdateInfo struct {
	Id             primitive.ObjectID `bson:"_id"`
	MacAddress     string             `json:"mac_address"`
	EmployeeID     string             `json:"employee_id"`
	CurrentVersion string             `json:"current_version"`
	TargetVersion  string             `json:"target_version"`
	Status         string             `json:"status"` // checking, downloading, installing, success, failed, rolled_back
	Error          string             `json:"error,omitempty"`
	Progress       int                `json:"progress"` // 0 to 100
	Timestamp      time.Time          `json:"timestamp"`
	Platform       string             `json:"platform"`
	Architecture   string             `json:"architecture"`
	LastHeartbeat  time.Time          `json:"last_heartbeat"`
}
