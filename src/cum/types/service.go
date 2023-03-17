package types

// Service represents a service interface
type Service interface {
	AddUser(user User) error
	RemoveUser(user User) error
	GetUsers() ([]User, error)

	AddTeam(team Team) error
	RemoveTeam(team Team) error
	GetTeams() ([]Team, error)

	GetID() string
	GetType() string
	String() string
}
