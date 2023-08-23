package endpoint

import (
	"context"
	services "goKit/Services"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndPoint(userService services.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := userService.GetName(r.Uid)

		return UserResponse{Result: result}, nil
	}
}
