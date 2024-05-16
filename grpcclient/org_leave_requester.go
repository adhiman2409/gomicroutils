package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) GetOrgLeaveConf(orgName, domain string) (OrgLeaveObj, error) {
	req := org.OrgLeaveRequest{
		OrgName: orgName,
		Domain:  domain,
	}
	res, err := a.client.GetOrgLeaveConf(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return OrgLeaveObj{}, err
	}
	orgLeaveList := []OrgLeaveConfiguration{}

	for _, l := range res.GetOrgLeaveConfigurations() {
		orgLeaveList = append(orgLeaveList, OrgLeaveConfiguration{
			OrgName:                     orgName,
			Title:                       l.GetTitle(),
			IsFixed:                     l.GetIsFixed(),
			Frequency:                   l.GetFrequency(),
			LeaveCount:                  float32(l.GetLeaveCount()),
			LeaveUnit:                   l.GetLeaveUnit(),
			MaxLeaveCount:               float32(l.GetMaxLeaveCount()),
			IsCarryForwardAllowed:       l.GetIsCarryForwardAllowed(),
			MaxCarryForwardCount:        float32(l.GetMaxCarryForwardCount()),
			BulkLeaveCount:              int(l.GetBulkLeaveCount()),
			BulkLeaveNoticeInDays:       int(l.GetBulkLeaveNoticeInDays()),
			IsEncashmentAllowed:         l.GetIsEncashmentAllowed(),
			ApplicableAfterWorkingDays:  int(l.GetApplicableAfterWorkingDays()),
			DocumentRequired:            l.GetDocumentRequired(),
			WeeklyOffAndHolidayIncluded: l.GetWeeklyOffAndHolidayIncluded(),
			IsActive:                    l.GetIsActive(),
			Description:                 l.GetDescription(),
			YearStartDate:               l.GetYearStartDate(),
			YearEndDate:                 l.GetYearEndDate(),
		})
	}
	leaveInfoObj := OrgLeaveObj{
		OrgName:                res.GetOrgName(),
		OrgLeaveConfigurations: orgLeaveList,
	}
	return leaveInfoObj, nil
}
