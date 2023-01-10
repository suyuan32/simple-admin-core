// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/tenant"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

// TenantCreate is the builder for creating a Tenant entity.
type TenantCreate struct {
	config
	mutation *TenantMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TenantCreate) SetCreatedAt(t time.Time) *TenantCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TenantCreate) SetNillableCreatedAt(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TenantCreate) SetUpdatedAt(t time.Time) *TenantCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TenantCreate) SetNillableUpdatedAt(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TenantCreate) SetStatus(u uint8) *TenantCreate {
	tc.mutation.SetStatus(u)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TenantCreate) SetNillableStatus(u *uint8) *TenantCreate {
	if u != nil {
		tc.SetStatus(*u)
	}
	return tc
}

// SetParentID sets the "parent_id" field.
func (tc *TenantCreate) SetParentID(u uuid.UUID) *TenantCreate {
	tc.mutation.SetParentID(u)
	return tc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (tc *TenantCreate) SetNillableParentID(u *uuid.UUID) *TenantCreate {
	if u != nil {
		tc.SetParentID(*u)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TenantCreate) SetName(s string) *TenantCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetLevel sets the "level" field.
func (tc *TenantCreate) SetLevel(u uint32) *TenantCreate {
	tc.mutation.SetLevel(u)
	return tc
}

// SetStartTime sets the "start_time" field.
func (tc *TenantCreate) SetStartTime(t time.Time) *TenantCreate {
	tc.mutation.SetStartTime(t)
	return tc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tc *TenantCreate) SetNillableStartTime(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetStartTime(*t)
	}
	return tc
}

// SetEndTime sets the "end_time" field.
func (tc *TenantCreate) SetEndTime(t time.Time) *TenantCreate {
	tc.mutation.SetEndTime(t)
	return tc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (tc *TenantCreate) SetNillableEndTime(t *time.Time) *TenantCreate {
	if t != nil {
		tc.SetEndTime(*t)
	}
	return tc
}

// SetContact sets the "contact" field.
func (tc *TenantCreate) SetContact(s string) *TenantCreate {
	tc.mutation.SetContact(s)
	return tc
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (tc *TenantCreate) SetNillableContact(s *string) *TenantCreate {
	if s != nil {
		tc.SetContact(*s)
	}
	return tc
}

// SetMobile sets the "mobile" field.
func (tc *TenantCreate) SetMobile(s string) *TenantCreate {
	tc.mutation.SetMobile(s)
	return tc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (tc *TenantCreate) SetNillableMobile(s *string) *TenantCreate {
	if s != nil {
		tc.SetMobile(*s)
	}
	return tc
}

// SetSortNo sets the "sort_no" field.
func (tc *TenantCreate) SetSortNo(u uint32) *TenantCreate {
	tc.mutation.SetSortNo(u)
	return tc
}

// SetNillableSortNo sets the "sort_no" field if the given value is not nil.
func (tc *TenantCreate) SetNillableSortNo(u *uint32) *TenantCreate {
	if u != nil {
		tc.SetSortNo(*u)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TenantCreate) SetID(u uuid.UUID) *TenantCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TenantCreate) SetNillableID(u *uuid.UUID) *TenantCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tc *TenantCreate) AddUserIDs(ids ...uuid.UUID) *TenantCreate {
	tc.mutation.AddUserIDs(ids...)
	return tc
}

// AddUsers adds the "users" edges to the User entity.
func (tc *TenantCreate) AddUsers(u ...*User) *TenantCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserIDs(ids...)
}

// SetParent sets the "parent" edge to the Tenant entity.
func (tc *TenantCreate) SetParent(t *Tenant) *TenantCreate {
	return tc.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Tenant entity by IDs.
func (tc *TenantCreate) AddChildIDs(ids ...uuid.UUID) *TenantCreate {
	tc.mutation.AddChildIDs(ids...)
	return tc
}

// AddChildren adds the "children" edges to the Tenant entity.
func (tc *TenantCreate) AddChildren(t ...*Tenant) *TenantCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildIDs(ids...)
}

// Mutation returns the TenantMutation object of the builder.
func (tc *TenantCreate) Mutation() *TenantMutation {
	return tc.mutation
}

// Save creates the Tenant in the database.
func (tc *TenantCreate) Save(ctx context.Context) (*Tenant, error) {
	var (
		err  error
		node *Tenant
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tenant)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TenantMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TenantCreate) SaveX(ctx context.Context) *Tenant {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TenantCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TenantCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TenantCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := tenant.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := tenant.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := tenant.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.StartTime(); !ok {
		v := tenant.DefaultStartTime()
		tc.mutation.SetStartTime(v)
	}
	if _, ok := tc.mutation.SortNo(); !ok {
		v := tenant.DefaultSortNo
		tc.mutation.SetSortNo(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := tenant.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TenantCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tenant.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Tenant.updated_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tenant.name"`)}
	}
	if _, ok := tc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Tenant.level"`)}
	}
	if _, ok := tc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Tenant.start_time"`)}
	}
	return nil
}

func (tc *TenantCreate) sqlSave(ctx context.Context) (*Tenant, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TenantCreate) createSpec() (*Tenant, *sqlgraph.CreateSpec) {
	var (
		_node = &Tenant{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tenant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tenant.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tenant.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tenant.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(tenant.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tenant.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Level(); ok {
		_spec.SetField(tenant.FieldLevel, field.TypeUint32, value)
		_node.Level = value
	}
	if value, ok := tc.mutation.StartTime(); ok {
		_spec.SetField(tenant.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := tc.mutation.EndTime(); ok {
		_spec.SetField(tenant.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := tc.mutation.Contact(); ok {
		_spec.SetField(tenant.FieldContact, field.TypeString, value)
		_node.Contact = value
	}
	if value, ok := tc.mutation.Mobile(); ok {
		_spec.SetField(tenant.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := tc.mutation.SortNo(); ok {
		_spec.SetField(tenant.FieldSortNo, field.TypeUint32, value)
		_node.SortNo = value
	}
	if nodes := tc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tenant.UsersTable,
			Columns: tenant.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tenant.ParentTable,
			Columns: []string{tenant.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tenant.ChildrenTable,
			Columns: []string{tenant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TenantCreateBulk is the builder for creating many Tenant entities in bulk.
type TenantCreateBulk struct {
	config
	builders []*TenantCreate
}

// Save creates the Tenant entities in the database.
func (tcb *TenantCreateBulk) Save(ctx context.Context) ([]*Tenant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tenant, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TenantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TenantCreateBulk) SaveX(ctx context.Context) []*Tenant {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TenantCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TenantCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
