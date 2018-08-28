package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidPayload Invalid payload
var ErrInvalidPayload = errors.New("Invalid payload")
