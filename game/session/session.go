package session

import (
	"sync"
	"sync/atomic"

	ssync "github.com/kunashu/simple.mud/sync"
)

type statePair struct {
	Manager StateManager
	State   State
}

type cmdChannel chan<- interface{}

type session struct {
	id         SessionId
	conn       net.Conn
	sessionCmd cmdChannel
	factoryDB  ssync.AsyncMap

	currPair statePair
	recv     chan string
	send     chan string
	done     chan bool
	wg       sync.WaitGroup
}

func newSession(id SessionId, conn net.Conn, cmd cmdChannel, factoryDB ssync.AsyncMap) *session {
	s := session{
		id: id, conn: conn, sessionCmd: cmd, factoryDB: factoryDB,
	}

	return &s
}

func (s *session) mainloop() {
Loop:
	for {
		select {
		case <-s.done:
			s.finalize()
		case send := <-s.send:
		case recv := <-s.recv:
		}
	}

}

func (s *session) finalize() {

}

func (s *session) Enter(pair statePair) {
	s.currPair = pair

	s.wg.Add(3)
	go s.mainloop()
	go s.sendloop()
	go s.recvloop()
}

func (s *session) Leave() statePair {

	return pair
}

func (s *session) Terminate() statePair {

	return pair
}

func (s *session) Id() SessionId {
	return s.id
}
func (s *session) ChangeState(id string, data interface{}) error {

}

func (s *session) Write(msg string) error {

}
func (s *session) Close() {

}
