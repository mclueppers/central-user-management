package types

import "errors"

var (
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")

	// ErrUsernameNotFound is returned when a username is not found
	ErrUsernameNotFound = errors.New("username not found")

	// ErrEmailNotFound is returned when an email is not found
	ErrEmailNotFound = errors.New("email not found")

	// ErrUserAlreadyExists is returned when a user already exists
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrGroupNotFound is returned when a group is not found
	ErrGroupNotFound = errors.New("group not found")

	// ErrGroupAlreadyExists is returned when a group already exists
	ErrGroupAlreadyExists = errors.New("group already exists")

	// ErrUserNotInGroup is returned when a user is not in a group
	ErrUserNotInGroup = errors.New("user not in group")

	// ErrGroupNotInGroup is returned when a group is not in a group
	ErrGroupNotInGroup = errors.New("group not in group")

	// ErrSessionNotFound is returned when a session is not found
	ErrSessionNotFound = errors.New("session not found")

	// ErrSessionAlreadyExists is returned when a session already exists
	ErrSessionAlreadyExists = errors.New("session already exists")

	// ErrSessionExpired is returned when a session is expired
	ErrSessionExpired = errors.New("session expired")

	// ErrSessionInvalid is returned when a session is invalid
	ErrSessionInvalid = errors.New("session invalid")

	// ErrInvalidMemberType is returned when a member is invalid
	ErrInvalidMemberType = errors.New("invalid member type")

	// ErrMemberNotFound is returned when a member is not found
	ErrMemberNotFound = errors.New("member not found")
)
