package errs

import (
	"fmt"
	"net/http"
)

func NewInternalServerErr(err error) Err {
	return errbase{
		Code:         http.StatusInternalServerError,
		Key:          InternalErrKey,
		Message:      "",
		ErrorMessage: err.Error(),
	}
}

func NewQueryNotExistErr(key string) Err {
	return errbase{
		Code:         http.StatusInternalServerError,
		Key:          InternalErrKey,
		Message:      "",
		ErrorMessage: fmt.Sprintf("Query %s does not exist", key),
	}
}

func NewParamNotExistErr(key string) Err {
	return errbase{
		Code:         http.StatusInternalServerError,
		Key:          InternalErrKey,
		Message:      "",
		ErrorMessage: fmt.Sprintf("Param %s does not exist", key),
	}
}
