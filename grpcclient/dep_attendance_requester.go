package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) GetDepAttendanceConf(depName, domain string) (DepAttendanceConf, error) {
	req := org.DepAttendanceRequest{
		DepName: depName,
		Domain:  domain,
	}
	res, err := a.client.GetDepAttendanceConf(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return DepAttendanceConf{}, err
	}
	attendanceInfo := DepAttendanceConf{
		DepId:                   res.DepId,
		Department:              res.DepName,
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
