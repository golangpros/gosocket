package gosocket

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
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

func (c *Conn) GetName() string {
	return c.name
}

func NewConn(c net.Conn, hbInterval time.Duration, hbTimeout time.Duration) *Conn {
	conn := &Conn{
		rawConn:    c,
		sendCh:     make(chan []byte, 100),
		done:       make(chan error),
		messageCh:  make(chan *Message, 100),
		hbInterval: hbInterval,
		hbTimeout:  hbTimeout,
	}

	conn.name = c.RemoteAddr().String()
	conn.hbTimer = time.NewTimer(conn.hbInterval)

	if conn.hbInterval == 0 {
		conn.hbTimer.Stop()
	}

	return conn
}

func (c *Conn) Close() {
	c.hbTimer.Stop()
	c.rawConn.Close()
}

func (c *Conn) SendMessage(msg *Message) error {
	pkg, err := Encode(msg)
	if err != nil {
		return err
	}

	c.sendCh <- pkg
	return nil
}

func (c *Conn) writeCoroutine(ctx context.Context) {
	hbData := make([]byte, 0)

	for {
		select {
		case <-ctx.Done():
			return

		case pkt := <-c.sendCh:

			if pkt == nil {
				continue
			}

			if _, err := c.rawConn.Write(pkt); err != nil {
				c.done <- err
			}

		case <-c.hbTimer.C:
			hbMessage := NewMessage(MsgHeartbeat, hbData)
			c.SendMessage(hbMessage)
			if c.hbInterval > 0 {
				c.hbTimer.Reset(c.hbInterval)
			}
		}
	}
}

func (c *Conn) readCoroutine(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return

		default:
			if c.hbInterval > 0 {
				err := c.rawConn.SetReadDeadline(time.Now().Add(c.hbTimeout))
				if err != nil {
					c.done <- err
					continue
				}
			}
			buf := make([]byte, 4)
			_, err := io.ReadFull(c.rawConn, buf)
			if err != nil {
				c.done <- err
				continue
			}

			bufReader := bytes.NewReader(buf)

			var dataSize int32
			err = binary.Read(bufReader, binary.LittleEndian, &dataSize)
			if err != nil {
				c.done <- err
				continue
			}

			databuf := make([]byte, dataSize)
			_, err = io.ReadFull(c.rawConn, databuf)
			if err != nil {
				c.done <- err
				continue
			}

			msg, err := Decode(databuf)
			if err != nil {
				c.done <- err
				continue
			}

			if msg.GetID() == MsgHeartbeat {
				continue
			}

			c.messageCh <- msg
		}
	}
}
