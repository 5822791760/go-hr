package apperr

import "net/http"

const (
	AuthorNameLength = "nameLength"
	AuthorNameExist  = "nameExist"
	AuthorNotFound   = "notFound"
)

func NewAuthorNotFoundErr(err error) Err {
	return &errBase{
		Code:         http.StatusNotFound,
		ErrorMessage: err.Error(),
		Context:      errContext{AuthorNotFound: "Author not found"},
	}
}

func NewAuthorValidateErr(errCtx errContext) Err {
	return &errBase{
		Code:         http.StatusBadRequest,
		ErrorMessage: "",
		Context:      errCtx,
	}
}

// ===== Context =====
func AddAuthorInvalidNameLengthContext(errCtx errContext) {
	errCtx[AuthorNameLength] = "Invalid Author name length"

}

func AddAuthorNameAlreadyExistContext(errCtx errContext) {
	errCtx[AuthorNameExist] = "Author name already exist"
}
