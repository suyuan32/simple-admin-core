package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/ent/token"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *core.TokenListReq) (*core.TokenListResp, error) {
	var tokens *ent.TokenPageList
	var err error
	if in.Username == "" && in.Uuid == "" && in.Nickname == "" && in.Email == "" {
		tokens, err = l.svcCtx.DB.Token.Query().Page(l.ctx, in.Page, in.PageSize)

		if err != nil {
			logx.Error(err.Error())
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	} else {
		var predicates []predicate.User

		if in.Uuid != "" {
			predicates = append(predicates, user.UUIDEQ(in.Uuid))
		}

		if in.Username != "" {
			predicates = append(predicates, user.Username(in.Username))
		}

		if in.Email != "" {
			predicates = append(predicates, user.EmailEQ(in.Email))
		}

		if in.Nickname != "" {
			predicates = append(predicates, user.NicknameEQ(in.Nickname))
		}

		u, err := l.svcCtx.DB.User.Query().Where(predicates...).First(l.ctx)

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

		tokens, err = l.svcCtx.DB.Token.Query().Where(token.UUIDEQ(u.UUID)).Page(l.ctx, in.Page, in.PageSize)

		if err != nil {
			logx.Error(err.Error())
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	resp := &core.TokenListResp{}
	resp.Total = tokens.PageDetails.Total

	for _, v := range tokens.List {
		resp.Data = append(resp.Data, &core.TokenInfo{
			Id:        v.ID,
			Uuid:      v.UUID,
			Token:     v.Token,
			Status:    uint64(v.Status),
			Source:    v.Source,
			ExpiredAt: v.ExpiredAt.UnixMilli(),
			CreatedAt: v.CreatedAt.UnixMilli(),
		})
	}

	return nil, nil
}
