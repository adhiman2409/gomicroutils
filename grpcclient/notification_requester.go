package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/notification"
)

func (a *NotificationClient) SendNotification(req NotificationRequest, domain string) (NotificationResponse, error) {

	r := notification.NotificationRequest{
		EmployeeId:     req.EmployeeId,
		Title:          req.Title,
		Body:           req.Body,
		ImageURL:       req.ImageUrl,
		DataMap:        req.DataMap,
		Domain:         domain,
		Group:          req.Group,
		Topic:          req.Topic,
		BroadcastToAll: req.BroadcastToAll,
	}

	res, err := a.client.SendNotification(context.Background(), &r)
	if err != nil {
		return NotificationResponse{}, err
	}

	return NotificationResponse{
		Status:  res.GetStatus(),
		Message: res.GetMessage(),
	}, nil
}
