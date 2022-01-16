package gosocket

type Message struct {
	msgSize  int32
	msgID    int32
	data     []byte
	checkSum uint32
}
