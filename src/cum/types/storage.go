package types

// Storage represents a storage for users, groups and sessions
type Storage interface {
	UserStorage
	GroupStorage
	SessionStorage
}

// StorageFactory represents a factory for storages
type StorageFactory interface {
	UserStorageFactory
	GroupStorageFactory
	SessionStorageFactory
}

// StorageFactoryFunc represents a factory function for storages
type StorageFactoryFunc func() (Storage, error)

// NewStorage creates a new storage
func (f StorageFactoryFunc) NewStorage() (Storage, error) {
	return f()
}

// NewStorage creates a new storage
func NewStorage(factory StorageFactory) (Storage, error) {
	userStorage, err := factory.NewUserStorage()
	if err != nil {
		return nil, err
	}
	groupStorage, err := factory.NewGroupStorage()
	if err != nil {
		return nil, err
	}
	sessionStorage, err := factory.NewSessionStorage()
	if err != nil {
		return nil, err
	}
	return &storage{
		userStorage:    userStorage,
		groupStorage:   groupStorage,
		sessionStorage: sessionStorage,
	}, nil
}

// storage represents a storage for users, groups and sessions
type storage struct {
	userStorage    UserStorage
	groupStorage   GroupStorage
	sessionStorage SessionStorage
}

// CreateUser creates a new user
func (s *storage) CreateUser(user *User) error {
	return s.userStorage.CreateUser(user)
}

// GetUserByID returns a user by ID
func (s *storage) GetUserByID(id string) (*User, error) {
	return s.userStorage.GetUserByID(id)
}

// GetUserByEmail returns a user by email
func (s *storage) GetUserByEmail(email string) (*User, error) {
	return s.userStorage.GetUserByEmail(email)
}

func (s *storage) GetUserByUsername(username string) (*User, error) {
	return s.userStorage.GetUserByUsername(username)
}

// UpdateUser updates a user
func (s *storage) UpdateUser(user *User) error {
	return s.userStorage.UpdateUser(user)
}

// DeleteUser deletes a user
func (s *storage) DeleteUser(id string) error {
	return s.userStorage.DeleteUser(id)
}

// CreateGroup creates a new group
func (s *storage) CreateGroup(group *Group) error {
	return s.groupStorage.CreateGroup(group)
}

// GetGroupByID returns a group by ID
func (s *storage) GetGroupByID(id string) (*Group, error) {
	return s.groupStorage.GetGroupByID(id)
}

// GetGroupByName returns a group by name
func (s *storage) GetGroupByName(name string) (*Group, error) {
	return s.groupStorage.GetGroupByName(name)
}

// UpdateGroup updates a group
func (s *storage) UpdateGroup(group *Group) error {
	return s.groupStorage.UpdateGroup(group)
}

// DeleteGroup deletes a group
func (s *storage) DeleteGroup(group *Group) error {
	return s.groupStorage.DeleteGroup(group)
}

// AddMemberToGroup adds a member to a group
func (s *storage) AddMemberToGroup(m Member, parentGroupID string) error {
	return s.groupStorage.AddMemberToGroup(m, parentGroupID)
}

// RemoveMemberFromGroup removes a member from a group
func (s *storage) RemoveMemberFromGroup(m *Member, parentGroupID string) error {
	return s.groupStorage.RemoveMemberFromGroup(m, parentGroupID)
}

// AddGroupToGroup adds a group to a group
func (s *storage) CreateSession(session *Session) error {
	return s.sessionStorage.CreateSession(session)
}

// GetSessionByID returns a session by ID
func (s *storage) GetSessionByID(id string) (*Session, error) {
	return s.sessionStorage.GetSessionByID(id)
}

// DeleteSession deletes a session
func (s *storage) DeleteSession(id string) error {
	return s.sessionStorage.DeleteSession(id)
}
