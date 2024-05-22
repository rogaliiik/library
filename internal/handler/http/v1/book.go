package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rogaliiik/library/internal/domain"
)

type bookCreateOutput struct {
	Id int `json:"id"`
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input *domain.Book
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	input.UserId = userId

	bookId, err := h.service.Book.Create(h.ctx, input)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, bookCreateOutput{bookId})
}

func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bookIdParam := chi.URLParam(r, "bookId")
	if bookIdParam == "" {
		sendErrorResponse(w, http.StatusBadRequest, ErrInvalidParameter.Error())
		return
	}

	bookId, err := strconv.Atoi(bookIdParam)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.Book.GetById(h.ctx, bookId, userId)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, book)
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	books, err := h.service.Book.GetAll(h.ctx, userId)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, books)
}

type statusOutput struct {
	Status string `json:"status"`
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input *domain.BookUpdateInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bookIdParam := chi.URLParam(r, "bookId")
	if bookIdParam == "" {
		sendErrorResponse(w, http.StatusBadRequest, ErrInvalidParameter.Error())
		return
	}

	bookId, err := strconv.Atoi(bookIdParam)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Book.Update(h.ctx, bookId, userId, input)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, statusOutput{"ok"})
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bookIdParam := chi.URLParam(r, "bookId")
	if bookIdParam == "" {
		sendErrorResponse(w, http.StatusBadRequest, ErrInvalidParameter.Error())
		return
	}

	bookId, err := strconv.Atoi(bookIdParam)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Book.Delete(h.ctx, bookId, userId)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusOK, statusOutput{"ok"})
}
