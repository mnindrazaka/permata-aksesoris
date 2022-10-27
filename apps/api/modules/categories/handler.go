package categories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	usecase Usecase
}

type Handler interface {
	getCategories(http.ResponseWriter, *http.Request)
	createCategory(http.ResponseWriter, *http.Request)
	updateCategory(http.ResponseWriter, *http.Request)
}

func NewHandler(usecase Usecase) Handler {
	return handler{usecase}
}

func (handler handler) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.usecase.getCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (handler handler) createCategory(w http.ResponseWriter, r *http.Request) {
	var categoryRequest Category
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := handler.usecase.createCategory(categoryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (handler handler) updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]

	var categoryRequest Category
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := handler.usecase.updateCategory(serial, categoryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}
