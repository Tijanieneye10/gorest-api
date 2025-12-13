package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tijanieneye10/restapi/internal/dtos"
	"github.com/tijanieneye10/restapi/internal/utils"
)

func (h Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dtos.UserDto
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.ErrorResponse(w, "invalid request", err.Error(), http.StatusBadRequest)
		}

		utils.SuccessResponse(w, "User Created successfully", req, http.StatusCreated)
	}
}
