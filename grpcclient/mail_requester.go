package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/mail"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func convertToAnyMap(input map[string]any) map[string]*anypb.Any {
	output := make(map[string]*anypb.Any)
	for key, value := range input {
		if protoMsg, ok := value.(protoreflect.ProtoMessage); ok {
			anyValue, err := anypb.New(protoMsg)
			if err == nil {
				output[key] = anyValue
			}
		}
	}
	return output
}

func (a *MailClient) SendMail(req MailRequest, domain string) (MailResponse, error) {

	r := mail.MailRequest{
		From:         req.From,
		To:           req.To,
		Subject:      req.Subject,
		DataMap:      convertToAnyMap(req.DataMap),
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
