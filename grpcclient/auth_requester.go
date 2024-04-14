package grpcclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
)

func (a *GrpcClient) Verify(accessToken string, routeName string) (AuthInfo, error) {

	req := auth.VerifyRequest{
		Token:     accessToken,
		RouteName: routeName,
	}
	res, err := a.client.Verify(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return AuthInfo{}, err
	}

	if res.IsAuthorised {
		return AuthInfo{
			Authorised:  true,
			Tenant:      res.Tenant,
			Domain:      res.Domain,
			Department:  res.Department,
			Name:        res.Name,
			EmailId:     res.EmailId,
			PhoneNumber: res.PhoneNumber,
			Role:        res.Role,
		}, nil
	} else {
		return AuthInfo{}, errors.New(res.Message)
	}
}
