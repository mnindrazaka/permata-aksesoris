package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules/products"
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

	databases.ProductsData[0].Category = databases.CategoriesData[0]
	databases.ProductsData[1].Category = databases.CategoriesData[1]

	databases.ProductsData[0].Images = []products.ProductImage{databases.ProductImagesData[0]}
	databases.ProductsData[1].Images = []products.ProductImage{databases.ProductImagesData[1]}

	var expectedData []map[string]interface{}
	data, err = json.Marshal(databases.ProductsData)
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
