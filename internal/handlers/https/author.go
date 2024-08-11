package https

import (
	"net/http"

	"github.com/5822791760/hr/internal/usecases"
	"github.com/5822791760/hr/pkg/errs"
)

type AuthorHandler struct {
	authorUsecase usecases.IAuthorUsecase
}

func NewAuthorHandler(authorService usecases.IAuthorUsecase) AuthorHandler {
	return AuthorHandler{
		authorUsecase: authorService,
	}
}

// CreateAuthor godoc
//
//	@Description	Create an Author
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		usecases.CreateAuthorBody	true	"Author update info"
//
//	@Success		200		{object}	repos.Author
//	@Failure		500,401	{object}	errs.errBase
//
//	@Router			/authors [post]
func (h AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	defer func() {
		WriteError(w, err)
	}()

	ctx := r.Context()

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

// QueryAuthors godoc
//
//	@Description	Find All Author
//	@Tags			authors
//	@Produce		json
//
//	@Success		200		{object}	[]repos.QueryAuthorGetAll
//	@Failure		500,401	{object}	errs.errBase
//
//	@Router			/authors [get]
func (h AuthorHandler) QueryAuthors(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	defer func() {
		WriteError(w, err)
	}()

	ctx := r.Context()

	resp, err := h.authorUsecase.QueryGetAll(ctx)
	if err != nil {
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h AuthorHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	defer func() {
		WriteError(w, err)
	}()

	ctx := r.Context()

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

// UpdateAuthor godoc
//
//	@Description	Update an Author
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		int							true	"Author ID"
//	@Param			request	body		usecases.UpdateAuthorBody	true	"Author update info"
//
//	@Success		200		{object}	repos.Author
//	@Failure		500,401	{object}	errs.errBase
//
//	@Router			/authors/{id} [put]
func (h AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var err errs.Err
	defer func() {
		WriteError(w, err)
	}()

	ctx := r.Context()

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
