# Bidirectional Streaming Implementation Guide

## Architecture Overview

```
┌──────────────────┐                    ┌──────────────────┐
│   Client Side    │                    │   Server Side    │
│                  │                    │                  │
│  ┌────────────┐  │                    │  ┌────────────┐  │
│  │ Stream     │  │◄──────────────────►│  │ Stream     │  │
│  │ Handler    │  │   Bidirectional    │  │ Manager    │  │
│  └────────────┘  │                    │  └────────────┘  │
│        │         │                    │        │         │
│        ├─────────┼────────────────────┼────────┤         │
│        │         │                    │        │         │
│  ┌─────▼──────┐  │                    │  ┌─────▼──────┐  │
│  │ Screenshot │  │                    │  │ Attendance │  │
│  │ Handler    │  │                    │  │ Publisher  │  │
│  └────────────┘  │                    │  └────────────┘  │
│                  │                    │        │         │
│  ┌────────────┐  │                    │  ┌─────▼──────┐  │
│  │ Config     │  │                    │  │ Command    │  │
│  │ Updater    │  │                    │  │ Queue      │  │
│  └────────────┘  │                    │  └────────────┘  │
└──────────────────┘                    └──────────────────┘
```

## Features Implemented

### 1. Real-time Attendance Updates (Server → Client Push)
- Instant notification when user checks in/out
- No more polling with heartbeat
- Reduces heartbeat frequency or eliminates it

### 2. On-Demand Screenshot (Server → Client Command)
- Admin can request screenshot anytime
- Client captures and uploads immediately
- Returns status back to server

### 3. Dynamic Configuration Updates
- Push config changes without restart
- Update monitoring settings on-the-fly

### 4. Remote Commands
- Start/Stop monitoring
- Flush logs immediately
- Restart client remotely

## Benefits vs Current Heartbeat Approach

| Feature | Current (Heartbeat Polling) | New (Bidirectional Stream) |
|---------|----------------------------|----------------------------|
| Latency | 30s average | <1s (instant) |
| Network calls | Every 30s | Only on change |
| Server load | High (constant polling) | Low (push on demand) |
| Bandwidth | Higher | 70-90% reduction |
| Scalability | Limited | High |
| Real-time | ❌ No | ✅ Yes |

## Implementation

### Server-Side Implementation

#### 1. Stream Manager (Connection Pool)

