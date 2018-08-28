package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//NewPostgres create connection
func NewPostgres(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.LogMode(false)
	db.SingularTable(true)

	return db, nil
}
