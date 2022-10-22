package logic

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
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
	result := l.svcCtx.DB.Model(&model.Token{}).Where("uuid = ?", in.UUID).Update("status", 0)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	var tokens []model.Token
	tokenData := l.svcCtx.DB.Where("uuid = ?", in.UUID).Where("status = ?", 0).
		Where("expire_at > ?", time.Now()).Find(&tokens)

	if tokenData.Error != nil && !errors.Is(tokenData.Error, gorm.ErrRecordNotFound) {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	for _, v := range tokens {
		err := l.svcCtx.Redis.Set("token_"+v.Token, "1")
		if err != nil {
			logx.Errorw(logmessage.RedisError, logx.Field("detail", err.Error()))
			return nil, status.Error(codes.Internal, errorx.RedisError)
		}
	}

	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
