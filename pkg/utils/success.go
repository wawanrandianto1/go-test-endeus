package utils

type Success struct {
	Success bool
	Message string
}

func NewSuccessMsg(msg string) Success {
	return Success{
		Success: true,
		Message: msg,
	}
}
