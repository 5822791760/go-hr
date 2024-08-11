package errs

type HttpErr struct {
	Code         uint   `json:"code"`
	Key          string `json:"key"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}
