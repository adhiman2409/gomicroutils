package grpcclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
)

func (a *GrpcClient) VerifyEmailToken(token, domain string) (EmailTokenInfo, error) {

	req := auth.VerifyEmailTokenRequest{
		Token:  token,
		Domain: domain,
	}
	res, err := a.client.VerifyEmailToken(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return EmailTokenInfo{}, err
	}

	if !res.IsError {
		return EmailTokenInfo{
			TenantId:     res.TenantId,
			Domain:       res.Domain,
			SubDomain:    res.SubDomain,
			EmailId:      res.EmailId,
			SecretId:     res.SecretId,
			Name:         res.Name,
			IsError:      false,
			ErrorMessage: "",
		}, nil
	} else {
		return EmailTokenInfo{}, errors.New(res.ErrorMessage)
	}
}

func (a *GrpcClient) GetNewEmailToken(tenantId, domain, subDomain, emailId, secretId, name string) (string, error) {

	req := auth.EmailTokenRequest{
		TenantId:  tenantId,
		Domain:    domain,
		SubDomain: subDomain,
		EmailId:   emailId,
		SecretId:  secretId,
		Name:      name,
	}
	res, err := a.client.GetNewEmailToken(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if !res.IsError {
		return res.Token, nil
	} else {
		return "", errors.New(res.ErrorMessage)
	}
}
