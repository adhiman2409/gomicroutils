package grpcclient

import (
	"context"

	"github.com/adhiman2409/gomicroutils/genproto/logger"
)

func (a *LoggerClient) Write(p []byte) (int, error) {
	r := logger.LogRequest{
		Log: string(p),
	}
	res, err := a.client.SendLog(context.Background(), &r)
	if err != nil {
		return 0, err
	}

	if res.Ok {
		return 1, nil
	}

	return 0, nil
}
