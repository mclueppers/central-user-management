package types

// Project represents a project
type Project struct {
	ID          string
	Name        string
	Description string
	Services    []*Service
	Teams       []*Team
}
