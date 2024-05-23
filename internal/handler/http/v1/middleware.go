package v1

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	requestIdCtx        = "X-Request-Id"

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

func (h *Handler) requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIdCtx, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type loggingResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (rw *loggingResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (h *Handler) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := h.log.With(slog.String("ID", r.Context().Value(requestIdCtx).(string)))
		log.Debug("Received request", slog.Any("method", r.Method), slog.Any("URL", r.RequestURI))

		now := time.Now()
		rw := &loggingResponseWriter{statusCode: http.StatusOK, ResponseWriter: w}
		next.ServeHTTP(rw, r)

		log.Debug("Handled request ", slog.Any("Method", r.Method), slog.Any("URL", r.RequestURI),
			slog.Any("Delay", time.Since(now)), slog.Any("Status code", rw.statusCode))
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
