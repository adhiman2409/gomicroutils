package grpcclient

import (
	"context"
	"fmt"
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/monitor"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UpdateCheckInCheckoutStatus updates employee check-in/checkout status via gRPC
func (a *MonitorClient) UpdateCheckInCheckoutStatus(req EmployeeCheckInCheckOutRequest, domain string) (EmployeeCheckInCheckOutResponse, error) {
	r := monitor.EmployeeCheckInCheckOutRequest{
		Day:               req.Day,
		Month:             req.Month,
		Year:              req.Year,
		EmployeeId:        req.EmployeeId,
		IsCheckedIn:       req.IsCheckedIn,
		CheckInTimestamp:  req.CheckInTimestamp,
		CheckInSource:     req.CheckInSource,
		IsCheckedOut:      req.IsCheckedOut,
		CheckOutTimestamp: req.CheckOutTimestamp,
		CheckOutSource:    req.CheckOutSource,
		Domain:            domain,
	}

	fmt.Printf("Sending gRPC request: %s\n", r.String())

	res, err := a.client.UpdateCheckInCheckoutStatus(context.Background(), &r)
	if err != nil {
		log.Printf("gRPC UpdateCheckInCheckoutStatus failed: %v\n", err)
		return EmployeeCheckInCheckOutResponse{}, err
	}

	return EmployeeCheckInCheckOutResponse{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
	}, nil
}

// Heartbeat sends a heartbeat to the server via gRPC
func (a *MonitorClient) Heartbeat(req HeartbeatRequest) (HeartbeatResponse, error) {
	r := monitor.HeartbeatRequest{
		EmployeeId: req.EmployeeId,
		Domain:     req.Domain,
	}

	log.Printf("Sending gRPC Heartbeat request for employee: %s\n", req.EmployeeId)

	res, err := a.client.Heartbeat(context.Background(), &r)
	if err != nil {
		log.Printf("gRPC Heartbeat failed: %v\n", err)
		return HeartbeatResponse{}, err
	}

	return HeartbeatResponse{
		Day:                    res.GetDay(),
		Month:                  res.GetMonth(),
		Year:                   res.GetYear(),
		IsCheckedIn:            res.GetIsCheckedIn(),
		CheckinTimestamp:       res.GetCheckinTimestamp(),
		CheckinSource:          res.GetCheckinSource(),
		IsCheckedOut:           res.GetIsCheckedOut(),
		CheckoutTimestamp:      res.GetCheckoutTimestamp(),
		CheckoutSource:         res.GetCheckoutSource(),
		IsOnLeave:              res.GetIsOnLeave(),
		IsHoliday:              res.GetIsHoliday(),
		IsWeeklyOffDay:         res.GetIsWeeklyOffDay(),
		IsAttendanceFetchError: res.GetIsAttendanceFetchError(),
		Message:                res.GetMessage(),
	}, nil
}

// FetchMonitoringConfig fetches monitoring configuration for a device via gRPC
func (a *MonitorClient) FetchMonitoringConfig(req MonitoringConfigRequest) (MonitoringConfigResponse, error) {
	r := monitor.MonitoringConfigRequest{
		MacAddress: req.MacAddress,
		Domain:     req.Domain,
	}

	log.Printf("Sending gRPC FetchMonitoringConfig request for MAC: %s\n", req.MacAddress)

	res, err := a.client.FetchMonitoringConfig(context.Background(), &r)
	if err != nil {
		log.Printf("gRPC FetchMonitoringConfig failed: %v\n", err)
		return MonitoringConfigResponse{}, err
	}

	return MonitoringConfigResponse{
		StartMonitoringAfterCheckInOnly: res.GetStartMonitoringAfterCheckInOnly(),
		IsCheckedIn:                     res.GetIsCheckedIn(),
		IsCheckedOut:                    res.GetIsCheckedOut(),
		UserId:                          res.GetUserId(),
		MacAddress:                      res.GetMacAddress(),
		Name:                            res.GetName(),
		Domain:                          res.GetDomain(),
		Department:                      res.GetDepartment(),
		Email:                           res.GetEmail(),
		MonitoringEnabled:               res.GetMonitoringEnabled(),
		IdleThreshold:                   res.GetIdleThreshold(),
		ScreenshotInterval:              res.GetScreenshotInterval(),
		CheckInterval:                   res.GetCheckInterval(),
		LocationInterval:                res.GetLocationInterval(),
		TrackApplications:               res.GetTrackApplications(),
		TrackLocation:                   res.GetTrackLocation(),
		TrackBrowser:                    res.GetTrackBrowser(),
		TrackDocuments:                  res.GetTrackDocuments(),
		TrackUSB:                        res.GetTrackUsb(),
		ScreenshotPath:                  res.GetScreenshotPath(),
	}, nil
}

// SendActivityLog sends a single activity log entry via gRPC
func (a *MonitorClient) SendActivityLog(req ActivityLogRequest) (ActivityLogResponse, error) {
	r := &monitor.ActivityLogRequest{
		Timestamp:  timestamppb.New(req.Timestamp),
		UserId:     req.UserId,
		Name:       req.Name,
		MacAddress: req.MacAddress,
		Message:    req.Message,
		EventType:  req.EventType,
		Hostname:   req.Hostname,
		IpAddress:  req.IpAddress,
		SessionId:  req.SessionId,
		Domain:     req.Domain,
	}

	// Convert nested structures
	if req.ActivityInfo != nil {
		r.ActivityInfo = &monitor.ActivityInfo{
			Timestamp:    timestamppb.New(req.ActivityInfo.Timestamp),
			AppName:      req.ActivityInfo.AppName,
			WindowTitle:  req.ActivityInfo.WindowTitle,
			DocumentName: req.ActivityInfo.DocumentName,
			BrowserUrl:   req.ActivityInfo.BrowserURL,
			IsIncognito:  req.ActivityInfo.IsIncognito,
			ProcessId:    req.ActivityInfo.ProcessID,
		}
	}

	if req.UserActivityInfo != nil {
		r.UserActivityInfo = &monitor.UserActivityInfo{
			Status:                 req.UserActivityInfo.Status,
			Timestamp:              timestamppb.New(req.UserActivityInfo.Timestamp),
			IdleDurationSeconds:    req.UserActivityInfo.IdleDurationSeconds,
			ActiveDurationSeconds:  req.UserActivityInfo.ActiveDurationSeconds,
			OfflineDurationSeconds: req.UserActivityInfo.OfflineDurationSeconds,
			SessionDurationSeconds: req.UserActivityInfo.SessionDurationSeconds,
			SleepDurationSeconds:   req.UserActivityInfo.SleepDurationSeconds,
		}
	}

	if req.SystemStatusInfo != nil {
		r.SystemStatusInfo = &monitor.SystemStatusInfo{
			Message: req.SystemStatusInfo.Message,
			Details: req.SystemStatusInfo.Details,
		}
	}

	if req.USBDeviceInfo != nil {
		r.UsbDeviceInfo = &monitor.USBDeviceInfo{
			VendorId:    req.USBDeviceInfo.VendorID,
			ProductId:   req.USBDeviceInfo.ProductID,
			VendorName:  req.USBDeviceInfo.VendorName,
			ProductName: req.USBDeviceInfo.ProductName,
			Action:      req.USBDeviceInfo.Action,
		}
	}

	if req.LocationInfo != nil {
		r.LocationInfo = &monitor.LocationInfo{
			City:      req.LocationInfo.City,
			Country:   req.LocationInfo.Country,
			Latitude:  req.LocationInfo.Latitude,
			Longitude: req.LocationInfo.Longitude,
			Accuracy:  req.LocationInfo.Accuracy,
			Timestamp: timestamppb.New(req.LocationInfo.Timestamp),
			Available: req.LocationInfo.Available,
		}
	}

	if req.NetworkStatus != nil {
		r.NetworkStatus = &monitor.NetworkStatus{
			Timestamp:              timestamppb.New(req.NetworkStatus.Timestamp),
			Ssid:                   req.NetworkStatus.SSID,
			UptimeSeconds:          req.NetworkStatus.UptimeSeconds,
			DowntimeSeconds:        req.NetworkStatus.DowntimeSeconds,
			SessionDurationSeconds: req.NetworkStatus.SessionDurationSeconds,
			IpAddress:              req.NetworkStatus.IpAddress,
		}
	}

	if req.ScreenshotInfo != nil {
		r.ScreenshotInfo = &monitor.ScreenshotInfo{
			Filename:   req.ScreenshotInfo.Filename,
			Timestamp:  timestamppb.New(req.ScreenshotInfo.Timestamp),
			MacAddress: req.ScreenshotInfo.MacAddress,
			UserId:     req.ScreenshotInfo.UserId,
			FilePath:   req.ScreenshotInfo.FilePath,
			Url:        req.ScreenshotInfo.URL,
			Domain:     req.ScreenshotInfo.Domain,
			Department: req.ScreenshotInfo.Department,
		}
	}

	if req.ApplicationUsageReport != nil {
		apps := make([]*monitor.ApplicationUsage, 0, len(req.ApplicationUsageReport.Applications))
		for _, app := range req.ApplicationUsageReport.Applications {
			apps = append(apps, &monitor.ApplicationUsage{
				AppName:               app.AppName,
				Timestamp:             app.Timestamp,
				TimeSpentTodaySeconds: app.TimeSpentTodaySeconds,
				LastActiveTime:        timestamppb.New(app.LastActiveTime),
				SessionCount:          app.SessionCount,
			})
		}

		topApps := make([]*monitor.ApplicationUsage, 0, len(req.ApplicationUsageReport.TopApps))
		for _, app := range req.ApplicationUsageReport.TopApps {
			topApps = append(topApps, &monitor.ApplicationUsage{
				AppName:               app.AppName,
				Timestamp:             app.Timestamp,
				TimeSpentTodaySeconds: app.TimeSpentTodaySeconds,
				LastActiveTime:        timestamppb.New(app.LastActiveTime),
				SessionCount:          app.SessionCount,
			})
		}

		r.ApplicationUsageReport = &monitor.ApplicationUsageReport{
			Applications:     apps,
			TotalTimeSeconds: req.ApplicationUsageReport.TotalTimeSeconds,
			TopApps:          topApps,
		}
	}

	log.Printf("Sending gRPC ActivityLog for event: %s\n", req.EventType)

	res, err := a.client.SendActivityLog(context.Background(), r)
	if err != nil {
		log.Printf("gRPC SendActivityLog failed: %v\n", err)
		return ActivityLogResponse{}, err
	}

	return ActivityLogResponse{
		Success:          res.GetSuccess(),
		Message:          res.GetMessage(),
		EntriesProcessed: res.GetEntriesProcessed(),
	}, nil
}

// SendActivityLogBatch sends multiple activity log entries via gRPC
func (a *MonitorClient) SendActivityLogBatch(req ActivityLogBatchRequest) (ActivityLogResponse, error) {
	entries := make([]*monitor.ActivityLogRequest, 0, len(req.Entries))

	for _, entry := range req.Entries {
		r := &monitor.ActivityLogRequest{
			Timestamp:  timestamppb.New(entry.Timestamp),
			UserId:     entry.UserId,
			Name:       entry.Name,
			MacAddress: entry.MacAddress,
			Message:    entry.Message,
			EventType:  entry.EventType,
			Hostname:   entry.Hostname,
			IpAddress:  entry.IpAddress,
			SessionId:  entry.SessionId,
			Domain:     entry.Domain,
		}

		// Convert nested structures for each entry (same as SendActivityLog)
		if entry.ActivityInfo != nil {
			r.ActivityInfo = &monitor.ActivityInfo{
				Timestamp:    timestamppb.New(entry.ActivityInfo.Timestamp),
				AppName:      entry.ActivityInfo.AppName,
				WindowTitle:  entry.ActivityInfo.WindowTitle,
				DocumentName: entry.ActivityInfo.DocumentName,
				BrowserUrl:   entry.ActivityInfo.BrowserURL,
				IsIncognito:  entry.ActivityInfo.IsIncognito,
				ProcessId:    entry.ActivityInfo.ProcessID,
			}
		}

		if entry.UserActivityInfo != nil {
			r.UserActivityInfo = &monitor.UserActivityInfo{
				Status:                 entry.UserActivityInfo.Status,
				Timestamp:              timestamppb.New(entry.UserActivityInfo.Timestamp),
				IdleDurationSeconds:    entry.UserActivityInfo.IdleDurationSeconds,
				ActiveDurationSeconds:  entry.UserActivityInfo.ActiveDurationSeconds,
				OfflineDurationSeconds: entry.UserActivityInfo.OfflineDurationSeconds,
				SessionDurationSeconds: entry.UserActivityInfo.SessionDurationSeconds,
				SleepDurationSeconds:   entry.UserActivityInfo.SleepDurationSeconds,
			}
		}

		if entry.SystemStatusInfo != nil {
			r.SystemStatusInfo = &monitor.SystemStatusInfo{
				Message: entry.SystemStatusInfo.Message,
				Details: entry.SystemStatusInfo.Details,
			}
		}

		if entry.USBDeviceInfo != nil {
			r.UsbDeviceInfo = &monitor.USBDeviceInfo{
				VendorId:    entry.USBDeviceInfo.VendorID,
				ProductId:   entry.USBDeviceInfo.ProductID,
				VendorName:  entry.USBDeviceInfo.VendorName,
				ProductName: entry.USBDeviceInfo.ProductName,
				Action:      entry.USBDeviceInfo.Action,
			}
		}

		if entry.LocationInfo != nil {
			r.LocationInfo = &monitor.LocationInfo{
				City:      entry.LocationInfo.City,
				Country:   entry.LocationInfo.Country,
				Latitude:  entry.LocationInfo.Latitude,
				Longitude: entry.LocationInfo.Longitude,
				Accuracy:  entry.LocationInfo.Accuracy,
				Timestamp: timestamppb.New(entry.LocationInfo.Timestamp),
				Available: entry.LocationInfo.Available,
			}
		}

		if entry.NetworkStatus != nil {
			r.NetworkStatus = &monitor.NetworkStatus{
				Timestamp:              timestamppb.New(entry.NetworkStatus.Timestamp),
				Ssid:                   entry.NetworkStatus.SSID,
				UptimeSeconds:          entry.NetworkStatus.UptimeSeconds,
				DowntimeSeconds:        entry.NetworkStatus.DowntimeSeconds,
				SessionDurationSeconds: entry.NetworkStatus.SessionDurationSeconds,
				IpAddress:              entry.NetworkStatus.IpAddress,
			}
		}

		if entry.ScreenshotInfo != nil {
			r.ScreenshotInfo = &monitor.ScreenshotInfo{
				Filename:   entry.ScreenshotInfo.Filename,
				Timestamp:  timestamppb.New(entry.ScreenshotInfo.Timestamp),
				MacAddress: entry.ScreenshotInfo.MacAddress,
				UserId:     entry.ScreenshotInfo.UserId,
				FilePath:   entry.ScreenshotInfo.FilePath,
				Url:        entry.ScreenshotInfo.URL,
				Domain:     entry.ScreenshotInfo.Domain,
				Department: entry.ScreenshotInfo.Department,
			}
		}

		if entry.VideoInfo != nil {
			r.VideoInfo = &monitor.VideoInfo{
				Filename:        entry.VideoInfo.Filename,
				Timestamp:       timestamppb.New(entry.VideoInfo.Timestamp),
				MacAddress:      entry.VideoInfo.MacAddress,
				UserId:          entry.VideoInfo.UserId,
				FilePath:        entry.VideoInfo.FilePath,
				Url:             entry.VideoInfo.URL,
				Domain:          entry.VideoInfo.Domain,
				Department:      entry.VideoInfo.Department,
				DurationSeconds: entry.VideoInfo.DurationSeconds,
			}
		}

		if entry.ApplicationUsageReport != nil {
			apps := make([]*monitor.ApplicationUsage, 0, len(entry.ApplicationUsageReport.Applications))
			for _, app := range entry.ApplicationUsageReport.Applications {
				apps = append(apps, &monitor.ApplicationUsage{
					AppName:               app.AppName,
					Timestamp:             app.Timestamp,
					TimeSpentTodaySeconds: app.TimeSpentTodaySeconds,
					LastActiveTime:        timestamppb.New(app.LastActiveTime),
					SessionCount:          app.SessionCount,
				})
			}

			topApps := make([]*monitor.ApplicationUsage, 0, len(entry.ApplicationUsageReport.TopApps))
			for _, app := range entry.ApplicationUsageReport.TopApps {
				topApps = append(topApps, &monitor.ApplicationUsage{
					AppName:               app.AppName,
					Timestamp:             app.Timestamp,
					TimeSpentTodaySeconds: app.TimeSpentTodaySeconds,
					LastActiveTime:        timestamppb.New(app.LastActiveTime),
					SessionCount:          app.SessionCount,
				})
			}

			r.ApplicationUsageReport = &monitor.ApplicationUsageReport{
				Applications:     apps,
				TotalTimeSeconds: entry.ApplicationUsageReport.TotalTimeSeconds,
				TopApps:          topApps,
			}
		}

		entries = append(entries, r)
	}

	batchReq := &monitor.ActivityLogBatchRequest{
		Entries: entries,
	}

	log.Printf("Sending gRPC ActivityLogBatch with %d entries\n", len(entries))

	res, err := a.client.SendActivityLogBatch(context.Background(), batchReq)
	if err != nil {
		log.Printf("gRPC SendActivityLogBatch failed: %v\n", err)
		return ActivityLogResponse{}, err
	}

	return ActivityLogResponse{
		Success:          res.GetSuccess(),
		Message:          res.GetMessage(),
		EntriesProcessed: res.GetEntriesProcessed(),
	}, nil
}
