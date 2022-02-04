package grpc_custom_transport

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

const Address = "localhost:19082"

func NewClientConn() *grpc.ClientConn {
	//https://segmentfault.com/a/1190000041066320
	clientConn, err := grpc.Dial(Address, grpc.WithContextDialer(func(ctx context.Context, address string) (net.Conn, error) {
		baseConn, err := net.Dial("tcp", address)
		if err != nil {
			panic(err)
		}
		return NewCodecConn(baseConn, &ABCodec{}), nil
	}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return clientConn
}
