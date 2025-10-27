package grpcclient

import "time"

// EmployeeCheckInCheckOutRequest represents check-in/check-out status update
type EmployeeCheckInCheckOutRequest struct {
	Day               int32  `json:"day"`
	Month             int32  `json:"month"`
	Year              int32  `json:"year"`
	EmployeeId        string `json:"employee_id"`
	IsCheckedIn       bool   `json:"is_checked_in"`
	CheckInTimestamp  string `json:"check_in_timestamp"`
	CheckInSource     string `json:"check_in_source"`
	IsCheckedOut      bool   `json:"is_checked_out"`
	CheckOutTimestamp string `json:"check_out_timestamp"`
	CheckOutSource    string `json:"check_out_source"`
	Domain            string `json:"domain"`
}

type EmployeeCheckInCheckOutResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// HeartbeatRequest represents a heartbeat request from the client
type HeartbeatRequest struct {
	EmployeeId string `json:"employee_id"`
	Domain     string `json:"domain"`
}

// HeartbeatResponse represents the server's response to a heartbeat
type HeartbeatResponse struct {
	Day                    int32  `json:"day"`
	Month                  int32  `json:"month"`
	Year                   int32  `json:"year"`
	IsCheckedIn            bool   `json:"is_checked_in"`
	CheckinTimestamp       string `json:"checkin_timestamp"`
	CheckinSource          string `json:"checkin_source"`
	IsCheckedOut           bool   `json:"is_checked_out"`
	CheckoutTimestamp      string `json:"checkout_timestamp"`
	CheckoutSource         string `json:"checkout_source"`
	IsOnLeave              bool   `json:"is_on_leave"`
	IsHoliday              bool   `json:"is_holiday"`
	IsWeeklyOffDay         bool   `json:"is_weekly_off_day"`
	IsAttendanceFetchError bool   `json:"is_attendance_fetch_error"`
	Message                string `json:"message"`
}

// MonitoringConfigRequest requests configuration for a device
type MonitoringConfigRequest struct {
	MacAddress string `json:"mac_address"`
	Domain     string `json:"domain"`
}

// MonitoringConfigResponse contains the monitoring configuration
type MonitoringConfigResponse struct {
	StartMonitoringAfterCheckInOnly bool   `json:"start_monitoring_after_check_in_only"`
	IsCheckedIn                     bool   `json:"is_checked_in"`
	IsCheckedOut                    bool   `json:"is_checked_out"`
	UserId                          string `json:"user_id"`
	MacAddress                      string `json:"mac_address"`
	Name                            string `json:"name"`
	Domain                          string `json:"domain"`
	FrontendURL                     string `json:"frontend_url"`
	Department                      string `json:"department"`
	Email                           string `json:"email"`
	MonitoringEnabled               bool   `json:"monitoring_enabled"`
	IdleThreshold                   int32  `json:"idle_threshold"`
	ScreenshotInterval              int32  `json:"screenshot_interval"`
	CheckInterval                   int32  `json:"check_interval"`
	LocationInterval                int32  `json:"location_interval"`
	TrackApplications               bool   `json:"track_applications"`
	TrackLocation                   bool   `json:"track_location"`
	TrackBrowser                    bool   `json:"track_browser"`
	TrackDocuments                  bool   `json:"track_documents"`
	TrackUSB                        bool   `json:"track_usb"`
	ScreenshotPath                  string `json:"screenshot_path"`
}

// ActivityLogRequest represents a single activity log entry
type ActivityLogRequest struct {
	Timestamp              time.Time               `json:"timestamp"`
	UserId                 string                  `json:"user_id"`
	Name                   string                  `json:"name"`
	MacAddress             string                  `json:"mac_address"`
	Message                string                  `json:"message"`
	EventType              string                  `json:"event_type"`
	Hostname               string                  `json:"hostname"`
	IpAddress              string                  `json:"ip_address"`
	SessionId              string                  `json:"session_id"`
	Domain                 string                  `json:"domain"`
	ActivityInfo           *ActivityInfo           `json:"activity_info,omitempty"`
	UserActivityInfo       *UserActivityInfo       `json:"user_activity_info,omitempty"`
	SystemStatusInfo       *SystemStatusInfo       `json:"system_status_info,omitempty"`
	USBDeviceInfo          *USBDeviceInfo          `json:"usb_device_info,omitempty"`
	LocationInfo           *LocationInfo           `json:"location_info,omitempty"`
	NetworkStatus          *NetworkStatus          `json:"network_status,omitempty"`
	ScreenshotInfo         *ScreenshotInfo         `json:"screenshot_info,omitempty"`
	VideoInfo              *VideoInfo              `json:"video_info,omitempty"`
	ApplicationUsageReport *ApplicationUsageReport `json:"application_usage_report,omitempty"`
}

