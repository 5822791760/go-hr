package apperr

type errContext = map[string]string

type errBase struct {
	Code         uint
	ErrorMessage string
	Context      errContext
}

type Err interface {
	ToHttp() HttpErr
	Error() string
}

func NewErrorContext() errContext {
	return errContext{}
}

func (e *errBase) Error() string {
	return e.ErrorMessage
}

func (e *errBase) AddContext(key string, msg string) {
	e.Context[key] = msg
}
