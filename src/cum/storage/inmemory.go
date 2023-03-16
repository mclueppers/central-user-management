package storage

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"cum/types"
)

// InMemoryStorage implements the Storage interface with an in-memory map
type InMemoryStorage struct {
	Users    map[string]*types.User
	Groups   map[string]*types.Group
	Sessions map[string]*types.Session
	mu       sync.Mutex
}

// NewInMemoryStorage creates a new InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Users:    make(map[string]*types.User),
		Groups:   make(map[string]*types.Group),
		Sessions: make(map[string]*types.Session),
	}
}

func (s *InMemoryStorage) NewUserStorage() (types.UserStorage, error) {
	return s, nil
}

func (s *InMemoryStorage) NewGroupStorage() (types.GroupStorage, error) {
	return s, nil
}

func (s *InMemoryStorage) NewSessionStorage() (types.SessionStorage, error) {
	return s, nil
}

// CreateUser creates a new user
func (s *InMemoryStorage) CreateUser(user *types.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Users[user.ID]; ok {
		return errors.New("user already exists")
	}
	s.Users[user.ID] = user
	return nil
}

// GetUserByID returns a user by its ID
func (s *InMemoryStorage) GetUserByID(id string) (*types.User, error) {
	if user, ok := s.Users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

// GetUserByUsername returns a user by its username
func (s *InMemoryStorage) GetUserByUsername(username string) (*types.User, error) {
	for _, user := range s.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("username not found")
}

// GetUserByEmail returns a user by its email
func (s *InMemoryStorage) GetUserByEmail(email string) (*types.User, error) {
	for _, user := range s.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// UpdateUser updates a user
func (s *InMemoryStorage) UpdateUser(user *types.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Users[user.ID]; !ok {
		return errors.New("user not found")
	}
	s.Users[user.ID] = user
	return nil
}

// DeleteUser deletes a user
func (s *InMemoryStorage) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Users[id]; !ok {
		return errors.New("user not found")
	}
	delete(s.Users, id)
	return nil
}

// CreateGroup creates a new group
func (s *InMemoryStorage) CreateGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Groups[group.ID]; ok {
		return errors.New("group already exists")
	}
	s.Groups[group.ID] = group
	return nil
}

// GetGroupByID returns a group by its ID
func (s *InMemoryStorage) GetGroupByID(id string) (*types.Group, error) {
	if group, ok := s.Groups[id]; ok {
		return group, nil
	}
	return nil, errors.New("group not found")
}

// GetGroupByName returns a group by its name
func (s *InMemoryStorage) GetGroupByName(name string) (*types.Group, error) {
	for _, group := range s.Groups {
		if group.Name == name {
			return group, nil
		}
	}
	return nil, errors.New("group not found")
}

// UpdateGroup updates a group
func (s *InMemoryStorage) UpdateGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Groups[group.ID]; !ok {
		return errors.New("group not found")
	}
	s.Groups[group.ID] = group
	return nil
}

// DeleteGroup deletes a group
func (s *InMemoryStorage) DeleteGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Groups[group.ID]; !ok {
		return errors.New("group not found")
	}
	delete(s.Groups, group.ID)
	return nil
}

// AddMemberToGroup adds a member to a group
func (s *InMemoryStorage) AddMemberToGroup(m types.Member, groupID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Groups[groupID]; !ok {
		return errors.New("group not found")
	}
	s.Groups[groupID].Members = append(s.Groups[groupID].Members, &m)
	return nil
}

// RemoveMemberFromGroup removes a member from a group
func (s *InMemoryStorage) RemoveMemberFromGroup(m *types.Member, groupID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Groups[groupID]; !ok {
		return errors.New("group not found")
	}
	for i, id := range s.Groups[groupID].Members {
		if (*id).GetID() == (*m).GetID() {
			s.Groups[groupID].Members = append(s.Groups[groupID].Members[:i], s.Groups[groupID].Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found")
}

// CreateSession creates a new session
func (s *InMemoryStorage) CreateSession(session *types.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Sessions[session.ID]; ok {
		return errors.New("session already exists")
	}
	s.Sessions[session.ID] = session
	return nil
}

// GetSessionByID returns a session by its ID
func (s *InMemoryStorage) GetSessionByID(id string) (*types.Session, error) {
	if session, ok := s.Sessions[id]; ok {
		return session, nil
	}
	return nil, errors.New("session not found")
}

// DeleteSession deletes a session
func (s *InMemoryStorage) DeleteSession(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Sessions[id]; !ok {
		return errors.New("session not found")
	}
	delete(s.Sessions, id)
	return nil
}

// String returns a string representation of the InMemoryStorage instance
func (s *InMemoryStorage) String() string {
	var sb strings.Builder
	sb.WriteString("In-memory storage:\n")
	sb.WriteString("\tUsers:\n")
	for _, user := range s.Users {
		sb.WriteString(fmt.Sprintf("\t\t%s\n", user))
	}
	sb.WriteString("\tGroups:\n")
	for _, group := range s.Groups {
		sb.WriteString(fmt.Sprintf("\t\t%s\n", strings.ReplaceAll(group.String(), "\n", "\n\t\t")))
	}
	sb.WriteString("\tSessions:\n")
	for _, session := range s.Sessions {
		sb.WriteString(fmt.Sprintf("\t\t%s\n", session))
	}
	return sb.String()
}

// Close closes the storage
func (s *InMemoryStorage) Close() error {
	return nil
}
