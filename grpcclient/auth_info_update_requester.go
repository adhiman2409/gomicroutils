package grpcclient

import (
	"context"
	"fmt"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
)

func (a *GrpcClient) UpdateInfo(id, department, designation, imageURL, role string) (string, error) {

	req := auth.InfoUpdateRequest{
		Id:          id,
		ImageURL:    imageURL,
		Role:        role,
		Designation: designation,
		Department:  department,
	}
	res, err := a.client.UpdateInfo(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return res.Message, err
	}

	return "success", nil
}
