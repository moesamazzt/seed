package biz

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{
		Code:    OK,
		Message: "success",
		Data:    data,
	}
}

func Fail(err *Error) *Result {
	return &Result{
		Code:    err.Code,
		Message: err.Message,
	}
}
