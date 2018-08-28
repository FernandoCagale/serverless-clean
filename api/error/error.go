package error

import (
	"net/http"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AddInternalServerError(message string) ResponseError {
	return ResponseError{http.StatusInternalServerError, message}
}

func AddNotFoundError(message string) ResponseError {
	return ResponseError{http.StatusNotFound, message}
}

func AddUnauthorizedError(message string) ResponseError {
	return ResponseError{http.StatusUnauthorized, message}
}

func AddBadRequestError(message string) ResponseError {
	return ResponseError{http.StatusBadRequest, message}
}

func AddMethodNotAllowedError(message string) ResponseError {
	return ResponseError{http.StatusMethodNotAllowed, message}
}
