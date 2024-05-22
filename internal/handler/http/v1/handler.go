package v1

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/rogaliiik/library/internal/service"
	"log/slog"
)

type Handler struct {
	ctx     context.Context
	service *service.Service
	log     *slog.Logger
}

func NewHandler(ctx context.Context, service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		ctx:     ctx,
		service: service,
		log:     log,
	}
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", h.signIn)
			r.Post("/sign-up", h.signUp)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(h.userVerifyMiddleware)
		r.Route("/book", func(r chi.Router) {
			r.Post("/", h.createBook)
			r.Get("/{bookId}", h.getBookById)
			r.Get("/", h.getAllBooks)
			r.Put("/{bookId}", h.updateBook)
			r.Delete("/{bookId}", h.deleteBook)
		})
	})

	return r
}
