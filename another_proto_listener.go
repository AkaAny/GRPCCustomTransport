package grpc_custom_transport

import (
	"fmt"
	"net"
)

type ABListener struct {
	baseListener net.Listener
}

func NewABListener(baseListener net.Listener) *ABListener {
	return &ABListener{baseListener: baseListener}
}

func (l ABListener) Accept() (net.Conn, error) {
	baseConn, err := l.baseListener.Accept()
	if err != nil {
		return baseConn, err
	}
	fmt.Println("remote addr:", baseConn.RemoteAddr())
	var codec = &ABCodec{}
	return NewCodecConn(baseConn, codec), nil
}

func (l ABListener) Close() error {
	return l.baseListener.Close()
}

func (l ABListener) Addr() net.Addr {
	return l.baseListener.Addr()
}
