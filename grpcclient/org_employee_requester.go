package grpcclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/org"
)

func (a *OrgClient) OrgAddEmployee(e OrgEmployeeInfo) (string, error) {

	req := org.EmployeeRequest{
		RequestType:     "add",
		Domain:          e.Domain,
		EmployeeId:      e.EmployeeId,
		Department:      e.Department,
		Designation:     e.Designation,
		Name:            e.Name,
		Status:          e.Status,
		ProfileImageURL: e.ProfileImageURL,
		EmailId:         e.EmailId,
	}
	res, err := a.client.OrgEmployee(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}

	if res.IsError {
		return "error", errors.New(res.Message)
	}

	return "success", nil
}

func (a *OrgClient) OrgRemoveEmployee(e OrgEmployeeInfo) (string, error) {

	req := org.EmployeeRequest{
		RequestType:     "remove",
		Domain:          e.Domain,
		EmployeeId:      e.EmployeeId,
		Department:      e.Department,
		Designation:     e.Designation,
		Name:            e.Name,
		Status:          e.Status,
		ProfileImageURL: e.ProfileImageURL,
		EmailId:         e.EmailId,
	}
	res, err := a.client.OrgEmployee(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}

	if res.IsError {
		return "error", errors.New(res.Message)
	}

	return "success", nil
}

func (a *OrgClient) OrgUpdateEmployee(e OrgEmployeeInfo) (string, error) {

	req := org.EmployeeRequest{
		RequestType:     "update",
		Domain:          e.Domain,
		EmployeeId:      e.EmployeeId,
		Department:      e.Department,
		Designation:     e.Designation,
		Name:            e.Name,
		Status:          e.Status,
		ProfileImageURL: e.ProfileImageURL,
		EmailId:         e.EmailId,
	}
	res, err := a.client.OrgEmployee(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}

	if res.IsError {
		return "error", errors.New(res.Message)
	}

	return "success", nil
}
