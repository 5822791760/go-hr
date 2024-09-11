package errs

import (
	"fmt"
	"net/http"
)

const (
	Internal = "internal"
)

func NewInternalServerErr(err error) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: err.Error(),
		Context:      errContext{Internal: ""},
	}
}

func NewInternalServerErrByString(message string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: message,
		Context:      errContext{Internal: ""},
	}
}

func NewNoRowAffectedErr() Err {
	return &errBase{
		Code:         http.StatusBadRequest,
		ErrorMessage: "No row affected",
		Context:      errContext{Internal: ""},
	}
}

func NewQueryNotExistErr(key string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: fmt.Sprintf("Query %s does not exist", key),
		Context:      errContext{Internal: ""},
	}
}

func NewParamNotExistErr(key string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: fmt.Sprintf("Param %s does not exist", key),
		Context:      errContext{Internal: ""},
	}
}
