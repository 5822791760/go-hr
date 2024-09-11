package https

import (
	"net/http"

	"github.com/5822791760/hr/internal/usecases"
	"github.com/5822791760/hr/pkg/errs"
	"github.com/5822791760/hr/pkg/helpers"
)

type AuthorHandler struct {
	db            helpers.Transactionable
	authorUsecase usecases.IAuthorUsecase
}

func NewAuthorHandler(db helpers.Transactionable, authorUsecase usecases.IAuthorUsecase) AuthorHandler {
	return AuthorHandler{
		db:            db,
		authorUsecase: authorUsecase,
	}
}

func (h AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	ctx, end, err := GetTxContext(r, h.db)

	defer func() {
		WriteError(w, end(err))
	}()

	var body usecases.CreateAuthorBody
	if err := ParseBody(r, &body); err != nil {
		return
	}

	author, err := h.authorUsecase.Create(ctx, body)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusCreated, author)
}

func (h AuthorHandler) QueryAuthors(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	ctx := GetContext(r, h.db)

	defer func() {
		WriteError(w, err)
	}()

	resp, err := h.authorUsecase.QueryGetAll(ctx)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h AuthorHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	ctx := GetContext(r, h.db)

	defer func() {
		WriteError(w, err)
	}()

	id, err := GetParamInt(r, "id")
	if err != nil {
		return
	}

	res, err := h.authorUsecase.FindOne(ctx, id)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusOK, res)
}

func (h AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	ctx, end, err := GetTxContext(r, h.db)

	defer func() {
		WriteError(w, end(err))
	}()

	id, err := GetParamInt(r, "id")
	if err != nil {
		return
	}

	var body usecases.UpdateAuthorBody
	if err := ParseBody(r, &body); err != nil {
		return
	}

	res, err := h.authorUsecase.Update(ctx, id, body)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusOK, res)
}

func (h AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	ctx, end, err := GetTxContext(r, h.db)

	defer func() {
		WriteError(w, end(err))
	}()

	id, err := GetParamInt(r, "id")
	if err != nil {
		return
	}

	res, err := h.authorUsecase.Delete(ctx, id)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusOK, res)
}
