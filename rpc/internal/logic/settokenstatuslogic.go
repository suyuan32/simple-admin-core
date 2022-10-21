package logic

import (
	"context"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenStatusLogic {
	return &SetTokenStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTokenStatusLogic) SetTokenStatus(in *core.SetStatusReq) (*core.BaseResp, error) {
	result := l.svcCtx.DB.Table("tokens").Where("id = ?", in.Id).Update("status", in.Status)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logx.Errorw("Update token status failed, please check the token id", logx.Field("TokenId", in.Id))
		return nil, status.Error(codes.InvalidArgument, errorx.UpdateFailed)
	}

	// add into redis
	if in.Status == 0 {
		var tokenData model.Token
		l.svcCtx.DB.Where("id = ?", in.Id).First(&tokenData)
		err := l.svcCtx.Redis.Setex("token_"+tokenData.Token, "1", int(tokenData.ExpireAt.Unix()-
			time.Now().Unix()))
		if err != nil {
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
		}
	} else if in.Status == 1 {
		var tokenData model.Token
		l.svcCtx.DB.Where("id = ?", in.Id).First(&tokenData)
		_, err := l.svcCtx.Redis.Del("token_" + tokenData.Token)
		if err != nil {
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
		}
	}

	logx.Infow("Update token status successfully", logx.Field("TokenId", in.Id),
		logx.Field("Status", in.Status))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
