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

	var response utils.Response
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	var expectedData []map[string]interface{}
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

type ResponseJWT struct {
	Data    map[string]string `json:"data"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
}

func TestCreateCategory(t *testing.T) {
	con, router, err := BeforeEach()
	if err != nil {
		log.Fatal(err)
		return
	}

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
		"title": "New Category",
		"slug":  "new-category",
		"icon":  "",
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

	r = httptest.NewRequest("POST", "/categories", bytes.NewReader(body))
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
		"title": "New Category",
		"slug":  "new-category",
		"icon":  "",
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

	r = httptest.NewRequest("PUT", "/categories/"+databases.CategoriesData[0].Serial, bytes.NewReader(body))
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

	r = httptest.NewRequest("DELETE", "/categories/"+databases.CategoriesData[0].Serial, nil)
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

	if err := AfterEach(con); err != nil {
		log.Fatal(err)
	}
}
