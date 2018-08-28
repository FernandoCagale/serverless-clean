package repository

import (
	"errors"
	"strconv"

	"gitlab.com/FernandoCagale/serverless-clean/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.Task
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Task{}
	return &IRepo{
		m: m,
	}
}

//Create a Task
func (r *IRepo) Create(e *entity.Task) error {
	if r.m[strconv.Itoa(e.ID)] != nil {
		return errors.New("exist")
	}
	r.m[strconv.Itoa(e.ID)] = e
	return nil
}

//Update a Task
func (r *IRepo) Update(id int, e *entity.Task) error {
	if r.m[strconv.Itoa(id)] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(id)] = e
	return nil
}

//Delete a Task
func (r *IRepo) Delete(id int) error {
	if r.m[strconv.Itoa(id)] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(id)] = nil
	return nil
}

//FindByID a Task
func (r *IRepo) FindByID(id int) (task *entity.Task, err error) {
	if r.m[strconv.Itoa(id)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(id)], nil
}

//FindAll a Task
func (r *IRepo) FindAll() (tasks []*entity.Task, err error) {
	for _, task := range r.m {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
