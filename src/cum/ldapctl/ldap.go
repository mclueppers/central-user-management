package ldapctl

// Module to manage LDAP groups and users
// This module is used to manage LDAP groups and users
// It is used to create, delete, modify and search LDAP groups and users
// It is also used to add and remove users from groups

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

// LDAP server connection
var l *ldap.Conn

// LDAP server connection error
var err error

// LDAP server address
var ldapServer string

// LDAP server port
var ldapPort string

// LDAP server base DN
var ldapBaseDN string

// LDAP server bind DN
var ldapBindDN string

// LDAP server bind password
var ldapBindPassword string

// LDAP server user search filter
var ldapUserSearchFilter string

// LDAP server group search filter
var ldapGroupSearchFilter string

// LDAP server user search base DN
var ldapUserSearchBaseDN string

// LDAP server group search base DN
var ldapGroupSearchBaseDN string

// LDAP server user search scope
var ldapUserSearchScope string

// LDAP server group search scope
var ldapGroupSearchScope string

// LDAP server user search attributes
var ldapUserSearchAttributes string

// LDAP server group search attributes
var ldapGroupSearchAttributes string

// LDAP server user search attributes
var ldapUserSearchAttributesArray []string

// LDAP server group search attributes
var ldapGroupSearchAttributesArray []string

// LDAP server user search scope
var ldapUserSearchScopeInt int

// LDAP server group search scope
var ldapGroupSearchScopeInt int

const (
	ldapUserDefaultPassword = "Cumulus"
)

