package errorx

import (
	"google.golang.org/grpc/codes"
)

const defaultRpcCode = 0

type RpcError struct {
	Code codes.Code `json:"code"`
	Msg  string     `json:"msg"`
}
type RpcErrorResponse struct {
	Code codes.Code `json:"code"`
	Msg  string     `json:"msg"`
}

func NewRpcError(code codes.Code, msg string) error {
	return &RpcError{Code: code, Msg: msg}
}
func NewDefaultRpcError(msg string) error {
	return NewRpcError(defaultRpcCode, msg)
}
func (e *RpcError) Error() string {
	return e.Msg
}
func (e *RpcError) Data() *RpcErrorResponse {
	return &RpcErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
