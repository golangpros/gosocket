package gosocket

import uuid "github.com/satori/go.uuid"

type Session struct {
	ID       string
	UID      string
	conn     *Conn
	settings map[string]interface{}
}

func newSession(conn *Conn) *Session {
	id, _ := uuid.NewV4()
	session := &Session{
		ID:       id.String(),
		UID:      "",
		conn:     conn,
		settings: make(map[string]interface{}),
	}

	return session
}

func (s *Session) getSessionID() string {
	return s.ID
}

func (s *Session) bindUserId(uid string) {
	s.UID = uid
}

func (s *Session) getUserId() string {
	return s.UID
}

func (s *Session) getConn() *Conn {
	return s.conn
}

func (s *Session) setConn(conn *Conn) {
	s.conn = conn
}
