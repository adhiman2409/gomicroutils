package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) GetAttendanceInfo(orgName, domain string) (AttendanceInfo, error) {
	req := org.AttendanceRequest{
		OrgName: orgName,
		Domain:  domain,
	}
	res, err := a.client.GetAttendanceInfo(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return AttendanceInfo{}, err
	}
	attendanceInfo := AttendanceInfo{
		OrgId:                   res.OrgId,
		OrgName:                 res.OrgName,
		CaptureCheckInLocation:  res.CaptureCheckInLocation,
		EnforceLocationCheckIn:  res.EnforceLocationCheckIn,
		EnforceLocationCheckOut: res.EnforceLocationCheckOut,
		RequiredLat:             res.RequiredLat,
		RequiredLng:             res.RequiredLng,
		FlexiHoursEnabled:       res.FlexiHoursEnabled,
		FlexiHourWindowInMin:    res.FlexiHourWindowInMin,
		OrgCheckInTime:          res.OrgCheckInTime,
		OrgCheckOutTime:         res.OrgCheckOutTime,
		CheckinMarginInMin:      res.CheckInMarginInMin,
		OfficeHours:             res.OfficeHours,
	}
	return attendanceInfo, nil
}
