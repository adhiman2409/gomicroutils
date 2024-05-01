package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) GetOrgAttendanceConf(orgName, domain string) (OrgAttendanceConf, error) {
	req := org.OrgAttendanceRequest{
		OrgName: orgName,
		Domain:  domain,
	}
	res, err := a.client.GetOrgAttendanceConf(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return OrgAttendanceConf{}, err
	}
	attendanceInfo := OrgAttendanceConf{
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
		WorkingDaysPerWeek:      res.WorkingDaysPerWeek,
		WeeklyOffDays:           res.WeeklyOffDays,
		DailyWorkingHours:       res.DailyWorkingHours,
	}
	return attendanceInfo, nil
}
