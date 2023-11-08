package base

import (
	"context"
	"errors"
	"time"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.ProjectConf.AllowInit {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.PermissionDeny)
	}

	result, err := l.svcCtx.CoreRpc.InitDatabase(l.ctx, &core.Empty{})
	if err != nil && !errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		return nil, err
	} else if errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		for {
			// wait 10 second for initialization
			time.Sleep(time.Second * 5)
			if initState, err := l.svcCtx.Redis.Get("database_init_state"); err == nil {
				if initState == "1" {
					return nil, errorx.NewCodeError(errorcode.InvalidArgument,
						l.svcCtx.Trans.Trans(l.ctx, i18n.AlreadyInit))
				}
			} else {
				return nil, errorx.NewCodeError(errorcode.Internal,
					l.svcCtx.Trans.Trans(l.ctx, i18n.RedisError))
			}

			if errMsg, err := l.svcCtx.Redis.Get("database_error_msg"); err == nil {
				if errMsg != "" {
					return nil, errorx.NewCodeError(errorcode.Internal, errMsg)
				}
			} else {
				return nil, errorx.NewCodeError(errorcode.Internal,
					l.svcCtx.Trans.Trans(l.ctx, i18n.RedisError))
			}
		}
	}

	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin Policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}

	// load role ban data
	err = l.svcCtx.LoadBanRoleData()
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
