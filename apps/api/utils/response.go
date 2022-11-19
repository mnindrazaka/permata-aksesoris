package utils

import (
	"encoding/json"
	"net/http"
)

type Response[T interface{}] struct {
	Data    T      `json:"data"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewResponse[T interface{}](data T, status string, err error) Response[T] {
	var res Response[T]

	res.Data = data
	res.Status = status
	res.Message = status

	if err != nil {
		res.Message = err.Error()
	}

	return res
}

func WriteSuccessResponse[T interface{}](w http.ResponseWriter, data T) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewResponse(data, http.StatusText(http.StatusOK), nil))
}

func WriteBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(NewResponse[interface{}](nil, http.StatusText(http.StatusBadRequest), err))
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(NewResponse[interface{}](nil, http.StatusText(http.StatusInternalServerError), err))
}

func WriteUnauthorizedResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(NewResponse[interface{}](nil, http.StatusText(http.StatusUnauthorized), err))
}
