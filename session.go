package gosocket

type Session struct {
	ID       string
	UID      string
	conn     *Conn
	settings map[string]interface{}
}
