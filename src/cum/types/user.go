package types

import "fmt"

// User represents a user entity
type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

// UserStorage represents a storage for users
type UserStorage interface {
	Close() error
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

// UserStorageFactory represents a factory for user storages
type UserStorageFactory interface {
	NewUserStorage() (UserStorage, error)
}

// UserStorageFactoryFunc represents a factory function for user storages
type UserStorageFactoryFunc func() (UserStorage, error)

// NewUserStorage creates a new user storage
func (f UserStorageFactoryFunc) NewUserStorage() (UserStorage, error) {
	return f()
}

// GetID returns the ID of the user
func (u *User) GetID() string {
	return u.ID
}

// GetType returns the type of the interface
func (u *User) GetType() string {
	return "user"
}

// GetUsername returns the username of the user
func (u *User) GetUsername() string {
	return u.Username
}

// GetEmail returns the email of the user
func (u *User) GetEmail() string {
	return u.Email
}

// String returns a string representation of the user
func (u *User) String() string {
	return fmt.Sprintf("User: %s, ID: %s, Email: %s", u.Username, u.ID, u.Email)
}
