// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/pantry"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PantryDelete is the builder for deleting a Pantry entity.
type PantryDelete struct {
	config
	hooks    []Hook
	mutation *PantryMutation
}

// Where appends a list predicates to the PantryDelete builder.
func (pd *PantryDelete) Where(ps ...predicate.Pantry) *PantryDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PantryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PantryDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PantryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(pantry.Table, sqlgraph.NewFieldSpec(pantry.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PantryDeleteOne is the builder for deleting a single Pantry entity.
type PantryDeleteOne struct {
	pd *PantryDelete
}

// Where appends a list predicates to the PantryDelete builder.
func (pdo *PantryDeleteOne) Where(ps ...predicate.Pantry) *PantryDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PantryDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pantry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PantryDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}