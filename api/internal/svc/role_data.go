package svc

import (
	"context"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

func (l *ServiceContext) LoadBanRoleData() error {
	if !l.Config.CoreRpc.Enabled {
		return errorx.NewCodeInternalError(i18n.ServiceUnavailable)
	}

	if l.BanRoleData == nil {
		l.BanRoleData = make(map[string]bool)
	}

	roleData, err := l.CoreRpc.GetRoleList(context.Background(), &core.RoleListReq{
		Page:     1,
		PageSize: 1000,
	})

	if err != nil {
		if strings.Contains(err.Error(), i18n.DatabaseError) {
			return nil
		}
		logx.Error("failed to load role data, please check if initialize the database")
		return errorx.NewCodeInternalError("failed to load role data")
	}

	var state bool
	for _, v := range roleData.Data {
		if uint8(*v.Status) == common.StatusNormal {
			state = false
		} else {
			state = true
		}

		l.BanRoleData[*v.Code] = state
	}

	return nil
}
