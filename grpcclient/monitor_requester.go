package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/monitor"
)

func (a *MonitorClient) SendNotification(req EmployeeCheckInCheckOutRequest, domain string) (EmployeeCheckInCheckOutResponse, error) {

	r := monitor.EmployeeCheckInCheckOutRequest{
		EmployeeId:   req.EmployeeId,
		IsCheckedIn:  req.IsCheckedIn,
		IsCheckedOut: req.IsCheckedOut,
		Domain:       domain,
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
