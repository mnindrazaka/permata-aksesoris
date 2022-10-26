package products

import (
	"encoding/json"
	"net/http"
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
		w.Write([]byte("Failed to get data"))
		return
	}
	json.NewEncoder(w).Encode(products)
}