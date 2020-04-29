package session

// Unique identifier of a Session within a Session Manager.
type SessionId string
type Session interface {
	// Id assigned to the session by the session manager.
	Id() SessionId
	// This will change the state of the session. The id will be used to call
	// the registered manager's NewState. The collee will not be able to Write,
	// nor it will be able to Close the session anymore and it wont be able to
	// change state again. If the id does not exist, then the session will be
	// closed.
	ChangeState(id string, data interface{}) error
	// Sends a message to the connections.
	Write(msg string) error
	// Closes the session. Once set to close, the session will no longer accept
	// any writes, nor it will accept any change states.
	Close()
}

// StateManager should handle the creation and destruction of a state.
type StateManager interface {
	// Should return a state for the given session.
	NewState(data) State
	DeleteState(State)
}

type State interface {
	// Handle incoming messages. The messages will be comming in serially.
	Handle(session Session, msg string)
	// This will be triggered once the session enters the state.
	Enter(session Session)
	// This will be triggered once the session leaves the state.
	Leave(session Session)
}
