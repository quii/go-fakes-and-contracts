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
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipe"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipeingredient"
)

// RecipeIngredientCreate is the builder for creating a RecipeIngredient entity.
type RecipeIngredientCreate struct {
	config
	mutation *RecipeIngredientMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetQuantity sets the "quantity" field.
func (ric *RecipeIngredientCreate) SetQuantity(i int) *RecipeIngredientCreate {
	ric.mutation.SetQuantity(i)
	return ric
}

// SetRecipeID sets the "recipe" edge to the Recipe entity by ID.
func (ric *RecipeIngredientCreate) SetRecipeID(id int) *RecipeIngredientCreate {
	ric.mutation.SetRecipeID(id)
	return ric
}

// SetNillableRecipeID sets the "recipe" edge to the Recipe entity by ID if the given value is not nil.
func (ric *RecipeIngredientCreate) SetNillableRecipeID(id *int) *RecipeIngredientCreate {
	if id != nil {
		ric = ric.SetRecipeID(*id)
	}
	return ric
}

// SetRecipe sets the "recipe" edge to the Recipe entity.
func (ric *RecipeIngredientCreate) SetRecipe(r *Recipe) *RecipeIngredientCreate {
	return ric.SetRecipeID(r.ID)
}

// AddIngredientIDs adds the "ingredient" edge to the Ingredient entity by IDs.
func (ric *RecipeIngredientCreate) AddIngredientIDs(ids ...int) *RecipeIngredientCreate {
	ric.mutation.AddIngredientIDs(ids...)
	return ric
}

// AddIngredient adds the "ingredient" edges to the Ingredient entity.
func (ric *RecipeIngredientCreate) AddIngredient(i ...*Ingredient) *RecipeIngredientCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ric.AddIngredientIDs(ids...)
}

// Mutation returns the RecipeIngredientMutation object of the builder.
func (ric *RecipeIngredientCreate) Mutation() *RecipeIngredientMutation {
	return ric.mutation
}

