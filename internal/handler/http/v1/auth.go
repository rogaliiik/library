package v1

import (
	"encoding/json"
	"net/http"

	"github.com/rogaliiik/library/internal/domain"
)

type signUpOutput struct {
	Id int `json:"id"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input *domain.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.Auth.CreateUser(h.ctx, input)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusCreated, signUpOutput{userId})
}

type signInOutput struct {
	Token string `json:"token"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input *domain.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Auth.GenerateToken(h.ctx, input.Username, input.Password)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, signInOutput{token})
}
