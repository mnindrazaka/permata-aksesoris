package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"
	"permata-aksesoris/apps/api/utils"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)
	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response utils.Response
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	expectedProductsData := []products.Product{
		{Serial: databases.ProductsData[0].Serial, Title: databases.ProductsData[0].Title, Slug: databases.ProductsData[0].Slug, Thumbnail: databases.ProductsData[0].Thumbnail, Description: databases.ProductsData[0].Description, CategorySerial: databases.ProductsData[0].CategorySerial, Category: databases.CategoriesData[0], Images: []products.ProductImage{databases.ProductImagesData[0]}},
		{Serial: databases.ProductsData[1].Serial, Title: databases.ProductsData[1].Title, Slug: databases.ProductsData[1].Slug, Thumbnail: databases.ProductsData[1].Thumbnail, Description: databases.ProductsData[1].Description, CategorySerial: databases.ProductsData[1].CategorySerial, Category: databases.CategoriesData[1], Images: []products.ProductImage{databases.ProductImagesData[1]}},
	}

	var expectedData []map[string]interface{}
	data, err = json.Marshal(expectedProductsData)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &expectedData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(expectedData), fmt.Sprint(response.Data))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}

func TestGetProductDetail(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("GET", "/products/"+databases.ProductsData[0].Serial, nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)
	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response utils.Response
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	expectedProductsData := []products.Product{
		{Serial: databases.ProductsData[0].Serial, Title: databases.ProductsData[0].Title, Slug: databases.ProductsData[0].Slug, Thumbnail: databases.ProductsData[0].Thumbnail, Description: databases.ProductsData[0].Description, CategorySerial: databases.ProductsData[0].CategorySerial, Category: databases.CategoriesData[0], Images: []products.ProductImage{databases.ProductImagesData[0]}},
		{Serial: databases.ProductsData[1].Serial, Title: databases.ProductsData[1].Title, Slug: databases.ProductsData[1].Slug, Thumbnail: databases.ProductsData[1].Thumbnail, Description: databases.ProductsData[1].Description, CategorySerial: databases.ProductsData[1].CategorySerial, Category: databases.CategoriesData[1], Images: []products.ProductImage{databases.ProductImagesData[1]}},
	}

	var expectedData map[string]interface{}
	data, err = json.Marshal(expectedProductsData[0])
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &expectedData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(expectedData), fmt.Sprint(response.Data))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}

func TestCreateProduct(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)

	userRepository := users.NewRepository(con)
	userUsecase := users.NewUsecase(userRepository)
	userHandler := users.NewHandler(userUsecase)
	users.NewRouter(userHandler, router)

	reqBody := map[string]interface{}{
		"email":    "admin@gmail.com",
		"password": "admin",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	reqBody = map[string]interface{}{
		"title":          "New Product",
		"slug":           "new-product",
		"thumbnail":      "",
		"description":    "This is a new product",
		"categorySerial": databases.CategoriesData[0].Serial,
	}
	body, err = json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response ResponseJWT
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	r = httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+response.Data["token"])
	w = httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}

func TestUpdateProduct(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)

	userRepository := users.NewRepository(con)
	userUsecase := users.NewUsecase(userRepository)
	userHandler := users.NewHandler(userUsecase)
	users.NewRouter(userHandler, router)

	reqBody := map[string]interface{}{
		"email":    "admin@gmail.com",
		"password": "admin",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	reqBody = map[string]interface{}{
		"title":          "New Product",
		"slug":           "new-product",
		"thumbnail":      "",
		"description":    "This is a new product",
		"categorySerial": databases.CategoriesData[0].Serial,
	}
	body, err = json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response ResponseJWT
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	r = httptest.NewRequest("PUT", "/products/"+databases.ProductsData[0].Serial, bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+response.Data["token"])
	w = httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}

func TestDeleteProduct(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)

	userRepository := users.NewRepository(con)
	userUsecase := users.NewUsecase(userRepository)
	userHandler := users.NewHandler(userUsecase)
	users.NewRouter(userHandler, router)

	reqBody := map[string]interface{}{
		"email":    "admin@gmail.com",
		"password": "admin",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response ResponseJWT
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	r = httptest.NewRequest("DELETE", "/products/"+databases.ProductsData[0].Serial, nil)
	r.Header.Set("Authorization", "Bearer "+response.Data["token"])
	w = httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}

func TestCreateProductImage(t *testing.T) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)

	userRepository := users.NewRepository(con)
	userUsecase := users.NewUsecase(userRepository)
	userHandler := users.NewHandler(userUsecase)
	users.NewRouter(userHandler, router)

	reqBody := map[string]interface{}{
		"email":    "admin@gmail.com",
		"password": "admin",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	reqBody = map[string]interface{}{
		"url": "https://picsum.photos/200/300",
	}
	body, err = json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response ResponseJWT
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	r = httptest.NewRequest("POST", "/products/"+databases.ProductsData[0].Serial+"/images", bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+response.Data["token"])
	w = httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err = ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}
