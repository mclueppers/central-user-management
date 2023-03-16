package storage

import (
	"cum/types"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/lib/pq"
)

// PostgresStorage implements the UserStorage interface with PostgreSQL
type PostgresStorage struct {
	config *PostgresStorageConfig
	db     *sql.DB
	mu     sync.Mutex
}

// PostgresStorageConfig is the configuration for a PostgresStorage
type PostgresStorageConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	Database           string
	SSLMode            string
	MaxIdleConnections int
	MaxOpenConnections int
}

// NewPostgresStorage creates a new PostgresStorage
func NewPostgresStorage(config *PostgresStorageConfig) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", "host="+config.Host+" port="+fmt.Sprint(config.Port)+" user="+config.User+" password="+config.Password+" dbname="+config.Database+" sslmode="+config.SSLMode)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetMaxOpenConns(config.MaxOpenConnections)

	// Create the users table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id VARCHAR(255) PRIMARY KEY, username VARCHAR(255) UNIQUE, email VARCHAR(255) UNIQUE, password VARCHAR(255))")
	if err != nil {
		return nil, fmt.Errorf("error creating users table: %v", err)
	}

	// Create the groups table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS groups (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255) UNIQUE)")
	if err != nil {
		return nil, fmt.Errorf("error creating groups table: %v", err)
	}

	// Create the sessions table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sessions (id VARCHAR(255) PRIMARY KEY, user_id VARCHAR(255), expires_at TIMESTAMP)")
	if err != nil {
		return nil, fmt.Errorf("error creating sessions table: %v", err)
	}

	// Create member type ENUM
	_, err = db.Exec(`DO $$
		BEGIN
			CREATE TYPE member_type_enum AS ENUM ('user', 'group');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END
		$$;`)
	if err != nil {
		return nil, fmt.Errorf("error creating member_type enum: %v", err)
	}

	// Create the group_members table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS group_members (group_id VARCHAR(255), member_id VARCHAR(255), member_type member_type_enum, PRIMARY KEY (group_id, member_id, member_type))")
	if err != nil {
		return nil, fmt.Errorf("error creating group_members table: %v", err)
	}

	return &PostgresStorage{
		db:     db,
		config: config,
	}, nil
}

func (s *PostgresStorage) NewUserStorage() (types.UserStorage, error) {
	return s, nil
}

func (s *PostgresStorage) NewGroupStorage() (types.GroupStorage, error) {
	return s, nil
}

func (s *PostgresStorage) NewSessionStorage() (types.SessionStorage, error) {
	return s, nil
}

// CreateUser creates a new user
func (s *PostgresStorage) CreateUser(user *types.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("INSERT INTO users(id, username, email, password) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			return errors.New("user already exists")
		}
		return err
	}
	return nil
}

// GetUserByID returns a user by its ID
func (s *PostgresStorage) GetUserByID(id string) (*types.User, error) {
	row := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id)
	user := &types.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUserByUsername returns a user by its username
func (s *PostgresStorage) GetUserByUsername(username string) (*types.User, error) {
	row := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username)
	user := &types.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUserByEmail returns a user by its email
func (s *PostgresStorage) GetUserByEmail(email string) (*types.User, error) {
	row := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email)
	user := &types.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user email not found")
		}
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user
func (s *PostgresStorage) UpdateUser(user *types.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("UPDATE users SET username = $2, email = $3, password = $4 WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func (s *PostgresStorage) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// CreateGroup creates a new group
func (s *PostgresStorage) CreateGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Create group
	stmt, err := tx.Prepare("INSERT INTO groups(id, name) VALUES($1, $2)")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(group.ID, group.Name)
	if err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			return errors.New("group already exists")
		}
		return err
	}

	// Add group members
	stmt, err = tx.Prepare("INSERT INTO group_members(group_id, member_id, member_type) VALUES($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, member := range group.Members {
		switch m := (*member).(type) {
		case *types.User:
			_, err = stmt.Exec(group.ID, m.ID, "user")
			if err != nil {
				tx.Rollback()
				return err
			}
		case *types.Group:
			_, err = stmt.Exec(group.ID, m.ID, "group")
			if err != nil {
				tx.Rollback()
				return err
			}
		default:
			tx.Rollback()
			return errors.New("invalid group member type")
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetGroupByID returns a group by its ID
func (s *PostgresStorage) GetGroupByID(id string) (*types.Group, error) {
	group := &types.Group{}
	err := s.db.QueryRow("SELECT id, name FROM groups WHERE id = $1", id).Scan(&group.ID, &group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("group not found")
		}
		return nil, err
	}

	rows, err := s.db.Query("SELECT member_id, member_type FROM group_members WHERE group_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var memberID string
		var memberType string
		err = rows.Scan(&memberID, &memberType)
		if err != nil {
			return nil, err
		}
		var member types.Member
		if memberType == "user" {
			member, err = s.GetUserByID(memberID)
		} else if memberType == "group" {
			member, err = s.GetGroupByID(memberID)
		} else {
			return nil, errors.New("invalid member type")
		}

		if err != nil {
			return nil, err
		}

		group.Members = append(group.Members, &member)
	}

	return group, nil
}

// GetGroupByName returns a group by its name
func (s *PostgresStorage) GetGroupByName(name string) (*types.Group, error) {
	group := &types.Group{}
	err := s.db.QueryRow("SELECT id, name FROM groups WHERE name = $1", name).Scan(&group.ID, &group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("group not found")
		}
		return nil, err
	}

	rows, err := s.db.Query("SELECT member_id, member_type FROM group_members WHERE group_id = $1", group.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var memberID string
		var memberType string
		err = rows.Scan(&memberID, &memberType)
		if err != nil {
			return nil, err
		}
		var member types.Member
		if memberType == "user" {
			member, err = s.GetUserByID(memberID)
		} else if memberType == "group" {
			member, err = s.GetGroupByID(memberID)
		} else {
			return nil, errors.New("invalid member type")
		}

		if err != nil {
			return nil, err
		}

		group.Members = append(group.Members, &member)
	}

	return group, nil
}

