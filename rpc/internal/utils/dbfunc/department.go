package dbfunc

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenAncestorsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenAncestorsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenAncestorsLogic {
	return &GenAncestorsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenAncestorsLogic) Department_ancestors(departmentID *uint64) *string {
	ancestors, err := l.svcCtx.DB.Debug().Department.Query().
		Modify(func(s *sql.Selector) {
			t1, t2 := sql.Table(department.Table), sql.Table(department.Table)
			with := sql.WithRecursive("ancestors")
			with.As(
				sql.Select(
					t1.C(department.FieldID),
					t1.C(department.FieldParentID),
				).AppendSelectExpr(sql.ExprFunc(func(b *sql.Builder) {
					b.Ident("1 as level")
				})).
					From(t1).
					Where(sql.EQ(department.FieldID, *departmentID)).
					UnionAll(
						sql.Select(t2.Columns(department.FieldID, department.FieldParentID)...).
							AppendSelectExpr(sql.ExprFunc(func(b *sql.Builder) {
								b.Ident(with.Name() + ".level + 1 as level")
							})).
							From(t2).
							Join(with).
							On(t2.C(department.FieldID), with.C(department.FieldParentID)),
					),
			)
			s.Prefix(with).Select(with.C(department.FieldID)).From(with)
		}).
		Select(department.FieldParentID).Strings(l.ctx)
	if err != nil {
		fmt.Println("error", err)
	}
	ancestorsStr := strings.Join(ancestors, ",")
	return &ancestorsStr
}
