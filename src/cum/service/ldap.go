package service

import (
	"cum/types"
	"fmt"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

// LDAPService is a service that uses LDAP for authentication.
type LDAPService struct {
	ID       string
	Type     string
	Host     string
	Port     int
	BindDN   string
	BindPass string
}

// GetType returns the type of the service.
func (s *LDAPService) GetType() string {
	return s.Type
}

// GetID returns the ID of the service.
func (s *LDAPService) GetID() string {
	return s.ID
}

// String returns a string representation of the service.
func (s *LDAPService) String() string {
	return fmt.Sprintf("%s-%s", s.GetID(), s.GetType())
}

// NewLDAPService creates a new LDAPService.
func NewLDAPService(id, host string, port int, bindDN, bindPass string) *LDAPService {
	return &LDAPService{
		ID:       id,
		Type:     "ldap",
		Host:     host,
		Port:     port,
		BindDN:   bindDN,
		BindPass: bindPass,
	}
}

// Authenticate authenticates a user against the LDAP server.
func (s *LDAPService) Authenticate(username, password string) (bool, error) {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return false, err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return false, err
	}

	return s.authenticate(l, username, password)
}

func (s *LDAPService) authenticate(l *ldap.Conn, username, password string) (bool, error) {
	// Search for the given username.
	err := l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return false, err
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,

		// The filter is the username.
		fmt.Sprintf("(uid=%s)", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, err
	}

	if len(sr.Entries) != 1 {
		return false, fmt.Errorf("user does not exist or too many entries returned")
	}

	userDN := sr.Entries[0].DN

	// Bind as the user to verify their password.
	err = l.Bind(userDN, password)
	if err != nil {
		if strings.Contains(err.Error(), "525") {
			return false, fmt.Errorf("invalid credentials")
		}

		return false, err
	}

	return true, nil
}

// AddUser adds a user to the LDAP server.
func (s *LDAPService) AddUser(user types.User) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Create the user.
	addRequest := ldap.NewAddRequest(fmt.Sprintf("uid=%s,ou=users,dc=example,dc=com", user.Email))
	addRequest.Attribute("objectClass", []string{"inetOrgPerson", "posixAccount", "shadowAccount"})
	addRequest.Attribute("cn", []string{user.FirstName + " " + user.LastName})
	addRequest.Attribute("sn", []string{user.LastName})
	addRequest.Attribute("uid", []string{user.Email})
	addRequest.Attribute("uidNumber", []string{"1000"})
	addRequest.Attribute("gidNumber", []string{"1000"})
	addRequest.Attribute("homeDirectory", []string{"/home/" + user.Email})
	addRequest.Attribute("loginShell", []string{"/bin/bash"})
	addRequest.Attribute("gecos", []string{user.FirstName + " " + user.LastName})
	addRequest.Attribute("userPassword", []string{user.Password})

	err = l.Add(addRequest)
	if err != nil {
		return err
	}

	return nil
}

// RemoveUser removes a user from the LDAP server.
func (s *LDAPService) RemoveUser(user types.User) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Delete the user.
	delRequest := ldap.NewDelRequest(fmt.Sprintf("uid=%s,ou=users,dc=example,dc=com", user.Email), nil)

	err = l.Del(delRequest)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user on the LDAP server.
func (s *LDAPService) UpdateUser(user types.User) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Modify the user.
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("uid=%s,ou=users,dc=example,dc=com", user.Email))
	modifyRequest.Replace("cn", []string{user.FirstName + " " + user.LastName})
	modifyRequest.Replace("sn", []string{user.LastName})
	modifyRequest.Replace("uid", []string{user.Email})
	modifyRequest.Replace("gecos", []string{user.FirstName + " " + user.LastName})
	modifyRequest.Replace("userPassword", []string{user.Password})

	err = l.Modify(modifyRequest)
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns a list of users from the LDAP server.
func (s *LDAPService) GetUsers() ([]types.User, error) {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return nil, err
	}

	// Search for the given username.
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,

		// The filter is the username.
		fmt.Sprintf("(objectClass=inetOrgPerson)"),
		[]string{"cn", "sn", "uid", "gecos"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	users := []types.User{}
	for _, entry := range sr.Entries {
		users = append(users, types.User{
			FirstName: entry.GetAttributeValue("cn"),
			LastName:  entry.GetAttributeValue("sn"),
			Email:     entry.GetAttributeValue("uid"),
			Password:  entry.GetAttributeValue("gecos"),
		})
	}

	return users, nil
}

// GetTeams returns a list of teams from the LDAP server.
func (s *LDAPService) GetTeams() ([]types.Team, error) {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return nil, err
	}

	// Search for the given username.
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,

		// The filter is the username.
		fmt.Sprintf("(objectClass=posixGroup)"),
		[]string{"cn", "gidNumber"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	teams := []types.Team{}
	for _, entry := range sr.Entries {
		teams = append(teams, types.Team{
			Name: entry.GetAttributeValue("cn"),
			ID:   entry.GetAttributeValue("gidNumber"),
		})
	}

	return teams, nil
}

// AddTeam adds a team to the LDAP server.
func (s *LDAPService) AddTeam(team types.Team) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Create the team.
	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=groups,dc=example,dc=com", team.Name))
	addRequest.Attribute("objectClass", []string{"posixGroup", "shadowAccount"})
	addRequest.Attribute("cn", []string{team.Name})
	addRequest.Attribute("gidNumber", []string{team.ID})

	err = l.Add(addRequest)
	if err != nil {
		return err
	}

	return nil
}

// RemoveTeam removes a team from the LDAP server.
func (s *LDAPService) RemoveTeam(team types.Team) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Delete the team.
	delRequest := ldap.NewDelRequest(fmt.Sprintf("cn=%s,ou=groups,dc=example,dc=com", team.Name), nil)

	err = l.Del(delRequest)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTeam updates a team on the LDAP server.
func (s *LDAPService) UpdateTeam(team types.Team) error {
	// Connect to the LDAP server.
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer l.Close()

	// Bind to the LDAP server.
	err = l.Bind(s.BindDN, s.BindPass)
	if err != nil {
		return err
	}

	// Modify the team.
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=groups,dc=example,dc=com", team.Name))
	modifyRequest.Replace("cn", []string{team.Name})
	modifyRequest.Replace("gidNumber", []string{team.ID})

	err = l.Modify(modifyRequest)
	if err != nil {
		return err
	}

	return nil
}
