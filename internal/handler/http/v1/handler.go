package v1

import (
	"context"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/rogaliiik/library/docs"
	"github.com/rogaliiik/library/internal/service"
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

	r.Use(middleware.Recoverer)
	//r.Use(middleware.Logger)
	r.Use(h.requestIdMiddleware)
	r.Use(h.logRequestMiddleware)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", h.signIn)
			r.Post("/sign-up", h.signUp)
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
	})

	return r
}
