// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chroma-core/chroma/go/pkg/coordinator/ent/predicate"
	"github.com/chroma-core/chroma/go/pkg/coordinator/ent/testbase"
)

// TestBaseDelete is the builder for deleting a TestBase entity.
type TestBaseDelete struct {
	config
	hooks    []Hook
	mutation *TestBaseMutation
}

// Where appends a list predicates to the TestBaseDelete builder.
func (tbd *TestBaseDelete) Where(ps ...predicate.TestBase) *TestBaseDelete {
	tbd.mutation.Where(ps...)
	return tbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tbd *TestBaseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tbd.sqlExec, tbd.mutation, tbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tbd *TestBaseDelete) ExecX(ctx context.Context) int {
	n, err := tbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tbd *TestBaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testbase.Table, sqlgraph.NewFieldSpec(testbase.FieldID, field.TypeInt))
	if ps := tbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tbd.mutation.done = true
	return affected, err
}

// TestBaseDeleteOne is the builder for deleting a single TestBase entity.
type TestBaseDeleteOne struct {
	tbd *TestBaseDelete
}

// Where appends a list predicates to the TestBaseDelete builder.
func (tbdo *TestBaseDeleteOne) Where(ps ...predicate.TestBase) *TestBaseDeleteOne {
	tbdo.tbd.mutation.Where(ps...)
	return tbdo
}

// Exec executes the deletion query.
func (tbdo *TestBaseDeleteOne) Exec(ctx context.Context) error {
	n, err := tbdo.tbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testbase.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tbdo *TestBaseDeleteOne) ExecX(ctx context.Context) {
	if err := tbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
