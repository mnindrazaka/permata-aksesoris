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
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response utils.Response[[]categories.Category]
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	var expectedData []categories.Category
	data, err = json.Marshal(databases.CategoriesData)
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

func TestCreateCategory(t *testing.T) {
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
		"title": "New Category",
		"slug":  "new-category",
		"icon":  "",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/categories", bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response utils.Response[interface{}]
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response[interface{}]{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(response))

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestUpdateCategory(t *testing.T) {
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
		"title": "New Category",
		"slug":  "new-category",
		"icon":  "",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("PUT", "/categories/"+databases.CategoriesData[0].Serial, bytes.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response[interface{}]
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response[interface{}]{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}

func TestDeleteCategory(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := Authenticate(router)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest("DELETE", "/categories/"+databases.CategoriesData[0].Serial, nil)
	r.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseCategory utils.Response[interface{}]
	if err := json.Unmarshal(data, &responseCategory); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, fmt.Sprint(utils.Response[interface{}]{
		Data:    nil,
		Status:  http.StatusText(http.StatusOK),
		Message: http.StatusText(http.StatusOK),
	}), fmt.Sprint(responseCategory))

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}
