package transport

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "goKit/Endpoint"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return endpoint.UserRequest{
			Uid: uid,
		}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(w).Encode(response)
}
