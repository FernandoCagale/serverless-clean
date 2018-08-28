package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/datastore"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/entity"
)

//GormRepository in memory repo
type GormRepository struct {
	connection string
}

//NewGormRepository create new repository
func NewGormRepository(connection string) *GormRepository {
	return &GormRepository{connection}
}

//Create a Task
func (r *GormRepository) Create(e *entity.Task) error {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return err
	}

	defer db.Close()

	return db.Create(e).Error
}

//Update a Task
func (r *GormRepository) Update(id int, e *entity.Task) error {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return err
	}

	defer db.Close()

	var task entity.Task

	if err := db.First(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return entity.ErrNotFound
		default:
			return err
		}
	}

	return db.Model(&task).Update("name", e.Name).Error
}

//Delete a Task
func (r *GormRepository) Delete(id int) error {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return err
	}

	defer db.Close()

	var task entity.Task

	if err := db.First(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return entity.ErrNotFound
		default:
			return err
		}
	}
	return db.Delete(task).Error
}

//FindByID a Task
func (r *GormRepository) FindByID(id int) (*entity.Task, error) {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var task entity.Task

	if err := db.First(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, entity.ErrNotFound
		default:
			return nil, err
		}
	}

	return &task, nil
}

//FindAll a Task
func (r *GormRepository) FindAll() (tasks []*entity.Task, err error) {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Find(&tasks).Error
	return tasks, err
}
