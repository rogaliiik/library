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

// @Summary Book Create
// @Security UserAuth
// @Tags bookApi
// @Description book create
// @ModuleID createBook
// @Accept  json
// @Produce  json
// @Param input body domain.Book true "book info"
// @Success 200 {object} bookCreateOutput
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/book/create [post]
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

// @Summary Book Get by id
// @Security UserAuth
// @Tags bookApi
// @Description book get by id
// @ModuleID getBookById
// @Accept  json
// @Produce  json
// @Param bookId path int true "book id"
// @Success 200 {object} domain.Book
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/book/{bookId} [get]
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

// @Summary Book Get all
// @Security UserAuth
// @Tags bookApi
// @Description get all books
// @ModuleID getAllBooks
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.Book
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/book [get]
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

// @Summary Book Update
// @Security UserAuth
// @Tags bookApi
// @Description update book
// @ModuleID updateBook
// @Accept  json
// @Produce  json
// @Param input body domain.BookUpdateInput true "book update info"
// @Param bookId path int true "book id"
// @Success 200 {object} statusOutput
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/book/{bookId} [put]
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

// @Summary Book Delete
// @Security UserAuth
// @Tags bookApi
// @Description delete book
// @ModuleID deleteBook
// @Accept  json
// @Produce  json
// @Param bookId path int true "book id"
// @Success 200 {object} statusOutput
// @Failure 400 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Router /v1/book/{bookId} [delete]
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
