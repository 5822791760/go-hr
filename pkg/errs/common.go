package errs

import (
	"fmt"
	"net/http"
)

func NewInternalServerErr(err error) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: err.Error(),
		Context: []errBaseContext{{
			Key:     InternalErrKey,
			Message: "",
		}},
	}
}

func NewInternalServerErrByString(message string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: message,
		Context: []errBaseContext{{
			Key:     InternalErrKey,
			Message: "",
		}},
	}
}

func NewNoRowAffectedErr() Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: "No row affected",
		Context: []errBaseContext{{
			Key:     InternalErrKey,
			Message: "",
		}},
	}
}

func NewQueryNotExistErr(key string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: fmt.Sprintf("Query %s does not exist", key),
		Context: []errBaseContext{{
			Key:     InternalErrKey,
			Message: "",
		}},
	}
}

func NewParamNotExistErr(key string) Err {
	return &errBase{
		Code:         http.StatusInternalServerError,
		ErrorMessage: fmt.Sprintf("Param %s does not exist", key),
		Context: []errBaseContext{{
			Key:     InternalErrKey,
			Message: "",
		}},
	}
}
