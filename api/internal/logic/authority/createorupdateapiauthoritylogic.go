package authority

import (
	"context"
	"net/http"
	"strconv"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CreateOrUpdateApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiAuthorityLogic {
	return &CreateOrUpdateApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateApiAuthorityLogic) CreateOrUpdateApiAuthority(req *types.CreateOrUpdateApiAuthorityReq) (resp *types.SimpleMsg, err error) {
	// clear old policies
	roleIdString := strconv.Itoa(int(req.RoleId))
	var oldPolicies [][]string
	oldPolicies = l.svcCtx.Casbin.GetFilteredPolicy(0, roleIdString)
	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, roleIdString)
		if err != nil {
			return nil, &httpx.ApiError{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			}
		}
		if !removeResult {
			return nil, httpx.NewApiError(http.StatusInternalServerError, "cannot clear old policies")
		}
	}
	// add new policies
	var policies [][]string
	for _, v := range req.Data {
		policies = append(policies, []string{roleIdString, v.Path, v.Method})
	}
	addResult, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil {
		return nil, err
	}
	if addResult {
		return &types.SimpleMsg{Msg: errorx.UpdateSuccess}, nil
	} else {
		return nil, httpx.NewApiError(http.StatusBadRequest, errorx.UpdateFailed)
	}
}
