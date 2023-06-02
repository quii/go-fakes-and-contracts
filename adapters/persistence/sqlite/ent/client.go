// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/pantry"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipe"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/recipeingredient"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Ingredient is the client for interacting with the Ingredient builders.
	Ingredient *IngredientClient
	// Pantry is the client for interacting with the Pantry builders.
	Pantry *PantryClient
	// Recipe is the client for interacting with the Recipe builders.
	Recipe *RecipeClient
	// RecipeIngredient is the client for interacting with the RecipeIngredient builders.
	RecipeIngredient *RecipeIngredientClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Ingredient = NewIngredientClient(c.config)
	c.Pantry = NewPantryClient(c.config)
	c.Recipe = NewRecipeClient(c.config)
	c.RecipeIngredient = NewRecipeIngredientClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Ingredient:       NewIngredientClient(cfg),
		Pantry:           NewPantryClient(cfg),
		Recipe:           NewRecipeClient(cfg),
		RecipeIngredient: NewRecipeIngredientClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Ingredient:       NewIngredientClient(cfg),
		Pantry:           NewPantryClient(cfg),
		Recipe:           NewRecipeClient(cfg),
		RecipeIngredient: NewRecipeIngredientClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Ingredient.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Ingredient.Use(hooks...)
	c.Pantry.Use(hooks...)
	c.Recipe.Use(hooks...)
	c.RecipeIngredient.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Ingredient.Intercept(interceptors...)
	c.Pantry.Intercept(interceptors...)
	c.Recipe.Intercept(interceptors...)
	c.RecipeIngredient.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *IngredientMutation:
		return c.Ingredient.mutate(ctx, m)
	case *PantryMutation:
		return c.Pantry.mutate(ctx, m)
	case *RecipeMutation:
		return c.Recipe.mutate(ctx, m)
	case *RecipeIngredientMutation:
		return c.RecipeIngredient.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// IngredientClient is a client for the Ingredient schema.
type IngredientClient struct {
	config
}

