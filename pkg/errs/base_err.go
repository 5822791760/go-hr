package errs

type errbase struct {
	Code         uint
	Key          string
	Message      string
	ErrorMessage string
}

type Err interface {
	ToHttp() HttpErr
	Error() string
}

func (e errbase) Error() string {
	return e.Message
}

// ====== HTTP =======

func (e errbase) ToHttp() HttpErr {
	return HttpErr{
		Code:         e.Code,
		Key:          e.Key,
		Message:      e.Message,
		ErrorMessage: e.ErrorMessage,
	}
}

// ====== HTTP =======
