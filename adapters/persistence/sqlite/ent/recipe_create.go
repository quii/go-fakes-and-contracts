// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipe"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipeingredient"
)

// RecipeCreate is the builder for creating a Recipe entity.
type RecipeCreate struct {
	config
	mutation *RecipeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (rc *RecipeCreate) SetName(s string) *RecipeCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RecipeCreate) SetDescription(s string) *RecipeCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetMealType sets the "meal_type" field.
func (rc *RecipeCreate) SetMealType(i int) *RecipeCreate {
	rc.mutation.SetMealType(i)
	return rc
}

// AddRecipeingredientIDs adds the "recipeingredient" edge to the RecipeIngredient entity by IDs.
func (rc *RecipeCreate) AddRecipeingredientIDs(ids ...int) *RecipeCreate {
	rc.mutation.AddRecipeingredientIDs(ids...)
	return rc
}

// AddRecipeingredient adds the "recipeingredient" edges to the RecipeIngredient entity.
func (rc *RecipeCreate) AddRecipeingredient(r ...*RecipeIngredient) *RecipeCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRecipeingredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (rc *RecipeCreate) Mutation() *RecipeMutation {
	return rc.mutation
}

// Save creates the Recipe in the database.
func (rc *RecipeCreate) Save(ctx context.Context) (*Recipe, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecipeCreate) SaveX(ctx context.Context) *Recipe {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecipeCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecipeCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecipeCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Recipe.name"`)}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Recipe.description"`)}
	}
	if _, ok := rc.mutation.MealType(); !ok {
		return &ValidationError{Name: "meal_type", err: errors.New(`ent: missing required field "Recipe.meal_type"`)}
	}
	return nil
}

func (rc *RecipeCreate) sqlSave(ctx context.Context) (*Recipe, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RecipeCreate) createSpec() (*Recipe, *sqlgraph.CreateSpec) {
	var (
		_node = &Recipe{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(recipe.Table, sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(recipe.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(recipe.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.MealType(); ok {
		_spec.SetField(recipe.FieldMealType, field.TypeInt, value)
		_node.MealType = value
	}
	if nodes := rc.mutation.RecipeingredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
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
//	client.Recipe.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (rc *RecipeCreate) OnConflict(opts ...sql.ConflictOption) *RecipeUpsertOne {
	rc.conflict = opts
	return &RecipeUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RecipeCreate) OnConflictColumns(columns ...string) *RecipeUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RecipeUpsertOne{
		create: rc,
	}
}

type (
	// RecipeUpsertOne is the builder for "upsert"-ing
	//  one Recipe node.
	RecipeUpsertOne struct {
		create *RecipeCreate
	}

	// RecipeUpsert is the "OnConflict" setter.
	RecipeUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *RecipeUpsert) SetName(v string) *RecipeUpsert {
	u.Set(recipe.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateName() *RecipeUpsert {
	u.SetExcluded(recipe.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *RecipeUpsert) SetDescription(v string) *RecipeUpsert {
	u.Set(recipe.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateDescription() *RecipeUpsert {
	u.SetExcluded(recipe.FieldDescription)
	return u
}

// SetMealType sets the "meal_type" field.
func (u *RecipeUpsert) SetMealType(v int) *RecipeUpsert {
	u.Set(recipe.FieldMealType, v)
	return u
}

// UpdateMealType sets the "meal_type" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateMealType() *RecipeUpsert {
	u.SetExcluded(recipe.FieldMealType)
	return u
}

// AddMealType adds v to the "meal_type" field.
func (u *RecipeUpsert) AddMealType(v int) *RecipeUpsert {
	u.Add(recipe.FieldMealType, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecipeUpsertOne) UpdateNewValues() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Recipe.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RecipeUpsertOne) Ignore() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeUpsertOne) DoNothing() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeCreate.OnConflict
// documentation for more info.
func (u *RecipeUpsertOne) Update(set func(*RecipeUpsert)) *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RecipeUpsertOne) SetName(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateName() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RecipeUpsertOne) SetDescription(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateDescription() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateDescription()
	})
}

// SetMealType sets the "meal_type" field.
func (u *RecipeUpsertOne) SetMealType(v int) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetMealType(v)
	})
}

// AddMealType adds v to the "meal_type" field.
func (u *RecipeUpsertOne) AddMealType(v int) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.AddMealType(v)
	})
}

// UpdateMealType sets the "meal_type" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateMealType() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateMealType()
	})
}

// Exec executes the query.
func (u *RecipeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RecipeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RecipeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RecipeCreateBulk is the builder for creating many Recipe entities in bulk.
type RecipeCreateBulk struct {
	config
	builders []*RecipeCreate
	conflict []sql.ConflictOption
}

// Save creates the Recipe entities in the database.
func (rcb *RecipeCreateBulk) Save(ctx context.Context) ([]*Recipe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recipe, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecipeCreateBulk) SaveX(ctx context.Context) []*Recipe {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecipeCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecipeCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Recipe.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (rcb *RecipeCreateBulk) OnConflict(opts ...sql.ConflictOption) *RecipeUpsertBulk {
	rcb.conflict = opts
	return &RecipeUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RecipeCreateBulk) OnConflictColumns(columns ...string) *RecipeUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RecipeUpsertBulk{
		create: rcb,
	}
}

// RecipeUpsertBulk is the builder for "upsert"-ing
// a bulk of Recipe nodes.
type RecipeUpsertBulk struct {
	create *RecipeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RecipeUpsertBulk) UpdateNewValues() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RecipeUpsertBulk) Ignore() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeUpsertBulk) DoNothing() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeCreateBulk.OnConflict
// documentation for more info.
func (u *RecipeUpsertBulk) Update(set func(*RecipeUpsert)) *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *RecipeUpsertBulk) SetName(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateName() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *RecipeUpsertBulk) SetDescription(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateDescription() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateDescription()
	})
}

// SetMealType sets the "meal_type" field.
func (u *RecipeUpsertBulk) SetMealType(v int) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetMealType(v)
	})
}

// AddMealType adds v to the "meal_type" field.
func (u *RecipeUpsertBulk) AddMealType(v int) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.AddMealType(v)
	})
}

// UpdateMealType sets the "meal_type" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateMealType() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateMealType()
	})
}

// Exec executes the query.
func (u *RecipeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RecipeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
