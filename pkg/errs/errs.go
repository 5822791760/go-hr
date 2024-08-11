package errs

type httpErrContext struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type HttpErr struct {
	Code         uint             `json:"code"`
	ErrorMessage string           `json:"error_message"`
	Context      []httpErrContext `json:"context"`
}
