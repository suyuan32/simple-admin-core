package tenant

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/tenant"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTenantLogic {
	return &CreateOrUpdateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateTenantLogic) CreateOrUpdateTenant(in *core.CreateOrUpdateTenantReq) (*core.BaseResp, error) {

	// 企业级别
	var tenantLevel uint32

	// 结束时间设置：如果没有传值，则有效期至3000年12月31日0时0分0秒（时间戳：32535100800）
	if in.EndTime == 0 {
		in.EndTime = 32535100800
	}

	// 如果Pid不为0，则需要提前设置租户的level；否则设置为企业的根级
	if in.Pid != 0 {
		parent, err := l.svcCtx.DB.Tenant.Query().Where(tenant.Pid(in.Pid)).First(l.ctx)
		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}
		tenantLevel = parent.Level + 1
	} else {
		in.Pid = 1
		tenantLevel = 1
	}

	// 如果是新增租户
	if in.Id == 0 {
		//  end_time :=sql.NullTime{time.Unix(in.EndTime, 0), in.EndTime == 0}
		err := l.svcCtx.DB.Tenant.Create().
			SetParentID(in.Pid).
			SetLevel(tenantLevel).
			SetName(in.Name).
			SetAccount(in.Account).
			SetStartTime(time.Unix(in.StartTime, 0)).
			SetEndTime(time.Unix(in.EndTime, 0)).
			SetContact(in.Contact).
			SetMobile(in.Mobile).
			SetSortNo(in.SortNo).
			Exec(l.ctx)
		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: i18n.CreateSuccess}, nil

	} else {
		// 	// 如果是更新租户信息
		// 	// 判断指定的父级标识是否在系统中存在，如果存在则不进行更新
		// 	exist, err := l.svcCtx.DB.Tenant.Query().Where(tenant.IDEQ(in.Pid)).Exist(l.ctx)
		// 	if err != nil {
		// 		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		// 		return nil, err
		// 	}
		// 	if !exist {
		// 		logx.Errorw("menu not found", logx.Field("menuId", in.Id))
		// 		return nil, status.Error(codes.InvalidArgument, "menu.menuNotExists")
		// 	}
		// 	l.svcCtx.DB.Tenant.UpdateOneID(in.Id).
		// 		SetParentID(in.Pid).
		// 		SetLevel(tenantLevel).
		// 		SetName(in.Name).
		// 		SetAccount(in.Account).
		// 		SetEndTime(time.Unix(in.EndTime, 0)).
		// 		SetContact(in.Contact).
		// 		SetMobile(in.Mobile).
		// 		SetSortNo(in.SortNo).
		// 		Exec(l.ctx)

		return &core.BaseResp{}, nil

	}
}
