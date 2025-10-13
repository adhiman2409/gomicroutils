package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/monitor"
)

func (a *MonitorClient) UpdateCheckInCheckoutStatus(req EmployeeCheckInCheckOutRequest, domain string) (EmployeeCheckInCheckOutResponse, error) {

	r := monitor.EmployeeCheckInCheckOutRequest{
		Day:               req.Day,
		Month:             req.Month,
		Year:              req.Year,
		EmployeeId:        req.EmployeeId,
		IsCheckedIn:       req.IsCheckedIn,
		CheckInTimestamp:  req.CheckInTimestamp,
		IsCheckedOut:      req.IsCheckedOut,
		CheckOutTimestamp: req.CheckOutTimestamp,
		Domain:            domain,
	}

	res, err := a.client.UpdateCheckInCheckoutStatus(context.Background(), &r)
	if err != nil {
		return EmployeeCheckInCheckOutResponse{}, err
	}

	return EmployeeCheckInCheckOutResponse{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
	}, nil
}
