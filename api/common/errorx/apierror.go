package errorx

const defaultApiCode = 1001

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type ApiErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewApiError(code int, msg string) error {
	return &ApiError{Code: code, Msg: msg}
}
func NewDefaultApiError(msg string) error {
	return NewApiError(defaultApiCode, msg)
}
func (e *ApiError) Error() string {
	return e.Msg
}
func (e *ApiError) Data() *ApiErrorResponse {
	return &ApiErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
