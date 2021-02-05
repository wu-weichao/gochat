package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"gochat/message"
	"net"
	"os"
)

func main() {
	coon, _ := net.Dial("tcp", "127.0.0.1:7001")
	fmt.Println("connect successful...")
	go func() {
		str := make([]byte, 4096)
		for {
			n, _ := os.Stdin.Read(str)
			// 加密
			//coon.Write(str[:n])
			pack := message.NewPackageV1(str[:n])
			fmt.Printf("pack: %v\n", pack)
			pack.Pack(coon)
		}
	}()
	go func() {
		scanner := bufio.NewScanner(coon)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return
			}
			if !atEOF && data[0] == 'v' {
				if len(data) <= message.LenEnd {
					return
				}
				var msgLen int32
				_ = binary.Read(bytes.NewReader(data[message.LenStart:message.LenEnd]), binary.BigEndian, &msgLen)
				if int(msgLen) <= len(data) {
					return int(msgLen) + message.LenEnd, data[:int(msgLen)+message.LenEnd], nil
				}
			}
			return
		})
		for scanner.Scan() {
			pack := message.Package{}
			pack.Unpack(bytes.NewReader(scanner.Bytes()))
			fmt.Printf("recv msg %s\n", pack.Msg)
		}
	}()
	for {

	}
}
