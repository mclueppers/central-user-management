package service

import (
	"cum/types"
)

// SlackService represents a slack service
type SlackService struct {
	token string
	ID    string
	Type  string
}

// NewSlackService creates a new slack service
func NewSlackService(token string) *SlackService {
	return &SlackService{
		token: token,
		ID:    "slack",
		Type:  "slack",
	}
}

// AddUser adds a user to Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) AddUser(user types.User) error {
	payload := map[string]interface{}{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"token":      s.token,
	}

	err := post("https://slack.com/api/users.admin.invite", payload)
	if err != nil {
		// TODO: handle error
		return err
	}

	return nil
}

// RemoveUser removes a user from Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) RemoveUser(user types.User) error {
	payload := map[string]interface{}{
		"email": user.Email,
		"token": s.token,
	}

	err := post("https://slack.com/api/users.admin.invite", payload)
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns a list of users from Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) GetUsers() ([]types.User, error) {
	payload := map[string]interface{}{
		"token": s.token,
	}

	users, err := get("https://slack.com/api/users.list", payload)
	if err != nil {
		return nil, err
	}

	return users.([]types.User), nil
}

// AddTeam adds a team to Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) AddTeam(team types.Team) error {
	payload := map[string]interface{}{
		"team":  team.Name,
		"token": s.token,
	}

	err := post("https://slack.com/api/team.create", payload)
	if err != nil {
		return err
	}

	return nil
}

// RemoveTeam removes a team from Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) RemoveTeam(team types.Team) error {
	payload := map[string]interface{}{
		"team":  team.Name,
		"token": s.token,
	}

	err := post("https://slack.com/api/team.create", payload)
	if err != nil {
		return err
	}

	return nil
}

// GetTeams returns a list of teams from Slack workspace using Slack API
// and returns an error if any
func (s *SlackService) GetTeams() ([]types.Team, error) {
	payload := map[string]interface{}{
		"token": s.token,
	}

	teams, err := get("https://slack.com/api/team.list", payload)
	if err != nil {
		return nil, err
	}

	return teams.([]types.Team), nil
}

// GetID returns the ID of the service
func (s *SlackService) GetID() string {
	return s.ID
}

// GetType returns the type of the service
func (s *SlackService) GetType() string {
	return s.Type
}

// String returns a string representation of the service
func (s *SlackService) String() string {
	return s.ID
}

// post sends a POST request to the given URL with the given payload
// and returns an error if any
func post(url string, payload map[string]interface{}) error {
	return nil
}

// get sends a GET request to the given URL with the given payload
// and returns an error if any
func get(url string, payload map[string]interface{}) (interface{}, error) {
	return nil, nil
}
