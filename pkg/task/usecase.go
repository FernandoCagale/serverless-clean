package task

import "gitlab.com/FernandoCagale/serverless-clean/pkg/entity"

//UseCase use case interface
type UseCase interface {
	Create(task *entity.Task) (err error)
	Update(id int, task *entity.Task) (err error)
	Delete(id int) (err error)
	FindByID(id int) (task *entity.Task, err error)
	FindAll() (tasks []*entity.Task, err error)
}
