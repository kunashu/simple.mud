package telnet

type SessionHandler interface {
	Handle(msg string)
	HandleErr(err error)
}

type Session interface {
	SetHandler(SessionHandler) error
	Write([]byte) (int64, error)
	Close()
}

type pSession struct {
}

func (s *pSession) SetHandler(handler SessionHandler) error {

}

func (s *pSession) closeConn() {

}
