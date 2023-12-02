package token

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/rpc/ent/token"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserAllTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserAllTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserAllTokenLogic {
	return &BlockUserAllTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserAllTokenLogic) BlockUserAllToken(in *core.UUIDReq) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Token.Update().Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).SetStatus(0).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	tokenData, err := l.svcCtx.DB.Token.Query().
		Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).
		Where(token.StatusEQ(0)).
		Where(token.ExpiredAtGT(time.Now())).
		All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	for _, v := range tokenData {
		expiredTime := int(v.ExpiredAt.Unix() - time.Now().Unix())
		if expiredTime > 0 {
			err = l.svcCtx.Redis.Setex("token_"+v.Token, "1", expiredTime)
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return nil, errorx.NewInternalError(i18n.RedisError)
			}
		}
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
