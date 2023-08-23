package main

import (
	"fmt"
	endpoint "goKit/Endpoint"
	services "goKit/Services"
	transport "goKit/Transport"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	user := services.UserService{}
	endp := endpoint.GenUserEndPoint(user)

	serverHandler := httptransport.NewServer(endp, transport.DecodeUserRequest, transport.EncodeUserResponse)
	err := http.ListenAndServe(":8082", serverHandler)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("\"服务已启动：http://localhost:8082")

}
