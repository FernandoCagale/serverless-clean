package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/FernandoCagale/serverless-clean/api/error"
	"gitlab.com/FernandoCagale/serverless-clean/api/render"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/entity"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/logger"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/task"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/task/repository"
	"gitlab.com/FernandoCagale/serverless-clean/pkg/task/usecase"
)

//MakeTaskHandlers make url handlers
func MakeTaskHandlers(r *mux.Router, service task.UseCase) {
	r.Handle("/v1/api/task/{id}", findByID(service)).Methods("GET")
	r.Handle("/v1/api/task/{id}", updateByID(service)).Methods("PUT")
	r.Handle("/v1/api/task/{id}", deleteByID(service)).Methods("DELETE")
	r.Handle("/v1/api/task", findAll(service)).Methods("GET")
	r.Handle("/v1/api/task", create(service)).Methods("POST")
}

//MakeTaskGorm database postgres
func MakeTaskGorm() task.UseCase {
	return usecase.NewService(repository.NewGormRepository(os.Getenv("DATASTORE_URL")))
}

//MakeTaskInmemory database memory
func MakeTaskInmemory() task.UseCase {
	return usecase.NewService(repository.NewInmemRepository())
}

func create(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var task *entity.Task

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			render.ResponseError(w, error.AddBadRequestError("Invalid request payload"))
			return
		}

		logger.WithFields(logger.Fields{
			"ID":   task.ID,
			"Name": task.Name,
		}).Info("create")

		defer r.Body.Close()

		if err := service.Create(task); err != nil {
			switch err {
			case entity.ErrInvalidPayload:
				render.ResponseError(w, error.AddBadRequestError(err.Error()))
			default:
				render.ResponseError(w, error.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, task, http.StatusCreated)
	})
}

func updateByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
			return
		}

		var task entity.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			render.ResponseError(w, error.AddBadRequestError("Invalid request payload"))
			return
		}

		logger.WithFields(logger.Fields{
			"ID":   id,
			"Name": task.Name,
		}).Info("updateByID")

		defer r.Body.Close()

		err = service.Update(id, &task)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, error.AddNotFoundError(err.Error()))
			case entity.ErrInvalidPayload:
				render.ResponseError(w, error.AddBadRequestError(err.Error()))
			default:
				render.ResponseError(w, error.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, map[string]string{"updated": "true"}, http.StatusOK)
	})
}

func deleteByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
			return
		}

		err = service.Delete(id)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, error.AddNotFoundError(err.Error()))
			default:
				render.ResponseError(w, error.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, map[string]string{"deleted": "true"}, http.StatusOK)
	})
}

func findByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, error.AddBadRequestError("Invalid task ID"))
			return
		}

		task, err := service.FindByID(id)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, error.AddNotFoundError(err.Error()))
			default:
				render.ResponseError(w, error.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, task, http.StatusOK)
	})
}

func findAll(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tasks, err := service.FindAll()
		if err != nil {
			render.ResponseError(w, error.AddInternalServerError(err.Error()))
			return
		}

		render.Response(w, tasks, http.StatusOK)
	})
}
