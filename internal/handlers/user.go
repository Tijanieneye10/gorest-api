package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tijanieneye10/restapi/internal/dtos"
	"github.com/tijanieneye10/restapi/internal/store"
	"github.com/tijanieneye10/restapi/internal/utils"
)

func (h Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dtos.UserDto
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.ErrorResponse(w, "invalid request", err.Error(), http.StatusBadRequest)
		}

		password, err := utils.HashPassword(req.Password)

		user, err := h.Queries.CreateUser(ctx, store.CreateUserParams{
			Username: req.Username,
			Password: password,
			Email:    req.Email,
		})

		if err != nil {
			log.Fatalf("error %v", err.Error())
		}

		utils.SuccessResponse(w, "User Created successfully", user, http.StatusCreated)
	}
}
