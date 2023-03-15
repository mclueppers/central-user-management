package types

import "fmt"

// Session represents a session entity
type Session struct {
	ID        string
	UserID    string
	ExpiresAt int64
}

// SessionStorage represents a storage for sessions
type SessionStorage interface {
	CreateSession(session *Session) error
	GetSessionByID(id string) (*Session, error)
	DeleteSession(id string) error
}

// SessionStorageFactory represents a factory for session storages
type SessionStorageFactory interface {
	NewSessionStorage() (SessionStorage, error)
}

// SessionStorageFactoryFunc represents a factory function for session storages
type SessionStorageFactoryFunc func() (SessionStorage, error)

// NewSessionStorage creates a new session storage
func (f SessionStorageFactoryFunc) NewSessionStorage() (SessionStorage, error) {
	return f()
}

func (s *Session) GetID() string {
	return s.ID
}

func (s *Session) String() string {
	return fmt.Sprintf("Session ID: %s, User: %s, Expires at: %d", s.ID, s.UserID, s.ExpiresAt)
}