```go
// File: /Users/adhiman/GitHub/activity.monitor.unirms/service/stream_manager.go
package service

import (
    "fmt"
    "sync"
    "github.com/adhiman2409/gomicroutils/genproto/monitor"
)

type ClientStream struct {
    UserID     string
    MacAddress string
    Domain     string
    Stream     monitor.MonitorService_MonitorStreamServer
    Connected  bool
    mu         sync.Mutex
}

type StreamManager struct {
    clients map[string]*ClientStream // key: userID or macAddress
    mu      sync.RWMutex
}

func NewStreamManager() *StreamManager {
    return &StreamManager{
        clients: make(map[string]*ClientStream),
    }
}

// Register a new client connection
func (sm *StreamManager) Register(userID, macAddress, domain string, stream monitor.MonitorService_MonitorStreamServer) {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    key := fmt.Sprintf("%s_%s", userID, macAddress)
    sm.clients[key] = &ClientStream{
        UserID:     userID,
        MacAddress: macAddress,
        Domain:     domain,
        Stream:     stream,
        Connected:  true,
    }
    fmt.Printf("[STREAM] Client registered: %s (MAC: %s)\n", userID, macAddress)
}

// Unregister a client
func (sm *StreamManager) Unregister(userID, macAddress string) {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    key := fmt.Sprintf("%s_%s", userID, macAddress)
    delete(sm.clients, key)
    fmt.Printf("[STREAM] Client unregistered: %s\n", key)
}

// Send attendance update to specific client
func (sm *StreamManager) SendAttendanceUpdate(userID string, update *monitor.AttendanceUpdate) error {
    sm.mu.RLock()
    defer sm.mu.RUnlock()

    // Find client by userID
    for key, client := range sm.clients {
        if client.UserID == userID && client.Connected {
            client.mu.Lock()
            err := client.Stream.Send(&monitor.ServerMessage{
                MessageId: fmt.Sprintf("att_%d", time.Now().Unix()),
                Timestamp: timestamppb.Now(),
                Message:   &monitor.ServerMessage_AttendanceUpdate{AttendanceUpdate: update},
            })
            client.mu.Unlock()

            if err != nil {
                fmt.Printf("[STREAM] Failed to send attendance update to %s: %v\n", key, err)
                client.Connected = false
                return err
            }
            fmt.Printf("[STREAM] Sent attendance update to %s\n", key)
            return nil
        }
    }
    return fmt.Errorf("client not found or disconnected: %s", userID)
}

// Request screenshot from specific client
func (sm *StreamManager) RequestScreenshot(userID, macAddress, requestID, reason string) error {
    sm.mu.RLock()
    defer sm.mu.RUnlock()

    key := fmt.Sprintf("%s_%s", userID, macAddress)
    client, exists := sm.clients[key]
    if !exists || !client.Connected {
        return fmt.Errorf("client not found or disconnected: %s", key)
    }

    client.mu.Lock()
    err := client.Stream.Send(&monitor.ServerMessage{
        MessageId: requestID,
        Timestamp: timestamppb.Now(),
        Message: &monitor.ServerMessage_ScreenshotCommand{
            ScreenshotCommand: &monitor.ScreenshotCommand{
                RequestId: requestID,
                Immediate: true,
                Reason:    reason,
            },
        },
    })
    client.mu.Unlock()

    if err != nil {
        fmt.Printf("[STREAM] Failed to request screenshot from %s: %v\n", key, err)
        client.Connected = false
        return err
    }

    fmt.Printf("[STREAM] Screenshot request sent to %s (RequestID: %s)\n", key, requestID)
    return nil
}

// Broadcast message to all connected clients
func (sm *StreamManager) BroadcastConfigUpdate(config *monitor.MonitoringConfigResponse) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()

    for key, client := range sm.clients {
        if client.Connected {
            client.mu.Lock()
            err := client.Stream.Send(&monitor.ServerMessage{
                MessageId: fmt.Sprintf("cfg_%d", time.Now().Unix()),
                Timestamp: timestamppb.Now(),
                Message: &monitor.ServerMessage_ConfigUpdate{
                    ConfigUpdate: &monitor.ConfigUpdate{
                        Config: config,
                        Reason: "Configuration updated by admin",
                    },
                },
            })
            client.mu.Unlock()

            if err != nil {
                fmt.Printf("[STREAM] Failed to send config update to %s: %v\n", key, err)
                client.Connected = false
            }
        }
    }
}

// Get list of connected clients
func (sm *StreamManager) GetConnectedClients() []string {
    sm.mu.RLock()
    defer sm.mu.RUnlock()

    clients := make([]string, 0, len(sm.clients))
    for key, client := range sm.clients {
        if client.Connected {
            clients = append(clients, key)
        }
    }
    return clients
}
```

#### 2. gRPC Stream Handler

