package v1

import (
	"encoding/json"
	"net/http"

	"github.com/rogaliiik/library/internal/domain"
)

type signUpOutput struct {
	Id int `json:"id"`
}

// @Summary User SignUp
// @Tags user-auth
// @Description user sign up
// @ModuleID signUp
// @Accept  json
// @Produce  json
// @Param input body domain.User true "user info"
// @Success 201 {object} signUpOutput
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/auth/sign-up [post]
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

// @Summary User SignIn
// @Tags user-auth
// @Description user sign in
// @ModuleID signIn
// @Accept  json
// @Produce  json
// @Param input body domain.User true "user info"
// @Success 200 {object} signInOutput
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/auth/sign-in [post]
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
