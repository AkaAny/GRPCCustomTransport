package grpc_custom_transport

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

type ABCodec struct {
}

var DefaultByteOrder = binary.BigEndian

const Magic = "hermesab"

func (c ABCodec) OnRead(baseConn net.Conn, b []byte) (n int, err error) {
	//magic
	var magic = make([]byte, len(Magic)) //byte
	{
		n, err = io.ReadAtLeast(baseConn, magic, len(magic))
		if err != nil {
			return n, err
		}
		if string(magic) != Magic {
			return 8, errors.New("invalid magic")
		}
	}
	//data length
	var dataLength = uint32(0)
	{
		err = binary.Read(baseConn, DefaultByteOrder, &dataLength)
		if err != nil {
			return 0, err
		}
	}
	//data
	{
		n, err = io.ReadAtLeast(baseConn, b, int(dataLength))
	}
	return n, nil
}

func (c ABCodec) OnWrite(baseConn net.Conn, b []byte) (n int, err error) {
	var buffer = bytes.NewBuffer(nil)
	buffer.Write([]byte(Magic))
	var dataLength = uint32(len(b))
	err = binary.Write(buffer, DefaultByteOrder, dataLength)
	if err != nil {
		return buffer.Len(), err
	}
	buffer.Write(b)
	actualWrite, err := buffer.WriteTo(baseConn)
	fmt.Println("actual write:", actualWrite)
	if err != nil {
		return int(actualWrite), nil
	}
	//返回的n是预期往baseConn里写入的n
	//grpc的transport层对这个有检查，internal/transport/http2_client.go:373
	return int(dataLength), err
}

func (c ABCodec) OnClose(baseConn net.Conn) error {
	fmt.Println("on close")
	return baseConn.Close()
}
