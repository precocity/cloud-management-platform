package request

import "cloud-management-platform/pkg/resources"

//Service To Manage Resource Requests
type Service struct {
	repo *Repository
}

//CreateService creates and instance of the Resource Request Service
func CreateService(repo *Repository) *Service {

	s := &Service{
		repo: repo,
	}

	return s
}

//CreateRequest will create a request in the system
func (s Service) CreateRequest(name string, description string, resourceType resources.ResourceType) error {

	//TODO: Send Resource Request To Datastore

	//TODO: Send Email to Supervisor To Review Request

	return nil
}
