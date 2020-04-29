package telnet

import (
	"net"
)

type SessionManager struct {
	attributes ManagerAttr
	sessions   map[string]Session
	count      int
}

type ManagerAttr struct {
	MaxDataRate   int
	SentTimeout   int
	MaxBufferSize int
	MaxSessions   int
}

func NewSessionManager(attr ManagerAttr) *SessionManager {
	return &SessionManager{
		attributes: attr,
		sessions:   map[string]Session{},
	}
}

func (s *SessionManager) NewSessison(conn net.Conn) (Session, error) {
	if s.AvailabelSessions() > 0 {
		session := NewSession().(pSession)

		return nil
	}

	return fmt.Errorf("No room availabe for the new connection.")
}

func (s *SessionManager) Available() int {

}

func (s *SessionManager) Total() int {

}

func (s *SessionManager) CloseAll() int64 {

	return count
}
