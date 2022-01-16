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

func (c *Conn) getName() string {
	return c.name
}

func NewConn(c net.Conn, hbInterval time.Duration, hbTimeout time.Duration) *Conn {
	conn := &Conn{}
	conn.name = c.RemoteAddr().String()
	conn.hbTimer = time.NewTimer(conn.hbInterval)

	if conn.hbInterval == 0 {
		conn.hbTimer.Stop()
	}
	return conn
}

func (c *Conn) close() {
	c.hbTimer.Stop()
	c.rawConn.Close()
}

func (c *Conn) sendMessage(msg *Message) error {
	pkg, err := Encode(msg)
	if err != nil {
		return err
	}

	c.sendCh <- pkg
	return nil
}