// UpdateGroup updates a group
func (s *PostgresStorage) UpdateGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Update group
	stmt, err := tx.Prepare("UPDATE groups SET name = $2 WHERE id = $1")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(group.ID, group.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete group members
	stmt, err = tx.Prepare("DELETE FROM group_members WHERE group_id = $1")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(group.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add group members
	stmt, err = tx.Prepare("INSERT INTO group_members(group_id, member_id, member_type) VALUES($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, member := range group.Members {
		_, err = stmt.Exec(group.ID, (*member).GetID(), (*member).GetType())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// AddMemberToGroup adds a member to a group
func (s *PostgresStorage) AddMemberToGroup(member types.Member, groupID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch member := member.(type) {
	case *types.User:
		user := member
		_, err := s.db.Exec("INSERT INTO group_members(group_id, member_id, member_type) VALUES($1, $2, $3)", groupID, user.ID, "user")
		if err != nil {
			return err
		}
	case *types.Group:
		subgroup := member
		_, err := s.db.Exec("INSERT INTO group_members(group_id, member_id, member_type) VALUES($1, $2, $3)", groupID, subgroup.ID, "group")
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown member type")
	}
	return nil
}

// RemoveMemberFromGroup removes a member from a group
func (s *PostgresStorage) RemoveMemberFromGroup(member *types.Member, groupID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch (*member).(type) {
	case *types.User:
		user := (*member).(*types.User)
		_, err := s.db.Exec("DELETE FROM group_members WHERE group_id = $1 AND member_id = $2 AND member_type = $3", groupID, user.ID, "user")
		if err != nil {
			return err
		}
	case *types.Group:
		subgroup := (*member).(*types.Group)
		_, err := s.db.Exec("DELETE FROM group_members WHERE group_id = $1 AND member_id = $2 AND member_type = $3", groupID, subgroup.ID, "group")
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown member type")
	}
	return nil
}

// DeleteGroup deletes a group
func (s *PostgresStorage) DeleteGroup(group *types.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("DELETE FROM groups WHERE id = $1", group.ID)
	if err != nil {
		return err
	}
	return nil
}

// CreateSession creates a new session
func (s *PostgresStorage) CreateSession(session *types.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("INSERT INTO sessions(id, user_id, expires_at) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.ID, session.UserID, session.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionByID returns a session by its ID
func (s *PostgresStorage) GetSessionByID(id string) (*types.Session, error) {
	row := s.db.QueryRow("SELECT id, user_id, expires_at FROM sessions WHERE id = $1", id)
	session := &types.Session{}
	err := row.Scan(&session.ID, &session.UserID, &session.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetSessionByUserID returns a session by its user ID
func (s *PostgresStorage) GetSessionByUserID(userID string) (*types.Session, error) {
	row := s.db.QueryRow("SELECT id, user_id, expires_at FROM sessions WHERE user_id = $1", userID)
	session := &types.Session{}
	err := row.Scan(&session.ID, &session.UserID, &session.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// UpdateSession updates a session
func (s *PostgresStorage) UpdateSession(session *types.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("UPDATE sessions SET user_id = $2, expires_at = $3 WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.ID, session.UserID, session.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession deletes a session
func (s *PostgresStorage) DeleteSession(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stmt, err := s.db.Prepare("DELETE FROM sessions WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database connection
func (s *PostgresStorage) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.db.Close()
	if err != nil {
		// Check if the connection is already closed
		if err.Error() != "sql: database is closed" {
			return err
		}
	}
	return nil
}
