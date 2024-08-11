package errs

import "net/http"

func NewAuthorNotFoundErr(err error) Err {
	return errbase{
		Code:         http.StatusNotFound,
		Key:          AuthorNotFoundErrKey,
		Message:      "Author not found",
		ErrorMessage: err.Error(),
	}
}