// ActivityLogBatchRequest contains multiple activity log entries
type ActivityLogBatchRequest struct {
	Entries []ActivityLogRequest `json:"entries"`
}

// ActivityLogResponse confirms log receipt
type ActivityLogResponse struct {
	Success          bool   `json:"success"`
	Message          string `json:"message"`
	EntriesProcessed int32  `json:"entries_processed"`
}

// ActivityInfo represents current activity information
type ActivityInfo struct {
	Timestamp    time.Time `json:"timestamp"`
	AppName      string    `json:"app_name"`
	WindowTitle  string    `json:"window_title"`
	DocumentName string    `json:"document_name"`
	BrowserURL   string    `json:"browser_url"`
	IsIncognito  bool      `json:"is_incognito"`
	ProcessID    int32     `json:"process_id"`
}

// UserActivityInfo represents user activity status
type UserActivityInfo struct {
	Status                 string    `json:"status"`
	Timestamp              time.Time `json:"timestamp"`
	IdleDurationSeconds    float64   `json:"idle_duration_seconds"`
	ActiveDurationSeconds  float64   `json:"active_duration_seconds"`
	OfflineDurationSeconds float64   `json:"offline_duration_seconds"`
	SessionDurationSeconds float64   `json:"session_duration_seconds"`
	SleepDurationSeconds   float64   `json:"sleep_duration_seconds"`
}

// SystemStatusInfo represents system status information
type SystemStatusInfo struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// USBDeviceInfo represents USB device information
type USBDeviceInfo struct {
	VendorID    int32  `json:"vendor_id"`
	ProductID   int32  `json:"product_id"`
	VendorName  string `json:"vendor_name"`
	ProductName string `json:"product_name"`
	Action      string `json:"action"`
}

// LocationInfo represents device location
type LocationInfo struct {
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Accuracy  float64   `json:"accuracy"`
	Timestamp time.Time `json:"timestamp"`
	Available bool      `json:"available"`
}

// NetworkStatus represents network connectivity status
type NetworkStatus struct {
	Timestamp              time.Time `json:"timestamp"`
	SSID                   string    `json:"ssid"`
	UptimeSeconds          float64   `json:"uptime_seconds"`
	DowntimeSeconds        float64   `json:"downtime_seconds"`
	SessionDurationSeconds float64   `json:"session_duration_seconds"`
	IpAddress              string    `json:"ip_address"`
}

// ScreenshotInfo represents screenshot metadata
type ScreenshotInfo struct {
	Filename   string    `json:"filename"`
	Timestamp  time.Time `json:"timestamp"`
	MacAddress string    `json:"mac_address"`
	UserId     string    `json:"user_id"`
	FilePath   string    `json:"file_path"`
	URL        string    `json:"url"`
	Domain     string    `json:"domain"`
	Department string    `json:"department"`
}

// VideoInfo represents video metadata
type VideoInfo struct {
	Filename        string    `json:"filename"`
	Timestamp       time.Time `json:"timestamp"`
	MacAddress      string    `json:"mac_address"`
	UserId          string    `json:"user_id"`
	FilePath        string    `json:"file_path"`
	URL             string    `json:"url"`
	Domain          string    `json:"domain"`
	Department      string    `json:"department"`
	DurationSeconds int32     `json:"duration_seconds"`
}

// ApplicationUsage represents time spent on an application
type ApplicationUsage struct {
	AppName               string    `json:"app_name"`
	Timestamp             float64   `json:"timestamp"`
	TimeSpentTodaySeconds float64   `json:"time_spent_today_seconds"`
	LastActiveTime        time.Time `json:"last_active_time"`
	SessionCount          int32     `json:"session_count"`
}

// ApplicationUsageReport represents a summary of application usage
type ApplicationUsageReport struct {
	Applications     []ApplicationUsage `json:"applications"`
	TotalTimeSeconds float64            `json:"total_time_seconds"`
	TopApps          []ApplicationUsage `json:"top_apps"`
}
