package gosocket

import (
	uuid "github.com/satori/go.uuid"
)

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

func (s *Session) bindUserID(uid string) {
	s.UID = uid
}

func (s *Session) getUserID() string {
	return s.UID
}

func (s *Session) getConn() *Conn {
	return s.conn
}

func (s *Session) setConn(conn *Conn) {
	s.conn = conn
}

func (s *Session) getSetting(key string) interface{} {

	if v, ok := s.settings[key]; ok {
		return v
	}

	return nil
}

func (s *Session) setSetting(key string, value interface{}) {
	s.settings[key] = value
}
