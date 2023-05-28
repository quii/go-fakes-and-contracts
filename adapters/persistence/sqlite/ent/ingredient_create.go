// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/pantry"
)

// IngredientCreate is the builder for creating a Ingredient entity.
type IngredientCreate struct {
	config
	mutation *IngredientMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ic *IngredientCreate) SetName(s string) *IngredientCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetVegan sets the "vegan" field.
func (ic *IngredientCreate) SetVegan(b bool) *IngredientCreate {
	ic.mutation.SetVegan(b)
	return ic
}

// SetNillableVegan sets the "vegan" field if the given value is not nil.
func (ic *IngredientCreate) SetNillableVegan(b *bool) *IngredientCreate {
	if b != nil {
		ic.SetVegan(*b)
	}
	return ic
}

// SetPantryID sets the "pantry" edge to the Pantry entity by ID.
func (ic *IngredientCreate) SetPantryID(id int) *IngredientCreate {
	ic.mutation.SetPantryID(id)
	return ic
}

// SetNillablePantryID sets the "pantry" edge to the Pantry entity by ID if the given value is not nil.
func (ic *IngredientCreate) SetNillablePantryID(id *int) *IngredientCreate {
	if id != nil {
		ic = ic.SetPantryID(*id)
	}
	return ic
}

// SetPantry sets the "pantry" edge to the Pantry entity.
func (ic *IngredientCreate) SetPantry(p *Pantry) *IngredientCreate {
	return ic.SetPantryID(p.ID)
}

// Mutation returns the IngredientMutation object of the builder.
func (ic *IngredientCreate) Mutation() *IngredientMutation {
	return ic.mutation
}

// Save creates the Ingredient in the database.
func (ic *IngredientCreate) Save(ctx context.Context) (*Ingredient, error) {
	ic.defaults()
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

// defaults sets the default values of the builder before save.
func (ic *IngredientCreate) defaults() {
	if _, ok := ic.mutation.Vegan(); !ok {
		v := ingredient.DefaultVegan
		ic.mutation.SetVegan(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IngredientCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Ingredient.name"`)}
	}
	if _, ok := ic.mutation.Vegan(); !ok {
		return &ValidationError{Name: "vegan", err: errors.New(`ent: missing required field "Ingredient.vegan"`)}
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
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(ingredient.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Vegan(); ok {
		_spec.SetField(ingredient.FieldVegan, field.TypeBool, value)
		_node.Vegan = value
	}
	if nodes := ic.mutation.PantryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   ingredient.PantryTable,
			Columns: []string{ingredient.PantryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pantry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pantry_ingredient = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Ingredient.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IngredientUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ic *IngredientCreate) OnConflict(opts ...sql.ConflictOption) *IngredientUpsertOne {
	ic.conflict = opts
	return &IngredientUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *IngredientCreate) OnConflictColumns(columns ...string) *IngredientUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &IngredientUpsertOne{
		create: ic,
	}
}

type (
	// IngredientUpsertOne is the builder for "upsert"-ing
	//  one Ingredient node.
	IngredientUpsertOne struct {
		create *IngredientCreate
	}

	// IngredientUpsert is the "OnConflict" setter.
	IngredientUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *IngredientUpsert) SetName(v string) *IngredientUpsert {
	u.Set(ingredient.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *IngredientUpsert) UpdateName() *IngredientUpsert {
	u.SetExcluded(ingredient.FieldName)
	return u
}

// SetVegan sets the "vegan" field.
func (u *IngredientUpsert) SetVegan(v bool) *IngredientUpsert {
	u.Set(ingredient.FieldVegan, v)
	return u
}

// UpdateVegan sets the "vegan" field to the value that was provided on create.
func (u *IngredientUpsert) UpdateVegan() *IngredientUpsert {
	u.SetExcluded(ingredient.FieldVegan)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *IngredientUpsertOne) UpdateNewValues() *IngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *IngredientUpsertOne) Ignore() *IngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IngredientUpsertOne) DoNothing() *IngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IngredientCreate.OnConflict
// documentation for more info.
func (u *IngredientUpsertOne) Update(set func(*IngredientUpsert)) *IngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IngredientUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *IngredientUpsertOne) SetName(v string) *IngredientUpsertOne {
	return u.Update(func(s *IngredientUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *IngredientUpsertOne) UpdateName() *IngredientUpsertOne {
	return u.Update(func(s *IngredientUpsert) {
		s.UpdateName()
	})
}

// SetVegan sets the "vegan" field.
func (u *IngredientUpsertOne) SetVegan(v bool) *IngredientUpsertOne {
	return u.Update(func(s *IngredientUpsert) {
		s.SetVegan(v)
	})
}

// UpdateVegan sets the "vegan" field to the value that was provided on create.
func (u *IngredientUpsertOne) UpdateVegan() *IngredientUpsertOne {
	return u.Update(func(s *IngredientUpsert) {
		s.UpdateVegan()
	})
}

// Exec executes the query.
func (u *IngredientUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IngredientCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IngredientUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *IngredientUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *IngredientUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// IngredientCreateBulk is the builder for creating many Ingredient entities in bulk.
type IngredientCreateBulk struct {
	config
	builders []*IngredientCreate
	conflict []sql.ConflictOption
}

// Save creates the Ingredient entities in the database.
func (icb *IngredientCreateBulk) Save(ctx context.Context) ([]*Ingredient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Ingredient, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
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
					spec.OnConflict = icb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Ingredient.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IngredientUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (icb *IngredientCreateBulk) OnConflict(opts ...sql.ConflictOption) *IngredientUpsertBulk {
	icb.conflict = opts
	return &IngredientUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *IngredientCreateBulk) OnConflictColumns(columns ...string) *IngredientUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &IngredientUpsertBulk{
		create: icb,
	}
}

// IngredientUpsertBulk is the builder for "upsert"-ing
// a bulk of Ingredient nodes.
type IngredientUpsertBulk struct {
	create *IngredientCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *IngredientUpsertBulk) UpdateNewValues() *IngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Ingredient.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *IngredientUpsertBulk) Ignore() *IngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IngredientUpsertBulk) DoNothing() *IngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IngredientCreateBulk.OnConflict
// documentation for more info.
func (u *IngredientUpsertBulk) Update(set func(*IngredientUpsert)) *IngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IngredientUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *IngredientUpsertBulk) SetName(v string) *IngredientUpsertBulk {
	return u.Update(func(s *IngredientUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *IngredientUpsertBulk) UpdateName() *IngredientUpsertBulk {
	return u.Update(func(s *IngredientUpsert) {
		s.UpdateName()
	})
}

// SetVegan sets the "vegan" field.
func (u *IngredientUpsertBulk) SetVegan(v bool) *IngredientUpsertBulk {
	return u.Update(func(s *IngredientUpsert) {
		s.SetVegan(v)
	})
}

// UpdateVegan sets the "vegan" field to the value that was provided on create.
func (u *IngredientUpsertBulk) UpdateVegan() *IngredientUpsertBulk {
	return u.Update(func(s *IngredientUpsert) {
		s.UpdateVegan()
	})
}

// Exec executes the query.
func (u *IngredientUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the IngredientCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IngredientCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IngredientUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
