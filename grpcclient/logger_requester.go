package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/logger"
)

func (a *LoggerClient) SendLog(req LogRequest, domain string) (LogResponse, error) {

	r := logger.LogRequest{
		Log:    req.Log,
		Domain: domain,
	}

	res, err := a.client.SendLog(context.Background(), &r)
	if err != nil {
		return LogResponse{}, err
	}

	return LogResponse{
		Ok: res.Ok,
	}, nil
}
