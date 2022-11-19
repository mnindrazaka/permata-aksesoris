package categories

import (
	"encoding/json"
	"net/http"
	"permata-aksesoris/apps/api/utils"

	"github.com/gorilla/mux"
)

type handler struct {
	usecase Usecase
}

type Handler interface {
	getCategories(http.ResponseWriter, *http.Request)
	createCategory(http.ResponseWriter, *http.Request)
	updateCategory(http.ResponseWriter, *http.Request)
	deleteCategory(http.ResponseWriter, *http.Request)
}

func NewHandler(usecase Usecase) Handler {
	return handler{usecase}
}

func (handler handler) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.usecase.getCategories()
	if err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse(w, categories)
}

func (handler handler) createCategory(w http.ResponseWriter, r *http.Request) {
	var categoryRequest Category
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	if err := handler.usecase.createCategory(categoryRequest); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse[interface{}](w, nil)
}

func (handler handler) updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]

	var categoryRequest Category
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	if err := handler.usecase.updateCategory(serial, categoryRequest); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse[interface{}](w, nil)
}

func (handler handler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]

	if err := handler.usecase.deleteCategory(serial); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse[interface{}](w, nil)
}
