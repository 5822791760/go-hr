package errs

import "net/http"

func NewAuthorNotFoundErr(err error) Err {
	return &errBase{
		Code:         http.StatusNotFound,
		ErrorMessage: err.Error(),
		Context: []errBaseContext{{
			Key:     AuthorNotFoundErrKey,
			Message: "Author not found",
		}},
	}
}

func NewAuthorValidateErr(errContexts []errBaseContext) Err {
	return &errBase{
		Code:         http.StatusBadRequest,
		ErrorMessage: "",
		Context:      errContexts,
	}
}

// ===== Context =====
func NewAuthorInvalidNameLengthContext() errBaseContext {
	return errBaseContext{
		Key:     AuthorInvalidNameLength,
		Message: "Invalid Author name length",
	}
}

func NewAuthorNameAlreadyExistContext() errBaseContext {
	return errBaseContext{
		Key:     AuthorNameAlreadyExist,
		Message: "Author name already exist",
	}
}
