package main

import (
	"context"
	"fmt"
	"grpc_custom_transport"
	"grpc_custom_transport/grpc/def"
)

func main() {
	var clientConn = grpc_custom_transport.NewClientConn()
	var client = def.NewUserServiceClient(clientConn)
	resp, err := client.GetUserInfo(context.Background(), &def.GetUserInfoRequest{
		UserID: "user1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
