package v1

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"

	bearerMethod = "Bearer"
)

func (h *Handler) userVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			sendErrorResponse(w, http.StatusUnauthorized, ErrEmptyAuthHeader.Error())
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != bearerMethod {
			sendErrorResponse(w, http.StatusUnauthorized, ErrInvalidAuthHeader.Error())
			return
		}

		userId, err := h.service.Auth.ParseToken(h.ctx, headerParts[1])
		if err != nil {
			sendErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserId(r *http.Request) (int, error) {
	userIdParam := r.Context().Value(userCtx)
	userId, ok := userIdParam.(int)
	if !ok {
		return 0, ErrInvalidIdType
	}
	return userId, nil
}
