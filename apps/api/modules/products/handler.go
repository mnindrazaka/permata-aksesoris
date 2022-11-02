package products

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
	getProducts(http.ResponseWriter, *http.Request)
	getProductDetail(http.ResponseWriter, *http.Request)
	createProduct(http.ResponseWriter, *http.Request)
	updateProduct(http.ResponseWriter, *http.Request)
	deleteProduct(http.ResponseWriter, *http.Request)

	createProductImage(http.ResponseWriter, *http.Request)
	deleteProductImage(http.ResponseWriter, *http.Request)
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

func (handler handler) getProductDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]
	product, err := handler.usecase.getProductDetail(serial)
	if err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}
	utils.WriteSuccessResponse(w, product)
}

func (handler handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest Product
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	if err := handler.usecase.createProduct(productRequest); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, nil)
}

func (handler handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]

	var productRequest Product
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	if err := handler.usecase.updateProduct(serial, productRequest); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, nil)
}

func (handler handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serial := vars["serial"]

	if err := handler.usecase.deleteProduct(serial); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, nil)
}

func (handler handler) createProductImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSerial := vars["productSerial"]

	var productImageRequest ProductImage
	if err := json.NewDecoder(r.Body).Decode(&productImageRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	if err := handler.usecase.createProductImage(productSerial, productImageRequest); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, nil)
}

func (handler handler) deleteProductImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageSerial := vars["imageSerial"]
	if err := handler.usecase.deleteProductImage(imageSerial); err != nil {
		utils.WriteInternalServerErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse(w, nil)
}
