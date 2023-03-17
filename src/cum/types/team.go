package types

// Team represents a team of a project
type Team struct {
	ID          string
	Name        string
	Description string
	TeamLead    User
	Members     []Member
}
