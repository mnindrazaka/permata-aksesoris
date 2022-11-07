package users

import (
	"encoding/json"
	"net/http"
	"permata-aksesoris/apps/api/utils"
)

type handler struct {
	usecase Usecase
}

type Handler interface {
	login(w http.ResponseWriter, r *http.Request)
}

func NewHandler(usecase Usecase) Handler {
	return handler{usecase}
}

func (handler handler) login(w http.ResponseWriter, r *http.Request) {
	var userRequest User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.WriteBadRequestResponse(w, err)
		return
	}

	token, err := handler.usecase.login(userRequest.Email, userRequest.Password)
	if err != nil {
		utils.WriteUnauthorizedResponse(w, err)
		return
	}

	response := map[string]string{"token": token}
	utils.WriteSuccessResponse(w, response)
}
