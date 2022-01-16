package gosocket

import (
	"net"
	"time"
)

type Conn struct {
	sid        string
	rawConn    net.Conn
	sendCh     chan []byte
	done       chan error
	hbTimer    *time.Timer
	name       string
	messageCh  chan *Message
	hbInterval time.Duration
	hbTimeout  time.Duration
}
