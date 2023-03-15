package storage

import (
	"cum/types"

	"github.com/go-redis/redis"
)

// RedisUserStorage implements the UserStorage interface with Redis
type RedisUserStorage struct {
	client *redis.Client
}

// NewRedisUserStorage creates a new RedisUserStorage
func NewRedisUserStorage(client *redis.Client) *RedisUserStorage {
	return &RedisUserStorage{
		client: client,
	}
}

// CreateUser creates a new user
func (s *RedisUserStorage) CreateUser(user *types.User) error {
	return nil
}

// GetUserByID returns a user by its ID
func (s *RedisUserStorage) GetUserByID(id string) (*types.User, error) {
	return nil, nil
}

// GetUserByUsername returns a user by its username
func (s *RedisUserStorage) GetUserByUsername(username string) (*types.User, error) {
	return nil, nil
}

// GetUserByEmail returns a user by its email
func (s *RedisUserStorage) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

// UpdateUser updates a user
func (s *RedisUserStorage) UpdateUser(user *types.User) error {
	return nil
}

// DeleteUser deletes a user
func (s *RedisUserStorage) DeleteUser(id string) error {
	return nil
}
