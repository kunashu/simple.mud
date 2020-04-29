package session

import (
	"net"
	"sync"
	"sync/atom"

	ssync "github.com/kunashu/simple.mud/sync"

	"github.com/satori/go.uuid"
)

type SessionManager struct {
	factoryDB    ssync.AsyncMap
	sessions     map[SessionId]Session
	defaultState ssync.SetOnce
	cmd          chan interface{}
}

func NewSessionManager() *SessionManager {
	mgr := &SessionManager{
		factoryDB: NewAsyncMap(),
		sessions:  map[SessionId]Session{},
		cmd:       make(chan interface{}, 100),
	}
	mgr.start()
	return mgr
}

func (s *SessionManager) start() {
	for cmd := range s.cmd {
		switch cmd.(type) {
		case newSessionCmd:
			cmd := cmd.(newSessionCmd)
			state := cmd.Mgr.NewState(nil)
			if state != nil {
				session := newSession(cmd.Id, cmd.Conn, s.cmd, s.factoryDB)

				s.session[cmd.Id] = session
				session.Enter(statePair{cmd.Mgr, state})
			}
		case changeStateCmd:
			cmd := cmd.(changeStateCmd)
			session := s.sessions[cmd.Id]
			new_state := cmd.Mgr.NewState(cmd.Data)
			pair := session.Leave()

			session.Enter(statePair{cmd.Mgr, new_state})
			pair.Mgr.DeleteState(pair.State)
		case closeSessionCmd:
			cmd := cmd.(closeSessionCmd)
			session := s.sessions[cmd.Id]
			pair := session.Terminate()

			pair.Mgr.DeleteState(pair.State)
			delete(s.sessions, cmd.Id)
		}
	}
}

// Create a new session with a given connection.
func (s *SessionManager) NewSession(conn net.Conn) error {
	if conn == nil {
		return fmt.Errorf("Trying to create a session without a connection.")
	}

	state, ok := s.defaultState.Get()
	if !ok {
		return fmt.Errorf("Can't create a session if there are no default states set.")
	}

	sessionId := SessionId(uuid.NewV4().String())

	s.cmd <- newSessionCmd{sessionId, state, conn}

	return nil
}

// Register a state manager. If the id already exists, then the function will
// panic.
func (s *SessionManager) RegisterState(id string, manager StateManager) error {
	if manager == nil {
		return fmt.Errorf("Trying to register id %s, with a nil value", id)
	}

	if ok := s.factoryDB.Set(id, manager); !ok {
		return fmt.Errorf("A state manager is already registered unded id %s", id)
	}

	return nil
}

// Set default state. This will be called every time a new session is created.
// Once set, it can't be set again.
func (s *SessionManager) DefaultState(id string) error {
	if _, ok := s.factoryDB.Get(id); ok {
		if s.defaultState.Set(id) {
			return nil
		}
		return fmt.Errorf("A default state has already been set, %s", id)
	}
	return fmt.Errorf("The provided state id does not exist, %s", id)
}