```go
// File: /Users/adhiman/GitHub/activity.monitor.unirms/adapters/grpc/stream_handler.go
package grpc

import (
    "fmt"
    "io"
    "github.com/adhiman2409/gomicroutils/genproto/monitor"
)

func (a *GrpcAdapter) MonitorStream(stream monitor.MonitorService_MonitorStreamServer) error {
    var userID, macAddress, domain string
    var registered bool

    // Cleanup on disconnect
    defer func() {
        if registered {
            a.streamManager.Unregister(userID, macAddress)
        }
    }()

    fmt.Println("[STREAM] New client connection")

    for {
        // Receive message from client
        msg, err := stream.Recv()
        if err == io.EOF {
            fmt.Printf("[STREAM] Client disconnected: %s\n", userID)
            return nil
        }
        if err != nil {
            fmt.Printf("[STREAM] Error receiving message: %v\n", err)
            return err
        }

        // Handle different message types
        switch m := msg.Message.(type) {
        case *monitor.ClientMessage_Connect:
            // Client initial connection
            userID = m.Connect.UserId
            macAddress = m.Connect.MacAddress
            domain = m.Connect.Domain

            fmt.Printf("[STREAM] Client connecting: UserID=%s, MAC=%s, Domain=%s, Version=%s\n",
                userID, macAddress, domain, m.Connect.ClientVersion)

            // Register client in stream manager
            a.streamManager.Register(userID, macAddress, domain, stream)
            registered = true

            // Send initial config
            config, err := a.monitorService.GetMonitorConfigByMac(macAddress, domain)
            if err == nil {
                configProto := convertConfigToProto(config)
                stream.Send(&monitor.ServerMessage{
                    MessageId: "init_config",
                    Timestamp: timestamppb.Now(),
                    Message: &monitor.ServerMessage_ConfigUpdate{
                        ConfigUpdate: &monitor.ConfigUpdate{
                            Config: configProto,
                            Reason: "Initial configuration",
                        },
                    },
                })
            }

        case *monitor.ClientMessage_Heartbeat:
            // Lightweight heartbeat (optional, stream connection itself indicates alive)
            fmt.Printf("[STREAM] Heartbeat from %s\n", userID)

        case *monitor.ClientMessage_ScreenshotResponse:
            // Screenshot response from client
            fmt.Printf("[STREAM] Screenshot response from %s: RequestID=%s, Success=%v\n",
                userID, m.ScreenshotResponse.RequestId, m.ScreenshotResponse.Success)

            // Store screenshot info or notify admin
            a.monitorService.HandleScreenshotResponse(userID, m.ScreenshotResponse)

        case *monitor.ClientMessage_Ack:
            // Client acknowledgment
            fmt.Printf("[STREAM] Ack from %s: MessageID=%s\n", userID, m.Ack.MessageId)

        default:
            fmt.Printf("[STREAM] Unknown message type from %s\n", userID)
        }
    }
}
```

#### 3. Integrate with Attendance Updates

```go
// File: /Users/adhiman/GitHub/activity.monitor.unirms/service/activity_service.go
// Add this method to your service

func (s *DefaultActivityMonitorService) NotifyAttendanceChange(userID string, isCheckedIn, isCheckedOut bool, timestamp time.Time, source string) {
    update := &monitor.AttendanceUpdate{
        UserId:            userID,
        IsCheckedIn:       isCheckedIn,
        IsCheckedOut:      isCheckedOut,
        CheckInTimestamp:  timestamppb.New(timestamp),
        CheckInSource:     source,
        Message:           fmt.Sprintf("Attendance updated: CheckedIn=%v, CheckedOut=%v", isCheckedIn, isCheckedOut),
    }

    // Send to stream manager
    err := s.streamManager.SendAttendanceUpdate(userID, update)
    if err != nil {
        fmt.Printf("[SERVICE] Failed to send attendance update: %v\n", err)
    }
}

// Call this whenever attendance changes
// In UpdateCheckInCheckoutStatus or similar methods:
func (s *DefaultActivityMonitorService) UpdateCheckInCheckoutStatus(req grpcclient.EmployeeCheckInCheckOutRequest) (string, *errs.RestError) {
    // ... existing code ...

    // After successful update, notify via stream
    go s.NotifyAttendanceChange(
        req.EmployeeId,
        req.IsCheckedIn,
        req.IsCheckedOut,
        time.Now(),
        req.CheckInSource,
    )

    return "Status updated", nil
}
```

### Client-Side Implementation

#### 1. Stream Client Handler

