package errs

type errBaseContext struct {
	Key     string
	Message string
}

type errBase struct {
	Code         uint
	ErrorMessage string
	Context      []errBaseContext
}

type Err interface {
	ToHttp() HttpErr
	Error() string
}

func NewErrorContext() []errBaseContext {
	return []errBaseContext{}
}

func (e *errBase) Error() string {
	return e.ErrorMessage
}

func (e *errBase) AddContext(errContext errBaseContext) {
	e.Context = append(e.Context, errContext)
}

// ====== HTTP =======

func (e *errBase) ToHttp() HttpErr {
	err := HttpErr{
		Code:         e.Code,
		ErrorMessage: e.ErrorMessage,
		Context:      []httpErrContext{},
	}

	for _, v := range e.Context {
		err.Context = append(err.Context, httpErrContext{
			Key:     v.Key,
			Message: v.Message,
		})
	}

	return err
}

// ====== HTTP =======
