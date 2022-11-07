package logic

import (
	"context"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenStatusLogic {
	return &UpdateTokenStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenStatusLogic) UpdateTokenStatus(in *core.StatusCodeReq) (*core.BaseResp, error) {
	token, err := l.svcCtx.DB.Token.UpdateOneID(in.Id).SetStatus(uint8(in.Status)).Save(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
		default:
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	// add into redis
	if in.Status == 0 {
		err = l.svcCtx.Redis.Setex("token_"+token.Token, "1", int(token.ExpiredAt.Unix()-
			time.Now().Unix()))
		if err != nil {
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
		}
	} else if in.Status == 1 {
		_, err = l.svcCtx.Redis.Del("token_" + token.Token)
		if err != nil {
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
		}
	}

	logx.Infow("Update token status successfully", logx.Field("TokenId", in.Id),
		logx.Field("Status", in.Status))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