```go
// File: /Users/adhiman/GitHub/activity.monitor.client.mac.go/internal/network/stream_client.go
package network

import (
    "context"
    "fmt"
    "io"
    "log"
    "time"

    "github.com/adhiman2409/gomicroutils/genproto/monitor"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type StreamClient struct {
    grpcClient      *GRPCClient
    stream          monitor.MonitorService_MonitorStreamClient
    userID          string
    macAddress      string
    domain          string
    screenshotChan  chan *monitor.ScreenshotCommand
    configChan      chan *monitor.MonitoringConfigResponse
    attendanceChan  chan *monitor.AttendanceUpdate
    commandChan     chan *monitor.MonitoringCommand
    connected       bool
    cancel          context.CancelFunc
}

func NewStreamClient(grpcClient *GRPCClient, userID, macAddress, domain string) *StreamClient {
    return &StreamClient{
        grpcClient:     grpcClient,
        userID:         userID,
        macAddress:     macAddress,
        domain:         domain,
        screenshotChan: make(chan *monitor.ScreenshotCommand, 10),
        configChan:     make(chan *monitor.MonitoringConfigResponse, 10),
        attendanceChan: make(chan *monitor.AttendanceUpdate, 10),
        commandChan:    make(chan *monitor.MonitoringCommand, 10),
    }
}

// Connect establishes bidirectional stream
func (sc *StreamClient) Connect() error {
    ctx, cancel := context.WithCancel(context.Background())
    sc.cancel = cancel

    stream, err := sc.grpcClient.client.MonitorStream(ctx)
    if err != nil {
        return fmt.Errorf("failed to create stream: %w", err)
    }
    sc.stream = stream
    sc.connected = true

    // Send initial connect message
    err = stream.Send(&monitor.ClientMessage{
        Message: &monitor.ClientMessage_Connect{
            Connect: &monitor.ClientConnect{
                UserId:        sc.userID,
                MacAddress:    sc.macAddress,
                Domain:        sc.domain,
                ClientVersion: "2.0.0",
            },
        },
    })
    if err != nil {
        return fmt.Errorf("failed to send connect message: %w", err)
    }

    log.Printf("[STREAM] Connected to server: UserID=%s, MAC=%s\n", sc.userID, sc.macAddress)

    // Start receiving messages
    go sc.receiveMessages()

    // Start heartbeat (lightweight, every 30s)
    go sc.sendHeartbeat()

    return nil
}

// Receive messages from server
func (sc *StreamClient) receiveMessages() {
    for sc.connected {
        msg, err := sc.stream.Recv()
        if err == io.EOF {
            log.Println("[STREAM] Server closed connection")
            sc.connected = false
            return
        }
        if err != nil {
            log.Printf("[STREAM] Error receiving message: %v\n", err)
            sc.connected = false
            return
        }

        // Send acknowledgment
        sc.sendAck(msg.MessageId, true)

        // Handle message
        switch m := msg.Message.(type) {
        case *monitor.ServerMessage_AttendanceUpdate:
            log.Printf("[STREAM] 🔔 Attendance Update: CheckedIn=%v, CheckedOut=%v\n",
                m.AttendanceUpdate.IsCheckedIn, m.AttendanceUpdate.IsCheckedOut)
            sc.attendanceChan <- m.AttendanceUpdate

        case *monitor.ServerMessage_ScreenshotCommand:
            log.Printf("[STREAM] 📸 Screenshot requested: RequestID=%s, Reason=%s\n",
                m.ScreenshotCommand.RequestId, m.ScreenshotCommand.Reason)
            sc.screenshotChan <- m.ScreenshotCommand

        case *monitor.ServerMessage_ConfigUpdate:
            log.Printf("[STREAM] ⚙️ Config update received: %s\n", m.ConfigUpdate.Reason)
            sc.configChan <- m.ConfigUpdate.Config

        case *monitor.ServerMessage_MonitoringCommand:
            log.Printf("[STREAM] 🎮 Command received: %v\n", m.MonitoringCommand.Command)
            sc.commandChan <- m.MonitoringCommand
        }
    }
}

// Send heartbeat periodically
func (sc *StreamClient) sendHeartbeat() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for sc.connected {
        select {
        case <-ticker.C:
            err := sc.stream.Send(&monitor.ClientMessage{
                Message: &monitor.ClientMessage_Heartbeat{
                    Heartbeat: &monitor.ClientHeartbeat{
                        Timestamp: timestamppb.Now(),
                    },
                },
            })
            if err != nil {
                log.Printf("[STREAM] Failed to send heartbeat: %v\n", err)
                sc.connected = false
                return
            }
        }
    }
}

// Send screenshot response
func (sc *StreamClient) SendScreenshotResponse(requestID, path string, success bool, errorMsg string) error {
    return sc.stream.Send(&monitor.ClientMessage{
        Message: &monitor.ClientMessage_ScreenshotResponse{
            ScreenshotResponse: &monitor.ScreenshotResponse{
                RequestId:      requestID,
                Success:        success,
                ScreenshotPath: path,
                ErrorMessage:   errorMsg,
            },
        },
    })
}

// Send acknowledgment
func (sc *StreamClient) sendAck(messageID string, success bool) error {
    return sc.stream.Send(&monitor.ClientMessage{
        Message: &monitor.ClientMessage_Ack{
            Ack: &monitor.ClientAck{
                MessageId: messageID,
                Success:   success,
            },
        },
    })
}

// Get channels for event handling
func (sc *StreamClient) ScreenshotChannel() <-chan *monitor.ScreenshotCommand {
    return sc.screenshotChan
}

func (sc *StreamClient) AttendanceChannel() <-chan *monitor.AttendanceUpdate {
    return sc.attendanceChan
}

func (sc *StreamClient) ConfigChannel() <-chan *monitor.MonitoringConfigResponse {
    return sc.configChan
}

func (sc *StreamClient) CommandChannel() <-chan *monitor.MonitoringCommand {
    return sc.commandChan
}

// Close stream
func (sc *StreamClient) Close() {
    sc.connected = false
    if sc.cancel != nil {
        sc.cancel()
    }
    log.Println("[STREAM] Connection closed")
}
```

