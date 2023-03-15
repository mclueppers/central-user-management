package types

import (
	"fmt"
	"strings"
)

// Group represents a group entity
type Group struct {
	ID          string
	Name        string
	Description string
	OwnerID     *Member
	Members     []*Member
}

// GroupStorage represents a storage for groups
type GroupStorage interface {
	AddMemberToGroup(m Member, parentGroupID string) error
	CreateGroup(group *Group) error
	DeleteGroup(group *Group) error
	GetGroupByID(id string) (*Group, error)
	GetGroupByName(name string) (*Group, error)
	UpdateGroup(group *Group) error
	RemoveMemberFromGroup(m *Member, parentGroupID string) error
}

// GroupStorageFactory represents a factory for group storages
type GroupStorageFactory interface {
	NewGroupStorage() (GroupStorage, error)
}

// GroupStorageFactoryFunc represents a factory function for group storages
type GroupStorageFactoryFunc func() (GroupStorage, error)

// NewGroupStorage creates a new group storage
func (f GroupStorageFactoryFunc) NewGroupStorage() (GroupStorage, error) {
	return f()
}

func (g *Group) GetID() string {
	return g.ID
}

func (g *Group) GetName() string {
	return g.Name
}

func (g *Group) GetDescription() string {
	return g.Description
}

func (g *Group) GetOwnerID() *Member {
	return g.OwnerID
}

func (g *Group) GetMembers() []*Member {
	return g.Members
}

func (g *Group) GetType() string {
	return "group"
}

func (g *Group) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Group: %s, ID: %s, Members:\n", g.Name, g.ID))
	for _, member := range g.Members {
		sb.WriteString(fmt.Sprintf("\t%s\n", *member))
	}
	return sb.String()
}
