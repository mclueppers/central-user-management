package storage

import (
	"cum/types"
	"fmt"

	"github.com/go-redis/redis"
)

// RedisStorage is a storage backend that uses Redis as a backend
type RedisStorage struct {
	client *redis.Client
}

// RedisStorageConfig is the configuration for the RedisStorage
type RedisStorageConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// NewRedisStorage creates a new RedisStorage
func NewRedisStorage(config *RedisStorageConfig) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + fmt.Sprint(config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return &RedisStorage{
		client: client,
	}
}

// NewUserStorage creates a new user storage
func (r *RedisStorage) NewUserStorage() (types.UserStorage, error) {
	return r, nil
}

// NewGroupStorage creates a new group storage
func (r *RedisStorage) NewGroupStorage() (types.GroupStorage, error) {
	return r, nil
}

// NewSessionStorage creates a new session storage
func (r *RedisStorage) NewSessionStorage() (types.SessionStorage, error) {
	return r, nil
}

// Get returns the value for a given key
func (r *RedisStorage) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// CreateUser creates a new user
func (r *RedisStorage) CreateUser(user *types.User) error {
	return r.client.Set(user.ID, user, 0).Err()
}

// GetUserByEmail returns a user by its email
func (r *RedisStorage) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

// GetUserByID returns a user by its ID
func (r *RedisStorage) GetUserByID(id string) (*types.User, error) {
	user := &types.User{}
	err := r.client.Get(id).Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates a user
func (r *RedisStorage) UpdateUser(user *types.User) error {
	return r.client.Set(user.ID, user, 0).Err()
}

// GetUserByUsername returns a user by its username
func (r *RedisStorage) GetUserByUsername(username string) (*types.User, error) {
	return nil, nil
}

// DeleteUser deletes a user
func (r *RedisStorage) DeleteUser(user string) error {
	return r.client.Del(user).Err()
}

// CreateGroup creates a new group
func (r *RedisStorage) CreateGroup(group *types.Group) error {
	return r.client.Set(group.ID, group, 0).Err()
}

// GetGroupByID returns a group by its ID
func (r *RedisStorage) GetGroupByID(id string) (*types.Group, error) {
	group := &types.Group{}
	err := r.client.Get(id).Scan(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// GetGroupByName returns a group by its name
func (r *RedisStorage) GetGroupByName(name string) (*types.Group, error) {
	return nil, nil
}

// UpdateGroup updates a group
func (r *RedisStorage) UpdateGroup(group *types.Group) error {
	return r.client.Set(group.ID, group, 0).Err()
}

// AddMemberToGroup adds a member to a group
func (r *RedisStorage) AddMemberToGroup(m types.Member, parentGroupId string) error {
	return nil
}

// RemoveMemberFromGroup removes a member from a group
func (r *RedisStorage) RemoveMemberFromGroup(m *types.Member, parentGroupId string) error {
	return nil
}

// DeleteGroup deletes a group
func (r *RedisStorage) DeleteGroup(group *types.Group) error {
	return r.client.Del(group.ID).Err()
}

// CreateSession creates a new session
func (r *RedisStorage) CreateSession(session *types.Session) error {
	return r.client.Set(session.ID, session, 0).Err()
}

// GetSessionByID returns a session by its ID
func (r *RedisStorage) GetSessionByID(id string) (*types.Session, error) {
	session := &types.Session{}
	err := r.client.Get(id).Scan(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// UpdateSession updates a session
func (r *RedisStorage) UpdateSession(session *types.Session) error {
	return r.client.Set(session.ID, session, 0).Err()
}

// DeleteSession deletes a session
func (r *RedisStorage) DeleteSession(session string) error {
	return r.client.Del(session).Err()
}

// Close closes the storage
func (r *RedisStorage) Close() error {
	return r.client.Close()
}
