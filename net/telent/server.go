package telnet

import (
	"net"
)

type Handler interface {
	// Handle the incoming new Session
	Handle(handler Session)
	HandleError(err error)
}

type Server struct {
	handler Handler
	mgr     *SessionManager
}

func NewServerWithListener(l net.Listener, handler Handler) *Server {
	return &Server{
		handler: handler,
		mgr: SessionManager{
			MaxDataRate:   128,
			SentTimeout:   60,
			MaxBufferSize: 65536,
			MaxSessions:   100,
		},
	}
}

func (srv *Server) Serve() error {
	manager := srv.mgr
	handler := srv.handler

	for {
		if conn, err := l.Accept(); err == nil {
			err := manager.NewSessison(conn, func(s Session) {
				handler.Handle(s)
			})
			if err != nil {
				handler.HandleErr(err)
			}
		} else {
			// TODO: Need to add logging interface
		}
	}
}
