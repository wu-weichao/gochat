package connector

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"gochat/message"

	"fmt"
	"gochat/config"
	"net"
)

type ConnectTcp struct {
	buf []byte
}

func TcpRun() error {
	// tcp server
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Tcp.Host, config.Tcp.Port))
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	fmt.Printf("tcp %v started\n", addr)
	var tcp = ConnectTcp{}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			// listener error

		}
		go tcp.readHandler(conn)
		go tcp.WriteHandler(conn, []byte("连接成功"))
	}
}

func (tcp *ConnectTcp) readHandler(conn *net.TCPConn) {
	// close
	defer conn.Close()

	// read
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return
		}
		if !atEOF && data[0] == byte('v') {
			if len(data) <= message.LenEnd {
				return
			}
			// get message
			var msgLen int32
			_ = binary.Read(bytes.NewReader(data[message.LenStart:message.LenEnd]), binary.BigEndian, &msgLen)
			if int(msgLen) <= len(data) {
				return int(msgLen) + message.LenEnd, data[:int(msgLen)+message.LenEnd], nil
			}
		}
		return
	})
	// 协议 [2][4][...]
	for scanner.Scan() {
		pack := &message.Package{}
		err := pack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {

		}
		/******** 解析内容 ********/
		// test login
		if string(pack.Msg) == "login" {
			// 记录用户登录信息

		}

		fmt.Printf("recv msg %s\n", pack.Msg)
	}
}

func (tcp *ConnectTcp) WriteHandler(conn *net.TCPConn, msg []byte) {
	pack := message.NewPackageV1(msg)
	_ = pack.Pack(conn)
}
