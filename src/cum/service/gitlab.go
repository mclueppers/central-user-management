package service

import (
	"cum/types"
)

// GitlabService represents a Gitlab service
type GitlabService struct {
	ID   string
	Type string

	instance GitlabServiceConfig
}

// GitlabServiceConfig represents a Gitlab service configuration
type GitlabServiceConfig struct {
	Host  string
	Token string
}

// NewGitlabService creates a new Gitlab service
func NewGitlabService(instance GitlabServiceConfig) *GitlabService {
	return &GitlabService{instance: instance}
}

// AddUser adds a user to the service
func (s *GitlabService) AddUser(user types.User) error {
	// TODO: implement
	return nil
}

// RemoveUser removes a user from the service
func (s *GitlabService) RemoveUser(user types.User) error {
	// TODO: implement
	return nil
}

// GetUsers gets all users from the service
func (s *GitlabService) GetUsers() ([]types.User, error) {
	// TODO: implement
	return nil, nil
}

// AddTeam adds a team to the service
func (s *GitlabService) AddTeam(team types.Team) error {
	// TODO: implement
	return nil
}

// RemoveTeam removes a team from the service
func (s *GitlabService) RemoveTeam(team types.Team) error {
	// TODO: implement
	return nil
}

// GetTeams gets all teams from the service
func (s *GitlabService) GetTeams() ([]types.Team, error) {
	// TODO: implement
	return nil, nil
}

// GetID gets the ID of the service
func (s *GitlabService) GetID() string {
	return s.ID
}

// GetType gets the type of the service
func (s *GitlabService) GetType() string {
	return s.Type
}

// String returns a string representation of the service
func (s *GitlabService) String() string {
	return s.ID
}
