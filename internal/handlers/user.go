package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tijanieneye10/restapi/internal/dtos"
	"github.com/tijanieneye10/restapi/internal/store"
	"github.com/tijanieneye10/restapi/internal/utils"
)

func (h Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dtos.UserDto

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			utils.ErrorResponse(w, "invalid credentials", err.Error(), http.StatusBadRequest)
		}

		user, err := h.Queries.GetUserByEmailOrUsername(ctx, req.Email)
		if err != nil {
			return
		}

		isMatched := utils.CheckPasswordHash(user.Password, req.Password)
		if !isMatched {
			utils.ErrorResponse(w, "invalid credentials", fmt.Sprintf("invalid credentials: %s", req.Email), http.StatusUnauthorized)
		}

		token, _ := utils.CreateToken(int64(user.ID), user.Username, []byte(os.Getenv("SECRET")))

		utils.SuccessResponse(w, "Login successful",
			map[string]any{
				"token": token,
				"user":  user,
			}, http.StatusOK)
	}
}

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
			utils.ErrorResponse(w, "invalid request", err.Error(), http.StatusBadRequest)
			return
		}

		utils.SuccessResponse(w, "User Created successfully", user, http.StatusCreated)
	}
}
