package main

import (
	"google.golang.org/grpc"
	"grpc_custom_transport"
	"grpc_custom_transport/grpc/def"
	"grpc_custom_transport/grpc/impl"
	"net"
)

func main() {
	var serviceImpl = new(impl.UserServiceImpl)
	var server = grpc.NewServer()
	def.RegisterUserServiceServer(server, serviceImpl)
	baseListener, err := net.Listen("tcp", grpc_custom_transport.Address)
	if err != nil {
		panic(err)
	}
	var listener = grpc_custom_transport.NewABListener(baseListener)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
