package grpcclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) InitOrganization(e OrgInitInfo) (string, error) {

	req := org.InitOrgRequest{
		Id:                e.Id,
		Eid:               e.EID,
		Domain:            e.Domain,
		OrgName:           e.OrgName,
		AdminEmailId:      e.AdminEmailId,
		AdminPhoneNumber:  e.AdminPhoneNumber,
		AdminName:         e.AdminName,
		Department:        e.Department,
		Designation:       e.Designation,
		EmailVerified:     e.EmailVerified,
		Password:          e.Password,
		Role:              e.Role,
		FirstLoginPending: e.FirstLoginPending,
		Status:            e.Status,
		UseGoogleOAuth:    e.UseGoogleOAuth,
	}
	res, err := a.client.InitOrganization(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}

	if res.IsError {
		return "error", errors.New(res.Message)
	}

	return "success", nil
}