// LDAP server connection
func ldapConnect() {
	// Connect to LDAP server
	l, err = ldap.Dial("tcp", ldapServer+":"+ldapPort)
	if err != nil {
		log.Fatal(err)
	}

	// Bind to LDAP server
	err = l.Bind(ldapBindDN, ldapBindPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// LDAP server disconnection
func ldapDisconnect() {
	l.Close()
}

// LDAP server user search
func ldapUserSearch(user string) *ldap.SearchResult {
	// Connect to LDAP server
	ldapConnect()

	// Search for the given user
	searchRequest := ldap.NewSearchRequest(
		ldapUserSearchBaseDN,
		ldapUserSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(ldapUserSearchFilter, user),
		ldapUserSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	return sr
}

// LDAP server group search
func ldapGroupSearch(group string) *ldap.SearchResult {
	// Connect to LDAP server
	ldapConnect()

	// Search for the given group
	searchRequest := ldap.NewSearchRequest(
		ldapGroupSearchBaseDN,
		ldapGroupSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(ldapGroupSearchFilter, group),
		ldapGroupSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	return sr
}

// LDAP server user add
func ldapUserAdd(user string, password string) {
	// Connect to LDAP server
	ldapConnect()

	// Add the given user
	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), nil)
	addRequest.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"})
	addRequest.Attribute("cn", []string{user})
	addRequest.Attribute("sn", []string{user})
	addRequest.Attribute("userPassword", []string{password})

	err = l.Add(addRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user delete
func ldapUserDelete(user string) {
	// Connect to LDAP server
	ldapConnect()

	// Delete the given user
	delRequest := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), nil)

	err = l.Del(delRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user modify
func ldapUserModify(user string, password string) {
	// Connect to LDAP server
	ldapConnect()

	// Modify the given user
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), nil)
	modifyRequest.Replace("userPassword", []string{password})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server group add
func ldapGroupAdd(group string) {
	// Connect to LDAP server
	ldapConnect()

	// Add the given group
	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)
	addRequest.Attribute("objectClass", []string{"top", "posixGroup"})
	addRequest.Attribute("cn", []string{group})
	addRequest.Attribute("gidNumber", []string{"1000"})

	err = l.Add(addRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server group delete
func ldapGroupDelete(group string) {
	// Connect to LDAP server
	ldapConnect()

	// Delete the given group
	delRequest := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)

	err = l.Del(delRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user add to group
func ldapUserAddToGroup(user string, group string) {
	// Connect to LDAP server
	ldapConnect()

	// Add the given user to the given group
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)
	modifyRequest.Add("memberUid", []string{user})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user delete from group
func ldapUserDeleteFromGroup(user string, group string) {
	// Connect to LDAP server
	ldapConnect()

	// Delete the given user from the given group
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)
	modifyRequest.Delete("memberUid", []string{user})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user check
func ldapUserCheck(user string, password string) bool {
	// Connect to LDAP server
	ldapConnect()

	// Check the given user
	err = l.Bind(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), password)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	return true
}

// LDAP server group check
func ldapGroupCheck(user string, group string) bool {
	// Connect to LDAP server
	ldapConnect()

	// Check the given group
	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN),
		ldapGroupSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(ldapGroupSearchFilter, user),
		ldapGroupSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	if len(sr.Entries) > 0 {
		return true
	}

	return false
}

// LDAP server user list
func ldapUserList() []string {
	// Connect to LDAP server
	ldapConnect()

	// List all users
	searchRequest := ldap.NewSearchRequest(
		ldapUserSearchBaseDN,
		ldapUserSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		ldapUserSearchFilter,
		ldapUserSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	var users []string

	for _, entry := range sr.Entries {
		users = append(users, entry.GetAttributeValue("cn"))
	}

	return users
}

// LDAP server group list
func ldapGroupList() []string {
	// Connect to LDAP server
	ldapConnect()

	// List all groups
	searchRequest := ldap.NewSearchRequest(
		ldapGroupSearchBaseDN,
		ldapGroupSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		ldapGroupSearchFilter,
		ldapGroupSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	var groups []string

	for _, entry := range sr.Entries {
		groups = append(groups, entry.GetAttributeValue("cn"))
	}

	return groups
}

// LDAP server user list from group
func ldapUserListFromGroup(group string) []string {
	// Connect to LDAP server
	ldapConnect()

	// List all users from the given group
	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN),
		ldapGroupSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		ldapGroupSearchFilter,
		ldapGroupSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	var users []string

	for _, entry := range sr.Entries {
		users = append(users, entry.GetAttributeValue("memberUid"))
	}

	return users
}

// LDAP server group list from user
func ldapGroupListFromUser(user string) []string {
	// Connect to LDAP server
	ldapConnect()

	// List all groups from the given user
	searchRequest := ldap.NewSearchRequest(
		ldapGroupSearchBaseDN,
		ldapGroupSearchScopeInt,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(ldapGroupSearchFilter, user),
		ldapGroupSearchAttributesArray,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	var groups []string

	for _, entry := range sr.Entries {
		groups = append(groups, entry.GetAttributeValue("cn"))
	}

	return groups
}

// LDAP server user password change
func ldapUserPasswordChange(user string, password string) {
	// Connect to LDAP server
	ldapConnect()

	// Change the given user password
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), nil)
	modifyRequest.Replace("userPassword", []string{password})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user password reset
func ldapUserPasswordReset(user string) {
	// Connect to LDAP server
	ldapConnect()

	// Reset the given user password
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), nil)
	modifyRequest.Replace("userPassword", []string{ldapUserDefaultPassword})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server user password check
func ldapUserPasswordCheck(user string, password string) bool {
	// Connect to LDAP server
	ldapConnect()

	// Check the given user password
	err = l.Bind(fmt.Sprintf("cn=%s,%s", user, ldapUserSearchBaseDN), password)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()

	return true
}

// LDAP server group add user
func ldapGroupAddUser(group string, user string) {
	// Connect to LDAP server
	ldapConnect()

	// Add the given user to the given group
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)
	modifyRequest.Add("memberUid", []string{user})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}

// LDAP server group delete user
func ldapGroupDeleteUser(group string, user string) {
	// Connect to LDAP server
	ldapConnect()

	// Delete the given user from the given group
	modifyRequest := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", group, ldapGroupSearchBaseDN), nil)
	modifyRequest.Delete("memberUid", []string{user})

	err = l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from LDAP server
	ldapDisconnect()
}