// NewIngredientClient returns a client for the Ingredient from the given config.
func NewIngredientClient(c config) *IngredientClient {
	return &IngredientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ingredient.Hooks(f(g(h())))`.
func (c *IngredientClient) Use(hooks ...Hook) {
	c.hooks.Ingredient = append(c.hooks.Ingredient, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ingredient.Intercept(f(g(h())))`.
func (c *IngredientClient) Intercept(interceptors ...Interceptor) {
	c.inters.Ingredient = append(c.inters.Ingredient, interceptors...)
}

// Create returns a builder for creating a Ingredient entity.
func (c *IngredientClient) Create() *IngredientCreate {
	mutation := newIngredientMutation(c.config, OpCreate)
	return &IngredientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ingredient entities.
func (c *IngredientClient) CreateBulk(builders ...*IngredientCreate) *IngredientCreateBulk {
	return &IngredientCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ingredient.
func (c *IngredientClient) Update() *IngredientUpdate {
	mutation := newIngredientMutation(c.config, OpUpdate)
	return &IngredientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IngredientClient) UpdateOne(i *Ingredient) *IngredientUpdateOne {
	mutation := newIngredientMutation(c.config, OpUpdateOne, withIngredient(i))
	return &IngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IngredientClient) UpdateOneID(id int) *IngredientUpdateOne {
	mutation := newIngredientMutation(c.config, OpUpdateOne, withIngredientID(id))
	return &IngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ingredient.
func (c *IngredientClient) Delete() *IngredientDelete {
	mutation := newIngredientMutation(c.config, OpDelete)
	return &IngredientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IngredientClient) DeleteOne(i *Ingredient) *IngredientDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IngredientClient) DeleteOneID(id int) *IngredientDeleteOne {
	builder := c.Delete().Where(ingredient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IngredientDeleteOne{builder}
}

// Query returns a query builder for Ingredient.
func (c *IngredientClient) Query() *IngredientQuery {
	return &IngredientQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIngredient},
		inters: c.Interceptors(),
	}
}

// Get returns a Ingredient entity by its id.
func (c *IngredientClient) Get(ctx context.Context, id int) (*Ingredient, error) {
	return c.Query().Where(ingredient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IngredientClient) GetX(ctx context.Context, id int) *Ingredient {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPantry queries the pantry edge of a Ingredient.
func (c *IngredientClient) QueryPantry(i *Ingredient) *PantryQuery {
	query := (&PantryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ingredient.Table, ingredient.FieldID, id),
			sqlgraph.To(pantry.Table, pantry.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ingredient.PantryTable, ingredient.PantryColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRecipeingredient queries the recipeingredient edge of a Ingredient.
func (c *IngredientClient) QueryRecipeingredient(i *Ingredient) *RecipeIngredientQuery {
	query := (&RecipeIngredientClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ingredient.Table, ingredient.FieldID, id),
			sqlgraph.To(recipeingredient.Table, recipeingredient.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ingredient.RecipeingredientTable, ingredient.RecipeingredientPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IngredientClient) Hooks() []Hook {
	return c.hooks.Ingredient
}

// Interceptors returns the client interceptors.
func (c *IngredientClient) Interceptors() []Interceptor {
	return c.inters.Ingredient
}

func (c *IngredientClient) mutate(ctx context.Context, m *IngredientMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IngredientCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IngredientUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IngredientDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Ingredient mutation op: %q", m.Op())
	}
}

// PantryClient is a client for the Pantry schema.
type PantryClient struct {
	config
}

// NewPantryClient returns a client for the Pantry from the given config.
func NewPantryClient(c config) *PantryClient {
	return &PantryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pantry.Hooks(f(g(h())))`.
func (c *PantryClient) Use(hooks ...Hook) {
	c.hooks.Pantry = append(c.hooks.Pantry, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `pantry.Intercept(f(g(h())))`.
func (c *PantryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Pantry = append(c.inters.Pantry, interceptors...)
}

// Create returns a builder for creating a Pantry entity.
func (c *PantryClient) Create() *PantryCreate {
	mutation := newPantryMutation(c.config, OpCreate)
	return &PantryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pantry entities.
func (c *PantryClient) CreateBulk(builders ...*PantryCreate) *PantryCreateBulk {
	return &PantryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pantry.
func (c *PantryClient) Update() *PantryUpdate {
	mutation := newPantryMutation(c.config, OpUpdate)
	return &PantryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PantryClient) UpdateOne(pa *Pantry) *PantryUpdateOne {
	mutation := newPantryMutation(c.config, OpUpdateOne, withPantry(pa))
	return &PantryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PantryClient) UpdateOneID(id int) *PantryUpdateOne {
	mutation := newPantryMutation(c.config, OpUpdateOne, withPantryID(id))
	return &PantryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pantry.
func (c *PantryClient) Delete() *PantryDelete {
	mutation := newPantryMutation(c.config, OpDelete)
	return &PantryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PantryClient) DeleteOne(pa *Pantry) *PantryDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PantryClient) DeleteOneID(id int) *PantryDeleteOne {
	builder := c.Delete().Where(pantry.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PantryDeleteOne{builder}
}

// Query returns a query builder for Pantry.
func (c *PantryClient) Query() *PantryQuery {
	return &PantryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePantry},
		inters: c.Interceptors(),
	}
}

// Get returns a Pantry entity by its id.
func (c *PantryClient) Get(ctx context.Context, id int) (*Pantry, error) {
	return c.Query().Where(pantry.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PantryClient) GetX(ctx context.Context, id int) *Pantry {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIngredient queries the ingredient edge of a Pantry.
func (c *PantryClient) QueryIngredient(pa *Pantry) *IngredientQuery {
	query := (&IngredientClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pantry.Table, pantry.FieldID, id),
			sqlgraph.To(ingredient.Table, ingredient.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, pantry.IngredientTable, pantry.IngredientColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PantryClient) Hooks() []Hook {
	return c.hooks.Pantry
}

// Interceptors returns the client interceptors.
func (c *PantryClient) Interceptors() []Interceptor {
	return c.inters.Pantry
}

func (c *PantryClient) mutate(ctx context.Context, m *PantryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PantryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PantryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PantryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PantryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Pantry mutation op: %q", m.Op())
	}
}

// RecipeClient is a client for the Recipe schema.
type RecipeClient struct {
	config
}

// NewRecipeClient returns a client for the Recipe from the given config.
func NewRecipeClient(c config) *RecipeClient {
	return &RecipeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `recipe.Hooks(f(g(h())))`.
func (c *RecipeClient) Use(hooks ...Hook) {
	c.hooks.Recipe = append(c.hooks.Recipe, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `recipe.Intercept(f(g(h())))`.
func (c *RecipeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Recipe = append(c.inters.Recipe, interceptors...)
}

// Create returns a builder for creating a Recipe entity.
func (c *RecipeClient) Create() *RecipeCreate {
	mutation := newRecipeMutation(c.config, OpCreate)
	return &RecipeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Recipe entities.
func (c *RecipeClient) CreateBulk(builders ...*RecipeCreate) *RecipeCreateBulk {
	return &RecipeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Recipe.
func (c *RecipeClient) Update() *RecipeUpdate {
	mutation := newRecipeMutation(c.config, OpUpdate)
	return &RecipeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecipeClient) UpdateOne(r *Recipe) *RecipeUpdateOne {
	mutation := newRecipeMutation(c.config, OpUpdateOne, withRecipe(r))
	return &RecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RecipeClient) UpdateOneID(id int) *RecipeUpdateOne {
	mutation := newRecipeMutation(c.config, OpUpdateOne, withRecipeID(id))
	return &RecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Recipe.
func (c *RecipeClient) Delete() *RecipeDelete {
	mutation := newRecipeMutation(c.config, OpDelete)
	return &RecipeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RecipeClient) DeleteOne(r *Recipe) *RecipeDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RecipeClient) DeleteOneID(id int) *RecipeDeleteOne {
	builder := c.Delete().Where(recipe.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RecipeDeleteOne{builder}
}

// Query returns a query builder for Recipe.
func (c *RecipeClient) Query() *RecipeQuery {
	return &RecipeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRecipe},
		inters: c.Interceptors(),
	}
}

// Get returns a Recipe entity by its id.
func (c *RecipeClient) Get(ctx context.Context, id int) (*Recipe, error) {
	return c.Query().Where(recipe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecipeClient) GetX(ctx context.Context, id int) *Recipe {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRecipeingredient queries the recipeingredient edge of a Recipe.
func (c *RecipeClient) QueryRecipeingredient(r *Recipe) *RecipeIngredientQuery {
	query := (&RecipeIngredientClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(recipe.Table, recipe.FieldID, id),
			sqlgraph.To(recipeingredient.Table, recipeingredient.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, recipe.RecipeingredientTable, recipe.RecipeingredientColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RecipeClient) Hooks() []Hook {
	return c.hooks.Recipe
}

// Interceptors returns the client interceptors.
func (c *RecipeClient) Interceptors() []Interceptor {
	return c.inters.Recipe
}

func (c *RecipeClient) mutate(ctx context.Context, m *RecipeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RecipeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RecipeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RecipeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Recipe mutation op: %q", m.Op())
	}
}

// RecipeIngredientClient is a client for the RecipeIngredient schema.
type RecipeIngredientClient struct {
	config
}

// NewRecipeIngredientClient returns a client for the RecipeIngredient from the given config.
func NewRecipeIngredientClient(c config) *RecipeIngredientClient {
	return &RecipeIngredientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `recipeingredient.Hooks(f(g(h())))`.
func (c *RecipeIngredientClient) Use(hooks ...Hook) {
	c.hooks.RecipeIngredient = append(c.hooks.RecipeIngredient, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `recipeingredient.Intercept(f(g(h())))`.
func (c *RecipeIngredientClient) Intercept(interceptors ...Interceptor) {
	c.inters.RecipeIngredient = append(c.inters.RecipeIngredient, interceptors...)
}

// Create returns a builder for creating a RecipeIngredient entity.
func (c *RecipeIngredientClient) Create() *RecipeIngredientCreate {
	mutation := newRecipeIngredientMutation(c.config, OpCreate)
	return &RecipeIngredientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RecipeIngredient entities.
func (c *RecipeIngredientClient) CreateBulk(builders ...*RecipeIngredientCreate) *RecipeIngredientCreateBulk {
	return &RecipeIngredientCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RecipeIngredient.
func (c *RecipeIngredientClient) Update() *RecipeIngredientUpdate {
	mutation := newRecipeIngredientMutation(c.config, OpUpdate)
	return &RecipeIngredientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecipeIngredientClient) UpdateOne(ri *RecipeIngredient) *RecipeIngredientUpdateOne {
	mutation := newRecipeIngredientMutation(c.config, OpUpdateOne, withRecipeIngredient(ri))
	return &RecipeIngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RecipeIngredientClient) UpdateOneID(id int) *RecipeIngredientUpdateOne {
	mutation := newRecipeIngredientMutation(c.config, OpUpdateOne, withRecipeIngredientID(id))
	return &RecipeIngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RecipeIngredient.
func (c *RecipeIngredientClient) Delete() *RecipeIngredientDelete {
	mutation := newRecipeIngredientMutation(c.config, OpDelete)
	return &RecipeIngredientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RecipeIngredientClient) DeleteOne(ri *RecipeIngredient) *RecipeIngredientDeleteOne {
	return c.DeleteOneID(ri.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RecipeIngredientClient) DeleteOneID(id int) *RecipeIngredientDeleteOne {
	builder := c.Delete().Where(recipeingredient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RecipeIngredientDeleteOne{builder}
}

// Query returns a query builder for RecipeIngredient.
func (c *RecipeIngredientClient) Query() *RecipeIngredientQuery {
	return &RecipeIngredientQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRecipeIngredient},
		inters: c.Interceptors(),
	}
}

// Get returns a RecipeIngredient entity by its id.
func (c *RecipeIngredientClient) Get(ctx context.Context, id int) (*RecipeIngredient, error) {
	return c.Query().Where(recipeingredient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecipeIngredientClient) GetX(ctx context.Context, id int) *RecipeIngredient {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRecipe queries the recipe edge of a RecipeIngredient.
func (c *RecipeIngredientClient) QueryRecipe(ri *RecipeIngredient) *RecipeQuery {
	query := (&RecipeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ri.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(recipeingredient.Table, recipeingredient.FieldID, id),
			sqlgraph.To(recipe.Table, recipe.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recipeingredient.RecipeTable, recipeingredient.RecipeColumn),
		)
		fromV = sqlgraph.Neighbors(ri.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIngredient queries the ingredient edge of a RecipeIngredient.
func (c *RecipeIngredientClient) QueryIngredient(ri *RecipeIngredient) *IngredientQuery {
	query := (&IngredientClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ri.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(recipeingredient.Table, recipeingredient.FieldID, id),
			sqlgraph.To(ingredient.Table, ingredient.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, recipeingredient.IngredientTable, recipeingredient.IngredientPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ri.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RecipeIngredientClient) Hooks() []Hook {
	return c.hooks.RecipeIngredient
}

// Interceptors returns the client interceptors.
func (c *RecipeIngredientClient) Interceptors() []Interceptor {
	return c.inters.RecipeIngredient
}

func (c *RecipeIngredientClient) mutate(ctx context.Context, m *RecipeIngredientMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RecipeIngredientCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RecipeIngredientUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RecipeIngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RecipeIngredientDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RecipeIngredient mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Ingredient, Pantry, Recipe, RecipeIngredient []ent.Hook
	}
	inters struct {
		Ingredient, Pantry, Recipe, RecipeIngredient []ent.Interceptor
	}
)
