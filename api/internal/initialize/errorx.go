package initialize

import (
	"github.com/suyuan32/simple-admin-core/api/common/errorx"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// initialize the error handler
func InitErrorHandler() {
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.ApiError:
			switch e.Code {
			case 0:
				return http.StatusOK, &types.SimpleMsg{
					Msg: status.Convert(err).Message(),
				}
			default:
				return e.Code, &types.SimpleMsg{
					Msg: status.Convert(err).Message(),
				}
			}
		default:
			return errorx.CodeFromGrpcError(e), &types.SimpleMsg{
				Msg: status.Convert(err).Message(),
			}
		}
	})
}
