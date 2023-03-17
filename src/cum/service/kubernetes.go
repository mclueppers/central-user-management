package service

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"cum/types"
)

// KubernetesService represents a Kubernetes service.
type KubernetesService struct {
	clientset *kubernetes.Clientset
}

// NewKubernetesService creates a new Kubernetes service.
func NewKubernetesService() (*KubernetesService, error) {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubernetesService{
		clientset: clientset,
	}, nil
}

// AddUser adds a user to the service.
func (s *KubernetesService) AddUser(user types.User) error {
	return nil
}

// RemoveUser removes a user from the service.
func (s *KubernetesService) RemoveUser(user types.User) error {
	return nil
}

// GetUsers gets all users from the service.
func (s *KubernetesService) GetUsers() ([]types.User, error) {
	return nil, nil
}

// AddTeam adds a team to the service.
func (s *KubernetesService) AddTeam(team types.Team) error {
	return nil
}

// RemoveTeam removes a team from the service.
func (s *KubernetesService) RemoveTeam(team types.Team) error {
	return nil
}

// GetTeams gets all teams from the service.
func (s *KubernetesService) GetTeams() ([]types.Team, error) {
	return nil, nil
}

// GetID gets the ID of the service.
func (s *KubernetesService) GetID() string {
	return ""
}

// GetType gets the type of the service.
func (s *KubernetesService) GetType() string {
	return ""
}

// String gets the string representation of the service.
func (s *KubernetesService) String() string {
	return ""
}