#### 2. Integrate Stream into Main Client

```go
// File: /Users/adhiman/GitHub/activity.monitor.client.mac.go/main.go
// Add to main function

func main() {
    // ... existing setup ...

    // Initialize stream client after gRPC is initialized
    var streamClient *network.StreamClient
    if networkClient.IsUsingGRPC() && cfgData.UserID != "" {
        streamClient = network.NewStreamClient(
            networkClient.GetGRPCClient(),
            cfgData.UserID,
            macAddress,
            cfgData.Domain,
        )

        err := streamClient.Connect()
        if err != nil {
            logger.Warnf("Failed to connect stream: %v", err)
        } else {
            logger.Info("Stream connected successfully")
            defer streamClient.Close()

            // Start stream event handlers
            go handleStreamEvents(streamClient, monitor, logger, networkClient)
        }
    }

    // ... rest of main ...
}

// Handle stream events
func handleStreamEvents(stream *network.StreamClient, monitor *monitoring.Monitor, logger *logging.Logger, client *network.Client) {
    for {
        select {
        case update := <-stream.AttendanceChannel():
            logger.Infof("🔔 Real-time attendance update: CheckedIn=%v, CheckedOut=%v",
                update.IsCheckedIn, update.IsCheckedOut)

            // Update local config
            cfg.UpdateCheckInStatus(update.IsCheckedIn, update.IsCheckedOut)

            // Flush logs if checked out
            if update.IsCheckedOut {
                logger.Info("User checked out - flushing logs")
                client.FlushBuffer()
            }

        case cmd := <-stream.ScreenshotChannel():
            logger.Infof("📸 Screenshot requested: %s (Reason: %s)",
                cmd.RequestId, cmd.Reason)

            // Take screenshot immediately
            path, err := monitor.TakeScreenshotOnDemand()
            if err != nil {
                logger.Errorf("Failed to take screenshot: %v", err)
                stream.SendScreenshotResponse(cmd.RequestId, "", false, err.Error())
            } else {
                logger.Infof("Screenshot taken: %s", path)
                stream.SendScreenshotResponse(cmd.RequestId, path, true, "")
            }

        case config := <-stream.ConfigChannel():
            logger.Info("⚙️ Configuration updated via stream")
            // Apply new config
            updateConfigFromProto(cfg, config)

        case cmd := <-stream.CommandChannel():
            logger.Infof("🎮 Received command: %v", cmd.Command)
            handleCommand(cmd, monitor, logger)
        }
    }
}

func handleCommand(cmd *monitor.MonitoringCommand, monitor *monitoring.Monitor, logger *logging.Logger) {
    switch cmd.Command {
    case monitor.MonitoringCommand_STOP_MONITORING:
        logger.Info("Stopping monitoring via command")
        monitor.Stop()
    case monitor.MonitoringCommand_FLUSH_LOGS:
        logger.Info("Flushing logs via command")
        client.FlushBuffer()
    // ... handle other commands
    }
}
```

## API Endpoints for Admin Dashboard

Add REST endpoints to trigger commands:

