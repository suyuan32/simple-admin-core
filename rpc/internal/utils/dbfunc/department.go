package dbfunc

import (
	"context"
	"fmt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
)

func GetDepartmentAncestors(departmentID *uint64, db *ent.Client, logger logx.Logger, ctx context.Context) (*string, error) {
	ancestors, err := db.Department.Query().
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
		Select(department.FieldParentID).Strings(ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(logger, err, fmt.Sprintf("failed to get the department ancestors of %d", departmentID))
	}

	if len(ancestors) == 0 {
		return nil, nil
	}

	return pointy.GetPointer(strings.Join(ancestors, ",")), nil
}
