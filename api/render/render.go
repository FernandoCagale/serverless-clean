package render

import (
	"encoding/json"
	"net/http"

	"gitlab.com/FernandoCagale/serverless-clean/api/error"
)

func Response(w http.ResponseWriter, response interface{}, code int) {
	json, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}
	addHeaderDefaults(w, code)
	w.Write(json)
}

func ResponseError(w http.ResponseWriter, response error.ResponseError) {
	json, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, error.AddInternalServerError(err.Error()))
		return
	}
	addHeaderDefaults(w, response.Code)
	w.Write(json)
}

func addHeaderDefaults(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
