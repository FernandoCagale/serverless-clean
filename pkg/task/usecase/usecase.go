package usecase

import (
	"gitlab.com/FernandoCagale/serverless-clean/pkg/entity"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/task"
)

//Service struct
type Service struct {
	repo task.Repository
}

//NewService new service
func NewService(r task.Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create a Task
func (s *Service) Create(e *entity.Task) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidPayload
	}
	return s.repo.Create(e)
}

//Update a Task
func (s *Service) Update(id int, e *entity.Task) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidPayload
	}
	return s.repo.Update(id, e)
}

//Delete a Task
func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

//FindByID a Task
func (s *Service) FindByID(id int) (task *entity.Task, err error) {
	return s.repo.FindByID(id)
}

//FindAll a Task
func (s *Service) FindAll() (tasks []*entity.Task, err error) {
	return s.repo.FindAll()
}
