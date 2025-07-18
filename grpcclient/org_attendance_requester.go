package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) GetOrgAttendanceConf(orgName, domain, country string) (OrgAttendanceConf, error) {
	req := org.OrgAttendanceRequest{
		OrgName: orgName,
		Domain:  domain,
		Country: country,
	}
	res, err := a.client.GetOrgAttendanceConf(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return OrgAttendanceConf{}, err
	}
	attendanceInfo := OrgAttendanceConf{
		OrgId:                   res.GetOrgId(),
		OrgName:                 res.GetOrgName(),
		CaptureCheckInLocation:  res.GetCaptureCheckInLocation(),
		EnforceLocationCheckIn:  res.GetEnforceLocationCheckIn(),
		EnforceLocationCheckOut: res.GetEnforceLocationCheckOut(),
		RequiredLat:             res.GetRequiredLat(),
		RequiredLng:             res.GetRequiredLng(),
		FlexiHoursEnabled:       res.GetFlexiHoursEnabled(),
		FlexiHourWindowInMin:    res.GetFlexiHourWindowInMin(),
		OrgCheckInTime:          res.GetOrgCheckInTime(),
		OrgCheckOutTime:         res.GetOrgCheckOutTime(),
		CheckinMarginInMin:      res.GetCheckInMarginInMin(),
		WorkingDaysPerWeek:      res.GetWorkingDaysPerWeek(),
		WeeklyOffDays:           res.GetWeeklyOffDays(),
		DailyWorkingHours:       res.GetDailyWorkingHours(),
		Country:                 res.GetCountry(),
	}
	return attendanceInfo, nil
}
