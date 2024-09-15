package httpv1

import (
	"net/http"

	"github.com/5822791760/hr/internal/backend/usecases/authorusecase"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/5822791760/hr/pkg/coreutil"
)

type AuthorHandler struct {
	db            coreutil.Transactionable
	authorUsecase authorusecase.IAuthorUsecase
}

func NewAuthorHandler(db coreutil.Transactionable, authorUsecase authorusecase.IAuthorUsecase) AuthorHandler {
	return AuthorHandler{
		db:            db,
		authorUsecase: authorUsecase,
	}
}

func (h AuthorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx, end, err := coreutil.GetTxContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, end(err))
	}()

	var body authorusecase.CreateBody
	if err := coreutil.ParseBody(r, &body); err != nil {
		return
	}

	author, err := h.authorUsecase.Create(ctx, body)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusCreated, author)
}

func (h AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx := coreutil.GetContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, err)
	}()

	resp, err := h.authorUsecase.GetAll(ctx)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusOK, resp)
}

func (h AuthorHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx := coreutil.GetContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, err)
	}()

	id, err := coreutil.GetParamInt(r, "id")
	if err != nil {
		return
	}

	res, err := h.authorUsecase.GetOne(ctx, id)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusOK, res)
}

func (h AuthorHandler) Update(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx, end, err := coreutil.GetTxContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, end(err))
	}()

	id, err := coreutil.GetParamInt(r, "id")
	if err != nil {
		return
	}

	var body authorusecase.UpdateBody
	if err := coreutil.ParseBody(r, &body); err != nil {
		return
	}

	res, err := h.authorUsecase.Update(ctx, id, body)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusOK, res)
}

func (h AuthorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx, end, err := coreutil.GetTxContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, end(err))
	}()

	id, err := coreutil.GetParamInt(r, "id")
	if err != nil {
		return
	}

	res, err := h.authorUsecase.Delete(ctx, id)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusOK, res)
}
