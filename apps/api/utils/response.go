package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

func NewResponse(data interface{}, status string, err error) Response {
	var res Response

	res.Data = data
	res.Status = status
	res.Message = status

	if err != nil {
		res.Message = err.Error()
	}

	return res
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewResponse(data, http.StatusText(http.StatusOK), nil))
}

func WriteBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(NewResponse(nil, http.StatusText(http.StatusBadRequest), err))
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(NewResponse(nil, http.StatusText(http.StatusInternalServerError), err))
}
