// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
)

// IngredientCreate is the builder for creating a Ingredient entity.
type IngredientCreate struct {
	config
	mutation *IngredientMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *IngredientCreate) SetName(s string) *IngredientCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetQuantity sets the "quantity" field.
func (ic *IngredientCreate) SetQuantity(u uint) *IngredientCreate {
	ic.mutation.SetQuantity(u)
	return ic
}

// Mutation returns the IngredientMutation object of the builder.
func (ic *IngredientCreate) Mutation() *IngredientMutation {
	return ic.mutation
}

// Save creates the Ingredient in the database.
func (ic *IngredientCreate) Save(ctx context.Context) (*Ingredient, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IngredientCreate) SaveX(ctx context.Context) *Ingredient {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IngredientCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IngredientCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IngredientCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Ingredient.name"`)}
	}
	if _, ok := ic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Ingredient.quantity"`)}
	}
	return nil
}

func (ic *IngredientCreate) sqlSave(ctx context.Context) (*Ingredient, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IngredientCreate) createSpec() (*Ingredient, *sqlgraph.CreateSpec) {
	var (
		_node = &Ingredient{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(ingredient.Table, sqlgraph.NewFieldSpec(ingredient.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(ingredient.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Quantity(); ok {
		_spec.SetField(ingredient.FieldQuantity, field.TypeUint, value)
		_node.Quantity = value
	}
	return _node, _spec
}

// IngredientCreateBulk is the builder for creating many Ingredient entities in bulk.
type IngredientCreateBulk struct {
	config
	builders []*IngredientCreate
}

// Save creates the Ingredient entities in the database.
func (icb *IngredientCreateBulk) Save(ctx context.Context) ([]*Ingredient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Ingredient, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IngredientMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IngredientCreateBulk) SaveX(ctx context.Context) []*Ingredient {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IngredientCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IngredientCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
