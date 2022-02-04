package grpc_custom_transport

import (
	"net"
	"time"
)

type DataChan chan []byte

type BaseReadFunc func(b []byte) (n int, err error)

type Codec interface {
	OnRead(baseConn net.Conn, b []byte) (n int, err error) //这里返回的n是baseConn读取的n
	OnWrite(baseConn net.Conn, b []byte) (n int, err error)
	OnClose(baseConn net.Conn) error
}

type CodecConn struct {
	baseConn net.Conn
	codec    Codec
}

func NewCodecConn(baseConn net.Conn, codec Codec) *CodecConn {
	return &CodecConn{
		baseConn: baseConn,
		codec:    codec,
	}
}

func (c CodecConn) Read(b []byte) (n int, err error) {
	return c.codec.OnRead(c.baseConn, b)
}

func (c CodecConn) Write(b []byte) (n int, err error) {
	return c.codec.OnWrite(c.baseConn, b)
}

func (c *CodecConn) Close() error {
	return c.codec.OnClose(c.baseConn)
}

func (c CodecConn) LocalAddr() net.Addr {
	return c.baseConn.LocalAddr()
}

func (c CodecConn) RemoteAddr() net.Addr {
	return c.baseConn.RemoteAddr()
}

func (c *CodecConn) SetDeadline(t time.Time) error {
	return c.baseConn.SetDeadline(t)
}

func (c *CodecConn) SetReadDeadline(t time.Time) error {
	return c.baseConn.SetReadDeadline(t)
}

func (c *CodecConn) SetWriteDeadline(t time.Time) error {
	return c.baseConn.SetWriteDeadline(t)
}
