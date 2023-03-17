package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"

	"cum/types"
)

// AWSService represents a AWS service
type AWSService struct {
	iam *iam.IAM
}

// NewAWSService creates a new AWS service
func NewAWSService() *AWSService {
	return &AWSService{
		iam: iam.New(session.New(), aws.NewConfig().WithRegion("us-east-1")),
	}
}

// AddUser adds a user to the service
func (s *AWSService) AddUser(user types.User) error {
	_, err := s.iam.CreateUser(&iam.CreateUserInput{
		UserName: aws.String(user.Email),
	})
	if err != nil {
		return err
	}

	return nil
}

// RemoveUser removes a user from the service
func (s *AWSService) RemoveUser(user types.User) error {
	_, err := s.iam.DeleteUser(&iam.DeleteUserInput{
		UserName: aws.String(user.Email),
	})
	if err != nil {
		return err
	}

	return nil
}

// GetUsers gets all users from the service
func (s *AWSService) GetUsers() ([]types.User, error) {
	users := []types.User{}

	resp, err := s.iam.ListUsers(&iam.ListUsersInput{})
	if err != nil {
		return nil, err
	}

	for _, user := range resp.Users {
		users = append(users, types.User{
			Email: *user.UserName,
		})
	}

	return users, nil
}

// AddTeam adds a team to the service
func (s *AWSService) AddTeam(team types.Team) error {
	return nil
}

// RemoveTeam removes a team from the service
func (s *AWSService) RemoveTeam(team types.Team) error {
	return nil
}

// GetTeams gets all teams from the service
func (s *AWSService) GetTeams() ([]types.Team, error) {
	return nil, nil
}

// GetID gets the service ID
func (s *AWSService) GetID() string {
	return "aws"
}

// GetType gets the service type
func (s *AWSService) GetType() string {
	return "iam"
}

// String returns a string representation of the service
func (s *AWSService) String() string {
	return fmt.Sprintf("%s-%s", s.GetID(), s.GetType())
}
