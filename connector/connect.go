package connector

import "net"

type Connect struct {
	uid     interface{}
	connTcp *net.TCPConn
	//connWs
}
