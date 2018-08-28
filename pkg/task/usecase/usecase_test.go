package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/entity"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/task/repository"
)

var (
	repo    *repository.IRepo
	service *Service
)

func init() {
	repo = repository.NewInmemRepository()
	service = NewService(repo)
}

func TestCreate(t *testing.T) {
	entity := &entity.Task{
		ID:   1,
		Name: "Fernando",
	}

	err := service.Create(entity)
	assert.Nil(t, err)
}

func TestCreateInvalidPayload(t *testing.T) {
	entity := &entity.Task{
		ID:   1,
		Name: "Fer",
	}

	err := service.Create(entity)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	entity := &entity.Task{
		ID:   1,
		Name: "Fernando Cagale",
	}

	err := service.Update(1, entity)
	assert.Nil(t, err)
}

func TestUpdateInvalidPayload(t *testing.T) {
	entity := &entity.Task{
		ID:   1,
		Name: "Fer",
	}

	err := service.Create(entity)
	assert.NotNil(t, err)
}

func TestUpdateNotFound(t *testing.T) {
	entity := &entity.Task{
		ID:   99,
		Name: "Invalid",
	}

	err := service.Update(99, entity)
	assert.NotNil(t, err)
}

func TestFindByID(t *testing.T) {
	task, err := service.FindByID(1)
	assert.Nil(t, err)
	assert.Equal(t, task.ID, 1)
	assert.Equal(t, task.Name, "Fernando Cagale")
}

func TestFindByIDNotFound(t *testing.T) {
	task, err := service.FindByID(99)
	assert.NotNil(t, err)
	assert.Nil(t, task)
}

func TestFindAll(t *testing.T) {
	tasks, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, tasks[0].ID, 1)
	assert.Equal(t, tasks[0].Name, "Fernando Cagale")
}

func TestDelete(t *testing.T) {
	err := service.Delete(1)
	assert.Nil(t, err)
}

func TestDeleteNotFound(t *testing.T) {
	err := service.Delete(99)
	assert.NotNil(t, err)
}
