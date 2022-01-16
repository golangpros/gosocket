package gosocket

import "net"

type Conn struct {
	sid     string
	rawConn net.Conn
	sendCh  chan []byte
	done    chan error
}
