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
	"permata-aksesoris/apps/api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestGetProductDetail(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := httptest.NewRequest("GET", "/products/"+databases.ProductsData[0].Serial, nil)
	w := httptest.NewRecorder()
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestCreateProduct(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := map[string]interface{}{
		"title":          "New Product",
		"slug":           "new-product",
		"thumbnail":      "",
		"description":    "This is a new product",
		"categorySerial": databases.CategoriesData[0].Serial,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestUpdateProduct(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := map[string]interface{}{
		"title":          "New Product",
		"slug":           "new-product",
		"thumbnail":      "",
		"description":    "This is a new product",
		"categorySerial": databases.CategoriesData[0].Serial,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("PUT", "/products/"+databases.ProductsData[0].Serial, bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestDeleteProduct(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("DELETE", "/products/"+databases.ProductsData[0].Serial, nil)
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestCreateProductImage(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := map[string]interface{}{
		"url": "https://picsum.photos/200/300",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/products/"+databases.ProductsData[0].Serial+"/images", bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestDeleteProductImage(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("DELETE", "/products/"+databases.ProductsData[0].Serial+"/images/"+databases.ProductImagesData[0].Serial, nil)
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}
