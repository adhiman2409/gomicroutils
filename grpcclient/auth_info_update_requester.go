package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
)

func (a *GrpcClient) UpdateInfo(id, imageURL, role, designation, department, domain, country, timezone string) (string, error) {

	req := auth.InfoUpdateRequest{
		Id:          id,
		ImageURL:    imageURL,
		Role:        role,
		Designation: designation,
		Department:  department,
		Domain:      domain,
		Country:     country,
		TZone:       timezone,
	}
	res, err := a.client.UpdateInfo(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return res.Message, err
	}

	return "success", nil
}
