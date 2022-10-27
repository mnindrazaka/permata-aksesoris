package products

import (
	"net/http"
	"permata-aksesoris/apps/api/utils"
)

type handler struct {
	usecase Usecase
}

type Handler interface {
	getProducts(http.ResponseWriter, *http.Request)
}

func NewHandler(usecase Usecase) Handler {
	return handler{usecase}
}

func (handler handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.usecase.getProducts()
	if err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse(w, products)
}
