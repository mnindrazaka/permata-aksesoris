package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/utils"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
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

	r := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(true)
	categoryRepository := categories.NewRepository(con)
	categoryUsecase := categories.NewUsecase(categoryRepository)
	categoryHandler := categories.NewHandler(categoryUsecase)
	categories.NewRouter(categoryHandler, router)
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

	if err := databases.Unseed(con); err != nil {
		log.Fatal(err)
	}
}
