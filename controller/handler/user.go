package handler

import (
	"github.com/yutify/gengo-api/domain/model"
	"github.com/yutify/gengo-api/usecase/service"
	"log"
	"net/http"
)

type userHandler struct {
	UserService service.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(ua service.UserService) UserHandler {
	return &userHandler{
		UserService: ua,
	}
}

type RegisterUserRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var ur RegisterUserRequest
	if err := decodeRequest(r.Body, &ur); err != nil {
		log.Fatal(err)
	}
	u := &model.User{
		ID:      ur.ID,
		Name:    ur.Name,
		IconURL: ur.IconURL,
	}
	if err := uh.UserService.CreateUser(ctx, u); err != nil {
		log.Fatal(err)
	}
	rendering.JSON(w, http.StatusOK, "ok.")
}
