package infra

import (
	"net/http"

	"github.com/gorilla/mux"
	errors "gitlab.com/FernandoCagale/serverless-clean/api/error"
	"gitlab.com/FernandoCagale/serverless-clean/api/render"
)

func methodNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddNotFoundError("Not Found"))
	})
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddMethodNotAllowedError("Method Not Allowed"))
	})
}

//NewRouter infra
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = methodNotFoundHandler()
	router.MethodNotAllowedHandler = methodNotAllowedHandler()
	return router
}
