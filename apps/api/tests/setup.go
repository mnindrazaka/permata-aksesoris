package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules"
	"permata-aksesoris/apps/api/utils"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func BeforeEach() (*gorm.DB, http.HandlerFunc, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return nil, nil, err
	}

	con, err := databases.NewTestDBConnection()
	if err != nil {
		return nil, nil, err
	}

	if err := databases.Migrate(con); err != nil {
		return nil, nil, err
	}

	if err := databases.Unseed(con); err != nil {
		return nil, nil, err
	}

	if err := databases.Seed(con); err != nil {
		return nil, nil, err
	}

	router := modules.NewApp(con)

	return con, router, nil
}

func AfterEach(con *gorm.DB) error {
	return databases.Unseed(con)
}

func Authenticate(router http.HandlerFunc) (string, error) {
	reqBody := map[string]interface{}{
		"email":    "admin@gmail.com",
		"password": "admin",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	r := httptest.NewRequest("POST", "/users/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		return "", err
	}

	var response utils.Response[map[string]string]
	if err := json.Unmarshal(data, &response); err != nil {
		return "", err
	}

	return response.Data["token"], nil
}
