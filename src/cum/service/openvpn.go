package service

import (
	"cum/types"
)

// OpenVPNService is a service that manages OpenVPN.
type OpenVPNService struct {
	// The path to the OpenVPN binary.
	Binary string
	// The path to the OpenVPN configuration file.
	Config string
	// The path to the OpenVPN log file.
	Log string
	// The path to the OpenVPN PID file.
	PID string
}

// NewOpenVPNService creates a new OpenVPNService.
func NewOpenVPNService() (*OpenVPNService, error) {
	return &OpenVPNService{
		Binary: "/usr/sbin/openvpn",
		Config: "/etc/openvpn/server.conf",
		Log:    "/var/log/openvpn.log",
		PID:    "/var/run/openvpn.pid",
	}, nil
}

// AddUser adds a user to the OpenVPN configuration file.
func (s *OpenVPNService) AddUser(user types.User) error {
	// TODO: implement
	return nil
}

// RemoveUser removes a user from the OpenVPN configuration file.
func (s *OpenVPNService) RemoveUser(user types.User) error {
	// TODO: implement
	return nil
}

// Users returns a list of users.
func (s *OpenVPNService) Users() ([]types.User, error) {
	// TODO: implement
	return nil, nil
}

// AddTeam adds a team to the OpenVPN configuration file.
func (s *OpenVPNService) AddTeam(team types.Team) error {
	// TODO: implement
	return nil
}

// RemoveTeam removes a team from the OpenVPN configuration file.
func (s *OpenVPNService) RemoveTeam(team types.Team) error {
	// TODO: implement
	return nil
}

// Teams returns a list of teams.
func (s *OpenVPNService) Teams() ([]types.Team, error) {
	// TODO: implement
	return nil, nil
}
