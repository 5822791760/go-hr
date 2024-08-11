package handlers

import (
	"net/http"

	"github.com/5822791760/hr/internal/usecases"
)

type AuthorHandler struct {
	authorUsecase usecases.IAuthorUsecase
}

func NewAuthorHandler(authorService usecases.IAuthorUsecase) AuthorHandler {
	return AuthorHandler{
		authorUsecase: authorService,
	}
}

func (h AuthorHandler) QueryAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.authorUsecase.QueryGetAll(ctx)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h AuthorHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := GetParamInt(r, "id")
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.authorUsecase.FindOne(ctx, id)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteJSON(w, http.StatusOK, res)
}
