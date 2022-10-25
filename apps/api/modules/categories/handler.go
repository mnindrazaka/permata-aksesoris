package categories

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	usecase Usecase
}

type Handler interface {
	getCategories(http.ResponseWriter, *http.Request)
}

func NewHandler(usecase Usecase) Handler {
	return handler{usecase}
}

func (handler handler) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.usecase.getCategories()
	if err != nil {
		w.Write([]byte("Failed to get data"))
		return
	}
	json.NewEncoder(w).Encode(categories)
}
