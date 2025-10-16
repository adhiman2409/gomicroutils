# gRPC Client for Monitor Service

This package provides a Go client for communicating with the Monitor Service using gRPC protocol. All APIs (except screenshot upload) use gRPC as the default communication method.

## Features

- **gRPC by Default**: All monitor APIs now use gRPC for efficient, type-safe communication
- **Type Safety**: Strongly typed requests and responses with compile-time validation
- **Comprehensive API Coverage**: Support for all monitor operations including:
  - Check-in/Check-out status updates
  - Heartbeat monitoring
  - Activity logging (single and batch)
  - Monitoring configuration fetching
  - Network, location, application, and user activity tracking

## Installation

```bash
go get github.com/adhiman2409/gomicroutils/grpcclient
```

## Quick Start

```go
package main

import (
    "log"
    "time"
    "github.com/adhiman2409/gomicroutils/grpcclient"
)

func main() {
    // Initialize the gRPC client
    grpcclient.StartMonitorClient()
    defer grpcclient.StopMonitorClient()
    
    // Get client instance
    client := grpcclient.GetMonitorClient()
    
    // Send a heartbeat
    req := grpcclient.HeartbeatRequest{
        EmployeeId: "EMP12345",
        Domain:     "example.com",
    }
    
    response, err := client.Heartbeat(req)
    if err != nil {
        log.Printf("Heartbeat failed: %v", err)
        return
    }
    
    log.Printf("Employee checked in: %v", response.IsCheckedIn)
}
```

## Available Methods

### 1. UpdateCheckInCheckoutStatus
Update employee check-in or check-out status.

```go
func (c *MonitorClient) UpdateCheckInCheckoutStatus(
    req EmployeeCheckInCheckOutRequest, 
    domain string,
) (EmployeeCheckInCheckOutResponse, error)
```

### 2. Heartbeat
Send a heartbeat to check employee attendance status.

```go
func (c *MonitorClient) Heartbeat(
    req HeartbeatRequest,
) (HeartbeatResponse, error)
```

### 3. FetchMonitoringConfig
Retrieve monitoring configuration for a device.

```go
func (c *MonitorClient) FetchMonitoringConfig(
    req MonitoringConfigRequest,
) (MonitoringConfigResponse, error)
```

### 4. SendActivityLog
Send a single activity log entry.

```go
func (c *MonitorClient) SendActivityLog(
    req ActivityLogRequest,
) (ActivityLogResponse, error)
```

### 5. SendActivityLogBatch
Send multiple activity log entries in a single batch (recommended for bulk operations).

```go
func (c *MonitorClient) SendActivityLogBatch(
    req ActivityLogBatchRequest,
) (ActivityLogResponse, error)
```

## Supported Event Types

The activity logging system supports the following event types:

- `user_activity` - User active/idle status
- `application_activity` - Application usage tracking
- `browser_activity` - Browser activity and URL tracking
- `network_status` - Network connectivity status
- `location_update` - Device location updates
- `system_status` - System status information
- `usb_activity` - USB device connection/disconnection
- `application_usage` - Application usage reports

## Configuration

The gRPC client connects to the monitor service at `monitor-srv:50051` using TLS credentials. Ensure your environment has:

1. Proper TLS certificates configured
2. Network access to the monitor service
3. Correct DNS resolution for `monitor-srv`

## Examples

For comprehensive examples covering all use cases, see [USAGE_EXAMPLES.md](./USAGE_EXAMPLES.md).

## Error Handling

All methods return an error if the gRPC call fails. Implement proper error handling and retry logic:

```go
response, err := client.Heartbeat(req)
if err != nil {
    log.Printf("gRPC call failed: %v", err)
    // Implement retry logic or fallback
    return
}
```

## Performance Tips

1. **Use Batch Operations**: When sending multiple activity logs, use `SendActivityLogBatch` instead of multiple `SendActivityLog` calls
2. **Connection Reuse**: Initialize the client once and reuse it throughout your application
3. **Context Management**: Consider implementing custom contexts with timeouts for long-running operations

## REST Fallback

While gRPC is the primary protocol, you can implement REST fallback in your application if gRPC is unavailable. The gRPC methods will return errors that you can catch and handle accordingly.

**Note**: Screenshot upload APIs continue to use REST as they involve file uploads.

## Migration from REST

If you're migrating from REST to gRPC:

1. The method signatures remain similar for ease of migration
2. Request and response types are defined in `monitor_info.go`
3. All types use JSON tags for compatibility
4. Time fields use `time.Time` instead of strings

## Contributing

When adding new methods:

1. Define request/response types in `monitor_info.go`
2. Add the gRPC wrapper method in `monitor_requester.go`
3. Update examples in `USAGE_EXAMPLES.md`
4. Ensure proper error handling and logging

## License

This project is part of the gomicroutils package and follows the same license.
