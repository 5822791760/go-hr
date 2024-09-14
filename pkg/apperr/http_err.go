package apperr

type HttpErr struct {
	Code         uint       `json:"code"`
	ErrorMessage string     `json:"error_message"`
	Context      errContext `json:"context"`
}

func (e *errBase) ToHttp() HttpErr {
	err := HttpErr{
		Code:         e.Code,
		ErrorMessage: e.ErrorMessage,
		Context:      e.Context,
	}

	return err
}
