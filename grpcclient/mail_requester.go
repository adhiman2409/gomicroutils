package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/mail"
)

func (a *MailClient) SendMail(req MailRequest, domain string) (MailResponse, error) {

	r := mail.MailRequest{
		From:         req.From,
		To:           req.To,
		Subject:      req.Subject,
		DataMap:      req.DataMap,
		TemplateName: req.TemplateName,
		AckRequired:  req.AckRequired,
		Priority:     req.Priority,
		Domain:       domain,
	}

	res, err := a.client.SendMail(context.Background(), &r)
	if err != nil {
		return MailResponse{}, err
	}

	return MailResponse{
		Status:        res.GetStatus(),
		Message:       res.GetMessage(),
		CorrelationId: res.GetCorrelationId(),
	}, nil
}
