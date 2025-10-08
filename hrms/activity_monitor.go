package hrms

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// ApplicationUsage represents time spent on an application
type ApplicationUsage struct {
	AppName        string    `bson:"app_name"`
	TimeSpent      float64   `bson:"time_spent_seconds"`       // Total time in current session
	TimeSpentToday float64   `bson:"time_spent_today_seconds"` // Total time today
	LastActiveTime time.Time `bson:"last_active_time"`
	SessionCount   int       `bson:"session_count"` // Number of times app was activated
}

// ApplicationUsageReport represents a summary of application usage
type ApplicationUsageReport struct {
	Applications []ApplicationUsage `bson:"applications"`
	TotalTime    float64            `bson:"total_time_seconds"`
	TopApps      []ApplicationUsage `bson:"top_apps,omitempty"` // Top 5 apps by time
}

// ActivityInfo represents current activity information
type ActivityInfo struct {
	AppName      string `bson:"app_name,omitempty"`
	WindowTitle  string `bson:"window_title,omitempty"`
	DocumentName string `bson:"document_name,omitempty"`
	BrowserURL   string `bson:"browser_url,omitempty"`
	IsIncognito  bool   `bson:"is_incognito,omitempty"`
	ProcessID    int    `bson:"process_id,omitempty"`
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
	Latitude  float64   `bson:"latitude"`
	Longitude float64   `bson:"longitude"`
	Accuracy  float64   `bson:"accuracy"`
	Timestamp time.Time `bson:"timestamp"`
	Available bool      `bson:"available"`
}

// UserActivityInfo represents user activity status
type UserActivityInfo struct {
	Status               string  `bson:"status"` // "Active" or "Idle"
	TotalIdleTimeToday   float64 `bson:"total_idle_time_today,omitempty"`
	TotalActiveTimeToday float64 `bson:"total_active_time_today,omitempty"`
	LastActivity         string  `bson:"last_activity,omitempty"`
}

// NetworkStatus represents network connectivity status
type NetworkStatus struct {
	IsOnline           bool    `bson:"is_online"`
	SSID               string  `bson:"ssid,omitempty"`
	TotalUpTimeToday   float64 `bson:"total_uptime_today,omitempty"`
	TotalDownTimeToday float64 `bson:"total_downtime_today,omitempty"`
	IPAddress          string  `bson:"ip_address"`
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
