// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent/pantry"
)

// Ingredient is the model entity for the Ingredient schema.
type Ingredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Vegan holds the value of the "vegan" field.
	Vegan bool `json:"vegan,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IngredientQuery when eager-loading is set.
	Edges             IngredientEdges `json:"edges"`
	pantry_ingredient *int
	selectValues      sql.SelectValues
}

// IngredientEdges holds the relations/edges for other nodes in the graph.
type IngredientEdges struct {
	// Pantry holds the value of the pantry edge.
	Pantry *Pantry `json:"pantry,omitempty"`
	// Recipeingredient holds the value of the recipeingredient edge.
	Recipeingredient []*RecipeIngredient `json:"recipeingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PantryOrErr returns the Pantry value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IngredientEdges) PantryOrErr() (*Pantry, error) {
	if e.loadedTypes[0] {
		if e.Pantry == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: pantry.Label}
		}
		return e.Pantry, nil
	}
	return nil, &NotLoadedError{edge: "pantry"}
}

// RecipeingredientOrErr returns the Recipeingredient value or an error if the edge
// was not loaded in eager-loading.
func (e IngredientEdges) RecipeingredientOrErr() ([]*RecipeIngredient, error) {
	if e.loadedTypes[1] {
		return e.Recipeingredient, nil
	}
	return nil, &NotLoadedError{edge: "recipeingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ingredient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ingredient.FieldVegan:
			values[i] = new(sql.NullBool)
		case ingredient.FieldID:
			values[i] = new(sql.NullInt64)
		case ingredient.FieldName:
			values[i] = new(sql.NullString)
		case ingredient.ForeignKeys[0]: // pantry_ingredient
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ingredient fields.
func (i *Ingredient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case ingredient.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case ingredient.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case ingredient.FieldVegan:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field vegan", values[j])
			} else if value.Valid {
				i.Vegan = value.Bool
			}
		case ingredient.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field pantry_ingredient", value)
			} else if value.Valid {
				i.pantry_ingredient = new(int)
				*i.pantry_ingredient = int(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Ingredient.
// This includes values selected through modifiers, order, etc.
func (i *Ingredient) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryPantry queries the "pantry" edge of the Ingredient entity.
func (i *Ingredient) QueryPantry() *PantryQuery {
	return NewIngredientClient(i.config).QueryPantry(i)
}

// QueryRecipeingredient queries the "recipeingredient" edge of the Ingredient entity.
func (i *Ingredient) QueryRecipeingredient() *RecipeIngredientQuery {
	return NewIngredientClient(i.config).QueryRecipeingredient(i)
}

// Update returns a builder for updating this Ingredient.
// Note that you need to call Ingredient.Unwrap() before calling this method if this Ingredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Ingredient) Update() *IngredientUpdateOne {
	return NewIngredientClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Ingredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Ingredient) Unwrap() *Ingredient {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ingredient is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Ingredient) String() string {
	var builder strings.Builder
	builder.WriteString("Ingredient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("vegan=")
	builder.WriteString(fmt.Sprintf("%v", i.Vegan))
	builder.WriteByte(')')
	return builder.String()
}

// Ingredients is a parsable slice of Ingredient.
type Ingredients []*Ingredient
