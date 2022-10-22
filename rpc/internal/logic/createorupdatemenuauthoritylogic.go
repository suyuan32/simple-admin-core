package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	// delete the data create before
	deleteString := fmt.Sprintf("DELETE from role_menus where role_id = %d", in.RoleId)
	result := l.svcCtx.DB.Exec(deleteString)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	var insertString strings.Builder
	insertString.WriteString("insert into role_menus values ")
	for i := 0; i < len(in.MenuId); i++ {
		if i != len(in.MenuId)-1 {
			insertString.WriteString(fmt.Sprintf("(%d, %d),", in.MenuId[i], in.RoleId))
		} else {
			insertString.WriteString(fmt.Sprintf("(%d, %d);", in.MenuId[i], in.RoleId))
		}
	}
	result = l.svcCtx.DB.Exec(insertString.String())
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, errorx.DatabaseError)
	}

	logx.Infow(logmessage.UpdateSuccess, logx.Field("authorityRelation", insertString.String()))
	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}
