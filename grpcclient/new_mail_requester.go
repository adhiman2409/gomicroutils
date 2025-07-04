package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/mail"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func (a *MailClient) NewSendMail(req NewMailRequest, domain string) (MailResponse, error) {

	// Convert req.DataMap (map[string]any) to map[string]*anypb.Any
	dataMap := make(map[string]*anypb.Any)
	for k, v := range req.DataMap {
		msg, ok := v.(proto.Message)
		if !ok {
			return MailResponse{}, fmt.Errorf("value for key %s does not implement proto.Message", k)
		}
		anyVal, err := anypb.New(msg)
		if err != nil {
			return MailResponse{}, err
		}
		dataMap[k] = anyVal
	}

	r := mail.NewMailRequest{
		From:         req.From,
		To:           req.To,
		Subject:      req.Subject,
		DataMap:      dataMap,
		TemplateName: req.TemplateName,
		AckRequired:  req.AckRequired,
		Priority:     req.Priority,
		Domain:       domain,
	}

	res, err := a.client.NewSendMail(context.Background(), &r)
	if err != nil {
		return MailResponse{}, err
	}

	return MailResponse{
		Status:        res.GetStatus(),
		Message:       res.GetMessage(),
		CorrelationId: res.GetCorrelationId(),
	}, nil
}
