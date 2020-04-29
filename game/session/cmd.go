package session

import (
	"net"
)

type baseCmd struct {
	Id  SessionId
	Mgr StateManager
}

type newSessionCmd struct {
	baseCmd
	Conn net.Conn
}

type changeStateCmd struct {
	baseCmd
	data interface{}
}

type closeSessionCmd baseCmd