// Save creates the RecipeIngredient in the database.
func (ric *RecipeIngredientCreate) Save(ctx context.Context) (*RecipeIngredient, error) {
	return withHooks(ctx, ric.sqlSave, ric.mutation, ric.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ric *RecipeIngredientCreate) SaveX(ctx context.Context) *RecipeIngredient {
	v, err := ric.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ric *RecipeIngredientCreate) Exec(ctx context.Context) error {
	_, err := ric.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ric *RecipeIngredientCreate) ExecX(ctx context.Context) {
	if err := ric.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ric *RecipeIngredientCreate) check() error {
	if _, ok := ric.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "RecipeIngredient.quantity"`)}
	}
	return nil
}

func (ric *RecipeIngredientCreate) sqlSave(ctx context.Context) (*RecipeIngredient, error) {
	if err := ric.check(); err != nil {
		return nil, err
	}
	_node, _spec := ric.createSpec()
	if err := sqlgraph.CreateNode(ctx, ric.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ric.mutation.id = &_node.ID
	ric.mutation.done = true
	return _node, nil
}

func (ric *RecipeIngredientCreate) createSpec() (*RecipeIngredient, *sqlgraph.CreateSpec) {
	var (
		_node = &RecipeIngredient{config: ric.config}
		_spec = sqlgraph.NewCreateSpec(recipeingredient.Table, sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ric.conflict
	if value, ok := ric.mutation.Quantity(); ok {
		_spec.SetField(recipeingredient.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if nodes := ric.mutation.RecipeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipeingredient.RecipeTable,
			Columns: []string{recipeingredient.RecipeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.recipe_recipeingredient = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ric.mutation.IngredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipeingredient.IngredientTable,
			Columns: recipeingredient.IngredientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RecipeIngredient.Create().
//		SetQuantity(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeIngredientUpsert) {
//			SetQuantity(v+v).
//		}).
//		Exec(ctx)
func (ric *RecipeIngredientCreate) OnConflict(opts ...sql.ConflictOption) *RecipeIngredientUpsertOne {
	ric.conflict = opts
	return &RecipeIngredientUpsertOne{
		create: ric,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ric *RecipeIngredientCreate) OnConflictColumns(columns ...string) *RecipeIngredientUpsertOne {
	ric.conflict = append(ric.conflict, sql.ConflictColumns(columns...))
	return &RecipeIngredientUpsertOne{
		create: ric,
	}
}

type (
	// RecipeIngredientUpsertOne is the builder for "upsert"-ing
	//  one RecipeIngredient node.
	RecipeIngredientUpsertOne struct {
		create *RecipeIngredientCreate
	}

	// RecipeIngredientUpsert is the "OnConflict" setter.
	RecipeIngredientUpsert struct {
		*sql.UpdateSet
	}
)

// SetQuantity sets the "quantity" field.
func (u *RecipeIngredientUpsert) SetQuantity(v int) *RecipeIngredientUpsert {
	u.Set(recipeingredient.FieldQuantity, v)
	return u
}

// UpdateQuantity sets the "quantity" field to the value that was provided on create.
func (u *RecipeIngredientUpsert) UpdateQuantity() *RecipeIngredientUpsert {
	u.SetExcluded(recipeingredient.FieldQuantity)
	return u
}

// AddQuantity adds v to the "quantity" field.
func (u *RecipeIngredientUpsert) AddQuantity(v int) *RecipeIngredientUpsert {
	u.Add(recipeingredient.FieldQuantity, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecipeIngredientUpsertOne) UpdateNewValues() *RecipeIngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RecipeIngredientUpsertOne) Ignore() *RecipeIngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeIngredientUpsertOne) DoNothing() *RecipeIngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeIngredientCreate.OnConflict
// documentation for more info.
func (u *RecipeIngredientUpsertOne) Update(set func(*RecipeIngredientUpsert)) *RecipeIngredientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeIngredientUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuantity sets the "quantity" field.
func (u *RecipeIngredientUpsertOne) SetQuantity(v int) *RecipeIngredientUpsertOne {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.SetQuantity(v)
	})
}

// AddQuantity adds v to the "quantity" field.
func (u *RecipeIngredientUpsertOne) AddQuantity(v int) *RecipeIngredientUpsertOne {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.AddQuantity(v)
	})
}

// UpdateQuantity sets the "quantity" field to the value that was provided on create.
func (u *RecipeIngredientUpsertOne) UpdateQuantity() *RecipeIngredientUpsertOne {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.UpdateQuantity()
	})
}

// Exec executes the query.
func (u *RecipeIngredientUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeIngredientCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeIngredientUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RecipeIngredientUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RecipeIngredientUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RecipeIngredientCreateBulk is the builder for creating many RecipeIngredient entities in bulk.
type RecipeIngredientCreateBulk struct {
	config
	builders []*RecipeIngredientCreate
	conflict []sql.ConflictOption
}

// Save creates the RecipeIngredient entities in the database.
func (ricb *RecipeIngredientCreateBulk) Save(ctx context.Context) ([]*RecipeIngredient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ricb.builders))
	nodes := make([]*RecipeIngredient, len(ricb.builders))
	mutators := make([]Mutator, len(ricb.builders))
	for i := range ricb.builders {
		func(i int, root context.Context) {
			builder := ricb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeIngredientMutation)
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
					_, err = mutators[i+1].Mutate(root, ricb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ricb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ricb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ricb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ricb *RecipeIngredientCreateBulk) SaveX(ctx context.Context) []*RecipeIngredient {
	v, err := ricb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ricb *RecipeIngredientCreateBulk) Exec(ctx context.Context) error {
	_, err := ricb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ricb *RecipeIngredientCreateBulk) ExecX(ctx context.Context) {
	if err := ricb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RecipeIngredient.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeIngredientUpsert) {
//			SetQuantity(v+v).
//		}).
//		Exec(ctx)
func (ricb *RecipeIngredientCreateBulk) OnConflict(opts ...sql.ConflictOption) *RecipeIngredientUpsertBulk {
	ricb.conflict = opts
	return &RecipeIngredientUpsertBulk{
		create: ricb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ricb *RecipeIngredientCreateBulk) OnConflictColumns(columns ...string) *RecipeIngredientUpsertBulk {
	ricb.conflict = append(ricb.conflict, sql.ConflictColumns(columns...))
	return &RecipeIngredientUpsertBulk{
		create: ricb,
	}
}

// RecipeIngredientUpsertBulk is the builder for "upsert"-ing
// a bulk of RecipeIngredient nodes.
type RecipeIngredientUpsertBulk struct {
	create *RecipeIngredientCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecipeIngredientUpsertBulk) UpdateNewValues() *RecipeIngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RecipeIngredient.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RecipeIngredientUpsertBulk) Ignore() *RecipeIngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeIngredientUpsertBulk) DoNothing() *RecipeIngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeIngredientCreateBulk.OnConflict
// documentation for more info.
func (u *RecipeIngredientUpsertBulk) Update(set func(*RecipeIngredientUpsert)) *RecipeIngredientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeIngredientUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuantity sets the "quantity" field.
func (u *RecipeIngredientUpsertBulk) SetQuantity(v int) *RecipeIngredientUpsertBulk {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.SetQuantity(v)
	})
}

// AddQuantity adds v to the "quantity" field.
func (u *RecipeIngredientUpsertBulk) AddQuantity(v int) *RecipeIngredientUpsertBulk {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.AddQuantity(v)
	})
}

// UpdateQuantity sets the "quantity" field to the value that was provided on create.
func (u *RecipeIngredientUpsertBulk) UpdateQuantity() *RecipeIngredientUpsertBulk {
	return u.Update(func(s *RecipeIngredientUpsert) {
		s.UpdateQuantity()
	})
}

// Exec executes the query.
func (u *RecipeIngredientUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RecipeIngredientCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeIngredientCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeIngredientUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