```go
// File: /Users/adhiman/GitHub/activity.monitor.unirms/adapters/rest/stream_endpoints.go

// POST /api/admin/screenshot/request
func (h *RestHandler) RequestScreenshot(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID     string `json:"user_id"`
        MacAddress string `json:"mac_address"`
        Reason     string `json:"reason"`
    }

    json.NewDecoder(r.Body).Decode(&req)

    requestID := uuid.New().String()
    err := h.streamManager.RequestScreenshot(req.UserID, req.MacAddress, requestID, req.Reason)

    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "request_id": requestID,
        "status": "requested",
    })
}

// GET /api/admin/clients/connected
func (h *RestHandler) GetConnectedClients(w http.ResponseWriter, r *http.Request) {
    clients := h.streamManager.GetConnectedClients()
    json.NewEncoder(w).Encode(map[string]interface{}{
        "count": len(clients),
        "clients": clients,
    })
}
```

## Migration Strategy

### Phase 1: Add Stream Support (Parallel with Heartbeat)
- ✅ Implement streaming proto
- ✅ Add server stream handler
- ✅ Add client stream handler
- Keep existing heartbeat working
- Test with subset of clients

### Phase 2: Migrate Attendance to Stream
- Push attendance updates via stream
- Keep heartbeat for backwards compatibility
- Monitor for 1-2 weeks

### Phase 3: Reduce Heartbeat Frequency
- Increase heartbeat from 30s to 5 minutes (just for health check)
- Stream handles all real-time updates

### Phase 4: Optional - Remove Heartbeat
- If stream is stable, remove heartbeat entirely
- Stream connection indicates client health

## Advantages Summary

### 1. Real-time Attendance Updates
**Before**: Client polls every 30s → 28.8KB/day/client
**After**: Server pushes on change → ~0.5KB/day/client
**Savings**: ~95% bandwidth reduction

### 2. On-Demand Screenshots
**Before**: Not possible, wait for next scheduled screenshot (5-15 min)
**After**: Instant screenshot (<2 seconds)

### 3. Scalability
**Before**: 1000 clients = 1000 heartbeat requests every 30s = 33 req/sec
**After**: 1000 clients = 1000 persistent connections, ~0 req/sec (push only)

### 4. Latency
**Before**: 0-30s delay for attendance updates
**After**: <1s instant push notification

## Testing

### Test Attendance Push
```bash
# Trigger check-in via web/API
curl -X POST http://localhost:8080/api/attendance/checkin \
  -d '{"employee_id":"PT24002","domain":"unirms.com"}'

# Client should immediately receive push (no waiting for heartbeat)
```

### Test Screenshot Request
```bash
# Request screenshot from admin dashboard
curl -X POST http://localhost:8080/api/admin/screenshot/request \
  -d '{"user_id":"PT24002","mac_address":"1697e92ca2fc","reason":"security_check"}'

# Client captures screenshot immediately and uploads
```

## Monitoring & Debugging

### Server Logs
```
[STREAM] Client registered: PT24002 (MAC: 1697e92ca2fc)
[STREAM] Sent attendance update to PT24002_1697e92ca2fc
[STREAM] Screenshot request sent to PT24002_1697e92ca2fc (RequestID: abc123)
[STREAM] Screenshot response from PT24002: RequestID=abc123, Success=true
```

### Client Logs
```
[STREAM] Connected to server: UserID=PT24002, MAC=1697e92ca2fc
[STREAM] 🔔 Attendance Update: CheckedIn=true, CheckedOut=false
[STREAM] 📸 Screenshot requested: RequestID=abc123, Reason=security_check
[STREAM] Screenshot taken: /path/to/screenshot.png
```

## Next Steps

1. Generate protobuf code: `protoc --go_out=. --go-grpc_out=. proto/monitor/monitor.proto`
2. Implement `StreamManager` on server
3. Implement `MonitorStream` gRPC handler
4. Implement `StreamClient` on client side
5. Add admin API endpoints
6. Test with single client
7. Gradually roll out to all clients

## Questions?

This is a complete implementation guide. You can:
- Start with just attendance updates (easiest)
- Add screenshot on-demand next
- Later add other commands as needed

The beauty of this design is it's **backward compatible** - existing clients using heartbeat will continue to work while new clients use streaming!
