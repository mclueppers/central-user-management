package types

// Member represents a member of a group
type Member interface {
	GetID() string
	GetType() string
	String() string
}
