# Monitor Client gRPC Usage Examples

This document provides examples for using the Monitor Client gRPC methods. All methods use gRPC as the primary communication protocol with automatic fallback to REST if gRPC is unavailable.

## Table of Contents
- [Initialization](#initialization)
- [Check-In/Check-Out Status](#check-incheck-out-status)
- [Heartbeat](#heartbeat)
- [Fetch Monitoring Configuration](#fetch-monitoring-configuration)
- [Send Activity Log](#send-activity-log)
- [Send Activity Log Batch](#send-activity-log-batch)

## Initialization

Before using any monitor client methods, initialize the client:

```go
import "github.com/adhiman2409/gomicroutils/grpcclient"

// Start the monitor client (typically done once at application startup)
grpcclient.StartMonitorClient()

// Get the client instance
client := grpcclient.GetMonitorClient()

// Don't forget to stop the client when done
defer grpcclient.StopMonitorClient()
```

## Check-In/Check-Out Status

Update employee check-in or check-out status:

```go
req := grpcclient.EmployeeCheckInCheckOutRequest{
    Day:               16,
    Month:             10,
    Year:              2025,
    EmployeeId:        "EMP12345",
    IsCheckedIn:       true,
    CheckInTimestamp:  "2025-10-16T09:00:00Z",
    CheckInSource:     "desktop_app",
    IsCheckedOut:      false,
    CheckOutTimestamp: "",
    CheckOutSource:    "",
    Domain:            "example.com",
}

response, err := client.UpdateCheckInCheckoutStatus(req, "example.com")
if err != nil {
    log.Printf("Failed to update check-in status: %v", err)
    return
}

log.Printf("Status: %v, Message: %s", response.Status, response.Message)
```

## Heartbeat

Send a heartbeat to check employee attendance status:

```go
req := grpcclient.HeartbeatRequest{
    EmployeeId: "EMP12345",
    Domain:     "example.com",
}

response, err := client.Heartbeat(req)
if err != nil {
    log.Printf("Heartbeat failed: %v", err)
    return
}

log.Printf("Employee Status: CheckedIn=%v, OnLeave=%v, Holiday=%v", 
    response.IsCheckedIn, 
    response.IsOnLeave, 
    response.IsHoliday)
```

## Fetch Monitoring Configuration

Retrieve monitoring configuration for a device:

```go
req := grpcclient.MonitoringConfigRequest{
    MacAddress: "00:1B:44:11:3A:B7",
    Domain:     "example.com",
}

config, err := client.FetchMonitoringConfig(req)
if err != nil {
    log.Printf("Failed to fetch config: %v", err)
    return
}

log.Printf("Monitoring Enabled: %v", config.MonitoringEnabled)
log.Printf("Screenshot Interval: %d seconds", config.ScreenshotInterval)
log.Printf("Idle Threshold: %d seconds", config.IdleThreshold)
log.Printf("Track Applications: %v", config.TrackApplications)
```

## Send Activity Log

Send a single activity log entry:

### User Activity Example
```go
req := grpcclient.ActivityLogRequest{
    Timestamp:  time.Now(),
    UserId:     "EMP12345",
    Name:       "John Doe",
    MacAddress: "00:1B:44:11:3A:B7",
    EventType:  "user_activity",
    Hostname:   "johns-macbook",
    IpAddress:  "192.168.1.100",
    SessionId:  "session-123",
    Domain:     "example.com",
    UserActivityInfo: &grpcclient.UserActivityInfo{
        Status:                 "Active",
        Timestamp:              time.Now(),
        ActiveDurationSeconds:  300.0,
        IdleDurationSeconds:    0.0,
        OfflineDurationSeconds: 0.0,
        SessionDurationSeconds: 300.0,
    },
}

response, err := client.SendActivityLog(req)
if err != nil {
    log.Printf("Failed to send activity log: %v", err)
    return
}

log.Printf("Success: %v, Entries Processed: %d", response.Success, response.EntriesProcessed)
```

### Application Activity Example
```go
req := grpcclient.ActivityLogRequest{
    Timestamp:  time.Now(),
    UserId:     "EMP12345",
    Name:       "John Doe",
    MacAddress: "00:1B:44:11:3A:B7",
    EventType:  "application_activity",
    Hostname:   "johns-macbook",
    Domain:     "example.com",
    ActivityInfo: &grpcclient.ActivityInfo{
        Timestamp:    time.Now(),
        AppName:      "Visual Studio Code",
        WindowTitle:  "main.go - gomicroutils",
        ProcessID:    12345,
    },
}

response, err := client.SendActivityLog(req)
// Handle response...
```

### Browser Activity Example
```go
req := grpcclient.ActivityLogRequest{
    Timestamp:  time.Now(),
    UserId:     "EMP12345",
    Name:       "John Doe",
    MacAddress: "00:1B:44:11:3A:B7",
    EventType:  "browser_activity",
    Hostname:   "johns-macbook",
    Domain:     "example.com",
    ActivityInfo: &grpcclient.ActivityInfo{
        Timestamp:    time.Now(),
        AppName:      "Google Chrome",
        WindowTitle:  "GitHub - adhiman2409/gomicroutils",
        BrowserURL:   "https://github.com/adhiman2409/gomicroutils",
        IsIncognito:  false,
    },
}

response, err := client.SendActivityLog(req)
// Handle response...
```

### Network Status Example
```go
req := grpcclient.ActivityLogRequest{
    Timestamp:  time.Now(),
    UserId:     "EMP12345",
    Name:       "John Doe",
    MacAddress: "00:1B:44:11:3A:B7",
    EventType:  "network_status",
    Hostname:   "johns-macbook",
    Domain:     "example.com",
    NetworkStatus: &grpcclient.NetworkStatus{
        Timestamp:              time.Now(),
        SSID:                   "Office-WiFi",
        UptimeSeconds:          60.0,
        DowntimeSeconds:        0.0,
        SessionDurationSeconds: 60.0,
        IpAddress:              "192.168.1.100",
    },
}

response, err := client.SendActivityLog(req)
// Handle response...
```

### Location Update Example
```go
req := grpcclient.ActivityLogRequest{
    Timestamp:  time.Now(),
    UserId:     "EMP12345",
    Name:       "John Doe",
    MacAddress: "00:1B:44:11:3A:B7",
    EventType:  "location_update",
    Hostname:   "johns-macbook",
    Domain:     "example.com",
    LocationInfo: &grpcclient.LocationInfo{
        City:      "San Francisco",
        Country:   "USA",
        Latitude:  37.7749,
        Longitude: -122.4194,
        Accuracy:  10.0,
        Timestamp: time.Now(),
        Available: true,
    },
}

response, err := client.SendActivityLog(req)
// Handle response...
```

## Send Activity Log Batch

Send multiple activity log entries in a single request for better performance:

```go
entries := []grpcclient.ActivityLogRequest{
    {
        Timestamp:  time.Now().Add(-5 * time.Minute),
        UserId:     "EMP12345",
        Name:       "John Doe",
        MacAddress: "00:1B:44:11:3A:B7",
        EventType:  "user_activity",
        Hostname:   "johns-macbook",
        Domain:     "example.com",
        UserActivityInfo: &grpcclient.UserActivityInfo{
            Status:                 "Active",
            Timestamp:              time.Now().Add(-5 * time.Minute),
            ActiveDurationSeconds:  300.0,
        },
    },
    {
        Timestamp:  time.Now().Add(-4 * time.Minute),
        UserId:     "EMP12345",
        Name:       "John Doe",
        MacAddress: "00:1B:44:11:3A:B7",
        EventType:  "application_activity",
        Hostname:   "johns-macbook",
        Domain:     "example.com",
        ActivityInfo: &grpcclient.ActivityInfo{
            Timestamp:    time.Now().Add(-4 * time.Minute),
            AppName:      "Slack",
            WindowTitle:  "Engineering Team",
        },
    },
    {
        Timestamp:  time.Now(),
        UserId:     "EMP12345",
        Name:       "John Doe",
        MacAddress: "00:1B:44:11:3A:B7",
        EventType:  "network_status",
        Hostname:   "johns-macbook",
        Domain:     "example.com",
        NetworkStatus: &grpcclient.NetworkStatus{
            Timestamp:              time.Now(),
            SSID:                   "Office-WiFi",
            UptimeSeconds:          300.0,
            IpAddress:              "192.168.1.100",
        },
    },
}

batchReq := grpcclient.ActivityLogBatchRequest{
    Entries: entries,
}

response, err := client.SendActivityLogBatch(batchReq)
if err != nil {
    log.Printf("Failed to send batch: %v", err)
    return
}

log.Printf("Batch sent successfully. Entries processed: %d", response.EntriesProcessed)
```

## Error Handling

All methods return an error if the gRPC call fails. The methods automatically log errors but you should handle them appropriately in your application:

```go
response, err := client.Heartbeat(req)
if err != nil {
    // gRPC call failed - implement fallback logic or retry
    log.Printf("gRPC call failed: %v", err)
    
    // Option 1: Retry
    time.Sleep(5 * time.Second)
    response, err = client.Heartbeat(req)
    
    // Option 2: Use REST API as fallback (implement separately)
    // response = restClient.Heartbeat(req)
    
    return
}

// Process successful response
log.Printf("Heartbeat successful: %v", response)
```

## Best Practices

1. **Batch Requests**: When sending multiple activity logs, use `SendActivityLogBatch` instead of multiple `SendActivityLog` calls for better performance.

2. **Error Handling**: Always check for errors and implement appropriate retry logic or fallback mechanisms.

3. **Connection Management**: Initialize the client once at application startup and reuse it throughout your application's lifecycle.

4. **Logging**: The client logs important events automatically. Configure your logging framework appropriately.

5. **Timeouts**: The gRPC client uses default timeouts. For long-running operations, consider implementing custom context with timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Use ctx for gRPC calls (requires modifying the methods to accept context)
```

## Notes on Screenshot Upload

As per requirements, screenshot upload continues to use REST API and is not available via gRPC. Use the appropriate REST endpoint for screenshot uploads.
