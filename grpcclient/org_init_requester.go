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
		Country:           e.Country,
		TimeZone:          e.TimeZone,
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

func (a *OrgClient) DeductExpenseFromBudget(e DeductExpenseInfo) (DeductExpenseInfoResponse, error) {

	req := org.DeductExpenseRequest{
		Domain:       e.Domain,
		DepartmentId: e.DepartmentId,
		FiscalYear:   e.FiscalYear,
		Amount:       e.Amount,
		Currency:     e.Currency,
		ExpenseId:    e.ExpenseId,
	}
	res, err := a.client.DeductExpenseFromBudget(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())

		return DeductExpenseInfoResponse{
			IsError: true,
			Message: err.Error(),
		}, err
	}

	if res.IsError {
		return DeductExpenseInfoResponse{
			IsError: true,
			Message: res.Message,
		}, err
	}

	return DeductExpenseInfoResponse{
		IsError: false,
		Message: "success",
	}, nil
}
