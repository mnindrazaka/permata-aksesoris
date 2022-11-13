package categories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/utils"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	r := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	categoryRepository := NewRepository(con)
	categoryUsecase := NewUsecase(categoryRepository)
	categoryHandler := NewHandler(categoryUsecase)
	NewRouter(categoryHandler, router)
	router.ServeHTTP(w, r)

	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	var response utils.Response
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	expectedData := []map[string]interface{}{
		{
			"serial": "CAT-1de488c4-fe14-4b3d-ac17-23dd8fe6383f",
			"title":  "Pencil",
			"slug":   "pencil",
			"icon":   "",
		},
		{
			"serial": "CAT-55f49826-0869-4475-8342-def5cf59f8a7",
			"title":  "tas",
			"slug":   "tas",
			"icon":   "",
		},
		{
			"serial": "CAT-6dc8366b-2499-4dcd-a50b-6bfa6ab7f949",
			"title":  "pita-pita",
			"slug":   "pita-pita",
			"icon":   "",
		},
		{
			"serial": "CAT-bde3d484-04ff-48cb-b339-a0c582684372",
			"title":  "jam tangan",
			"slug":   "jam-tangan",
			"icon":   "",
		},
		{
			"serial": "CAT-DZsZDebw",
			"title":  "kacamata",
			"slug":   "kacamata",
			"icon":   "",
		},
	}

	assert.Equal(t, fmt.Sprint(expectedData), fmt.Sprint(response.Data))
}
