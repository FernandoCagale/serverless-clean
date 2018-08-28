package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

//Task entity
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Validate struct
func (e Task) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(5, 50)),
	)
}
